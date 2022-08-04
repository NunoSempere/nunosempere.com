How much would I have to run to lose 20 kilograms?
===================================================

In short, from my estimates, I would have to run 70-ish to 280-ish 5km runs, which would take me between half a year and a bit over two years. But my gut feeling is telling me that it would take me twice as long, say, between a year and four.

![](https://i.imgur.com/WvP7i1M.png)

I came up with that estimate because was recently doing some exercise and I didn't like the machine's calorie loss calculations, so I rolled some calculations of my own, in [Squiggle](https://www.squiggle-language.com/playground)

```
// Calculations
meters_run = 5k // meters
weight_displaced = 120 // kilograms
gravity_constant = 9.81 // m/s^2
jules = meters_run * weight_displaced * gravity_constant
jules_to_calories(jules) = jules / 4184 // <https://en.wikipedia.org/wiki/Calorie>
calories = jules_to_calories(jules)

// Display
calories
```

So if I feed this calculation into [squiggle](https://www.squiggle-language.com/playground) this calculation tells me that I spent 1410 calories, as opposed to the 500-ish that I get from the machine, or similar quantities from [random internet sites](https://caloriesburnedhq.com/calories-burned-running/). The site linked even has a calculator, which is pretty wrong.

Anyways, I also want to add uncertainty, both because

- I don't really know my weight (and, to be honest, I don't really want to know)
- and because I'm not actually displacing my weight in a vaccum—e.g., I'm not pushing myself in a straight line but rather jumping a bit when I run.


```
// Calculations
meters_run = 5k // meters
weight_displaced = 118 to 125 // kilograms
gravity_constant = 9.81 // m/s^2
fudge_factor = mx(1 to 1.05, 1 to 1.1, [0.5, 0.5])
jules = meters_run * weight_displaced * gravity_constant * fudge_factor
jules_to_calories(jules) = jules / 4184 // <https://en.wikipedia.org/wiki/Calorie>
calories = jules_to_calories(jules)

// Display
calories
```

This produces the following calorie estimates:

![](https://i.imgur.com/InNQEN0.png)

But I don't just want to know the number of calories, I want to know how this relates to weight. [This random hackernews comment](https://news.ycombinator.com/item?id=10145690#:~:text=That%20Ignores%20Hunger-,1%20kg%20of%20fat%20is%20about%207700%20calories.,about%201%20kg%20of%20fat.) mentions that 1kg of fat is 7700 calories, and the number is likewise spread across the internet, but I somehow don't have that much confidence in it. After some searches on Google Scholar, I settle on "calories per kilogram of adipose tissue" as the key term I'm looking for, and I find [this paper](https://www.ncbi.nlm.nih.gov/pmc/articles/PMC4035446/). The paper mentions that after an initial weight loss period, the body becomes more lethargic, which does intutively make sense. So because your body adapts to lower caloric consumption (or to higher energy expenditure), these calculations don't really make sense directly.

Still, I think that they can establish a rough lower bound of how much exercise I would have to do. So, continuing with that model:

```

// Physical calculations
meters_run = 5k // meters
weight_displaced = 118 to 125 // kilograms
gravity_constant = 9.81 // m/s^2
fudge_factor = mx(1 to 1.05, 1 to 1.1, [0.5, 0.5])
jules_in_run = meters_run * weight_displaced * gravity_constant * fudge_factor
jules_to_calories(jules) = jules / 4184 // <https://en.wikipedia.org/wiki/Calorie>
calories_expended_in_run = jules_to_calories(jules_in_run)

// Fake-ish calculatio/i/s
calories_to_restrict_or_spend_for_one_kilogram_weight_loss = 5k to 20k // calories per kg
// lower because calculations could be off, or I could become buffed
// higher because of things like lower energy expenditure, increased apetite, etc.
weight_loss_target = 20 // kg
calories_to_kill = weight_loss_target * calories_to_restrict_or_spend_for_one_kilogram_weight_loss

// Algebra over fake numbers
runs_needed = calories_to_kill / calories_expended_in_run
runs_per_week = 2 to 3
years_needed = runs_needed / (runs_per_week * 52)

// Display

{
    runs_needed: runs_needed,
    years_needed: years_needed
}
```

This produces the following output:

![](https://i.imgur.com/WvP7i1M.png)

So I would have to run between 70-ish and 280-ish 5km run to loose 20kg, which would take me half a year to a bit over two years. And this is if the calorie conversion estimation method can be trusted. In practice, my gut feeling informed by the above is telling me that this is an underestimate, and that I would need somewhere between, say, one and four years to loose that weight, so I'm in for the long haul.

In conversations with people on [Twitter](https://twitter.com/NunoSempere/status/1552636309025562625) and elsewhere, they bring up the following points:

- Exercise improves apetite, and as such it's probably not a great way to loose weight
- Reducing calorie intake is "vastly easier" than increasing calorie use
- Even a little bit of strength training helps a lot because it increases your resting burn rate.
- "Abs are made in the kitchen"
- I get recommendations for [The Hacker's Diet](https://www.fourmilab.ch/hackdiet/) and [MacroFactor](https://macrofactorapp.com/). I install the later on my phone.
