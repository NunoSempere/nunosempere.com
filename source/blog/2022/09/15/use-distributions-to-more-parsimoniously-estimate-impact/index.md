Use distributions to more parsimoniously estimate impact
========================================================

## Summary

By incorporating uncertainty into its estimates, GiveWell would produce better estimates. This is best done by working with distributions, as opposed to point estimates. For example, "$294 per doubling of consumption" is a point estimate[^1], but the following is a distribution:

![](https://i.imgur.com/jEP0tE9.png)

This post is an entry to GiveWell's [Change Our Mind Contest](https://www.givewell.org/research/change-our-mind-contest).

## Index

- Summary
- Prior work
- Advantages
  - Bottom line may change
  - Value of information calculations become easier
  - it becomes possible to (attempt to adjust) for the optimizers' curse
  - It is the parsimonious way to do this, and thus I expect a better treatment of uncertainty to have further accuracy gains
- How to do it
  - Which language to use
  - To do it in-house or to delegate it
  - To do it all at once or bit by bit
- Conclusion


## Prior work

People have been pointing out this problem for a while:

- [Uncertainty and sensitivity analyses of GiveWell's cost-effectiveness analyses](https://forum.effectivealtruism.org/posts/gMxTEMvh8RttX9Nt4/uncertainty-and-sensitivity-analyses-of-givewell-s-cost)
- [Quantifying Uncertainty in GiveWell's GiveDirectly Cost-Effectiveness Analysis](https://forum.effectivealtruism.org/posts/4Qdjkf8PatGBsBExK/adding-quantified-uncertainty-to-givewell-s-cost)
- [Adding Quantified Uncertainty to GiveWell's Cost Effectiveness Analysis of the Against Malaria Foundation](https://forum.effectivealtruism.org/posts/4Qdjkf8PatGBsBExK/adding-quantified-uncertainty-to-givewell-s-cost)
- [A critical review of GiveWell's 2022 cost-effectiveness model](https://forum.effectivealtruism.org/posts/6dtwkwBrHBGtc3xes/a-critical-review-of-givewell-s-2022-cost-effectiveness)

This is because adding distributional estimates is relatively low-hanging fruit that GiveWell may have been procrastinating on.

## Advantages

### Bottom line may change

The key reason to use distributional estimates is that the bottom line may change.

One key example I keep pointing to is the paper [Dissolving the Fermi Paradox](https://arxiv.org/abs/1806.02404). They provide a clear example where using point estimates leads to a very different outcome than using distributions. When using point estimates, one arrives at the [Fermi paradox](https://en.wikipedia.org/wiki/Fermi_paradox), whereas when using distributions one arrives at the reasonable conclusion that we have a high chance of being alone in the galaxy.

Similarly, GiveWell's estimates may be different if they use distributional estimates, particularly for distributions for which there is high uncertainty. This might be the case for individual interventions such as deworming.

### Value of information calculations become easier

With distributions, value of information calculations become more tractable. In particular, the probability that a given intervention will beat GiveWell's top charities upon further investigation becomes a calculable amount. This amount would depend on the current distribution of impact of the new opportunity, the expected narrowing with more investigation, the distribution of impact of current GW top charities, and the expected yearly amount of funding.

This point might might also affect meta-global-health and development: better treatment of uncertainty might suggest that incubators, meta-charities, etc., such as [IPA](https://www.poverty-action.org/) or [Charity Entrepreneurship](https://www.charityentrepreneurship.com/) are more or less neglected than currently considered. It would also ground the need for something like a [GiveWell incubation grant](https://www.givewell.org/research/incubation-grants) in terms of a new charity having a high enough chance of beating GiveWell charities, which could be calculated from its current expected impact distribution.

### It becomes possible to (attempt to adjust) for the optimizers' curse

The [Optimizers' curse](https://christiansmith.cc/2019/04/03/the-optimizers-curse-wrong-way-reductions/) seems like the biggest theoretical objection to GiveWell's current recommendations. It points out that opportunities estimated to be the best will be selected for having high impact, but also for having a high error rate in the estimation process. It's possible that this may have been the case for deworming.

The Optimizers' curse is not unfixable, and in particular there is a [Bayesian solution](https://www.lesswrong.com/posts/5gQLrJr2yhPzMCcni/the-optimizer-s-curse-and-how-to-beat-it). But for this solution to work, one has to construct a prior[^2], and the estimates have to be in distributional form.

Note that GiveWell may have previously been bit by the optimizers' curse: in [their answer to HLI on deworming, they write](https://forum.effectivealtruism.org/posts/MKiqGvijAXfcBHCYJ/deworming-and-decay-replicating-givewell-s-cost?commentId=Qt26uR9ZT6ru8xDqi#comments), they write:

> Once we do this work, our best guess is that we will reduce our estimate of the cost-effectiveness of deworming by 10%-30%. Had we made this change in 2019 when KLPS-4 was released, we would have recommended $2-$8m less in grants to deworming (out of $55m total) since 2019.

Correcting for this by explicitly accounting for the optimizers' curse may have made this error less severe—though GiveWell probably did adjust, if perhaps on a more ad-hoc way. See some discussion of a related point [here](https://forum.effectivealtruism.org/posts/MKiqGvijAXfcBHCYJ/deworming-and-decay-replicating-givewell-s-cost?commentId=mTWzsjgDhGCyAu24z).

### It is the parsimonious way to do this, and thus I expect a better treatment of uncertainty to have further accuracy gains

Ultimately, the "true shape" of impact estimates—and of our knowledge about the world more generally—seems like it is uncertain, probabilistic, and distributional. As such, I expect there would be additional advantages besides the ones I've pointed above, like:

- More easily identifying the parts of a calculation which are more crucial to its result
- Allowing donors to make decisions which better fit their own risk profiles
- Making estimates, such as the ones for AMF which depend on the age of the beneficiary, way less clunky[^3]
- Comparisons against other cause areas (e.g., animal welfare) may become a bit more accurate

## How to do it

### Which tools to use

Because of GiveWell's model of doing its analysis and then publishing it, I think that mostly any programming language with support for manipulating probabilities (R, Python, Turing.jl, Stata, etc.) would be ok. In particular, the key factor seems like ease of use for current GiveWell staff-members.

I'd put in a good word for [Squiggle](https://www.squiggle-language.com/), the language which my org is developing. It may be particularly suitable for making calculators or tools to allow users to later tweak estimates in the browser.

But I think that any choice of tool that supports distribution would do a-ok.

### To do it in-house or to delegate it

I think that doing this kind of thing well would require having plenty of access to GiveWell's internal knowledge and access to GiveWell staff's intuitions. And an outside effort would have to be so tightly integrated with the GiveWell team that you might as well do it in-house.

### To do it all at once or bit by bit

It is possible that translating all models to distributions might be too much work for one year—I can't say since I'm not familiar with GiveWell's internal procedures. Because of the uncertainty in how long this would take, it seems to me that creating a distributional estimate alongside the traditional pointwise estimate for one top charity next year seems like a prudent way to start out.

## In conclusion

In conclusion, I suggest that GiveWell should be creating distributional rather than pointwise estimates, because this would lead to more accurate results and a better treatment of uncertainty.

[^1]: And in fact, I think this corresponds to GiveWell's current estimate, obtained by dividing over [this cell](https://docs.google.com/spreadsheets/d/1Kq6iHSQFr3eRz1p9KclHuJTQiaJYkpViOyneSd9KCJc/edit#gid=1680005064), because one "unit of value" is one doubling of consumption for one year.

[^2]: One shortcut when considering a _new_ charity might be to use a guess as to the distribution of cost-effectiveness for all past charities GiveWell has considered in the past.

[^3]: Right now, AMF assigns a homogeneous value to lives within a bucket of ~5 years, and then estimates what proportion of recipients belong to each bucket. This is complicated and clunky; one could instead do something like what's hinted [here](https://forum.effectivealtruism.org/posts/4Qdjkf8PatGBsBExK/adding-quantified-uncertainty-to-givewell-s-cost?commentId=q6EQPSmMCg8cTiEDN), i.e., define continuous functions for each of the components of the estimation, and then combine them. In particular, the value of a life of age _t_ can be estimated as the life expectancy for someone that age, without the need for bucketing.
