// Command webserver serves nunosempere.com: a plain Go HTTP server that
// reproduces the behaviour of the werc "anti-framework" setup the site used
// to run on (markdown pages, sidebar navigation tree, blog index and feeds,
// sitemap, directory listings, static files), without CGI or plan9port.
//
// Run it from the site root, or point -root at it:
//
//	cd _server && go build -o webserver . && ./webserver -root .. -addr 127.0.0.1:8080
package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

//go:embed templates/*.html
var templateFS embed.FS

//go:embed pub
var embeddedPub embed.FS

// Server holds the site configuration and shared state.
type Server struct {
	root     string // absolute path of the site directory
	baseURL  string // e.g. https://nunosempere.com
	pub      fs.FS  // files served under /pub/
	page     *template.Template
	notFound *template.Template
	blogs    []string // blog roots as URL paths with trailing slash, e.g. "/blog/"
	maxPosts int

	cfg   *configCache
	cache *renderCache
}

func main() {
	root := flag.String("root", "", "site directory (defaults to the current directory, or its parent when run from _server/)")
	addr := flag.String("addr", "127.0.0.1:8080", "address to listen on")
	baseURL := flag.String("base-url", "https://nunosempere.com", "public base URL, used in feeds and the sitemap")
	pubDir := flag.String("pub", "", "directory served under /pub/ (defaults to the files embedded in the binary)")
	maxPosts := flag.Int("max-posts", 32, "maximum number of posts listed in feeds")
	flag.Parse()

	r, err := resolveRoot(*root)
	if err != nil {
		log.Fatal(err)
	}

	s := &Server{
		root:     r,
		baseURL:  strings.TrimRight(*baseURL, "/"),
		maxPosts: *maxPosts,
		cfg:      &configCache{},
		cache:    newRenderCache(),
	}
	if *pubDir != "" {
		s.pub = os.DirFS(*pubDir)
	} else {
		sub, err := fs.Sub(embeddedPub, "pub")
		if err != nil {
			log.Fatal(err)
		}
		s.pub = sub
	}
	s.page = template.Must(template.ParseFS(templateFS, "templates/page.html"))
	s.notFound = template.Must(template.ParseFS(templateFS, "templates/404.html"))
	s.blogs = s.findBlogs()
	s.initMarkdown()

	log.Printf("serving %s on http://%s (blogs: %v)", s.root, *addr, s.blogs)
	if err := s.writeSitemapFiles(); err != nil {
		log.Printf("sitemap: %v", err)
	}

	srv := &http.Server{
		Addr:              *addr,
		Handler:           logging(s),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       120 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

// resolveRoot picks the site directory: the flag, else the current directory,
// else the parent directory (so the binary works when started from _server/).
func resolveRoot(flagRoot string) (string, error) {
	candidates := []string{flagRoot}
	if flagRoot == "" {
		candidates = []string{".", ".."}
	}
	for _, c := range candidates {
		abs, err := filepath.Abs(c)
		if err != nil {
			continue
		}
		if _, err := os.Stat(filepath.Join(abs, "_werc")); err == nil {
			return abs, nil
		}
		if _, err := os.Stat(filepath.Join(abs, "index.md")); err == nil {
			return abs, nil
		}
	}
	return "", fmt.Errorf("could not find the site root (looked for _werc/ or index.md in %v); use -root", candidates)
}

// statusWriter records the status code for logging.
type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}
	return w.ResponseWriter.Write(b)
}

func logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := &statusWriter{ResponseWriter: w}
		h.ServeHTTP(sw, r)
		if sw.status == 0 {
			sw.status = http.StatusOK
		}
		log.Printf("%d %s %s - %s - %s", sw.status, r.Method, r.RequestURI, r.Referer(), r.UserAgent())
	})
}

// ServeHTTP mirrors werc.rc's request handling: sanitize the path, apply the
// canonical redirects, pick a body handler, then wrap it in the page template.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	reqPath := sanitizePath(r.URL.Path)

	// Never expose the server's own code or the git history.
	if first := firstComponent(reqPath); first == "_server" || first == ".git" {
		s.serve404(w, reqPath)
		return
	}

	if strings.HasSuffix(reqPath, "/index") {
		redirect(w, r, strings.TrimSuffix(reqPath, "index"))
		return
	}

	local := s.localPath(reqPath)
	switch {
	case strings.HasSuffix(reqPath, "/"):
		if isDir(local) {
			local = filepath.Join(local, "index")
		} else if matches, _ := filepath.Glob(escapeGlob(local) + "*"); len(matches) > 0 {
			// /foo/ when foo.md exists -> /foo
			redirect(w, r, strings.TrimRight(reqPath, "/"))
			return
		} else {
			s.serve404(w, reqPath)
			return
		}
	case strings.HasSuffix(reqPath, ".") || strings.HasSuffix(reqPath, ",") ||
		strings.HasSuffix(reqPath, ";") || strings.HasSuffix(reqPath, ":"):
		redirect(w, r, reqPath[:len(reqPath)-1])
		return
	case isDir(local):
		redirect(w, r, reqPath+"/")
		return
	}

	pg := &pageData{ReqPath: reqPath, PageTitle: titleFromPath(reqPath)}

	// Body handlers, in werc's order of precedence. Blog year/month/day
	// archives are generated and take precedence over hand-written indexes.
	var body string
	handled := true
	switch {
	case s.isArchive(s.blogFor(reqPath), reqPath) && strings.HasSuffix(local, string(filepath.Separator)+"index"):
		body, pg.PageTitle = s.blogArchive(s.blogFor(reqPath), reqPath)
	case isFile(local + ".md"):
		html, title := s.renderMarkdownFile(local + ".md")
		body = html
		if title != "" {
			pg.PageTitle = title
		}
	case isFile(local + ".html"):
		html, title := s.renderHTMLFile(local + ".html")
		body = html
		if title != "" {
			pg.PageTitle = title
		}
	case reqPath == "/sitemap":
		body = s.sitemapPage()
	case isFile(local + ".txt"):
		body = s.renderTxtFile(local + ".txt")
	default:
		handled = false
	}

	// Blog app: may override the main handler for the blog index and feeds.
	if blog := s.blogFor(reqPath); blog != "" {
		switch reqPath {
		case blog:
			body = s.blogIndex(blog)
			pg.ExtraHeaders = template.HTML(fmt.Sprintf(
				"<link rel=\"alternate\" type=\"application/atom+xml\" title=\"ATOM\" href=\"%[1]sindex.atom\" />\n"+
					"<link rel=\"alternate\" type=\"application/rss+xml\" title=\"RSS\" href=\"%[1]sindex.rss\" />\n"+
					"<link rel=\"alternate\" type=\"application/json\" title=\"JSON\" href=\"%[1]sfeed.json\" />", blog))
			handled = true
		case blog + "index.rss":
			s.serveRSS(w, blog, reqPath)
			return
		case blog + "index.atom":
			s.serveAtom(w, blog, reqPath)
			return
		case blog + "feed.json":
			s.serveJSONFeed(w, blog, reqPath)
			return
		}
	}

	if !handled {
		switch {
		case strings.HasSuffix(local, string(filepath.Separator)+"index"):
			body = s.dirListing(reqPath)
			handled = true
		case strings.HasSuffix(local, ".html") && isFile(local):
			redirect(w, r, strings.TrimSuffix(reqPath, ".html"))
			return
		case isFile(local):
			s.serveStatic(w, r, local)
			return
		case strings.HasPrefix(reqPath, "/pub/"):
			s.servePub(w, r, strings.TrimPrefix(reqPath, "/pub/"))
			return
		default:
			s.serve404(w, reqPath)
			return
		}
	}

	pg.Body = template.HTML(body)
	s.renderPage(w, http.StatusOK, pg)
}

func redirect(w http.ResponseWriter, r *http.Request, to string) {
	http.Redirect(w, r, to, http.StatusMovedPermanently)
}

func (s *Server) serve404(w http.ResponseWriter, reqPath string) {
	var sb strings.Builder
	_ = s.notFound.Execute(&sb, map[string]string{"URL": s.baseURL + reqPath})
	pg := &pageData{ReqPath: reqPath, PageTitle: titleFromPath(reqPath), Body: template.HTML(sb.String())}
	s.renderPage(w, http.StatusNotFound, pg)
}

// serveStatic serves a regular file. Like werc, unknown and text-like types
// are sent as text/plain so that scripts, markdown sources etc. display inline.
func (s *Server) serveStatic(w http.ResponseWriter, r *http.Request, abs string) {
	f, err := os.Open(abs)
	if err != nil {
		s.serve404(w, sanitizePath(r.URL.Path))
		return
	}
	defer f.Close()
	st, err := f.Stat()
	if err != nil || !st.Mode().IsRegular() {
		s.serve404(w, sanitizePath(r.URL.Path))
		return
	}
	w.Header().Set("Content-Type", contentType(abs))
	http.ServeContent(w, r, st.Name(), st.ModTime(), f)
}

func (s *Server) servePub(w http.ResponseWriter, r *http.Request, rel string) {
	rel = path.Clean("/" + rel)[1:]
	f, err := s.pub.Open(rel)
	if err != nil {
		s.serve404(w, sanitizePath(r.URL.Path))
		return
	}
	defer f.Close()
	st, err := f.Stat()
	if err != nil || st.IsDir() {
		s.serve404(w, sanitizePath(r.URL.Path))
		return
	}
	w.Header().Set("Content-Type", contentType(rel))
	if rs, ok := f.(interface {
		Read([]byte) (int, error)
		Seek(int64, int) (int64, error)
	}); ok {
		http.ServeContent(w, r, st.Name(), st.ModTime(), rs)
		return
	}
	data, err := fs.ReadFile(s.pub, rel)
	if err != nil {
		s.serve404(w, sanitizePath(r.URL.Path))
		return
	}
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	if r.Method != http.MethodHead {
		_, _ = w.Write(data)
	}
}

func contentType(name string) string {
	ext := strings.ToLower(filepath.Ext(name))
	ct := mime.TypeByExtension(ext)
	switch ext {
	case ".css":
		return "text/css; charset=utf-8"
	case ".js":
		return "text/javascript; charset=utf-8"
	case ".ico":
		return "image/x-icon"
	}
	if ct == "" || strings.HasPrefix(ct, "text/") {
		return "text/plain; charset=utf-8"
	}
	return ct
}
