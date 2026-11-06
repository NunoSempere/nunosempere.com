Hacking on rose
===============

The [rose browser](https://github.com/mini-rose/rose) is a minimal browser for Linux machines. I've immensely enjoyed hacking on it this last week, so I thought I'd leave some notes.

Rose is written in C, and it's based on Webkit and GTK. Webkit is the engine that drives Safari, and a fork of some previous open-source libraries, [KHTML](https://en.wikipedia.org/wiki/KHTML) and [KJS](https://en.wikipedia.org/wiki/KJS_%28software%29). [GTK](https://en.wikipedia.org/wiki/GTK) is a library for creating graphical interfaces. You can conveniently use the two together using [WebKitGTK](https://webkitgtk.org/).

![Image of this blogpost from the rose homepage](https://images.nunosempere.com/blog/2022/12/20/hacking-on-rose/rose-blog-post.png)
*<br>Pictured: An earlier version of this blogpost in the rose browser.*

Personally, I find the simplicity of the code soothing. At the same time, though, I've been adding some code to it, to make the browser more usable and for the fun of it. You can see my tweaks [here](https://github.com/NunoSempere/rose-browser).

My first tweak was to add a [reader mode](https://github.com/NunoSempere/rose-browser/tree/personal/plugins/readability). It, for instance, returns this beautiful page:

![Pictured: Good newsletter, where you can read the content](https://images.nunosempere.com/blog/2022/12/20/hacking-on-rose/matt-levine-clean.png)

when fed the aesthetic abomination that is Matt Levine's newsletter:

![Pictured: Bad newsletter, with lots of advertisements](https://images.nunosempere.com/blog/2022/12/20/hacking-on-rose/matt-levine-og.png)

I lifted the code for doing this from [Firefox](https://github.com/mozilla/readability), figured out how to integrate it with WebKit, and then improved it a bit.

My second big tweak was to add [redirection functionality](https://github.com/NunoSempere/rose-browser/tree/personal/plugins/libre_redirect) from annoying websites to their open source frontends, in the style of [LibRedirect](https://libredirect.codeberg.page/):

![Pictured: list of sites that I am redirecting](https://images.nunosempere.com/blog/2022/12/20/hacking-on-rose/libre_redirect.png)

And I've added some code for installing this browser on [Ubuntu 20.04](https://github.com/NunoSempere/rose-browser/raw/personal/rose), which is what I'm personally using. If you are a dear friend using Linux feel free to ask for help if you get stuck trying to install this.

<p><section id="isso-thread">
  <noscript>Javascript needs to be activated to view comments.</noscript>
</section></p>

