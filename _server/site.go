package main

import (
	"bufio"
	"bytes"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

// --- request path handling -------------------------------------------------

var (
	multiSlash = regexp.MustCompile(`/+`)
	multiDot   = regexp.MustCompile(`\.\.+`)
	wordDash   = regexp.MustCompile(`([a-z])-([a-z])`)
)

// allowedPathByte reports whether c may appear in a request path, matching
// werc's forbidden_uri_chars='[^a-zA-Z0-9_+\-\/\.,:]'.
func allowedPathByte(c byte) bool {
	switch {
	case c >= 'a' && c <= 'z', c >= 'A' && c <= 'Z', c >= '0' && c <= '9':
		return true
	}
	return strings.IndexByte("_+-/.,:", c) >= 0
}

// sanitizePath strips forbidden characters, collapses repeated slashes and
// runs of dots (which also neutralises ".." traversal), as werc.rc does.
func sanitizePath(p string) string {
	var b strings.Builder
	for i := 0; i < len(p); i++ {
		if allowedPathByte(p[i]) {
			b.WriteByte(p[i])
		}
	}
	s := multiSlash.ReplaceAllString(b.String(), "/")
	s = multiDot.ReplaceAllString(s, ".")
	if !strings.HasPrefix(s, "/") {
		s = "/" + s
	}
	return s
}

func firstComponent(p string) string {
	p = strings.TrimPrefix(p, "/")
	if i := strings.IndexByte(p, '/'); i >= 0 {
		return p[:i]
	}
	return p
}

// localPath maps a sanitized request path to a filesystem path under root.
func (s *Server) localPath(reqPath string) string {
	return filepath.Join(s.root, filepath.FromSlash(path.Clean("/"+reqPath)))
}

// relPath converts an absolute path under root back to a site URL path.
func (s *Server) relPath(abs string) string {
	rel, err := filepath.Rel(s.root, abs)
	if err != nil || rel == "." {
		return "/"
	}
	return "/" + filepath.ToSlash(rel)
}

func escapeGlob(p string) string {
	var b strings.Builder
	for _, c := range p {
		if strings.ContainsRune(`*?[\`, c) {
			b.WriteByte('\\')
		}
		b.WriteRune(c)
	}
	return b.String()
}

func isDir(p string) bool {
	st, err := os.Stat(p)
	return err == nil && st.IsDir()
}

func isFile(p string) bool {
	st, err := os.Stat(p)
	return err == nil && st.Mode().IsRegular()
}

// titleFromPath builds werc's fallback page title from the URL components:
// "blog - 2025 - 07 - 07 - dear forecaster".
func titleFromPath(reqPath string) string {
	parts := strings.FieldsFunc(reqPath, func(r rune) bool { return r == '/' })
	if len(parts) == 0 {
		return ""
	}
	t := strings.Join(parts, " - ")
	t = wordDash.ReplaceAllString(t, "$1 $2")
	return strings.ReplaceAll(t, "_", " ")
}

// --- directory entry filters (werc's $dirfilter) --------------------------

// hiddenName reports whether a path component is hidden from navigation,
// listings and the sitemap: names starting with "." or "_" (and _werc).
func hiddenName(name string) bool {
	return name == "" || name[0] == '.' || name[0] == '_'
}

// hiddenPath reports whether any component of a site-relative path is hidden.
// werc applied the filter to full paths, so listings inside hidden
// directories come out empty.
func hiddenPath(rel string) bool {
	for _, c := range strings.Split(rel, "/") {
		if c != "" && hiddenName(c) {
			return true
		}
	}
	return false
}

func forbiddenName(name string) bool {
	for i := 0; i < len(name); i++ {
		if !allowedPathByte(name[i]) {
			return true
		}
	}
	return false
}

// filteredFile reports whether a non-directory entry is dropped by dirfilter.
func filteredFile(name string) bool {
	switch name {
	case "index.md", "index.html", "index.txt", "index.tpl", "robots.txt", "sitemap.txt", "sitemap.xml":
		return true
	}
	return false
}

func hasPageExt(name string) bool {
	switch strings.ToLower(filepath.Ext(name)) {
	case ".md", ".html", ".txt":
		return true
	}
	return false
}

// stripPageExt removes a .md/.html/.txt extension (werc's $dirclean).
func stripPageExt(name string) string {
	if hasPageExt(name) {
		return strings.TrimSuffix(name, filepath.Ext(name))
	}
	return name
}

// --- titles --------------------------------------------------------------

func isSetextUnderline(line string) bool {
	line = strings.TrimRight(line, "\r")
	if len(line) < 2 {
		return false
	}
	return strings.Trim(line, "=") == ""
}

// mdTitle implements werc's get_md_title: the first line is the title when the
// second line is a run of "=" and the third line is blank (or absent).
func mdTitle(src []byte) string {
	sc := bufio.NewScanner(bytes.NewReader(src))
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	var lines []string
	for sc.Scan() && len(lines) < 3 {
		lines = append(lines, sc.Text())
	}
	if len(lines) < 2 || !isSetextUnderline(lines[1]) {
		return ""
	}
	if len(lines) == 3 && strings.TrimSpace(lines[2]) != "" {
		return ""
	}
	return strings.TrimSpace(lines[0])
}

// hasSetextTitle reports whether src starts with "Title\n===\n\n" so that the
// first three lines can be stripped (werc's strip_title_from_md_file).
func stripMdTitle(lines []string) []string {
	if len(lines) >= 3 && isSetextUnderline(lines[1]) && strings.TrimSpace(lines[2]) == "" {
		return lines[3:]
	}
	return lines
}

var (
	htmlTitleRe   = regexp.MustCompile(`(?i)<title> *([^<]+?) *(</title>.*)?$`)
	htmlLeadText  = regexp.MustCompile(`^(<[^>]+>)*([^<]+)`)
	htmlTitleScan = 32
)

// htmlTitle implements werc's get_html_title: a <title> within the first 32
// lines, else the first run of text not inside a tag.
func htmlTitle(src []byte) string {
	lines := strings.SplitN(string(src), "\n", htmlTitleScan+1)
	if len(lines) > htmlTitleScan {
		lines = lines[:htmlTitleScan]
	}
	for _, l := range lines {
		if m := htmlTitleRe.FindStringSubmatch(l); m != nil {
			return strings.TrimSpace(m[1])
		}
	}
	// Like werc's sed: only the first matching line counts, even if it is
	// just indentation (in which case there is no title).
	for _, l := range lines {
		if m := htmlLeadText.FindStringSubmatch(l); m != nil {
			return strings.TrimSpace(m[2])
		}
	}
	return ""
}

// fileTitle implements get_file_title for a file or a directory (via its index).
func fileTitle(abs string, dir bool) string {
	if dir {
		if b, err := os.ReadFile(filepath.Join(abs, "index.md")); err == nil {
			return mdTitle(b)
		}
		if b, err := os.ReadFile(filepath.Join(abs, "index.html")); err == nil {
			return htmlTitle(b)
		}
		return ""
	}
	switch strings.ToLower(filepath.Ext(abs)) {
	case ".md":
		if b, err := os.ReadFile(abs); err == nil {
			return mdTitle(b)
		}
	case ".html":
		if b, err := os.ReadFile(abs); err == nil {
			return htmlTitle(b)
		}
	}
	return ""
}

// --- _werc/config ----------------------------------------------------------

type siteConfig struct {
	Title    string
	SubTitle string
}

type configCache struct {
	mu      sync.Mutex
	modTime time.Time
	size    int64
	cfg     siteConfig
	loaded  bool
}

// config reads siteTitle/siteSubTitle from _werc/config, re-reading when the
// file changes (makeconfig.sh rotates the title periodically).
func (s *Server) config() siteConfig {
	c := s.cfg
	c.mu.Lock()
	defer c.mu.Unlock()
	p := filepath.Join(s.root, "_werc", "config")
	st, err := os.Stat(p)
	if err != nil {
		if !c.loaded {
			c.cfg = siteConfig{Title: "Measure is unceasing"}
			c.loaded = true
		}
		return c.cfg
	}
	if c.loaded && st.ModTime().Equal(c.modTime) && st.Size() == c.size {
		return c.cfg
	}
	b, err := os.ReadFile(p)
	if err != nil {
		return c.cfg
	}
	cfg := siteConfig{}
	for _, line := range strings.Split(string(b), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		k, v, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		v = rcUnquote(strings.TrimSpace(v))
		switch strings.TrimSpace(k) {
		case "siteTitle":
			cfg.Title = v
		case "siteSubTitle":
			cfg.SubTitle = v
		}
	}
	c.cfg, c.modTime, c.size, c.loaded = cfg, st.ModTime(), st.Size(), true
	return c.cfg
}

// rcUnquote removes rc-shell single quotes (” inside quotes is a literal ').
func rcUnquote(v string) string {
	if len(v) >= 2 && v[0] == '\'' && v[len(v)-1] == '\'' {
		return strings.ReplaceAll(v[1:len(v)-1], "''", "'")
	}
	return v
}

// findBlogs looks for directories whose _werc/config enables the blog app.
func (s *Server) findBlogs() []string {
	var blogs []string
	_ = filepath.WalkDir(s.root, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			switch d.Name() {
			case ".git", "_server":
				return filepath.SkipDir
			}
			return nil
		}
		if d.Name() != "config" || filepath.Base(filepath.Dir(p)) != "_werc" {
			return nil
		}
		b, err := os.ReadFile(p)
		if err != nil {
			return nil
		}
		for _, line := range strings.Split(string(b), "\n") {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "conf_enable_blog") {
				rel := s.relPath(filepath.Dir(filepath.Dir(p)))
				if !strings.HasSuffix(rel, "/") {
					rel += "/"
				}
				blogs = append(blogs, rel)
				break
			}
		}
		return nil
	})
	return blogs
}

// blogFor returns the blog root that contains reqPath, or "".
func (s *Server) blogFor(reqPath string) string {
	for _, b := range s.blogs {
		if strings.HasPrefix(reqPath, b) {
			return b
		}
	}
	return ""
}

// footer returns the contents of _werc/lib/footer.inc, if present.
func (s *Server) footer() string {
	b, err := os.ReadFile(filepath.Join(s.root, "_werc", "lib", "footer.inc"))
	if err != nil {
		return ""
	}
	return string(b)
}
