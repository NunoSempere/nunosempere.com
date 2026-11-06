package main

import (
	"fmt"
	"html"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

// navTree implements werc's nav_tree: list every ancestor directory of the
// request (root included) and render the union as a nested <ul>, marking the
// entries on the current path with class="thisPage".
func (s *Server) navTree(reqPath string) string {
	dirs := []string{"/"}
	acc := ""
	for _, c := range strings.Split(strings.Trim(reqPath, "/"), "/") {
		if c == "" {
			continue
		}
		acc += "/" + c
		dirs = append(dirs, acc)
	}

	set := map[string]bool{}
	for _, d := range dirs {
		if hiddenPath(d) {
			continue
		}
		abs := s.localPath(d)
		entries, err := os.ReadDir(abs)
		if err != nil {
			continue // last element may be a file, as in werc
		}
		for _, e := range entries {
			name := e.Name()
			if hiddenName(name) || forbiddenName(name) || filteredFile(name) {
				continue
			}
			if isDir(filepath.Join(abs, name)) {
				set[path.Join(d, name)+"/"] = true
			} else if hasPageExt(name) {
				set[path.Join(d, stripPageExt(name))] = true
			}
		}
	}
	paths := make([]string, 0, len(set))
	for p := range set {
		paths = append(paths, p)
	}
	sort.Strings(paths)

	var b strings.Builder
	b.WriteString("<ul>\n")
	last := 2
	for _, p := range paths {
		d := ""
		if strings.HasSuffix(p, "/") {
			d = "/"
		}
		bare := strings.TrimSuffix(p, "/")
		fields := strings.Split(bare, "/")
		nf := len(fields)
		for i := 0; i < nf-last; i++ {
			b.WriteString("<li><ul>\n")
		}
		for i := 0; i < last-nf; i++ {
			b.WriteString("</ul></li>\n")
		}
		last = nf
		bname := strings.NewReplacer("-", " ", "_", " ").Replace(fields[nf-1] + d)
		pa := p
		if d == "" {
			pa += "/"
		}
		if strings.HasPrefix(reqPath+"/", pa) {
			fmt.Fprintf(&b, "<li><a href=\"%s\" class=\"thisPage\">&raquo;<i> %s</i></a></li>\n", p, html.EscapeString(bname))
		} else {
			fmt.Fprintf(&b, "<li><a href=\"%s\">&rsaquo; %s</a></li>\n", p, html.EscapeString(bname))
		}
	}
	for i := 0; i < last-2; i++ {
		b.WriteString("</ul></li>\n")
	}
	b.WriteString("</ul>\n")
	return b.String()
}

// dirListing implements werc's dir_listing_handler for directories that have
// no index page.
func (s *Server) dirListing(reqPath string) string {
	d := strings.TrimSuffix(reqPath, "/")
	heading := strings.ReplaceAll(d, "/", " / ")

	var b strings.Builder
	fmt.Fprintf(&b, "<h1 class=\"dir-list-head\">%s</h1> <ul class=\"dir-list\">\n", html.EscapeString(heading))
	if !hiddenPath(reqPath) {
		abs := s.localPath(reqPath)
		entries, _ := os.ReadDir(abs)
		names := make([]string, 0, len(entries))
		for _, e := range entries {
			name := e.Name()
			if hiddenName(name) || forbiddenName(name) || filteredFile(name) {
				continue
			}
			if isDir(filepath.Join(abs, name)) {
				names = append(names, name+"/")
			} else {
				names = append(names, stripPageExt(name))
			}
		}
		sort.Strings(names)
		for _, n := range names {
			fmt.Fprintf(&b, "<li><a href=\"%s\">%s</a></li>\n", n, html.EscapeString(n))
		}
	}
	b.WriteString("</ul>\n")
	return b.String()
}
