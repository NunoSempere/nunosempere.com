// Command webserver renders the site in source/ as a static site (-build) or
// serves it directly (-addr). It reproduces the behaviour of the werc setup
// the site used to run on: markdown pages, sidebar navigation tree, blog index
// and feeds, sitemap, directory listings and static files.
//
//	webserver -root source -build build          # generate the static site
//	webserver -root source -addr 127.0.0.1:8080  # development server
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

	rootResolved   string // root with symlinks resolved, for the escape check
	followSymlinks bool   // allow symlinks that point outside root
	gen            *ttlCache

	cfg   *configCache
	cache *renderCache
}

func main() {
	root := flag.String("root", "", "site directory (defaults to ./source, ../source, . or ..)")
	addr := flag.String("addr", "127.0.0.1:8080", "address to listen on")
	buildDir := flag.String("build", "", "write the site as static files into this directory and exit")
	baseURL := flag.String("base-url", "https://nunosempere.com", "public base URL, used in feeds, sitemaps and 404 pages")
	pubDir := flag.String("pub", "", "directory served under /pub/ (defaults to the files embedded in the binary)")
	maxPosts := flag.Int("max-posts", 32, "maximum number of posts listed in feeds")
	followSymlinks := flag.Bool("follow-symlinks", false, "serve files reached through symlinks that point outside the site directory")
	cacheTTL := flag.Duration("cache-ttl", 60*time.Second, "how long generated pages (blog index, feeds, sitemaps) are cached")
	flag.Parse()

	r, err := resolveRoot(*root)
	if err != nil {
		log.Fatal(err)
	}
	resolved, err := filepath.EvalSymlinks(r)
	if err != nil {
		log.Fatal(err)
	}
	s := &Server{
		root:           r,
		rootResolved:   resolved,
		followSymlinks: *followSymlinks,
		baseURL:        strings.TrimRight(*baseURL, "/"),
		maxPosts:       *maxPosts,
		cfg:            &configCache{},
		cache:          newRenderCache(),
		gen:            newTTLCache(*cacheTTL),
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

	if *buildDir != "" {
		if err := s.build(*buildDir); err != nil {
			log.Fatal(err)
		}
		return
	}

	log.Printf("serving %s on http://%s (blogs: %v)", s.root, *addr, s.blogs)
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

// resolveRoot picks the site directory: the flag, else the first of
// ./source, ../source, . and .. that looks like a site.
func resolveRoot(flagRoot string) (string, error) {
	candidates := []string{flagRoot}
	if flagRoot == "" {
		candidates = []string{"source", "../source", ".", ".."}
	}
	for _, c := range candidates {
		abs, err := filepath.Abs(c)
		if err != nil {
			continue
		}
		if _, err := os.Stat(filepath.Join(abs, "index.md")); err == nil {
			return abs, nil
		}
		if _, err := os.Stat(filepath.Join(abs, "_werc")); err == nil {
			return abs, nil
		}
	}
	return "", fmt.Errorf("could not find the site root (looked for index.md in %v); use -root", candidates)
}

// response is the outcome of handling one URL path. Exactly one of redirect,
// file, pub or body is set. It is served over HTTP by ServeHTTP and written to
// disk by build.
type response struct {
	status      int
	redirect    string // target of a 301
	file        string // absolute path of a static file in the site
	pub         string // path of a file in the /pub/ filesystem
	contentType string
	body        []byte
}

func redirectTo(to string) response {
	return response{status: http.StatusMovedPermanently, redirect: to}
}

func htmlResponse(status int, body []byte) response {
	return response{status: status, contentType: "text/html; charset=utf-8", body: body}
}

// handle mirrors werc.rc's request handling: canonical redirects, then a body
// handler chosen by what exists on disk, wrapped in the page template.
func (s *Server) handle(reqPath string) response {
	// Never expose the git history.
	if firstComponent(reqPath) == ".git" {
		return s.notFoundResponse(reqPath)
	}
	if strings.HasSuffix(reqPath, "/index") {
		return redirectTo(strings.TrimSuffix(reqPath, "index"))
	}

	local := s.localPath(reqPath)
	switch {
	case strings.HasSuffix(reqPath, "/"):
		if s.browsable(local) {
			local = filepath.Join(local, "index")
		} else if s.servable(local+".md") || s.servable(local+".html") || s.servable(local+".txt") {
			// /foo/ when foo.md exists -> /foo. (werc globbed foo*, which let
			// anyone enumerate hidden names one prefix at a time.)
			return redirectTo(strings.TrimRight(reqPath, "/"))
		} else {
			return s.notFoundResponse(reqPath)
		}
	case strings.HasSuffix(reqPath, ".") || strings.HasSuffix(reqPath, ",") ||
		strings.HasSuffix(reqPath, ";") || strings.HasSuffix(reqPath, ":"):
		return redirectTo(reqPath[:len(reqPath)-1])
	case s.browsable(local):
		return redirectTo(reqPath + "/")
	}

	// Machine-readable sitemaps and a default robots.txt, generated on demand.
	// A real file of the same name in the site root wins.
	if !s.servable(local) {
		switch reqPath {
		case "/sitemap.txt":
			return response{status: 200, contentType: "text/plain; charset=utf-8", body: s.sitemapTxt()}
		case "/sitemap.xml":
			return response{status: 200, contentType: "application/xml; charset=utf-8", body: s.sitemapXML()}
		case "/sitemap.gz":
			return response{status: 200, contentType: "application/gzip", body: s.sitemapGz()}
		case "/robots.txt":
			return response{status: 200, contentType: "text/plain; charset=utf-8", body: s.robots()}
		}
	}

	pg := &pageData{ReqPath: reqPath, PageTitle: titleFromPath(reqPath)}
	isIndex := strings.HasSuffix(local, string(filepath.Separator)+"index")
	blog := s.blogFor(reqPath)

	// Body handlers, in werc's order of precedence. Blog year/month/day
	// archives are generated and take precedence over hand-written indexes.
	var body string
	handled := true
	switch {
	case s.isArchive(blog, reqPath) && isIndex:
		body, pg.PageTitle = s.blogArchive(blog, reqPath)
	case s.servable(local + ".md"):
		html, title := s.renderMarkdownFile(local + ".md")
		body = html
		if title != "" {
			pg.PageTitle = title
		}
	case s.servable(local + ".html"):
		html, title := s.renderHTMLFile(local + ".html")
		body = html
		if title != "" {
			pg.PageTitle = title
		}
	case reqPath == "/sitemap":
		body = s.sitemapPage()
	case s.servable(local + ".txt"):
		body = s.renderTxtFile(local + ".txt")
	default:
		handled = false
	}

	// Blog app: may override the main handler for the blog index and feeds.
	if blog != "" {
		switch reqPath {
		case blog:
			body = s.blogIndex(blog)
			pg.ExtraHeaders = template.HTML(fmt.Sprintf(
				"<link rel=\"alternate\" type=\"application/atom+xml\" title=\"ATOM\" href=\"%[1]sindex.atom\" />\n"+
					"<link rel=\"alternate\" type=\"application/rss+xml\" title=\"RSS\" href=\"%[1]sindex.rss\" />\n"+
					"<link rel=\"alternate\" type=\"application/json\" title=\"JSON\" href=\"%[1]sfeed.json\" />", blog))
			handled = true
		case blog + "index.rss":
			return response{status: 200, contentType: "text/xml; charset=utf-8", body: s.rss(blog, reqPath)}
		case blog + "index.atom":
			return response{status: 200, contentType: "application/atom+xml; charset=utf-8", body: s.atom(blog, reqPath)}
		case blog + "feed.json":
			return response{status: 200, contentType: "application/json; charset=utf-8", body: s.jsonFeed(blog, reqPath)}
		}
	}

	if !handled {
		switch {
		case isIndex:
			body = s.dirListing(reqPath)
		case strings.HasSuffix(local, ".html") && s.servable(local):
			return redirectTo(strings.TrimSuffix(reqPath, ".html"))
		case s.servable(local):
			return response{status: 200, file: local}
		case strings.HasPrefix(reqPath, "/pub/"):
			rel := path.Clean("/" + strings.TrimPrefix(reqPath, "/pub/"))[1:]
			if st, err := fs.Stat(s.pub, rel); err == nil && !st.IsDir() {
				return response{status: 200, pub: rel}
			}
			return s.notFoundResponse(reqPath)
		default:
			return s.notFoundResponse(reqPath)
		}
	}

	pg.Body = template.HTML(body)
	out, err := s.renderPage(pg)
	if err != nil {
		return response{status: http.StatusInternalServerError, contentType: "text/plain; charset=utf-8", body: []byte(err.Error())}
	}
	return htmlResponse(http.StatusOK, out)
}

func (s *Server) notFoundResponse(reqPath string) response {
	var sb strings.Builder
	_ = s.notFound.Execute(&sb, map[string]any{"URL": s.baseURL + reqPath, "Static": false})
	pg := &pageData{ReqPath: reqPath, PageTitle: titleFromPath(reqPath), Body: template.HTML(sb.String())}
	out, err := s.renderPage(pg)
	if err != nil {
		return response{status: http.StatusInternalServerError, contentType: "text/plain; charset=utf-8", body: []byte(err.Error())}
	}
	return htmlResponse(http.StatusNotFound, out)
}

// --- HTTP front end --------------------------------------------------------

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
		w.Header().Set("X-Content-Type-Options", "nosniff")
		sw := &statusWriter{ResponseWriter: w}
		h.ServeHTTP(sw, r)
		if sw.status == 0 {
			sw.status = http.StatusOK
		}
		// %q keeps a crafted User-Agent or Referer from forging log lines.
		log.Printf("%d %s %q - %q - %q", sw.status, r.Method, r.RequestURI, r.Referer(), r.UserAgent())
	})
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	reqPath := sanitizePath(r.URL.Path)
	resp := s.handle(reqPath)
	switch {
	case resp.redirect != "":
		http.Redirect(w, r, resp.redirect, http.StatusMovedPermanently)
	case resp.file != "":
		s.serveStatic(w, r, resp.file)
	case resp.pub != "":
		s.servePub(w, r, resp.pub)
	default:
		w.Header().Set("Content-Type", resp.contentType)
		w.WriteHeader(resp.status)
		_, _ = w.Write(resp.body)
	}
}

// serveStatic serves a regular file. Like werc, unknown and text-like types
// are sent as text/plain so that scripts, markdown sources etc. display inline.
func (s *Server) serveStatic(w http.ResponseWriter, r *http.Request, abs string) {
	f, err := os.Open(abs)
	if err != nil {
		s.writeResponse(w, s.notFoundResponse(sanitizePath(r.URL.Path)))
		return
	}
	defer f.Close()
	st, err := f.Stat()
	if err != nil || !st.Mode().IsRegular() {
		s.writeResponse(w, s.notFoundResponse(sanitizePath(r.URL.Path)))
		return
	}
	w.Header().Set("Content-Type", contentType(abs))
	http.ServeContent(w, r, st.Name(), st.ModTime(), f)
}

func (s *Server) servePub(w http.ResponseWriter, r *http.Request, rel string) {
	data, err := fs.ReadFile(s.pub, rel)
	if err != nil {
		s.writeResponse(w, s.notFoundResponse(sanitizePath(r.URL.Path)))
		return
	}
	w.Header().Set("Content-Type", contentType(rel))
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	if r.Method != http.MethodHead {
		_, _ = w.Write(data)
	}
}

func (s *Server) writeResponse(w http.ResponseWriter, resp response) {
	w.Header().Set("Content-Type", resp.contentType)
	w.WriteHeader(resp.status)
	_, _ = w.Write(resp.body)
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
