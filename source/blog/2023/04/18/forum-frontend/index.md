A Soothing Frontend for the Effective Altruism Forum 
====================================================

### About 

<a href="https://forum.nunosempere.com">forum.nunosempere.com</a> is a frontend for the [Effective Altruism Forum](https://forum.effectivealtruism.org). It aims to present EA Forum posts in a way which I personally find soothing. It achieves that that goal at the cost of pretty restricted functionality—like not having a frontpage, or not being able to make or upvote comments and posts.

### Usage

Instead of having a frontpage, this frontend merely has an endpoint:

``` 
https://forum.nunosempere.com/posts/[post id]/[optional post slug]
```

An [example post](https://forum.nunosempere.com/posts/fz6NdW4i25FZoRQyz/use-of-i-d-bet-on-the-ea-forum-is-mostly-metaphorical)---chosen to be a link post so that you can see both the article and the comments---looks like this: 

![](https://images.nunosempere.com/blog/2023/04/18/forum-frontend/forum-post.png)

### Notes 

- Clicking on an article's or on a comment's timestamp directs to their location on the original EA Forum. 
- I am choosing not to have a frontpage because it would just add additional development time to a hobby project, and because it doesn't support my use case. My usecase is just being subscribed via rss, and displaying forum posts which interest me in this new frontend. 
- On the topic of rss, this frontend has an rss endpoint at [forum.nunosempere.com/feed](https://forum.nunosempere.com/feed): it piggybacks off the [ea.greaterwrong.com RSS feed](https://ea.greaterwrong.com/?format=rss). 
- Other alternatives to the EA Forum known to me are the aforementioned [ea.greaterwrong.com](https://ea.greaterwrong.com/) as well as [eaforum.issarice.com](https://eaforum.issarice.com/). 

<p>
  <section id='isso-thread'>
  <noscript>Javascript needs to be activated to view comments.</noscript>
  </section>
</p>
