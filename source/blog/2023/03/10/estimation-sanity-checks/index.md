Estimation for sanity checks
============================

I feel very warmly about using relatively quick estimates to carry out sanity checks, i.e., to quickly check whether something is clearly off, whether some decision is clearly overdetermined, or whether someone is just bullshitting. This is in contrast to Fermi estimates, which aim to arrive at an estimate for a quantity of interest, and which I also feel warmly about but which aren't the subject of this post. In this post, I explain why I like quantitative sanity checks so much, and I give some examples.

### Why I like this so much

I like this so much because:

- It is very defensible. There are some cached arguments against more quantified estimation, but sanity checking cuts through most—if not all—of them. "Oh, well, I just think that estimation has some really nice benefits in terms of sanity checking and catching bullshit, and in particular in terms of defending against scope insensitivity. And I think we are not even at the point where we are deploying enough estimation to catch all the mistakes that would be obvious in hindsight after we did some estimation" is both something I believe and also just a really nice motte to retreat when I am tired, don't feel like defending a more ambitious estimation agenda, or don't want to alienate someone socially by having an argument. 
- It can be very cheap, a few minutes, a few Google searches. This means that you can practice quickly and build intuitions.
- They are useful, as we will see below.

### Some examples

Here are a few examples where I’ve found estimation to be useful for sanity-checking. I mention these because I think that the theoretical answer becomes stronger when paired with a few examples which display that dynamic in real life. 

#### Photo Patch Foundation

The [Photo Patch Foundation](https://photopatch.org/) is an organization which has received a [small amount of funding](https://www.openphilanthropy.org/grants/photo-patch-foundation-general-support-2019/) from Open Philanthropy:

> Photo Patch has a website and an app that allows kids with incarcerated parents to send letters and pictures to their parents in prison for free. This diminishes barriers, helps families remain in touch, and reduces the number of children who have not communicated with their parents in weeks, months, or sometimes years.

It takes [little digging](https://donorbox.org/patching-relationships-with-letters-photos-2) to figure out that their costs are $2.5/photo. If we take the [AMF numbers at all seriously](https://forum.effectivealtruism.org/posts/4Qdjkf8PatGBsBExK/adding-quantified-uncertainty-to-givewell-s-cost), it seems very likely that this is not a good deal. For example, for $2.5 you can deworm several kids in developing countries, or buy [a bit more](https://www.againstmalaria.com/DollarsPerNet.aspx) than one malaria net. Or, less intuitively, trading 0.05% chance of saving a statistical life for sending a photo to a prisoner seems like a pretty bad trade--0.05% of a statistical life corresponds to 0.05/100 × 70 years × 365 = 12 statistical days.

One can then do [somewhat more elaborate estimations](https://forum.effectivealtruism.org/posts/h2N9qEbvQ6RHABcae/a-critical-review-of-open-philanthropy-s-bet-on-criminal) about criminal justice reform.

#### Sanity-checking that supply chain accountability has enough scale

At some point in the past, I looked into [supply chain accountability](https://forum.effectivealtruism.org/posts/ME4zE34KBSYnt6hGp/new-cause-proposal-international-supply-chain-accountability), a cause area related to improving how multinational corporations treat labor. One quick sanity check is, well, how many people does this affect? You can check, and per [here](https://static.inditex.com/annual_report_2021/es/documentos/informe-de-gestion-integrado-2021.pdf)[^1], Inditex—a retailer which owns brands like Zara, Pull&Bear, Massimo Dutti, etc.—employed 3M people in its supply chain, as of 2021.

So scalability is large enough that this may warrant further analysis. One this simple sanity check is passed, one can then go on and do some more complex estimation about how cost-effective improving supply chain accountability is, like [here](https://www.getguesstimate.com/models/14645). 

[^1]: Sum the number of people in p. 391-392 [here](https://static.inditex.com/annual_report_2021/es/documentos/informe-de-gestion-integrado-2021.pdf): 19546 + 49647 + 90363 + 383032 + 435469 + 845778 + 134970 + 92146 + 652808 + 381607 + 8499 + 4989 = 3098854. 

#### Sanity checking the cost-effectiveness of the EA Wiki

In my analysis of the EA Wiki, I calculated how much the person behind the EA Wiki [was being paid per word](https://forum.effectivealtruism.org/posts/kTLR23dFRB5pJryvZ/external-evaluation-of-the-ea-wiki#Costs_per_word_compared_to_other_industries), and found that it was in the ballpark of other industries. If it had been egregiously low, my analysis could have been shorter, and maybe concluded that this was a really good bargain. If the amount had been egregiously high, maybe I would have had to dig in about why that was. 

As it was, the sanity check was passed, and I went on to look at [other considerations](https://forum.effectivealtruism.org/posts/kTLR23dFRB5pJryvZ/external-evaluation-of-the-ea-wiki#Evaluating_outcomes).

#### Optimistic estimation for early causes 

Occasionally, I've seen some optimistic cost-effectiveness estimates by advocates of a particular cause area or approach (e.g., [here](https://forum.effectivealtruism.org/posts/CcNY4MrT5QstNh4r7/cost-effectiveness-of-foods-for-global-catastrophes-even), [here](https://forum.effectivealtruism.org/posts/HqEmL7XAuuD5Pc4eg/evaluating-strongminds-how-strong-is-the-evidence), or [here](https://forum.effectivealtruism.org/posts/XpeamS2yTNhagxAip/remote-health-centers-in-uganda-a-cost-effective)). One possible concern here is that because it's the advocates that are doing this cost-effective estimates, they might be biased upwards. But even if they are biased upwards, they are not completely uninformative: they show that at least some assumptions and parameters, chosen by someone who is trying their best, under which the proposed intervention looks great. And then further research might reveal that the initial optimism is or isn't warranted. But that first hurdle isn't trivial.

#### Other examples

- You can see the revival of LessWrong pretty clearly if you look at the [number of votes per year](https://nunosempere.com/blog/2021/06/24/shallow-evaluations-of-longtermist-organizations/). Evaluating the value of that revival is much harder, but one first sanity check is to see whether there was some revival being done.
- When evaluating small purchases, sometimes the cost of the item is much lower than the cost of thinking about it, or the cost of the time one would spend using the item (e.g., for me, the cost of a hot chocolate is smaller than the cost of sitting down to enjoy a hot chocolate). I usually take this as a strong sign that the price shouldn't be the main consideration for those types of purchase, and that I should remember that I am no longer a poor student.
- Some causes, like rare diseases, are not going to pass a cost-effectiveness sanity check, because they affect too few people.
- If you spend a lot of time in front of a computer, or having calls, the cost of better computer equipment and a better microphone is most likely worth it. I wish I'd internalized this sooner.
- Raffles and lotteries (e.g., "make three forecasts and enter a lottery to win $300", or "answer this survey to enter a raffle to win $500") are usually not worth it, because they don't reveal the number of people who enter, and it's usually very high.
- etc.

### Conclusion

I explained why I like estimates as sanity checks: they are useful, cheap, and very defensible. I then gave several examples of dead-simple sanity checks, and in each case pointed to more elaborate follow-up estimates. 

PS: You can subscribe to these posts [here](https://nunosempere.com/.subscribe/)

<p>
  <section id='isso-thread'>
  <noscript>Javascript needs to be activated to view comments.</noscript>
  </section>
</p>


