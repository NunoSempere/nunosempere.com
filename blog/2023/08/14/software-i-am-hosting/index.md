Webpages I am making available to my corner of the internet
===========================================================

Here is a list of internet services that I make freely available to friends and allies, broadly defined—if you are reading this, you qualify. These are ordered roughly in order of usefulness. 

### search.nunosempere.com

[search.nunosempere.com](https://search.nunosempere.com/) is an instance of [Whoogle](https://github.com/benbusby/whoogle-search). It presents Google results as they were and as they should have been: without clutter and without advertisements.

Readers are welcome to make this their default search engine. The process to do this is a bit involved and depends on the browser, but can be found with a Whoogle search. In past years, I've had technical difficulties around once every six months, but tend to fix them quickly.

### forum.nunosempere.com

[forum.nunosempere.com](https://forum.nunosempere.com) is a frontend to the [Effective Altruism Forum](https://forum.effectivealtruism.org/) that I personally find soothing. It is *much* faster than the official frontend, more minimalistic, and offers an RSS endpoint for all posts [here](https://forum.nunosempere.com/feed).

```
$ time curl https://forum.effectivealtruism.org > /dev/null
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  847k    0  847k    0     0   439k      0 --:--:--  0:00:01 --:--:--  438k

real    0m1.945s
user    0m0.030s
sys     0m0.021s

$ time curl https://forum.nunosempere.com/frontpage > /dev/null
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 35091  100 35091    0     0   190k      0 --:--:-- --:--:-- --:--:--  190k

real    0m0.195s
user    0m0.025s
sys     0m0.004s
```

If you use the EA Forum with some frequency, I'd recommend giving it and [ea.greaterwrong.com](https://ea.greaterwrong.com/) a spin.

### shapleyvalue.com

[shapleyvalue.com](http://shapleyvalue.com/) is an online calculator for [Shapley Values](https://wikiless.northboot.xyz/wiki/Shapley_value?lang=en). I wrote it for [this explainer](https://forum.effectivealtruism.org/s/XbCaYR3QfDaeuJ4By/p/XHZJ9i7QBtAJZ6byW) after realizing that no other quick calculators exist.

### Find a beta distribution which fits a given confidence interval

[trastos.nunosempere.com/fit-beta](https://trastos.nunosempere.com/fit-beta) is a POST endpoint to find a beta distribution that fits a given confidence interval. 

```
curl -X POST -H "Content-Type: application/json" \
  -d '{"ci_lower": "0.2", "ci_upper":"0.8", "ci_length": "0.95"}' \
  https://trastos.nunosempere.com/fit-beta
```

I also provide a widget [here](https://nunosempere.com/blog/2023/03/15/fit-beta/) and an npm package [here](https://www.npmjs.com/package/fit-beta), which is probably more convenient than the endpoint.

### nunosempere.com/misc/proportional-approval-voting-calculator/

Proportional approval voting is a bit tricky to generalize to choosing candidates for more than one position, which is why little software for it exists. [This page](https://nunosempere.com/misc/proportional-approval-voting-calculator/) provides a samaple implementation. It was previously hosted [here](https://nunosempere.com/misc/proportional-approval-voting-calculator/).

### git.nunosempere.com

[git.nunosempere.com](https://git.nunosempere.com/) is my personal git server. It hosts some of my personal projects, and occasional backups of some open source projects worth preserving.

### video.nunosempere.com

[video.nunosempere.com](https://video.nunosempere.com) is a [peertube](https://github.com/Chocobozzz/PeerTube/) instance with some videos worth preserving. 

### royalroad.nunosempere.com

A frontend for [Royal Road](https://www.royalroad.com/), a site which hosts online fiction but which has grown pretty cluttered. Reuses a whole lot of the code from forum.nunosempere.com.

### gatitos.nunosempere.com

Shows a photo of two cute cats:

<img src="https://gatitos.nunosempere.com/">

### Also on this topic

- [Soothing Software](https://nunosempere.com/blog/2023/03/27/soothing-software/)
- [Hacking on rose](https://nunosempere.com/blog/2022/12/20/hacking-on-rose/)—in particular, readers might be interested in [this code](https://git.nunosempere.com/open.source/rosenrot/src/branch/master/plugins/style/style.js#L62) to block advertisements on Reddit and Twitter. It could be adapted for Firefox with an extension like [Stylus](https://addons.mozilla.org/en-US/firefox/addon/styl-us/).
- [Metaforecast](https://metaforecast.org/), which I started, and which is now maintained by Slava Matyuhin of QURI and myself.

<p>
  <section id='isso-thread'>
  <noscript>javascript needs to be activated to view comments.</noscript>
  </section>
</p>
