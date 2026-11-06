$5k challenge to quantify the impact of 80,000 hours' top career paths
==============

## Motivation

80,000 hours has identified a number of promising career paths. They have a fair amount of analysis behind their recommendations, and in particular, they have a list of [top ten priority paths.](https://80000hours.org/career-reviews/#our-priority-paths) 

However, 80,000 hours doesn’t quite[^1] have quantitative estimates of these paths' value. Although their usefulness would not be guaranteed, quantitative estimates could make it clearer:

*   how valuable their top career paths are relative to each other
*   how valuable their top career paths are relative to options further down their list
*   at which level of personal fit one should switch between different career paths[^2]
*   where the expected impact is coming from, and which variables we are most uncertain about
*   eventually, whether certain opportunities are valuable in themselves or for the value of information or career capital that they provide
*   etc.

**The Prize**

Following up on the [$1,000 Squiggle Experimentation Challenge](https://forum.effectivealtruism.org/posts/ZrWuy2oAxa6Yh3eAw/usd1-000-squiggle-experimentation-challenge) and the [Forecasting Innovation Prize](https://forum.effectivealtruism.org/posts/8Nwy3tX2WnDDSTRoi/announcing-the-forecasting-innovation-prize) we are offering a prize of $5k for quantitative estimates of the value of [80,000 hours' top 10 career paths](https://80000hours.org/career-reviews/#our-priority-paths). 

## Rules

**Step 1:** Make a public post online between now and December 1, 2022. Posts on the EA Forum (link posts are fine) are encouraged.  
  
**Step 2:** Complete [this submission form.](https://airtable.com/shrFX656rcOXaDioh)

**Further details**

*   Participants can use units or strategies of their choice—these might be QALYs, percentage points of reduction in existential risk, [basis points of the future](https://forum.effectivealtruism.org/posts/xmmqDdGqNZq5RELer/shallow-evaluations-of-longtermist-organizations#Notes), [basis points of existential risk reduced](https://forum.effectivealtruism.org/posts/cKPkimztzKoCkZ75r/how-many-ea-2021-usds-would-you-trade-off-against-a-0-01), career-dependent units, etc. Contestants could also use some other method, like [relative values](https://forum.effectivealtruism.org/posts/hrdxf5qdKmCZNWTvs/valuing-research-works-by-eliciting-comparisons-from-ea), estimating proxies, or some original option.
*   We are specifically looking for quantitative estimates that attempt to estimate some magnitude reasonably close to the real world, similar to the units above[^3]. So for example, assigning valuations from 0 to 5 stars would not fulfil the requirements of the contest, but estimates in terms of the units above would qualify.
*   Participants are free to estimate the value of one, several, or all ten career paths.
*   Participants are free to use whatever tool or language they want to produce these estimates. Some possible tooling might be: Excel, [Squiggle](https://www.squiggle-language.com/), [Guesstimate](https://getguesstimate.com/), probabilistic languages or libraries (e.g., Turing.jl, PyMC3, Stan), [Causal](https://causal.app/), working directly in a popular programming language, etc.
*   Participants _can_ provide point estimates of impact, but they are encouraged to provide their estimates as distributions instead. 
*   Participants are free to estimate the impact of a marginal person, of a marginal person with a good fit, the average value, etc. Participants are welcome to provide both average and marginal value—for example, they could provide a function which provides an estimate of marginal value at different levels of labor and capital.

We provide some examples of possible rough submissions in an appendix. We are also happy to comment on estimation strategies: feel free to leave a comment on this post or to send a message to Nuño Sempere using the EA forum message functionality.

## Judging

The judges will be Nuño Sempere, Eli Lifland, Alex Lawsen and Sam Nolan. These judges will judge on their personal capacities, and their stances do not represent their organizations.

Judges will estimate the quality and value of the entries, and we will distribute the prize amount of $5k[^4] in proportion to an equally weighted aggregate of those subjective estimates[^5].

To reduce our operational burden, we are looking to send out around three to five prizes. If there are more than five submissions, we plan to implement a lottery system. For example, a participant who would have won $100 would instead get a 10% chance of receiving $1k.

## Acknowledgements

<p><img src="https://i.imgur.com/7yuRrge.png" class="img-frontpage-center"></p>

This contest is a project of the Quantified Uncertainty Research Institute, which is providing the contest funds and administration. Thanks in advance to Eli Lifland, Alex Lawsen and Sam Nolan for their good judgments. Thanks to Ozzie Gooen for comments and suggestions.

## Appendix: Example models

### Example I: "[Founder of new projects tackling top problems](https://80000hours.org/career-reviews/founder-impactful-organisations/)"

The following is a crude example estimate for the career path of [Founder of new projects tackling top problems](https://80000hours.org/career-reviews/founder-impactful-organisations/), written in [Squiggle](https://www.squiggle-language.com/playground#code=eNq1lN9P2zAQx%2F8Vq0%2FplCZuuwCqtIdKMLQfXceKNqGFB5M6iYdjZ%2FYFGhD%2F%2B85OC90QUB72kF%2F23dd3n7vLbc%2BW%2BnrRVBUzbW8CpuGhXzpaCtBmsyKUAMHk4ncjikLyBRihit6kF8dkXnPDQGjFJMlKJiVXBbepwneV8Xk%2BBWDoropT01iYquUxB8DP941a4oO8IxccWJCEZEz7BAUhoBFNCGhCo3FIhv0HrbXrVM1NwZS48efO89OSHxuNehuxIUW%2FB7VxJ3bQiaUKl48siIoBJzonoqpZBqlqOTOy3ZZ%2BiLFaBQm9dDoJpZchGdHua0hnIUlm3Qa%2B%2FqTRHqZCo5G%2F0%2BQcTyzEFT8Uhmcg2%2B9MNpjJyfTzmf3KzaGWkhk8YBgHwz3qdEb7dB06u%2BSK5EZXxOqKkwWryBctmSJ8Hb2NUqW34vXis0aCqKXgpos7VcSFNfThIgL3RBQH%2FomgDvyp4xG%2B4m2T4nno%2Fd6GzhbvI%2FQ8T1Xf07vipiVIvChlS4SyNSa3JBfthJQAtZ3Eca5NU0U8zzFpzJ5J7CNhqwijjWttwcbHN9XHUSNOV2%2BLs7LOTj7F99aKWzsQdsAGmVa%2FGpW53AY6H1T3qWGD2ZLVkrd%2FpYvcu1onqZIi54gJYW%2BXtGPirfaxcl26m66YFoXhBYJNFWhgcl2sIwWG1xhWY2wp6g%2FKVw%2BVdu3xN47lzk3srZ9uRr%2F9ckt5s%2Beawxs8AcnvPQL8LyVyLaAkFVuJqqmwIhAshYXQrfSRDpZcZWj2TRTl9tYucH%2Bg8oytEPI2uv%2FAGYInSYduEGb9VwGH4Bnk7qfk9Nxwa5wh%2FFnejzIxOCQ4TZYrK9wQuM6Ekm%2FwkpIbHr1UMggeFW3rp3cobC1Zm6pb4ozRfOc6TNYer%2FFBlzu8end%2FAB%2BtOwA%3D).

```
// Operational challenges
chanceOfAttainingTrustAndGettingFunding = beta(5, 30) // t(0.05 to 0.3, 1)
chanceOfGettingAnOrganizationOfTheGround = beta(10, 10) // t(0.3 to 0.8, 1)

// Estimate of impact
yearlyOrganizationFunding = mx(50k to 500k, 200k to 10M, 5M to 50M, [0.65, 0.25, 0.05])
giveDirectlyValueOfQALYsPerDollar = 1/(160 to 2700)
// ^ taken from some Sam Nolan’s estimates: 
// <https://observablehq.com/@hazelfire/givewells-givedirectly-cost-effectiveness-analysis>
organizationValueMultiplier = mx(
  [0.1 to 1, 1 to 8, 8 to 80, 80 to 320, 320 to 500k],
  [4, 8, 4, 2, 1]
)
// very roughly inspired by:
// https://forum.effectivealtruism.org/posts/GzmJ2uiTx4gYhpcQK/
//  effectiveness-is-a-conjunction-of-multipliers
shapleyMultiplier = 0.2 to 0.5
lifetimeOfOrganization = mx(2 to 7, 5 to 50)
 
// Aggregate
totalValueOfEntrepeneurshipInQALYs = chanceOfAttainingTrustAndGettingFunding *
  chanceOfGettingAnOrganizationOfTheGround *
  yearlyOrganizationFunding *
  giveDirectlyValueOfQALYsPerDollar *
  organizationValueMultiplier *
  lifetimeOfOrganization *
  shapleyMultiplier
 
// Aggregate with maximums
t(dist, max) = truncateRight(dist, max)
totalValueOfEntrepeneurshipInQALYsWithMaxs =
  chanceOfAttainingTrustAndGettingFunding *
  chanceOfGettingAnOrganizationOfTheGround *
  t(yearlyOrganizationFunding, 500M) *
  giveDirectlyValueOfQALYsPerDollar *
  t(organizationValueMultiplier, 10M) * 
  // ^ overall estimate really sensitive to the maximum here.
  lifetimeOfOrganization *
  t(shapleyMultiplier, 1)

 // Display
{ 
    totalValueOfEntrepeneurshipInQALYsWithMaxs: 
      totalValueOfEntrepeneurshipInQALYsWithMaxs 
} 
```

Alone, the estimate might be too obscure, so it would be better if it were accompanied by some explanation about the estimation strategy it is using. So, its estimation strategy is:

*   To estimate the chance of getting funding and then getting an organization off the ground
    *   This is based on subjective guesses. Perhaps Charity Entrepreneurship, or EA funds if it kept data, could have better estimates
*   To estimate the value that an organization produces. This is the weakest part of the model, and it would be better if it were based on specific steps. Instead, we are using more of a "black box" model, and estimating:
    *   The funding that the organization would receive
    *   The QALYs per dollar that a reference organization—GiveDirectly—produces, taken from [Sam Nolan’s estimate thereof](https://observablehq.com/@hazelfire/givewells-givedirectly-cost-effectiveness-analysis).
    *   The advantage over GiveDirectly that the new organization would have. We are getting this estimate from [this EA forum post](https://forum.effectivealtruism.org/posts/GzmJ2uiTx4gYhpcQK/effectiveness-is-a-conjunction-of-multipliers)
*   To estimate some other factors to go from the above to the total output, again based on pretty subjective estimates:
    *   The lifetime of the organization
    *   The "Shapley multiplier" penalizes efforts which require more people. In this case, we are saying that the founder gets between 20% and 50% of the impact.

We also have to take care that not only the 90% confidence interval, but also the overall shape of the estimates was correct. For this reason, we have a step where we truncate some of them.

As mentioned, a key input of the model is the multiplier of impact over GiveDirectly, but this is based on black box reasoning. This could be a possible point of improvement. For example, we could improve it with an estimate of how many QALYs, or what percentage of the future is an speculative area like AI safety research worth.

### Example II: Value of global health charities

There are various distributional models of global health charities in the EA forum that participants may want to take some inspiration from, e.g.:

*   Dan Wahl’s estimate of the cost-effectiveness [of LEEP](https://danwahl.net/blog/leep-cea)
*   Sam Nolan’s [cost-effectiveness models](https://forum.effectivealtruism.org/users/hazelfire)

The advantage of these is that they can be pretty clean. The disadvantage is that they come from a different cause area.

### Example III: Value of the Centre for the Governance of AI

[Here](https://forum.effectivealtruism.org/posts/EPhDMkovGquHtFq3h/an-experiment-eliciting-relative-estimates-for-open#A_concrete_example), I give an estimate for the value of the Centre for the Governance of AI (GovAI) in terms of basis points of existential risk reduced. It might serve as a source of inspiration. One disadvantage is that it only considers one particular pathway to impact that GovAI might have, and it doesn't consider other pathways that might be more important—e.g., field-building.

### Example IV: Value of ALLFED

Historically, one of the few longtermist organizations which has made an attempt to estimate their own impact quantitatively is ALLFED. A past estimate of theirs can be seen [here](https://www.getguesstimate.com/models/9782). My sense is that the numeric estimates might have been on the optimistic side (some alternative numbers [here](https://forum.effectivealtruism.org/posts/xmmqDdGqNZq5RELer/shallow-evaluations-of-longtermist-organizations#Disagreements_and_Uncertainties)). But the estimation strategy of dividing their influence and impact depending on different steps might be something to take inspiration from.about  
 

[^1]: 80,000 hours, when thinking abou their own impaabouct, internally use "discounted impact-adjusted peak year" (DIPY). But this seems like a fairly coarse unit.
    
[^2]: This is actually more nuanced. There might be some frustration about people quickly/naïvely jumping to whatever cause or sub-cause has the best apparent marginal value at each point in time rather than committing to something. But this might be counterproductive if people have more impact staying in one place, or if impact is a combination of people working on different areas. For a specific example, suppose that impact is a [Cobb–Douglas](https://en.wikipedia.org/wiki/Cobb%E2%80%93Douglas_production_function) function of work in different areas, and that there is some coordination inefficiencies. Then focusing on attaining the optimal proportion of people in each area might be better than aiming to estimate marginal values through time.
    
[^3]: The criteria isn't exactly to have a unit such that 2x on that unit is twice as better. For example, percentage reductions of existential/catastrophic risk in the presence of several such risks aren't additive, but we would accept such estimates. Similarly, relative values can only be translated to magnitudes in an "additive" unit with a bit of work, but we would also accept such estimates. 
    
[^4]: Having a fixed pot is slightly less elegant than deciding beforehand on an amount to reward for a given level of quality, but it comes with an added operations burden/uncertainty.
    
[^5]: For example, if we get two submissions and we estimate the first one to be twice as valuable as the second one, the first submission would receive $3.33k and the second submission would receive 1.66k. Instead, if the first submission's individual estimates were estimated to be twice as valuable, but also were twice as many in number as those of the second submission, the first one would receive $4k and the second one would receive $1k.
