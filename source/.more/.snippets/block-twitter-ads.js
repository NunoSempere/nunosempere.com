if (document.domain == "twitter.com" ){
  styles = `
  /* hide promoted tweets */
  :has(meta[property="og:site_name"][content="Twitter"])
    [data-testid="cellInnerDiv"]:has(svg + [dir="auto"]) {
    display: none;
  }
  [data-testid^="placementTracking"] {
    display: none;
  }

  /* hide what's happening section */
  :has(meta[property="og:site_name"][content="Twitter"])
    [aria-label="Timeline: Trending now"] {
    display: none !important;
  }
  [data-testid^="sidebarColumn"] {
    display: none;
  }

  `
}
