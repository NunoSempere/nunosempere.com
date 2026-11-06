Introduction to Fermi estimates
===============================

*The following are my notes from an intro to [Fermi estimates](https://en.wikipedia.org/wiki/Fermi_problem) class I gave at [ESPR](https://espr-camp.org/), in preparation for a Fermithon, i.e., a Fermi estimates tournament.*

Fermi estimation is a method for arriving an estimate of an uncertain variable of interest. Given a variable of interest, sometimes you can decompose it into steps, and multiplying those steps together gives you a more accurate estimate than estimating the thing you want to know directly. I'll go through a proof sketch for this at the end of the post.<p><img src="https://i.imgur.com/5uk7LBH.png"  class="img-medium-center"></p>

If you want to take over the world, why should you care about this? Well, you may care about this if you hope that having better models of the world would lead you to make better decisions, and to better achieve your goals. And Fermi estimates are one way of training or showing off the skill of building models of the world. They have fast feedback loops, because you can in many cases then check the answer on the internet afterwards. But they are probably most useful in cases where you can't.

The rest of the class was a trial by fire: I presented some questions, students gave their own estimates, and then briefly discussed them. In case you want to give it a try before seeing the answers, the questions I considered were:

1. How many people have covid in the UK right now (2022-08-20)?
2. How many cumulative person years did people live in/under the Soviet Union?
3. How many intelligent species does the universe hold outside of Earth?
4. Are any staff members dating?
5. How many "state-based conflicts" are going on right now? ("state based conflict" = at least one party is a state, at least 25 deaths a year, massacres and genocides not included)
6. How much does ESPR (a summer camp) cost?
7. How many people are members of the Chinese communist party?
8. What is the US defense budget?
9. How many daily viewers does Tucker Carlson have?

## 1. How many people have covid in the UK right now (2022-08-20)?

### My own [answer](https://www.squiggle-language.com/playground#code=eNp1jrsKwjAUhl8lZBTpzRsUHJ2kdChuhRDatD2Y5sSYtIj47qYFF7XLGf7L958nvXc4Fq7vuXnQ1Bon1rN0qsGi%2BSigwAKXxc1B20pRWAOqpSnVBjUaC6hyZ%2FMmjiJyJFEQE4tkV6rG8Goy2dgh6%2FggWIUD1D7zpxj6UyqN2kk%2Bl7Bhl7PP7g9BkmTeEqil%2BGUtrazIN2yJEYZ%2B5Do9vQm2GX29AQ8jZBI%3D)

```
proportionOutOf100 = 0.1 to 5
fraction_who_have_covid = proportionOutOf100 /100
population_of_UK = 67.22M
people_who_have_covid = fraction_who_have_covid * population_of_UK
people_who_have_covid // 67k to 3.4M
```

![](https://i.imgur.com/SNNpfge.png)

### Some student guesses

- 130k
- 600k
- 1M to 2M

### [Check](https://www.squiggle-language.com/playground#code=eNp9j0FPhEAMhf%2FKhJOaDQwYjdnEo6c9ciWZTJgCTaDFobgS43%2B3oBs163prXtvvvfeWTB0fy3kYfFySvcQZdpv0FFA4nhQkFPR9%2BTxj2%2FZQSkRqk31SMzUYBwgugJduciNEF%2FxiHk1R0Rh55CjI5Lg5XURYNQh6YtPCCJu8ImHx%2FdeFLi5iM%2FMvs6LaT%2BAarzSUxUUvoDjhK5vaPMut3anpNlybTGH%2BtecXDKbuPLWaSTqc0opwGHtUd4KjW4k%2Fe%2F2KmplzQ32nBuotYphV0UH%2F7taqhf2Gf4I2vK4vW96Yc97fFG10aw%2Brz%2F2DPWiP5P0Dbqavjg%3D%3D)

To check, we can use the number of confirmed deaths:

<iframe src="https://ourworldindata.org/explorers/coronavirus-data-explorer?zoomToSelection=true&time=2020-03-01..latest&facet=none&pickerSort=asc&pickerMetric=location&hideControls=true&Metric=Confirmed+deaths&Interval=7-day+rolling+average&Relative+to+Population=true&Color+by+test+positivity=false&country=~GBR" loading="lazy" style="width: 100%; height: 600px; border: 0px none;"></iframe>

```
confirmed_deaths_per_day = 2
proportion_of_deaths_reported = 0.2 t 1
total_deaths = confirmed_deaths_per_day / proportion_of_deaths_reported
case_fatality_rate = to(0.01/100, 0.1/100) // paxlovid changed this.
implied_new_cases_per_day = total_deaths / case_fatality_rate
infection_duration = 5 to 20
implied_total_cases = implied_new_cases_per_day * infection_duration
implied_total_cases // 30K to 680K.
```

![](https://i.imgur.com/YK7vVKx.png)

A [reader](https://www.reddit.com/r/TheMotte/comments/wy58g2/comment/ilw2vnn/?context=3) points out that the UK's Office for National Statistics [carries out randomized surveys](https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/conditionsanddiseases/bulletins/coronaviruscovid19infectionsurveypilot/26august2022) ([a](https://web.archive.org/web/20220826225309/https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/conditionsanddiseases/bulletins/coronaviruscovid19infectionsurveypilot/26august2022)). For the week of August 16th, they reported an estimated 1.3M cases across the UK, which falls within our initial estimate's interval, but not within our check's (?!).

## 2. How many cumulative person years did people live in/under the Soviet Union?

### My own [answer](https://www.squiggle-language.com/playground#code=eNp9jb0KwjAUhV%2FlklGUJlWRFlwKjp2yFkLQEC%2B0SU1vKkV8d2PRQUGnw%2FnhfDc2nP1Vxq7TYWIlhWiWc3Q4IfnwTtAhoW7lJaK1rZEU0FlWMj1a1fs%2BtprQOzX4EQ2p6JKBPWx5DeRhzXnduFc3GR2G1ImiELBKInaNOyb%2B%2FDEa9bX7R4AFfKyzDDbVk5jn1e9Tdn8ACl1ZaA%3D%3D)

```
avg_population_soviet_union = 50M to 300M
soviet_years = 1991 - 1917
cumulative_soviet_years = avg_population_soviet_union  * soviet_years 
cumulative_soviet_years // 4B to 22B
```

![](https://i.imgur.com/4Z5Argk.png)

Students guessed pretty much the same.

### Check

Per [this graph](https://en.wikipedia.org/wiki/File:Demography_USSR.svg):

<p><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/d/db/Demography_USSR.svg/1200px-Demography_USSR.svg.png"  class="img-medium-center"></p>

the average population seems to have been around 200M, implying 14.8 billion years.

## 3. How many intelligent species does the galaxy hold?

### My own answer

Probably just one (!?).

### Student guesses

- 10^4
- 0.4T
- 0 to 1

### Check

The [dissolving the Fermi paradox](https://arxiv.org/abs/1806.02404) paper gives the following estimate:

<p><img src="./.images/dissolving-fermi.png"  class="img-medium-center"></p>

which gives something like a 40% chance of us being alone in the observable universe. I think the paper shows the importance of using distributions, rather than point-estimates: using point estimates results in the [Fermi paradox](https://en.wikipedia.org/wiki/Fermi_paradox). This is the reason why I've been multiplying distributions, rather than point estimates

## 4. Are any staff members dating?

Note: Question does not include Junior Counselors, because I don't know the answer to that.

### Own answer

The camp has 11 male and 3 women staff members. So the number of heterosexual combinations is (11 choose 1) × (3 choose 1) = 11 × 3 = 33 possible pairings. However, some of the pairings are not compatible, because they repeat the same person, so the number of monogamous heterosexual pairings is lower.

Instead, say I'm giving a 1% to 3% a priori probability for any man-woman pairing. How did I arrive at this? Essentially, 0.1% feels too low and 5% too high. That implies a cumulative ~10% to 30% probability that a given woman staff member is dating any man staff member. Note that this rounds off nontraditional pairings—there is 1 nonbinary Junior Counsellor, but none amongst instructors, that I recall.

If we run with the 30% chance for each woman staff member:

```
p(0) = 1 × 0.7 × 0.7 × 0.7 = 34.3%
p(1) = 3 × 0.3 × 0.7 × 0.7 = 44.1%
p(2) = 3 × 0.3 × 0.3 × 0.7 = 18.9%
p(3) = 1 × 0.3 × 0.3 × 0.3 = 2.7%
```

If 10%:

```
p(0) = 1 × 0.9 × 0.9 × 0.9 = 72.9%
p(1) = 3 × 0.1 × 0.9 × 0.9 = 24.3%
p(2) = 3 × 0.1 × 0.1 × 0.9 = 2.7%
p(3) = 1 × 0.1 × 0.1 × 0.1 = 0.1%
```

So the probability that there is at least one pairing ranges from 65% to 35%, which is almost symmetrical. So this implies a roughly 50% that there is at least one pairing.

### Student estimates

- 25% no couples, 1: 50%, ...
- 0:25, 1:25 , 2: 33%, ...
- 0: 0, 1: 1/6, 2: 2/6, >=3: 3/6
- 0: 1/4, 1: 1/4, ...

### Check

You don't get to know the answer to this!

## 5. How many "state-based" conflicts are going on right now?

- At least one party is a state
- At least 25 deaths a year.
- One sided violence, like massacres, are not included

### My own [answer](https://www.squiggle-language.com/playground#code=eNqFjssKwjAQRX9lyEpLqdPagggu%2FYIuVUqtsQaSSU0TtIj%2FbrQ%2BsAvdzQz3zLkX1h70KXdKlaZjc2scDx%2Bn5U5YbV4XQcKKUuZHJ%2Bpa8twaQTWbM3KqqLQjv%2FMWFqDOowQRrIYpYgjPOUUcr6kxutHGCk0fpHDU2nIruWcxwuyexmi6pv4x7aWobFs03LyTT7jrbbGXPBQhrDCahZ5ONuMB76PfRQP4XSaA%2F%2F6hYzKBOL1XiTNk1xvgWnoq)

```
num_countries = mx(200 to 300, 200 to 400)
proportion_countries_unstable = 0.05 to 0.3
num_conflicts_per_unstable_country = mx(1, 2 to 4, [0.8, 0.2])
num_conflicts = num_countries * proportion_countries_unstable * num_conflicts_per_unstable_country 
num_conflicts // 14 to 150
```

![](https://i.imgur.com/i5onQt7.png)

### Check

<iframe src="https://ourworldindata.org/grapher/the-number-of-active-state-based-conflicts?country=~OWID_WRL" loading="lazy" style="width: 100%; height: 600px; border: 0px none;"></iframe>

## 6. How much does ESPR cost?

### My own [answer](https://www.squiggle-language.com/playground#code=eNp1jcEKwjAMhl%2Bl9KRDWCfKZLCjpx17LYziyix0yWxTZIjvbjsY6MFD%2BMMXvj8vHu74lHGatF94Qz6aw4qugyX0G7FgyWonH9GOozOSvIWRNxzi1M8GZ2dYy061ghsGSsTnCQjrOuglXSshGGEOoSB7CYfMzwoISbs%2Buwl8dRbsX1%2FBtoofuyxZLbr8p65S7o4X0e0V8PcHUqtRhA%3D%3D)

```
num_people = 47
cost_per_person_per_day = 100 to 1000
num_days = 15
total_cost = num_people * cost_per_person_per_day * num_days
total_cost // 70K to 710K (280K)
```

![](https://i.imgur.com/5PJq1ch.png)

### Student estimates

Student estimates ranged from ~$50k to $150k

### Check

Per [this source](https://www.openphilanthropy.org/grants/european-summer-program-on-rationality-general-support/), ESPR cost $500k over two years. The money was used for a variety of events, including an online version in 2020 and a range of smaller events. If we count four events, the price tag per event is around $125k; if we only count two summer camps, it's up to $250k, but probably somewhere around $170k. Some of the students forgot about the cost of staff members, food, or common rooms.

## 7. How many people are members of the Chinese communist party

### My own [answer](https://www.squiggle-language.com/playground#code=eNqNjkELgkAUhP%2FKsseI3BAvQpegoyevwmK26QP3PX37lpDov6eCUJfoNsx8zMxTh44eZfS%2B5knnwtHtV%2BtyAyHeHEAQqPtyjNC2vSuFAVuda4zeNh2gC06d1PFc4cA0EAsQWrpvmQW0DXkfEYLYoWaZZtwcjMmU0CLSCteub8h656%2BOwwx%2FLu3Uvyu%2FW5NEZcVyIDVFhfr1Bj3QZNo%3D)

```
num_chinese = 1B
proportion_of_chinese_in_communist_party = 0.005 to 0.03
num_communist_party_members = num_chinese * proportion_of_chinese_in_communist_party
num_communist_party_members // 5M to 30M
```

![](https://i.imgur.com/qLRFz0q.png)

### Check

[Wikipedia](https://en.wikipedia.org/wiki/Chinese_Communist_Party#:~:text=As%20of%202022%2C%20the%20CCP,India's%20Bharatiya%20Janata%20Party.) reveals that there are 96M ~ 100M members, which around 10% of their population. Surprising.

## What is the US defense budget?

### My own [answer](https://www.squiggle-language.com/playground#code=eNp9jk%2BLwjAQxb%2FKkFOVYlNrl0XYi%2BDRUz0KodI0BtpMTBN2Rfa770Qq3T9lL3nkvXnzmzsbLvhehb6v3Y1tvQsyfVj7Rnt0T0cb7XXdVdeglepk5Z02im2ZaqzAVgy21gbeID%2BejEUbutprNN%2BTDT%2F8jsJAfsFjMK55OEm%2BKsEjFAtYwp%2FGEn4gM5jBEcihRTfVRCx5FI1spRkkUfqPhK94EUGkmzKNsh6%2F5eJkqDROi3NolPRUms6ky%2F5HzC3IMihyvouQ1xfSpKSXUOzzC9ZhiMw%3D)

```
gdp_of_spain = 1T
population_of_spain = 40M
population_of_us = 300M
gdp_of_us = (1.5 to 3) * population_of_us * gdp_of_spain / population_of_spain
proportion_of_us_gdp_to_defense = mx(0.03 to 0.045, 0.02 to 0.05)
us_defense_budget = gdp_of_us * proportion_of_us_gdp_to_defense
us_defense_budget // 310B to 860B (560B)
```

![](https://i.imgur.com/y9SwK2R.png)

### Check

$715B in 2022, per [Wikipedia](https://en.wikipedia.org/wiki/Military_budget_of_the_United_States#Budget_for_FY2022)

## How many daily viewers does Tucker Carlson have?

### My own [answer](https://www.squiggle-language.com/playground#code=eNp9jsEKgkAQhl9l2VNFpCYGCR07epJuwrLYqkO6a%2BNuJtG7t2lBWHYa%2BP9vZr4bbQrVxqaqOHY01GjEso%2F2R9AK3wlI0MDL%2BGwgz0sRawSZ05DWqjYl16AkUxk7xGRHfNeNEpkhT%2Fu45TotmBRtw1DkFsays5S7cgOilZ3%2BJ1yAhYeVTF17zh%2BwYALTJj0J7MnXwW0ih5BdQLQCmwJq23%2BpLhJJyH%2FPETKym26H%2F788HId4XkRm%2FmoTPXXXQTRPJL0%2FAIk%2BjW0%3D)

```
population_of_US = 300M
fraction_watch_news_regularly = 0.05 to 0.3
fraction_which_watch_fox = 0.3 to 0.5
fraction_which_watch_tucker = 0.5 to 0.9
tucker_viewership = population_of_US * 
  fraction_watch_news_regularly * 
  fraction_which_watch_fox * 
  fraction_which_watch_tucker
tucker_viewership // 11M (3.6M to 25M)
```

![](https://i.imgur.com/Ths8M1R.png)

### Check

Per [this source](<https://deadline.com/2022/08/fox-news-viewership-july-tucker-carlson-1235084361/>), Carlson gets about 3M viewers, so a bit outside my 90% range. I think I mis-estimated the fraction which watch news regularly, should probably have had a section for cable news specifically.

## Proof sketch that Fermi estimates produce better estimates than guessing directly

A [lognormal](https://en.wikipedia.org/wiki/Log-normal_distribution) distribution whose logarithm is a normal. It normally looks something like this:

<p><img src="https://i.imgur.com/AjNygK1.png"  class="img-medium-center"></p>

but if we display it with the axis in log scale—such that the distance between 30 and 40 is the same as the distance between 300 and 400, or between 3k and 4k, or such that the distance between 30 and 300 is the same as the distance between 300 and 3k—, it looks like a normal:

<p><img src="https://i.imgur.com/M7G5ZV2.png"  class="img-medium-center"></p>

I think that a lognormal is a nice way to capture and represent the type of uncertainty that we have when we have uncertainty over several orders of magnitude. One hint might be that, per the [central limit theorem](https://en.wikipedia.org/wiki/Central_limit_theorem), sums of independent variables will tend to normals, which means that products will tend to be lognormals. And in practice I've also found that lognormals tend to often capture the shape of my uncertainty.

So then here is our proof sketch that you can get better estimates by decomposing a problem into steps:

- Assume that guesses are independent lognormals, with a mean centered on the correct amount. This traditionally means that your guesses are lognormally distributed random variables, that is, that you are drawing your guess from a lognormal. But you could also think of giving your guess as a distribution in the shape of a lognormal, as I have done in the examples above.
- Observe that the multiplication of two lognormals is the sum of the logs of those variables, which are normally distributed. 
  - Per [well known math](https://en.wikipedia.org/wiki/Sum_of_normally_distributed_random_variables), the sum of two normals is a normal. And hence the product of two lognormals is a lognormal.
  - Look at the shape of _standard deviation / mean_ of the underlying normals.
  - E.g., consider _n_ lognormals with mean _m_ and standard deviation _s_ on their underlying normals.
  - Then that ratio is _s / m_
  - Then consider the product of _n_ such lognormals
  - Per [this](https://stats.stackexchange.com/questions/212690/the-product-of-two-lognormal-random-variables), the mean of the underlying normal of that product is going to be _n * m_, and the standard deviation of the underlying normal of that product is going to be _sqrt(n)*s_
  - Then the ratio is going to be _sqrt(n) * s / (n * m)_, which converges to 0 as _n_ grows higher.
  - We will get a similar phenomenon when we use lognormals that are not idential.
  - Then, if you have the ability to make estimates with a certain degree of relative error—i.e., such that the _standard deviation / mean_ of your estimates is roughly constant—then you are better off decomposing your estimate in chunks and estimating those chunks.

I then have to argue that the ratio of _standard deviation / mean_ is a meaningful quantity, because the proof wouldn't go through if we had used _mean / variance_. But the _standard deviation / mean_ is known as the [coefficient of variation](https://en.wikipedia.org/wiki/Coefficient_of_variation), and it doesn't depend on the unit of measurement used—e.g,. centimeters vs meters—unlike some equivalent measure based on the variance.

