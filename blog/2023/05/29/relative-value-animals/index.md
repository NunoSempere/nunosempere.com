Relative values for animal suffering and ACE Top Charities
==========================================================

tl;dr: In this post, I present relative estimates for animal suffering and 2022 top Animal Charity Evaluators (ACE) charities. I am doing this to showcase a new tool from the Quantified Uncertainty Research Institute (QURI) and to present an alternative to ACE's current rubric-based approach.

### Introduction and goals

At QURI, we're experimenting with using relative values to estimate the worth of various items and interventions. Instead of basing value on a specific unit, we ask how valuable each item in a list is, compared to each other item. You can see an overview of this approach [here](https://forum.nunosempere.com/posts/EFEwBvuDrTLDndqCt/relative-value-functions-a-flexible-new-format-for-value).

In this context, I thought it would be meaningful to estimate some items in animal welfare and suffering. I estimated the value of a few a few animal quality-adjusted life-years—fish, chicken, pigs and cows—relative to each other. Then I using those, I estimated the value of top and standout charities as chosen by ACE (Animal Charity Evaluators) in 2022. 

This exercise might perhaps be useful to ACE, not necessarily from the estimates themselves, which are admittedly mediocre, but rather by considering these estimates as a potential template for evaluating the value of uncertain interventions outside of global health and development.

![](https://images.nunosempere.com/blog/2023/05/29/estimation-mountain.png)

### Link to the model

You can view these relative estimates [here](https://relative-values-git-animals-2023-04-quantified-uncertainty.vercel.app/interfaces/relative-values-animals-2023-04/models/relative-values-animals-2023-04) ([a](https://web.archive.org/web/20230529163148/https://relative-values-git-animals-2023-04-quantified-uncertainty.vercel.app/interfaces/relative-values-animals-2023-04/models/relative-values-animals-2023-04)). The app in which they live has different views:

A view showing all of the estimates compared to each other:

![](https://images.nunosempere.com/blog/2023/05/29/relative-app-1.png)

A view showing items’ values compared to one reference item:

![](https://images.nunosempere.com/blog/2023/05/29/relative-app-2.png)

There is also a view plotting [uncertainty vs value](https://relative-values-git-animals-2023-04-quantified-uncertainty.vercel.app/interfaces/relative-values-animals-2023-04/models/relative-values-animals-2023-04/plot), and a view showing [the underlying code](https://relative-values-git-animals-2023-04-quantified-uncertainty.vercel.app/interfaces/relative-values-animals-2023-04/models/relative-values-animals-2023-04/edit), which is editable (though slow). 

### Discussion

#### Expected quality of the model

I expect these estimates to have numerous flaws. Previously, I worked on an aggregator for forecasts called Metaforecast, as part of which I assigned a “stars rating” to quickly signal the expected quality of probabilities from different platforms. If I applied that same rating here, these estimates would have a stars quality rating of one out of five possible stars, at most.

One particular flaw that reviewers highlighted is that these estimates are in terms of what I value, after short amount of reflection——as opposed to e.g., what the whole Effective Altruism community would value after a whole lot of reflection, or as opposed to any particular fleshed-out philosophical view. I mostly chose to do this for speed. Future estimates in this vain might want to something more sophisticated, like letting the user input their own values, or letting the user choose between several philosophical views, or aggregate them.

One key insufficiency of these estimates is that they estimate what I personally value after a short amount of reflection. They don't necessarily represent what the entire Effective Altruism community or any particular philosophical viewpoint might value after in-depth reflection. I chose this approach mainly for efficiency. Future iterations might adopt a more sophisticated approach, such as allowing users to input their own values, or selecting from several philosophical perspectives, or aggregating them.

### Methodology

I came up with these estimates in three steps:

1. Estimated the Quality-Adjusted Life Years (QALYs) value of a few animal species
2. Mechanistically estimated the value of three reference charities, in terms of QALYs for the species valued in step 1
3. Estimated the value of the remaining charities in terms of the reference charities in step 2.

#### Estimating relative value of animal QALYs

I started by estimating the relative value of a QALY for a few animal species. Then I derived estimates for QALY per kilogram and QALYs per calorie, which could later be useful for improving calculators like [this one](https://foodimpacts.org/).

Here is an example for cows:

```
// Add human QALY as a reference point
one_human_qaly = {
  id: "one_human_qaly", 
  name: "1 human QALY (quality-adjusted life-year)",
  value: normal(1, 0.01)
}

// Cows
value_happy_cow_year = 0.05 to 0.3 
// ^ in human qalys
value_tortured_cow_year = -(0.1 to 2)
value_farmed_cow_year = normal({ p10: -0.2, p90: 0.1 })
// ^ purely subjective estimates
// the thing is, it doesn't seem that unlikely to me
// that cows do lead net positive lives
weight_cow = mixture([450 to 1800, 360 to 1100], [1/2,1/2])
non_wastage_proportion_cow = (0.5 to 0.7) -> ss // should be a beta. 
lifetime_cow = (30 to 42) / 12
calories_cow = mixture(0.8M to 1.4M, (500k to 700k) * (weight_cow * non_wastage_proportion_cow)/1000) 
// ^ kilocalories, averaging two estimates from
// <https://www.reddit.com/r/theydidthemonstermath/comments/a8ha9r/how_many_calories_are_in_a_whole_cow/>

cow_estimates = {
  name: "cow",
  value_year: value_farmed_cow_year -> ss,
  weight: weight_cow,
  calories: calories_cow,
  lifetime: lifetime_cow -> ss
}
```

#### Coming up with mechanistic estimates for three reference projects

I then looked at three reference projects for which I thought a mechanistic estimate might be feasible: the Fish Welfare Initiative (FWI), Beyond Burgers, and the Open Wing Alliance. For each of those projects, I estimated how many specific animals did they affect, and by how much, and arrived at a wide subjective estimate of their impact.

For example, in the case of FWI I looked at their impact page for the number of animals they probably have helped. I then came up with an uncertain estimate for how much they had helped each animal. I took various shortcuts, for example, I pretended that the fish which FWI helped were salmon, because details about their life expectancy and caloric content were easy and quick to look up online. In fact, I expect the vast majority of fish that FWI helps to not be salmon, but I don't expect the difference to matter all that much when estimating total impact.

Here is how my estimate for the Fish Welfare Initiative looks:

```
fish_potentially_helped = 1M to 2M
shrimp_potentially_helped = 1M to 2M
improvement_as_proportion_of_lifetime = (0.05 to 0.5) -> ss
sign_flip_to_denote_improvement(x) = -x

value_fwi_fish = (
    fish_potentially_helped * 
    improvement_as_proportion_of_lifetime * 
    (
       salmon_estimates.value_year /
       Salmon_estimates.lifetime
    )
  ) -> sign_flip_to_denote_improvement

value_of_shrimp_in_fish = (0.3 to 1)
// ^ very uncertain, subjective
value_fwi_shrimp = (
    shrimp_potentially_helped * 
    improvement_as_proportion_of_lifetime * 
    (
       salmon_estimates.value_year /
       Salmon_estimates.lifetime
    ) *
    value_of_shrimp_in_fish
  ) -> sign_flip_to_denote_improvement

value_fwi_so_far = value_fwi_fish + value_fwi_shrimp
proportion_fwi_in_2022 = 1/4 to 1/2
value_fwi  = value_fwi_so_far * proportion_fwi_in_2022 

fwi_item = {
  name: "Fish Welfare Initiative",
  year: 2022, 
  slug: "fish_welfare_initiative",
  value: value_fwi -> ss 
}
```

#### Estimating other charities in terms of the reference projects

I estimated the value of the remaining projects in terms of the previous three. For example, here is my estimate of the Good Food Institute:

```
value_reference_top_animal_org = mixture(
  [
    fwi_value,
    open_wing_alliance_value, 
    beyond_meat_value/(10 to 1k)
    // ^ beyond meat seems significantly more scaled up than the avg org working to affect cows
  ], 
  [ 1/3, 1/3, 1/3 ]
) -> SampleSet.fromDist

beyond_meat_equivalents_gfi = 0.01 to 2
value_gfi = mixture(
  [ 
    beyond_meat_equivalents_gfi * beyond_meat_item.value, 
    value_reference_top_animal_org
  ], 
  [ 2/3, 1/3 ]
)
```

and here is my estimate for Compassion USA:

```
value_compassion_usa = mixture(
  [
    open_wing_alliance_value * 
      truncateRight(0.05 to 10, 100), 
    value_reference_top_animal_org * 
      truncate(0.05 to 10, 100)
  ], 
  [ 1/2, 1/2 ]
) 
```


#### A comment on maintaining correlations

These estimates are written in Squiggle, which aims to make it easy to do relative values through its functionality around sample sets. For example,

```
x = SampleSet.fromDist(1 to 100) 
y = 2 * x y/x 
// ^ is a distribution which is 2 everywhere
or
x = SampleSet.fromDist(1 to 100)
y = SampleSet.fromDist(2 to 200) 
z = x * y (z/x) / y 
// ^ is one everywhere
```

I’ve usually shortened SampleSet.fromDist to just “ss.”

As a note of caution, note that mantaining correlations while having mixtures of different distributions is more tricky.

### Conclusion

This post presents a model that starts with very rough estimates of the value of several types of animal suffering. It then uses these to build up mechanistic estimates of a few animal charities, and then uses those mechanistic estimates to give a guess as to the impact of all top ACE charities in 2022. 

The motives for doing that were:

- To showcase some tooling recently built at QURI
- To show one possible path for having quantified estimates for speculative projects—as opposed to the rubric-based approach that organizations like ACE or Charity Entrepeneurship use.

### Acknowledgements

<p><img src="https://images.nunosempere.com/quri/logo.png" style="width: 20%;">
<br>

This is a project of the Quantified Uncertainty Research Institute, from which I've since then taken a leave of absence. Thanks to Ozzie Gooen and Holly Elmore for their feedback.

<p>
  <section id='isso-thread'>
  <noscript>Javascript needs to be activated to view comments.</noscript>
  </section>
</p>
