package main

import "testing"

func TestSanitizePath(t *testing.T) {
	cases := map[string]string{
		"/":                       "/",
		"//blog///2025/":          "/blog/2025/",
		"/a/../b":                 "/a/./b",
		"/blog/post%20name":       "/blog/post20name",
		"/ok_+-.,:chars":          "/ok_+-.,:chars",
		"/niño/":                  "/nio/",
		"/with space/and<script>": "/withspace/andscript",
	}
	for in, want := range cases {
		if got := sanitizePath(in); got != want {
			t.Errorf("sanitizePath(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestTitleFromPath(t *testing.T) {
	cases := map[string]string{
		"/":                                   "",
		"/blog/":                              "blog",
		"/blog/2025/07/07/dear-forecaster/":   "blog - 2025 - 07 - 07 - dear forecaster",
		"/misc/proportional-approval-voting/": "misc - proportional approval voting",
		"/some_page":                          "some page",
	}
	for in, want := range cases {
		if got := titleFromPath(in); got != want {
			t.Errorf("titleFromPath(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestMdTitle(t *testing.T) {
	cases := map[string]string{
		"Title\n===\n\nBody":   "Title",
		"Title\n=====\n":       "Title",
		"Title\n===\nBody":     "",
		"# Title\n\nBody":      "",
		"Title\n=\n\nBody":     "",
		"Title\n---\n\nBody":   "",
		"Dear forecaster\n===": "Dear forecaster",
	}
	for in, want := range cases {
		if got := mdTitle([]byte(in)); got != want {
			t.Errorf("mdTitle(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestHtmlTitle(t *testing.T) {
	if got := htmlTitle([]byte("<html><head>\n<title> Hello </title>\n")); got != "Hello" {
		t.Errorf("got %q", got)
	}
	if got := htmlTitle([]byte("\n\n    <script src=x></script>\n<!--\n\tSources:\n-->")); got != "" {
		t.Errorf("expected no title for indentation-only first match, got %q", got)
	}
	if got := htmlTitle([]byte("<form><div><h3>Sign up</h3>")); got != "Sign up" {
		t.Errorf("got %q", got)
	}
}

func TestMarkdownCompat(t *testing.T) {
	(&Server{}).initMarkdown()
	cases := map[string]string{
		"tl;dr^2 and 10^(-900)":             "tl;dr<sup>2</sup> and 10<sup>-900</sup>",
		"^ not a sup":                       "^ not a sup",
		"$$ P(A \\& B) = \\frac{1}{2} $$":   "$$ P(A \\&amp; B) = \\frac{1}{2} $$",
		"costs $150 and \"quotes\" -- dash": "costs $150 and &ldquo;quotes&rdquo; &ndash; dash",
		"~~gone~~":                          "<del>gone</del>",
	}
	for in, want := range cases {
		got := renderMarkdown([]byte(in + "\n"))
		if got != "<p>"+want+"</p>\n" {
			t.Errorf("renderMarkdown(%q) = %q, want <p>%s</p>", in, got, want)
		}
	}
}
