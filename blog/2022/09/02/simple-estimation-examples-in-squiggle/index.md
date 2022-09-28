Simple estimation examples in Squiggle
==============

This post goes through several simple estimates, written in [Squiggle](https://www.squiggle-language.com/), a new estimation language. My hope is that it might make it easier to write more estimates of a similar sort, wider adoption of Squiggle itself, and ultimately better decisions. 

## Initial setup

One can use Squiggle in several ways. This blog post will cover using it on its website and in a Google Spreadsheet. An upcoming blog post will cover using it in more complicated setups.

### [squiggle-language.com](https://www.squiggle-language.com/playground)

The simplest way to use Squiggle is to open [squiggle-language.com/playground](https://www.squiggle-language.com/playground). You will see something like the following:

<p><img src="https://i.imgur.com/lqEe46S.png" class="img-medium-center"></p>

You can write your model on the editor in the left side, and the results will be displayed in the right side. As you make edits, the url in your browser bar will change so that you copy it and use it to refer people to your model.

### [Squiggle <> Google docs](https://docs.google.com/spreadsheets/d/1uKtQDlOQ0UYJznlzDNAwSdvwIB6bYv4tT8Z5gIQi7Fc/edit#gid=0)

When working with multiple models, I’ve also found it useful to use Squiggle in Google sheets. To do so, make a copy of [this spreadsheet](https://docs.google.com/spreadsheets/d/1uKtQDlOQ0UYJznlzDNAwSdvwIB6bYv4tT8Z5gIQi7Fc/edit#gid=1407581126), and allow app permissions. Edit the “Main” sheet, and click on “Squiggle” > “Feed into Squiggle” to compute models. If you have difficulties, read the “Instructions” sheet, or leave a comment. 


<p><img src="https://i.imgur.com/yCskQSU.png" class="img-medium-center"></p>

So without further ado, the simple example models:

## Partially replicating [Dissolving the Fermi Paradox](https://arxiv.org/abs/1806.02404) (complexity = 1/10)

Page 2 of the paper defines the factors for the [Drake equation](https://en.wikipedia.org/wiki/Drake_equation#Equation):

<p><img src="https://i.imgur.com/e5eAKsg.png" class="img-medium-center"></p>

Page 10 of the paper gives its estimates for the factors of the [Drake equation](https://en.wikipedia.org/wiki/Drake_equation#Equation):

<p><img src="https://i.imgur.com/F53wXSL.png)" class="img-medium-center"></p>

Because Squiggle doesn’t yet have the log-uniform probability distribution, we’re going to have to define it first. A [log-uniform](https://en.wikipedia.org/wiki/Reciprocal_distribution) is a probability distribution whose log is a uniform distribution.

```
// Define the log-uniform 
// (<https://en.wikipedia.org/wiki/Reciprocal_distribution>)
loguniform(a, b) = exp(uniform(log(a), log(b)))

// Estimates
rate_of_star_formation = loguniform(1,100)
fraction_of_stars_with_planets = loguniform(0.1, 1)
number_of_habitable_planets_per_star_system = loguniform(0.1, 1)
fraction_of_habitable_planets_in_which_any_life_appears = 1
// ^ more on this below
fraction_of_planets_with_life_in_which_intelligent_life_appears = loguniform(0.001, 1)
fraction_of_intelligent_planets_which_are_detectable_as_such = loguniform(0.01, 1)
longevity_of_detectable_civilizations = loguniform(100, 10000000000)

// Expected number of civilizations in the Milky way;
// see footnote 3 (p. 5)
n = rate_of_star_formation *
  fraction_of_stars_with_planets *
  number_of_habitable_planets_per_star_system * 
  fraction_of_habitable_planets_in_which_any_life_appears * 
  fraction_of_planets_with_life_in_which_intelligent_life_appears * 
  fraction_of_intelligent_planets_which_are_detectable_as_such *
  longevity_of_detectable_civilizations

// Display
n 
```

This produces the [following estimate](https://www.squiggle-language.com/playground#code=eNqdkk1vnDAQhv%2FKiBNE2wVU9bJtc0qOvXSvqJaBAUYxNrGHEFr1v9d4P7QfidotF4TH7%2FPO8M6vyHVm2o59L%2B0cbRqpHK7C2WNNbGy0YTv6E9LEJNX2eaS2VbhlS7qNNlGawgM2pBG4Q1Cm%2FTBqaoztodC%2BFn%2FpmAe3SVPU64meaMCa5NrYNl2%2B0u9Y0WBNJZWoyXloOTIZfZ8U2rP2qFiuoEzgK%2BDrEB%2FOfDmWyWqxjMsk8YJg%2BOiYesnoCm39S5hGOJZWLBq5oD3mhJyv8izz2sbKaqke7jsxEXdiUFIju3NNts5XkHuRHvsS7SLpZEksS4UHhRh8IRi72TH27xBOba8ZpMXUUdUJqWehqEEhhwF9c56Wh2l%2FQG8sgp%2BKO3JQojLTOfXACuMExpFKmlEpalHzJf2s1yx7o9tT8dFj16xFUSNjtRtGOuHGqruC7pnK6BZfiOcFeiKr6IUU%2FQyRXXTkE%2FPa7Pgco38dvBpr2OUCpoFzCumwpN9IPc0wyflzkDlEaIxhbRjhI8TDGj4t4XrTdzbortAAf9mYcOeWBbmDS%2BotC3Gt%2Fp%2Fgryk3xxzm%2FqdQ96k9kPPgefnhhY5%2B%2FwEFqbCw):

<p><img src="https://i.imgur.com/rKqyWoZ.png" class="img-medium-center"></p>

The estimate is fairly wide, but the model gives a 10%-ish chance that there is, in expectation, less than once civilization in the Milky Way. After updating on a bunch of observations, the paper raises that probability, hence the conclusion that the Fermi paradox has been “dissolved”. 

Why did we set *fraction_of_planets_in_which_any_life_appears* to 1? Well, the paper considers an estimate of 1−exp(−r), where r is distributed as a lognormal(1,50). But because r ranges from very small numbers to very large numbers, they get collapsed to either 0 or 1 when going through 1−exp(−r), which produces some numerical errors when multiplying by 0. In addition, that estimate has been [questioned](https://forum.effectivealtruism.org/posts/kvZshdx5FzTPjyhxG/the-fermi-paradox-has-not-been-dissolved). 

So following a similar move in the paper, we can set that factor to a high value (in this case, to 1, meaning that all planets capable of life do host life). And then, when we notice that the probability of no other life in the Milky Way is still significant, the Fermi paradox will still have been somewhat dissolved, though to a lesser extent.

From here on, we could tweak the rate of life formation to other plausible estimates, like:

- 0.01 to 1, per [this post](https://forum.effectivealtruism.org/posts/kvZshdx5FzTPjyhxG/the-fermi-paradox-has-not-been-dissolved#comments).
- 0.001 to 1
- 10^-100 to 0.1
- 1 - (10^(-36) to 1), which approximates 1 - exp(-lognormal(1,50))
- (10^(-36) to 1), which flips the above

We could also [question the other estimates](https://forum.effectivealtruism.org/posts/kvZshdx5FzTPjyhxG/the-fermi-paradox-has-not-been-dissolved#The_origin_of_life), or notice that the uncertainty over the rate of life formation probably dominates other uncertainties.

## Expected value for a list of things (complexity = 2/10)

Moving on to a slightly more complicated example, I was recently talking with someone who was thinking about various career pathways. To estimate the effectiveness of each opportunity, I created the following [spreadsheet](https://docs.google.com/spreadsheets/d/1BlTejl2qh9gOoQVVIkdimdj3oE5vgeUUfxMUigP-qOg/edit#gid=0) (project name and notes censored): 

<p><img src="https://i.imgur.com/RW8H2Wd.png" class="img-medium-center"></p>

That is, I decomposed each job opportunity into:

*   value if successful, 
*   probability of it being successful,
*   effort required. 

I estimated each of the values, combined them as strings using a Google Sheets function, and then fed it to Squiggle using a copy of [this Google spreadsheet](https://docs.google.com/spreadsheets/d/1uKtQDlOQ0UYJznlzDNAwSdvwIB6bYv4tT8Z5gIQi7Fc/edit?usp=sharing). The whole thing took 10 mins. And the complexity in this example comes from integrating Squiggle with Google docs, which was annoying to set up but which shouldn’t be as annoying to use.

One interesting feature of these estimates is that the value estimates were [relative](https://forum.effectivealtruism.org/posts/hrdxf5qdKmCZNWTvs/valuing-research-works-by-eliciting-comparisons-from-ea), rather than using a concrete unit like QALYs. That is, the thing that has meaning is the ratio between the different values in that column, not the absolute values. In this case, this makes estimates much, much quicker to produce, but more error-prone, and harder to interpret, improve, or disagree with.

## Probability of a nuclear explosion in London (complexity = 3/10)

A [logic model](https://en.wikipedia.org/wiki/Logic_model) (I also use pathway to impact as a synonym for this) is a chain between inputs and outcomes, which at its simplest might look as follows:  
 
<p><img src="https://i.imgur.com/ev5fz2m.png" class="img-medium-center"></p>

When doing quick estimates, we can keep this pathway to impact implicit. But when working with slightly more complex models, it can be worth it to write it explicitly. I’ve been using [excalidraw](https://excalidraw.com/) for this, and I would recommend it.  

For example, when calculating the chance that a nuclear bomb hits London, we can draw a logic model like:

<p><img src="https://i.imgur.com/Zzckq5Y.png" class="img-medium-center"></p>

To get a sense of the historical frequency, we can use Laplace’s rule of succession. Per Wikipedia on [Russia–NATO relations](https://en.wikipedia.org/wiki/Russia%E2%80%93NATO_relations), the one direct NATO/Russia incident was the [shootdown of a Russian plane by Turkish forces in 2015](https://en.wikipedia.org/wiki/2015_Russian_Sukhoi_Su-24_shootdown). Putting numbers on this, we have:

```
num_years_NATO_Russia_relations = 2022 - 1991
num_years_with_direct_confrontation = 1
laplace(successes, failures) = (successes + 1)/( successes + failures + 2) // this is a function!

[ 
  laplace(num_years_with_direct_confrontation,
  num_years_NATO_Russia_relations), laplace(0, num_years_NATO_Russia_relations)
] // this is an array

// [ 0.0588, 0.0303 ] = [ 5.88%, 3.03% ]
```

That is, if we count the Turkish incident as a “confrontation”, the base rate implied by Laplace’s rule would be 5.88%, and if we don’t, it would be 3.03%. See also some more discussion about possible confrontations [here](https://twitter.com/NunoSempere/status/1547245909930639361). I think that when I’m thinking of a direct NATO/Russia confrontation, I’m thinking more of NATO’s Article 4 being invoked.

In any case, giving some numbers I think are plausible, we would have something like:

```
chance_of_direct_nato_confrontation = 0.005 to 0.05
and_use_of_nuclear_weapons_in_Europe = 0.01 to 0.2
and_use_of_nuclear_weapons_in_London = 0.05 to 0.3


use_of_nuclear_weapons_in_London = chance_of_direct_nato_confrontation * 
  and_use_of_nuclear_weapons_in_Europe * 
  and_use_of_nuclear_weapons_in_London
  
yearly_probability = mean(use_of_nuclear_weapons_in_London)
monthly_probability = 1 - (1 - yearly_probability)^(1/12) // roughly
```

  
This produces an estimate of ~0.02% per year, or 0.0016% per month. We can then transform these into expected hours lost:

```
remaining_life_expectancy_in_years = 40 to 60
days_in_a_year = 365
hours_in_a_day = 24

expected_lost_hours = monthly_probability * remaining_life_expectancy_in_years * days_in_a_year * hours_in_a_day

{
    expected_lost_hours: expected_lost_hours, 
    mean_expected_lost_hours: mean(expected_lost_hours)
} // this is a record
```

The last four lines are a record. It is telling Squiggle two things:

*   To build a record, an object which contains other objects, similarly to how one would do so in Javascript
*   Because the record isn’t assigned to any variable, to only display that record

This produces the following output:

<p><img src="https://i.imgur.com/0tTT6bP.png" class="img-medium-center"></p>

From here, we could refine these estimates, or add some probability of getting out of London before a bomb hits. 

In this model, we had a few moving parts, so now seems like a good time to mention that we can [link to a squiggle model directly](https://www.squiggle-language.com/playground/#code=eNqNUctugzAQ%2FBWLU0BpC3kdkHrMrbdcra4cMGDJrKmxlaIo%2F17bpIc2lsLFsmZ3dmZ3rsnYqcvJ9j3TU1I2TI58HbBjLYzSSWm0dYhAYQSTpy8r2lbyk9EC26RMqo5hxUE1UAvNKwPIjIJKYaMVGmaEQvJO8tc83xOj%2FGdPkWENdgwstJXkTMOFs0HhCALhaLUa%2BMwqZtLmGedDYf2rdBfaUqS4gLFkg4xQJGSR7SWts7b3N7mKnGDQ6szOQgozOUc9Z7h6xk4p9s5f98AuyAtZ%2Bedxdvq5Kt6KTeqVNe%2BZCxVbkKLhwL8Ht7u7xOQlPHV0s3a5v%2BUhp1izKYizUHOl7cEF2Smr77BrcOhm52fPw3gNUo0GQpNfK2I3Iwt8ZOSfekb%2BCnvNqz%2F748plBFv71oibMgaG5shCZQxch%2FB9fhClhGQjFZfIjWJy%2BwGOw1Uj) by copying the link which the playground produces when edited.

## [Bilateral Ceasefire between Russia & Ukraine](https://www.metaculus.com/questions/10050/bilateral-ceasefire-between-russia--ukraine/#comment-95768), using Laplace’s rule of succession (complexity = 4/10)

Metaculus has a question on a [Bilateral Ceasefire between Russia & Ukraine](https://www.metaculus.com/questions/10050/bilateral-ceasefire-between-russia--ukraine/#comment-95768). The question asks about two different things:

*   Whether there will be a ceasefire
*   Whether it will last for at least 30 days.

We can use [Laplace’s rule of succession](https://en.wikipedia.org/wiki/Rule_of_succession) for both cases, first to estimate the likelihood that a ceasefire starts, and then to estimate whether it will last. For the latter, we will make the Laplace prediction, but also look at a [list of ceasefires](https://en.wikipedia.org/wiki/Category:Ceasefires) to estimate the [historical frequency of success](https://docs.google.com/spreadsheets/d/1ufNP_iuTghzKGUCRPMoLY4-c9PibiNlVFEaAxyAad-A/edit#gid=315186780).

For the likelihood that a ceasefire will start:  
 

```
// Likelihood that a ceasefire will start
numSuccesses = 0 // no ceasefire so far
numFailures = 138 // days since the 24th of February 
numFutureTrials = 172 // days until end of year

laplaceSuccessByNthTrial(successes, failures, numFutureTrials) = {
  numTrialsAlready = successes + failures
  numTrialsAtTheEnd = numTrialsAlready + numFutureTrials
  pAllFailures = (failures + 1) / ( numTrialsAtTheEnd + 1)
  // e.g., 10 trials, 10 failures, 0 successes
  // then probability of one additional failure is:
  // 11/12
  // and the probability of three additional failures is:
  // 11/12 * 12/13 * 13/14 = 11 / 14 =  (failures + 1) / ( numTrialsAtTheEnd + 1)
  // e.g.: 3/4 * 4/5 * 5/6 * 6/7 * 7/8 = 3/8
  1 - pAllFailures
}
chance_of_ceasefire = laplaceSuccessByNthTrial(numSuccesses, numFailures, numFutureTrials)
```

And for the likelihood that a ceasefire, once started, lasts, we are going to do some textbook outside view aggregation with the “mixture” function.

```
// Likelihood that a ceasefire lasts

// per laplace
laplace(successes, failures) = (successes + 1)/( successes + failures + 2)
chance_ceasefire_lasts_laplace = laplace(0,0)

// per dataset of ceasefires 
// (<https://docs.google.com/spreadsheets/d/1ufNP_iuTghzKGUCRPMoLY4-c9PibiNlVFEaAxyAad-A/edit#gid=315186780>)
chance_ceasefire_lasts_historical = 0.3130434783

// per dataset of ceasefires, restricted to Soviet conflicts (same link)
chance_ceasefire_lasts_historical_russian = 0.4

// per brief literature review (same link)
chance_ceasefire_lasts_lit_review = 0.6392156863

// overall guess
distribution_ceasefire_lasts = mixture(
  chance_ceasefire_lasts_laplace, 
  chance_ceasefire_lasts_historical,
  chance_ceasefire_lasts_historical_russian, 
  chance_ceasefire_lasts_lit_review, 
  [0.25, 0.3, 0.3, 0.15]) 
// instead of mixture(...), we can also use mx(...), which is the same but shorter.
chance_ceasefire_lasts = mean(distribution_ceasefire_lasts)
```

[Putting this all together](https://www.squiggle-language.com/playground/#code=eNqdVe9v0zAQ%2FVdORUgpZHXd9BcRQyqw8QGYJjqQEEWVmziNhWsX29lWEP875zRNu9JuE19c63Lvvbvzs%2Fu7YXN9My4WC2ZWjThj0vKwjJ2lwmnTiJ0pMCKUcILJ8c9CzOeSj50Rat6IG4TAB%2FGDS5FrnYLLmQMGCWeWZ8JwuBFSgnXMuIlSxWJcJAm3lls4hTYgVumdZKshY6ZMPGdCFqbMo9HQZ6ZsZcEKlXBU4dDpuhx0Bud8ZgosHdawwiHqymClJXTQqaGFckICV6lHrbjXmSjJlpIlvCrr9erC5SU4sJtCQyxpXUoIewJNVPg9UeDj68hIGs7SFcZrPDyvCe6muqucn2E1p%2F%2FCn%2B8reeRyJOXOVIINK2bTJhAIDnD7Tx6LM%2BCteSsE2gZXppTbbWftbcFVPs5YwdLoGZsJKdzKT00rDixFWwitmNzAQdi4AlFKaKfaM5WWB7XH4XLDD7HYfRp4BrRDaOR%2FI0K7%2Fjgp9lnu%2Fqv9GCLSRbou6eHaI31c%2B2SA64AMkTQiQ59N4eTOsCfqz0QlOUPnTXU23dr1FI7aZ9fqa98cNZG34QOXSDLrbJW35GYjW9v3kFu9OYNdF9ImCQ7aEredZt1hrTotVaeVxLbZoB22mzvFpMwhwvmzrbEWys%2FBy9y5pY0JSXViW3Ot8eVoJXpB7NI73eacO0tSQovs4nIqiqt5%2Fuv9u89vPl1%2B1B%2B%2Bdk%2BSF5diJi7kl%2FMzNrpdjVh6MiIcrfNkLtLTiPbosD8Ytl8dLz8XFh8xkaDR8MVpRTRqd6PuYBg91EEIuOBdSRzHE9Ew1tcCcxKtMolBi8NlCzwaoX48Rn1qCmsFU2UV3R3xmRE8QxrHDfOuQFkUunkUPaKmVbqn7UcvOrTXH%2FY3zelrJMX3d17gmU9UKnxDs8JfvH0uJFiIW19AcL8RQniw10ekbMZxNHXbWgjf2q1OL%2FSnVy%2B09%2F3oWHwrnKngvnY3Bv7EbSFdzVRdj6yQd675ocv%2F7EjlJa0paWPwhfeRoNNr9Z9Wkm%2BFxUGuJqr86zhAHR8KhjvJe4rxkfgu5FBj8T3fyjev8ecvmSzUuQ%3D%3D), we get:

```
// Result
chance_successful_ceasefire = chance_of_ceasefire * chance_ceasefire_lasts
// result: 0.256 = 25.6%

// Display
{
  chance_of_ceasefire: chance_of_ceasefire,
  chance_ceasefire_lasts: chance_ceasefire_lasts,
  chance_successful_ceasefire:chance_successful_ceasefire
}
```

<p><img src="https://i.imgur.com/iUmJSPk.png" class="img-medium-center"></p>

## Conclusion

We have looked at some simple models and produced some estimates of things we might care about. But this amount of complexity hasn’t exhausted the number of things we can estimate. 

In particular, we haven’t yet looked at:

*   Estimates of variables over time
*   Functions which produce estimates
*   Using Squiggle as a programmer rather than as a normal user.

An upcoming blog post will look at these complications. Readers might also want to browse the [Squiggle EA Forum tag](https://forum.effectivealtruism.org/topics/squiggle), which has begun to fill up with some neat estimates.

## Acknowledgements

<p><img src="https://i.imgur.com/7yuRrge.png" class="img-frontpage-center"></p>

This post is a project by the [Quantified Uncertainty Research Institute](https://quantifieduncertainty.org/) (QURI).
