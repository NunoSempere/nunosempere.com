package main

import (
	"compress/gzip"
	"fmt"
	"html"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

// child is an entry considered by the sitemap: a directory, or a .md/.html/.txt page.
type child struct {
	rel     string // site path; directories end in "/"
	abs     string
	isDir   bool
	lastmod string
}

// children implements the sitemap's get_children: subdirectories and pages of a
// directory, with hidden/underscore entries, index pages and sitemap/robots
// files removed, sorted by name.
func (s *Server) children(dirRel string) []child {
	abs := s.localPath(dirRel)
	entries, err := os.ReadDir(abs)
	if err != nil {
		return nil
	}
	var out []child
	for _, e := range entries {
		name := e.Name()
		if hiddenName(name) || forbiddenName(name) || filteredFile(name) {
			continue
		}
		full := filepath.Join(abs, name)
		st, err := os.Stat(full)
		if err != nil {
			continue
		}
		c := child{abs: full, isDir: st.IsDir(), lastmod: st.ModTime().Format("2006-01-02")}
		if c.isDir {
			c.rel = path.Join(dirRel, name) + "/"
		} else if hasPageExt(name) {
			c.rel = path.Join(dirRel, name)
		} else {
			continue
		}
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool { return path.Base(out[i].rel) < path.Base(out[j].rel) })
	return out
}

type sitemapResult struct {
	html string
	urls []string
	xml  []string
}

// sitemapDir implements the recursive listDir of werc's sitemap.tpl, including
// its quirks: single-child chains are joined with "·" instead of nesting.
func (s *Server) sitemapDir(dirRel string, r *sitemapResult, b *strings.Builder) {
	siblings := s.children(dirRel)
	n := len(siblings)
	if n != 0 && n != 1 {
		b.WriteString("<ul class=\"sitemap-list\" style=\"list-style-type: none; list-style-position: outside;\" >\n")
	}
	for _, c := range siblings {
		own := 0
		if c.isDir {
			own = len(s.children(c.rel))
		}
		title := fileTitle(c.abs, c.isDir)
		url := c.rel
		if !c.isDir {
			url = stripPageExt(url)
		}
		dirname := strings.NewReplacer("-", " ", "_", " ").Replace(path.Base(strings.TrimSuffix(url, "/")))

		if n != 1 {
			b.WriteString("<li style=\"list-style-type: none; list-style-position: outside; \">\n")
		}
		if title != "" {
			fmt.Fprintf(b, "<a href=\"%s\"  >%s</a>\n", url, title)
		} else if dirname != "images" {
			fmt.Fprintf(b, "<a href=\"%s\" style=\"text-decoration: none;\"  >%s</a>\n", url, html.EscapeString(dirname))
		}
		if own != 1 {
			b.WriteString("</li>\n")
		} else {
			b.WriteString("·\n")
		}
		full := s.baseURL + url
		r.urls = append(r.urls, full)
		r.xml = append(r.xml, fmt.Sprintf("<url><loc>%s</loc><lastmod>%s</lastmod></url>", html.EscapeString(full), c.lastmod))
		if c.isDir {
			s.sitemapDir(c.rel, r, b)
		}
	}
	if n != 0 && n != 1 {
		b.WriteString("</ul>\n")
	}
}

func (s *Server) buildSitemap() *sitemapResult {
	r := &sitemapResult{}
	var b strings.Builder
	b.WriteString("<h1>Site map</h1>\n\n")
	s.sitemapDir("/", r, &b)
	r.html = b.String()
	return r
}

// sitemapPage renders /sitemap and refreshes sitemap.txt and sitemap.gz in the
// site root, as werc's template did.
func (s *Server) sitemapPage() string {
	r := s.buildSitemap()
	if err := s.writeSitemap(r); err != nil {
		fmt.Fprintln(os.Stderr, "sitemap:", err)
	}
	return r.html
}

func (s *Server) writeSitemapFiles() error {
	return s.writeSitemap(s.buildSitemap())
}

func (s *Server) writeSitemap(r *sitemapResult) error {
	if len(r.urls) == 0 {
		return nil
	}
	txt := filepath.Join(s.root, "sitemap.txt")
	if err := writeAtomic(txt, []byte(strings.Join(r.urls, "\n")+"\n")); err != nil {
		return err
	}
	var gz strings.Builder
	zw := gzip.NewWriter(&gz)
	fmt.Fprintf(zw, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">\n")
	for _, x := range r.xml {
		fmt.Fprintln(zw, x)
	}
	fmt.Fprintln(zw, "</urlset>")
	if err := zw.Close(); err != nil {
		return err
	}
	return writeAtomic(filepath.Join(s.root, "sitemap.gz"), []byte(gz.String()))
}

func writeAtomic(dst string, data []byte) error {
	tmp := dst + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, dst)
}
