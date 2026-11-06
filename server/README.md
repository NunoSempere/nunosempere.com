# nunosempere.com site generator and server

A small Go program that turns `../source/` into a website. It replaces the
werc + plan9port + fcgiwrap stack the site used to run on, while producing the
same pages: Markdown rendered inside the same header/sidebar/article layout,
the same stylesheet, the blog front page and year archives, RSS/Atom/JSON
feeds, the sitemap, directory listings for directories without an index, and
static files.

It has two modes:

- `-build DIR` renders everything into `DIR` as static files and exits. This
  is what `make` at the repository root does, producing `../build/`, which can
  be deployed to any static host.
- `-addr HOST:PORT` serves the same pages live. Useful as a development server
  (`make serve`), or as a production server behind nginx if you prefer.

Nothing about the content layout changes. Pages are `index.md` files in
directories, blog posts live at `source/blog/YYYY/MM/DD/<slug>/index.md`, the
site title is read from `source/_werc/config` if present, and
`source/blog/_werc/config` marks `blog/` as the blog.

## Build and run

    make            # build the binary and generate ../build/
    make serve      # development server on http://127.0.0.1:8080
    make deploy     # wrangler pages deploy ../build/
    make test

Or directly:

    cd server && go build -o webserver .
    ./webserver -root ../source -build ../build
    ./webserver -root ../source -addr 127.0.0.1:8080

Flags:

| flag         | default                   | meaning                                            |
|--------------|---------------------------|----------------------------------------------------|
| `-root`      | `./source`, `../source`, `.` or `..` | site directory                          |
| `-build`     |                           | write the static site into this directory and exit |
| `-addr`      | `127.0.0.1:8080`          | listen address for the live server                 |
| `-base-url`  | `https://nunosempere.com` | absolute URL used in feeds, sitemaps and 404 pages |
| `-pub`       | embedded                  | directory served under `/pub/` instead of `pub/`   |
| `-max-posts` | `32`                      | posts in the feeds                                 |
| `-cache-ttl` | `60s`                     | how long generated pages are cached (live server)  |
| `-follow-symlinks` | off                 | serve files behind symlinks that point outside the site directory |

The stylesheet (`pub/style/style.css`) and the page template
(`templates/page.html`) are compiled into the binary; edit them and rebuild.

## What the static build contains

- `index.html` for every directory and every `foo.md`/`foo.html`/`foo.txt`
  page (at `foo/index.html`), rendered exactly as the live server would.
- Every other file copied as is, including `.md` sources, so raw URLs keep
  working. `.html` sources are not copied, since they were never served raw.
- `blog/index.rss`, `blog/index.atom`, `blog/feed.json`, `sitemap/index.html`,
  `sitemap.xml`, `sitemap.txt`, `sitemap.gz`, `robots.txt`, `404.html`,
  `_headers` (Cloudflare Pages / Netlify header rules) and `pub/`.
- A `.generated-by-webserver` marker. The build refuses to empty a non-empty
  directory that lacks it.

Files over 25 MB are copied but reported, because Cloudflare Pages rejects
them.

### Cloudflare

`../wrangler.jsonc` declares `build/` as the static assets of a Worker with no
code, with trailing-slash handling and `404.html` like the live server.

- Git integration: in the Cloudflare dashboard, Workers & Pages, Create,
  import the repository. Build command
  `cd server && go build -o webserver . && ./webserver -root ../source -build ../build`,
  deploy command `npx wrangler deploy`, environment variable `GO_VERSION=1.26`.
  Every push then redeploys. Only committed files are deployed, so gitignored
  directories such as `.secret/` are not published this way.
- From your machine: `npm install -g wrangler`, `wrangler login` once, then
  `make deploy`. This uploads whatever is in the working tree.

Cloudflare rejects files over 25 MB and deployments over 20,000 files.

## Behaviour carried over from werc

- `/foo/` serves `foo/index.md`; `/foo` redirects to `/foo/` when `foo` is a
  directory, and `/foo.html`, `/foo/index`, `/foo.` redirect to the canonical URL.
- Names starting with `.` or `_` are served but never listed in the sidebar,
  directory listings or sitemap (so `.secret/…` keeps working).
- `/sitemap` renders the site map page. `/sitemap.xml`, `/sitemap.txt`,
  `/sitemap.gz` and a default `/robots.txt` are generated on demand from the
  same tree walk; nothing is written to disk. A real file of the same name in
  the site root takes precedence.
- Raw files (`.md`, scripts, images, pdfs) are served directly, text-like
  files as `text/plain` like before.
- `.git/` is never served.

Markdown is rendered with goldmark configured to match the discount flags the
site used (`fencedcode`, `footnote`, `latex`): tables, strikethrough, footnotes,
smart quotes/dashes, `x^2` superscripts and verbatim `$$ … $$` blocks.

## Security notes

- Request paths are reduced to `[a-zA-Z0-9_+-/.,:]`, runs of dots are
  collapsed and the result is cleaned, so `..` can never leave the site
  directory. Only `GET` and `HEAD` are accepted.
- Files and directories reached through a symlink are served, listed or
  walked only if they resolve to a path inside the site directory, unless
  `-follow-symlinks` is given.
- `.git/` is never served. Everything else in the directory
  is reachable by URL, including dot-directories; that is the intended
  "unlisted but public" behaviour inherited from werc. Keep private material
  out of the tree.
- Generated pages are cached for `-cache-ttl`, so hammering `/sitemap.xml` or
  the feeds does not turn into repeated tree walks.
- Responses carry `X-Content-Type-Options: nosniff`. Log lines quote the
  request, referer and user agent so they cannot forge entries.
- Bind to loopback and let nginx terminate TLS; the server speaks plain
  HTTP/1.1 only. Rebuild with a current Go toolchain after Go security
  releases (`go get toolchain@go1.N.M`), since `net/http` is the whole attack
  surface.

## Deploying behind nginx (Alpine)

For a self-hosted static deployment, build on the server (or `rsync` the
`build/` directory to it) and use `contrib/nginx-site.conf`, which serves
`build/` with `try_files` and a `404.html` error page. `contrib/webserver.initd`
is an OpenRC service for the live-server mode, should you prefer that; pair it
with a `proxy_pass http://127.0.0.1:8080` location instead of the static root.
