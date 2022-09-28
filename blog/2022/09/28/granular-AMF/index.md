Use a less coarse analysis of AMF beneficiary age and consider counterfactual deaths
===================================================================================

**tl;dr**: GiveWell considers a fairly coarse division of beneficiary age, and groups children of 0 to 5 years old together. This may lead to inaccurate or inexact calculations. In addition, GiveWell doesn't completely account for counterfactual mortality: where a beneficiary is saved from dying of malaria but dies later anyways, although this inaccuracy is not as severe as it could be.

Following up on [Use distributions to more parsimoniously estimate impact](https://nunosempere.com/blog/2022/09/15/use-distributions-to-more-parsimoniously-estimate-impact/), I was looking at the population analysis of the AMF distributions, because a [previous attempt](https://forum.effectivealtruism.org/posts/4Qdjkf8PatGBsBExK/adding-quantified-uncertainty-to-givewell-s-cost) at adding uncertainty to the analysis was messier than I would have wished.

But following up on  that analysis, I realized that [the strategy GiveWell uses](https://docs.google.com/spreadsheets/d/1tytvmV_32H8XGGRJlUzRDTKTHrdevPIYmb_uc6aLeas/edit#gid=1364064522) is:

1. Divide the population in 5 year chunks (0 to 4, 5 to 9, 10 to 14)
2. Estimate the risk reduction due to malaria nets (applies mostly to the 0 to 4 chunk)
3. Multiply 1. and 2. to arrive at an estimate for total number of under-5 deaths avoided.
4. ...

Now, it's possible that the below is a result of me not understanding factors which have been implicitly considered in the model, but I see two problems with this. One could be seen as kind of a nitpick, and the second one seems like a more substantial point.

## The nitpick: The 0 to 4 years old bucket is too coarse, because mortality is meaningfully different

From [Age-Specific Malaria Mortality Rates in the KEMRI/CDC Health and Demographic Surveillance System in Western Kenya, 2003–2010](https://www.ncbi.nlm.nih.gov/pmc/articles/PMC4152016/), figure 1:

![](https://i.imgur.com/N46sNHe.jpeg)

and from [Patterns of Age-Specific Mortality in Children in Endemic Areas of Sub-Saharan Africa](https://www.ncbi.nlm.nih.gov/books/NBK1688/):

![](https://i.imgur.com/fXhoI9J.jpeg)

So my tentative understanding is that most of the mortality happens in the first two years, and that mortality decays reasonably fast.

Now, why has GiveWell used a coarser aggregation? My sense is that this was a totally reasonable decision, because the above data was annoying to find, and sources (such as the WHO), usually provide it aggregated.

So anyways, disaggregating by year might provide some marginal improvements, because much of the benefit is concentrated amongst younger beneficiaries, rather than being uniformly spread-out between beneficiaries of less than five years old. This might be considered a nitpick, except that...

## The more substantial point: The analysis is not considering counterfactual mortality, and so overestimates impact

As far as I understand—and let me reiterate that I could be wrong here—GiveWell's spreadsheet is considering reductions in mortality due to malaria. But then it is not considering the depressing case where:

- A child is protected from dying from malaria because of a LLIN distribution
- But later dies anyways because of a different disease

This is relevant because malaria is probably not the top cause of death in the regions considered:

<figure>V
![](https://i.imgur.com/Oc3JtrT.jpeg)
<figcaption>Cumulative probability of dying by given age. a: all cause mortality; b: malaria specific mortality</figcaption>
</figure>

<iframe src="https://ourworldindata.org/grapher/causes-of-death-in-children?country=~NGA" loading="lazy" style="width: 100%; height: 600px; border: 0px none;"></iframe>

and because mortality in the regions where AMF distributes nets is relatively high:

<iframe src="https://ourworldindata.org/grapher/child-mortality-around-the-world?country=COG~COD~GIN~TGO~NGA~TCD~SSD" loading="lazy" style="width: 100%; height: 600px; border: 0px none;"></iframe>

So in combination:

1. Much of the benefit of malaria prevention falls on recently born children, rather than being homogeneously distributed across 0 to 5 year-olds
2. But because there are other possible cause of deaths, and because malaria might not be the leading cause of death in the regions considered, young children which are protected from malaria could die of other causes.
3. These two factors together may reduce the expected benefit of LLINs by ~10%

## The solution

The solution is to apply a discount factor corresponding to the probability that a beneficiary who was saved from dying from malaria might die from other factors. To do this, it might be useful to decompose the calculation year by year—or even using narrower bands.

A quick and dirty adjustment would be to add \*(1-B40/100) to row B56 (resp. rows C to J) of the AMF calculation, as in [this spreadsheet](https://docs.google.com/spreadsheets/d/1Gtd5ga6UREGFjnz-V4HiXAZLhOM_ifrv4CchUGSsRf4/edit?usp=sharing). This ends up making a difference of 2.5% (in the case of DRC) to 22.5% (in the case of Chad). 

Besides the issue with child mortality being higher for younger children, I would also argue that children who would have otherwise died of malaria are differentially more likely to die of other causes. For instance, a sickly child is likelier to be saved from malaria, but also likelier to die of other causes. It's possible that this might add another 2 to 10% difference.

Conversely, the [Pryce et al.](http://ncbi.nlm.nih.gov/pmc/articles/PMC6418392/) Cochrane Review on which GiveWell's estimates are based does estimate reduction in *overall mortality*. This diminishes the impact of my argument above, because if a child was saved from malaria but died anyways of something else, it would be counted correctly in the tally of overall mortality. However, the lengths of the followups in the analyses considered by the review are: 2 years, 12 months, 6 months, 24 months, 12 months, 8 months, 12 months, 12 months, 12 months, 2 years, 1 year, 6 months, 2 years, 1 year, 12 months, 2 years, 2 years, 13 months, 6 months, 15 weeks, 1 year, 6 months, and 10 months, which isn't enough to deal with all of the under 5 mortality. 

Looking again at the mortality rates:

<figure>
![](https://i.imgur.com/Oc3JtrT.jpeg)
<figcaption>Cumulative probability of dying by given age. a: all cause mortality; b: malaria specific mortality</figcaption>
</figure>

that consideration probably roughly ~halves the potential adjustment.

The above post was written in response to [GiveWell's Change Our Mind Contest](https://www.givewell.org/research/change-our-mind-contest). But if you are reading this on my blog, you may want to: [Donate to GiveWell](https://secure.givewell.org/).
