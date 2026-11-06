# nunosempere.com

- `source/` — the site: Markdown pages in directories, blog posts under
  `source/blog/YYYY/MM/DD/<slug>/index.md`, images and other files next to them.
- `server/` — a small Go program that renders `source/` either as a static site
  (`make`, output in `build/`) or as a live development server (`make serve`).
- `build/` — generated output, not committed. Deploy this directory to any
  static host (`make deploy` for Cloudflare Pages).

See `server/README.md` for how pages are rendered and the available flags.
