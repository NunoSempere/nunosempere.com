package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"html"
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
	posts := s.allPosts(blog)
	if len(posts) > s.maxPosts {
		posts = posts[:s.maxPosts]
	}
	return posts
}

// allPosts returns every post under a blog root, newest first.
func (s *Server) allPosts(blog string) []post {
	pattern := filepath.Join(s.localPath(blog), "[0-9][0-9][0-9][0-9]", "[0-9][0-9]", "[0-9][0-9]", "*")
	matches, _ := filepath.Glob(pattern)
	var posts []post
	for _, m := range matches {
		name := filepath.Base(m)
		if hiddenName(name) || forbiddenName(name) || !s.browsable(m) {
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
	return posts
}

var archivePath = regexp.MustCompile(`^[0-9]{4}/([0-9]{2}/([0-9]{2}/)?)?$`)

// isArchive reports whether reqPath is a year, month or day directory of the
// blog, e.g. /blog/2025/ or /blog/2025/07/.
func (s *Server) isArchive(blog, reqPath string) bool {
	return blog != "" && archivePath.MatchString(strings.TrimPrefix(reqPath, blog))
}

// blogArchive lists the posts of a year, month or day in chronological order,
// in the style of blog/_werc/traverse.sh:
//
//	In 2025...
//	- 2025/07/07: [Dear forecaster](/blog/2025/07/07/dear-forecaster/)
func (s *Server) blogArchive(blog, reqPath string) (string, string) {
	period := strings.TrimSuffix(strings.TrimPrefix(reqPath, blog), "/")
	title := "In " + period + "..."
	var md strings.Builder
	md.WriteString(title + "\n===\n\n")
	var posts []post
	for _, p := range s.allPosts(blog) {
		if strings.HasPrefix(p.date+"/", period+"/") {
			posts = append(posts, p)
		}
	}
	for i := len(posts) - 1; i >= 0; i-- { // oldest first
		p := posts[i]
		fmt.Fprintf(&md, "- %s: [%s](%s)\n", p.date, postTitle(p), p.url)
	}
	if len(posts) == 0 {
		md.WriteString("Nothing here.\n")
	}
	return renderMarkdown([]byte(md.String())), title
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
	indexCut = 12 // lines of markdown kept per post preview (title, body, read-more)
	postCut  = 10 // lines taken from each post
)

// blogIndex implements blagh_body: a feed link line, then for each of the
// newest posts (conf_max_posts_per_page of them) a preview made of the post's
// first lines with the title turned into a link, followed by a "read more"
// pointer, and finally a link to the sitemap. Each preview is rendered on its
// own, as werc's per-iteration pipeline did, with headings demoted (h1->h3,
// h2+ -> i).
func (s *Server) blogIndex(blog string) string {
	return s.gen.get("index:"+blog, func() any { return s.blogIndexUncached(blog) }).(string)
}

func (s *Server) blogIndexUncached(blog string) string {
	var b strings.Builder
	b.WriteString("<p><a href=\"index.rss\">RSS Feed</a>, <a href=\"../sitemap\">all content</a></p>\n")

	for _, p := range s.postList(blog) {
		lines := postLines(p, postCut)
		if len(lines) > 0 {
			lines[0] = fmt.Sprintf("[%s](%s) (%s)", lines[0], p.url, p.date)
		}
		lines = append(lines, "",
			fmt.Sprintf("<p style=\"text-align: left; margin-right: 10%%; font-size: 1.12em;\"><a style=\"text-decoration: none;\" href=\"%s\"><b>&rarr;</b> &#42604; ...</a></p>", p.url),
			"", "")
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
	}
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

// rss renders the RSS 2.0 feed (werc's rss20.tpl).
func (s *Server) rss(blog, reqPath string) []byte {
	return s.gen.get("rss:"+reqPath, func() any { return s.rssUncached(blog, reqPath) }).([]byte)
}

func (s *Server) rssUncached(blog, reqPath string) []byte {
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
	return b.Bytes()
}

// atom renders the Atom feed (werc's atom.tpl).
func (s *Server) atom(blog, reqPath string) []byte {
	return s.gen.get("atom:"+reqPath, func() any { return s.atomUncached(blog, reqPath) }).([]byte)
}

func (s *Server) atomUncached(blog, reqPath string) []byte {
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
	return b.Bytes()
}

// jsonFeed renders the JSON Feed (werc's jsonfeed.tpl).
func (s *Server) jsonFeed(blog, reqPath string) []byte {
	return s.gen.get("json:"+reqPath, func() any { return s.jsonFeedUncached(blog, reqPath) }).([]byte)
}

func (s *Server) jsonFeedUncached(blog, reqPath string) []byte {
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
	out, _ := json.MarshalIndent(feed, "", "  ")
	return append(out, '\n')
}
