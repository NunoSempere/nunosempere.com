Hacking on the Rose browser
===========================

I recently had the pleasure of tinkering with the [rose browser](https://github.com/mini-rose/rose) for a few days. It's a Linux-based minimalist web browser, making it easy and fun to customize. To help others who may explore this project, I decided to leave some notes about my experience.

Rose is written in C and takes advantage of the WebKit and GTK libraries. WebKit, which drives Safari, is a fork of the open-source KHTML and KJS libraries. GTK, on the other hand, is a library for enabling graphical user interfaces. You can utilize both libraries simultaneously by using WebKitGTK.

Personally, I have found solace in the simplicity of this code. Nevertheless, for usability purposes and for my own enjoyment, I have been adding some code to it. You can have a look at my edits [here](https://github.com/NunoSempere/rose-browser).

I've made two main changes to the [Rose Browser](https://github.com/NunoSempere/rose-browser): 

1. I've implemented a [reader mode](https://github.com/NunoSempere/rose-browser/tree/personal/plugins/readability) 
2. I've added [redirection functionality](https://github.com/NunoSempere/rose-browser/tree/personal/plugins/libre_redirect), which redirects users to open source frontends of otherwise annoying websites.

The reader mode allows readers to access a page in a more visually-appealing way; it returns a page that is easy and enjoyable to read.


Check out this [picture](https://i.imgur.com/ExEoWqq.png) of a great newsletter, where you can read its content.

When presented with the unpleasant experience of reading Matt Levine's newsletter:



![An image of an overcrowded newsletter full of advertisements](https://i.imgur.com/uRDsnXB.png)

I obtained the code for this from [Firefox](https://github.com/mozilla/readability). Subsequently, I implemented it within WebKit and made a few modifications.

I offer an alternative version of [LibRedirect](https://libredirect.codeberg.page/), which redirects annoying webpages to better, open-source alternatives.

![List of sites to which I have redirected traffic](https://i.imgur.com/NRo9wrF.png)

I have added instructions for installing the Rose browser on [Ubuntu 20.04](https://github.com/NunoSempere/rose-browser/raw/personal/rose), which is the version I am currently using.

If you are a dear friend who is using Linux, please reach out for assistance if you run into any difficulties while attempting to install this.
