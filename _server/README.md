# nunosempere.com web server

A small Go program that serves this directory as a website. It replaces the
werc + plan9port + fcgiwrap stack the site used to run on, while producing the
same pages: Markdown rendered inside the same header/sidebar/article layout,
the same stylesheet, the blog front page, RSS/Atom/JSON feeds, the sitemap,
directory listings for directories without an index, and static files.

Nothing about the content layout changes. Pages are still `index.md` files in
directories, blog posts still live at `blog/YYYY/MM/DD/<slug>/index.md`, the
site title is still read from `_werc/config` (so `_werc/makeconfig.sh` keeps
working), and `blog/_werc/config` still marks `blog/` as the blog.

## Build and run

    cd _server
    go build -o webserver .
    ./webserver -root .. -addr 127.0.0.1:8080

Flags:

| flag         | default                   | meaning                                            |
|--------------|---------------------------|----------------------------------------------------|
| `-root`      | `.` or `..`               | site directory (auto-detected via `_werc/`)        |
| `-addr`      | `127.0.0.1:8080`          | listen address                                     |
| `-base-url`  | `https://nunosempere.com` | absolute URL used in feeds, sitemap and 404 pages  |
| `-pub`       | embedded                  | directory served under `/pub/` instead of `pub/`   |
| `-max-posts` | `32`                      | posts in the feeds                                 |

The stylesheet (`pub/style/style.css`) and the page template
(`templates/page.html`) are compiled into the binary; edit them and rebuild.

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
- `_server/` and `.git/` are never served.

Markdown is rendered with goldmark configured to match the discount flags the
site used (`fencedcode`, `footnote`, `latex`): tables, strikethrough, footnotes,
smart quotes/dashes, `x^2` superscripts and verbatim `$$ … $$` blocks.

## Deploying behind nginx (Alpine)

See `contrib/nginx-site.conf` for the server block (plain `proxy_pass`, no
FastCGI) and `contrib/webserver.initd` for an OpenRC service. Roughly:

    go build -o /home/www/bin/webserver ./_server      # or copy the binary over
    cp _server/contrib/webserver.initd /etc/init.d/webserver && chmod 755 /etc/init.d/webserver
    rc-update add webserver default && rc-service webserver start
    cp _server/contrib/nginx-site.conf /etc/nginx/sites-available/nunosempere.com
    nginx -t && rc-service nginx reload

The binary is static and needs nothing else installed: no plan9port, no rc,
no fcgiwrap, no discount.
