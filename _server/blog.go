package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

// post is one blog entry, found at <blog>/YYYY/MM/DD/<name>/index.md.
type post struct {
	url     string // site path with trailing slash, e.g. /blog/2025/07/07/dear-forecaster/
	date    string // YYYY/MM/DD
	abs     string // directory on disk
	modTime time.Time
}

// postList implements get_post_list: newest first, at most maxPosts.
func (s *Server) postList(blog string) []post {
	pattern := filepath.Join(s.localPath(blog), "[0-9][0-9][0-9][0-9]", "[0-9][0-9]", "[0-9][0-9]", "*")
	matches, _ := filepath.Glob(pattern)
	var posts []post
	for _, m := range matches {
		name := filepath.Base(m)
		if hiddenName(name) || forbiddenName(name) || !isDir(m) {
			continue
		}
		rel := s.relPath(m) + "/"
		st, err := os.Stat(m)
		if err != nil {
			continue
		}
		posts = append(posts, post{
			url:     rel,
			date:    strings.Join(strings.Split(strings.TrimPrefix(rel, blog), "/")[:3], "/"),
			abs:     m,
			modTime: st.ModTime(),
		})
	}
	sort.Slice(posts, func(i, j int) bool { return posts[i].url > posts[j].url })
	if len(posts) > s.maxPosts {
		posts = posts[:s.maxPosts]
	}
	return posts
}

// postLines returns the lines of a post's index.md, up to max (0 = all).
func postLines(p post, max int) []string {
	f, err := os.Open(filepath.Join(p.abs, "index.md"))
	if err != nil {
		return nil
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1024*1024), 4*1024*1024)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
		if max > 0 && len(lines) >= max {
			break
		}
	}
	return lines
}

func postTitle(p post) string {
	l := postLines(p, 1)
	if len(l) == 0 {
		return ""
	}
	return l[0]
}

var (
	h1Open   = regexp.MustCompile(`(?i)<h1(\s[^>]*)?>`)
	h1Close  = regexp.MustCompile(`(?i)</h1>`)
	hnOpen   = regexp.MustCompile(`(?i)<h[2-9](\s[^>]*)?>`)
	hnClose  = regexp.MustCompile(`(?i)</h[2-9]>`)
	indexCut = 12 // lines of markdown shown on the blog front page
	postCut  = 10 // lines taken from each post
)

// blogIndex implements blagh_body: a feed link line, the opening lines of the
// newest post(s) with the title turned into a link, a "read more" pointer, and
// a link to the sitemap. Headings are demoted as werc did (h1->h3, h2+->i).
func (s *Server) blogIndex(blog string) string {
	var b strings.Builder
	b.WriteString("<p style=\"font-size: 25px;\"><a href=\"index.rss\">RSS Feed</a>, <a href=\"../sitemap\">all content</a></p>\n")

	var lines []string
	for _, p := range s.postList(blog) {
		pl := postLines(p, postCut)
		if len(pl) > 0 {
			pl[0] = fmt.Sprintf("[%s](%s) (%s)", pl[0], p.url, p.date)
		}
		lines = append(lines, pl...)
		lines = append(lines, "",
			fmt.Sprintf("<p style=\"text-align: left; margin-right: 10%%; font-size: 28px;\"><a style=\"text-decoration: none;\" href=\"%s\"><b>&rarr;</b> &#42604; ...</a></p>", p.url),
			"", "")
		if len(lines) >= indexCut {
			break
		}
	}
	if len(lines) > indexCut {
		lines = lines[:indexCut]
	}
	out := renderMarkdown([]byte(strings.Join(lines, "\n") + "\n"))
	// Same order as werc's sed: h2..h9 first, then h1 -> h3.
	out = hnOpen.ReplaceAllString(out, "<i>")
	out = hnClose.ReplaceAllString(out, "</i>")
	out = h1Open.ReplaceAllString(out, "<h3>")
	out = h1Close.ReplaceAllString(out, "</h3>")
	b.WriteString(out)
	b.WriteString("\n\n\n\n\n")
	b.WriteString("<p><a href=\"../sitemap\">More content</a></p>\n")
	return b.String()
}

// summary renders the first lines of a post without its title (werc used the
// first 10 lines for RSS/Atom and the whole post for the JSON feed).
func summary(p post, max int) string {
	lines := stripMdTitle(postLines(p, max))
	return renderMarkdown([]byte(strings.Join(lines, "\n") + "\n"))
}

func (p post) pubDate() time.Time {
	t, err := time.Parse("2006/01/02", p.date)
	if err != nil {
		return p.modTime
	}
	return t.UTC()
}

func lastBuild(posts []post) time.Time {
	if len(posts) == 0 {
		return time.Now()
	}
	return posts[0].modTime
}

func cdata(s string) string {
	return "<![CDATA[" + strings.ReplaceAll(s, "]]>", "]]]]><![CDATA[>") + "]]>"
}

func (s *Server) serveRSS(w http.ResponseWriter, blog, reqPath string) {
	cfg := s.config()
	posts := s.postList(blog)
	self := s.baseURL + reqPath
	var b bytes.Buffer
	b.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n\n")
	b.WriteString("<rss version=\"2.0\" xmlns:atom=\"http://www.w3.org/2005/Atom\">\n    <channel>\n")
	fmt.Fprintf(&b, "        <atom:link href=\"%s\" rel=\"self\" type=\"application/rss+xml\" />\n", self)
	fmt.Fprintf(&b, "        <title>%s</title>\n", cdata(cfg.Title))
	fmt.Fprintf(&b, "        <link>%s</link>\n", self)
	b.WriteString("        <description><![CDATA[]]></description>\n")
	b.WriteString("        <language>en-us</language>\n")
	b.WriteString("        <generator><![CDATA[Go, in the spirit of Tom Duff's rc and Kris Maglione's clever hackery]]></generator>\n")
	fmt.Fprintf(&b, "<lastBuildDate>%s</lastBuildDate>\n", lastBuild(posts).UTC().Format(time.RFC1123Z))
	for _, p := range posts {
		url := s.baseURL + p.url
		b.WriteString("        <item>\n")
		fmt.Fprintf(&b, "            <title>%s</title>\n", cdata(postTitle(p)))
		b.WriteString("<author><![CDATA[nuno.sempere@protonmail.com (Nuño Sempere)]]></author>\n")
		fmt.Fprintf(&b, "            <link>%s</link>\n", url)
		fmt.Fprintf(&b, "            <guid isPermaLink=\"true\">%s</guid>\n", url)
		fmt.Fprintf(&b, "            <pubDate>%s</pubDate>\n", p.pubDate().Format(time.RFC1123Z))
		fmt.Fprintf(&b, "            <description> %s </description>\n", html.EscapeString(summary(p, postCut)))
		b.WriteString("        </item>\n")
	}
	b.WriteString("    </channel>\n</rss>\n")
	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	_, _ = w.Write(b.Bytes())
}

func (s *Server) serveAtom(w http.ResponseWriter, blog, reqPath string) {
	cfg := s.config()
	posts := s.postList(blog)
	self := s.baseURL + reqPath
	var b bytes.Buffer
	b.WriteString("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n\n")
	b.WriteString("<feed xmlns=\"http://www.w3.org/2005/Atom\"\n    xmlns:thr=\"http://purl.org/syndication/thread/1.0\">\n\n")
	fmt.Fprintf(&b, "    <link rel=\"self\" href=\"%s\"/>\n", self)
	fmt.Fprintf(&b, "    <id>%s</id>\n", self)
	b.WriteString("    <icon><![CDATA[/favicon.ico]]></icon>\n\n")
	fmt.Fprintf(&b, "    <title>%s</title>\n", cdata(cfg.Title))
	fmt.Fprintf(&b, "    <subtitle>%s</subtitle>\n\n", cdata(cfg.SubTitle))
	fmt.Fprintf(&b, "    <updated>%s</updated>\n", lastBuild(posts).UTC().Format(time.RFC3339))
	fmt.Fprintf(&b, "    <link href=\"%s\"/>\n\n", s.baseURL+blog)
	for _, p := range posts {
		url := s.baseURL + p.url
		b.WriteString("    <entry>\n")
		fmt.Fprintf(&b, "        <id>%s</id>\n", url)
		fmt.Fprintf(&b, "        <link href=\"%s\"/>\n", url)
		fmt.Fprintf(&b, "        <title>%s</title>\n", cdata(postTitle(p)))
		b.WriteString("        <author><name><![CDATA[Nuño Sempere]]></name></author>\n")
		fmt.Fprintf(&b, "        <content type=\"html\">%s</content>\n", cdata(summary(p, postCut)))
		fmt.Fprintf(&b, "        <published>%s</published>\n", p.pubDate().Format(time.RFC3339))
		fmt.Fprintf(&b, "        <updated>%s</updated>\n", p.modTime.UTC().Format(time.RFC3339))
		b.WriteString("    </entry>\n\n")
	}
	b.WriteString("</feed>\n")
	w.Header().Set("Content-Type", "application/atom+xml; charset=utf-8")
	_, _ = w.Write(b.Bytes())
}

func (s *Server) serveJSONFeed(w http.ResponseWriter, blog, reqPath string) {
	cfg := s.config()
	type item struct {
		ID          string `json:"id"`
		URL         string `json:"url"`
		Title       string `json:"title"`
		ContentHTML string `json:"content_html"`
		Published   string `json:"date_published"`
	}
	feed := struct {
		Version     string `json:"version"`
		Title       string `json:"title"`
		HomePageURL string `json:"home_page_url"`
		FeedURL     string `json:"feed_url"`
		Items       []item `json:"items"`
	}{
		Version:     "https://jsonfeed.org/version/1",
		Title:       cfg.Title,
		HomePageURL: s.baseURL,
		FeedURL:     s.baseURL + reqPath,
		Items:       []item{},
	}
	for _, p := range s.postList(blog) {
		url := s.baseURL + p.url
		feed.Items = append(feed.Items, item{
			ID: url, URL: url, Title: postTitle(p),
			ContentHTML: summary(p, 0),
			Published:   p.pubDate().Format(time.RFC3339),
		})
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	_ = enc.Encode(feed)
}
