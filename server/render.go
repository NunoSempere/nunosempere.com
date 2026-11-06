package main

import (
	"bytes"
	"html"
	"html/template"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	ghtml "github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

// pageData feeds templates/page.html.
type pageData struct {
	PageTitle    string
	ReqPath      string
	SiteTitle    string
	SiteSubTitle string
	ExtraHeaders template.HTML
	Nav          template.HTML
	TopNav       template.HTML
	Body         template.HTML
	Footer       template.HTML
}

// renderPage wraps a body in the page template: header, sidebar, footer.
func (s *Server) renderPage(pg *pageData) ([]byte, error) {
	cfg := s.config()
	pg.SiteTitle = cfg.Title
	pg.SiteSubTitle = cfg.SubTitle
	if pg.PageTitle == "" {
		pg.PageTitle = strings.TrimSpace(cfg.Title + " " + cfg.SubTitle)
	}
	pg.Nav = template.HTML(s.navTree(pg.ReqPath))
	pg.TopNav = template.HTML(s.topNav(pg.ReqPath))
	pg.Footer = template.HTML(s.footer())

	var buf bytes.Buffer
	if err := s.page.Execute(&buf, pg); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// --- markdown ---------------------------------------------------------------

var md goldmark.Markdown

// initMarkdown configures goldmark to behave like the discount markdown the
// site used (-f fencedcode -f footnote -f latex, with smartypants, tables,
// strikethrough and superscript on by default).
func (s *Server) initMarkdown() {
	md = goldmark.New(
		goldmark.WithExtensions(
			extension.Table,
			extension.Strikethrough,
			extension.Footnote,
			extension.Typographer,
			&discountExt{},
		),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		goldmark.WithRendererOptions(ghtml.WithUnsafe()),
	)
}

func renderMarkdown(src []byte) string {
	var buf bytes.Buffer
	if err := md.Convert(src, &buf); err != nil {
		return "<pre>" + html.EscapeString(string(src)) + "</pre>"
	}
	return buf.String()
}

// renderCache memoises rendered markdown files by path, mtime and size.
type renderCache struct {
	mu sync.Mutex
	m  map[string]cacheEntry
}

type cacheEntry struct {
	modTime time.Time
	size    int64
	html    string
	title   string
}

func newRenderCache() *renderCache { return &renderCache{m: map[string]cacheEntry{}} }

func (s *Server) renderMarkdownFile(abs string) (string, string) {
	st, err := os.Stat(abs)
	if err != nil {
		return "", ""
	}
	c := s.cache
	c.mu.Lock()
	e, ok := c.m[abs]
	c.mu.Unlock()
	if ok && e.modTime.Equal(st.ModTime()) && e.size == st.Size() {
		return e.html, e.title
	}
	src, err := os.ReadFile(abs)
	if err != nil {
		return "", ""
	}
	e = cacheEntry{modTime: st.ModTime(), size: st.Size(), html: renderMarkdown(src), title: mdTitle(src)}
	c.mu.Lock()
	c.m[abs] = e
	c.mu.Unlock()
	return e.html, e.title
}

var (
	bodyOpen  = regexp.MustCompile(`(?is)<body[^>]*>`)
	bodyClose = regexp.MustCompile(`(?is)</\s*body[^>]*>`)
)

// renderHTMLFile implements werc's html_handler: the contents of <body> when
// present, otherwise the whole file.
func (s *Server) renderHTMLFile(abs string) (string, string) {
	src, err := os.ReadFile(abs)
	if err != nil {
		return "", ""
	}
	body := src
	if m := bodyOpen.FindIndex(src); m != nil {
		body = src[m[1]:]
		if e := bodyClose.FindIndex(body); e != nil {
			body = body[:e[0]]
		}
	}
	return string(body), htmlTitle(src)
}

func (s *Server) renderTxtFile(abs string) string {
	src, err := os.ReadFile(abs)
	if err != nil {
		return ""
	}
	return "<pre>\n" + html.EscapeString(string(src)) + "</pre>\n"
}

// --- discount compatibility extension ------------------------------------

// discountExt adds two inline constructs discount supported: superscripts
// ("x^2", "10^(-900)") and verbatim "$$ ... $$" LaTeX passthrough.
type discountExt struct{}

func (e *discountExt) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(&supParser{}, 500),
		util.Prioritized(&mathParser{}, 500),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(&supRenderer{}, 500),
	))
}

var kindSup = ast.NewNodeKind("Superscript")

type supNode struct{ ast.BaseInline }

func (n *supNode) Kind() ast.NodeKind { return kindSup }
func (n *supNode) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

type supParser struct{}

func (p *supParser) Trigger() []byte { return []byte{'^'} }

func isAlnum(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func (p *supParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	line, seg := block.PeekLine()
	if len(line) < 2 {
		return nil
	}
	var start, end, consumed int
	if line[1] == '(' {
		close := bytes.IndexByte(line[2:], ')')
		if close < 0 {
			return nil
		}
		start, end, consumed = 2, 2+close, 2+close+1
	} else {
		prev := block.PrecendingCharacter()
		if !(prev < 128 && isAlnum(byte(prev))) {
			return nil
		}
		i := 1
		for i < len(line) && isAlnum(line[i]) {
			i++
		}
		if i == 1 {
			return nil
		}
		start, end, consumed = 1, i, i
	}
	n := &supNode{}
	n.AppendChild(n, ast.NewTextSegment(text.NewSegment(seg.Start+start, seg.Start+end)))
	block.Advance(consumed)
	return n
}

type supRenderer struct{}

func (r *supRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(kindSup, func(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			_, _ = w.WriteString("<sup>")
		} else {
			_, _ = w.WriteString("</sup>")
		}
		return ast.WalkContinue, nil
	})
}

type mathParser struct{}

func (p *mathParser) Trigger() []byte { return []byte{'$'} }

func (p *mathParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	line, _ := block.PeekLine()
	if !bytes.HasPrefix(line, []byte("$$")) {
		return nil
	}
	close := bytes.Index(line[2:], []byte("$$"))
	if close < 0 {
		return nil
	}
	n := ast.NewString(append([]byte(nil), line[:close+4]...))
	n.SetRaw(true)
	block.Advance(close + 4)
	return n
}
