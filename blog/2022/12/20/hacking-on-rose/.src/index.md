Hacking on the Rose browser
===========================

I recently spent a few days hacking on the [rose browser](https://github.com/mini-rose/rose). It's a minimalist browser for Linux machines which is easy to play with. I enjoyed the experience, and thought to leave some notes.

Rose is written in C, and uses the Webkit and GTK libraries. Webkit, which drives Safari, started out as a fork of some previous open-source libraries, [KHTML](https://en.wikipedia.org/wiki/KHTML) and [KJS](https://en.wikipedia.org/wiki/KJS_%28software%29). [GTK](https://en.wikipedia.org/wiki/GTK) is a library for creating graphical interfaces. You can use the two together using [WebKitGTK](https://webkitgtk.org/).

![Image of this blogpost from the rose homepage](https://i.imgur.com/0sVCaN7.png)
*<br>Pictured: An earlier version of this blogpost in the rose browser.*

Personally, I find the simplicity of the code soothing. Still, to improve usability and for my own enjoyment, I've been adding some code to it. You can have a look at my tweaks [here](https://github.com/NunoSempere/rose-browser). These implement two main changes:

1. I've implemented a [reader mode](https://github.com/NunoSempere/rose-browser/tree/personal/plugins/readability)
2. I've added [redirection functionality](https://github.com/NunoSempere/rose-browser/tree/personal/plugins/libre_redirect), which redirects users to open source frontends of otherwise annoying websites.

The reader mode, for instance, returns this beautiful to read page:

![Pictured: Good newsletter, where you can read the content](https://i.imgur.com/ExEoWqq.png)

when fed the cluttered aesthetic abomination that is Matt Levine's newsletter:

![Pictured: Bad newsletter, with lots of advertisements](https://i.imgur.com/uRDsnXB.png)

I lifted the code for doing this from [Firefox](https://github.com/mozilla/readability), figured out how to integrate it with WebKit, and then made a few modifications.

As mentioned, I also wrote a version of [LibRedirect](https://libredirect.codeberg.page/) for the rose browser. It which direct annoying websites into better open-source versions:

![Pictured: list of sites that I am redirecting](https://i.imgur.com/NRo9wrF.png)

Lastly, I've added some code for installing the Rose browser on [Ubuntu 20.04](https://github.com/NunoSempere/rose-browser/raw/personal/rose), which is what I'm personally using. 

If you are a close friend using Linux feel free to reach out if you get stuck trying to install this.


<p><section id="isso-thread">
  <noscript>Javascript needs to be activated to view comments.</noscript>
</section></p>
