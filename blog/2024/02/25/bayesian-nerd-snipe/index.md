A Bayesian Nerd-Snipe
======

Consider the number of people you know who share your birthday. This seems an unbiased estimate of the number of people who, if they had been born the same day of the year as you, you'd know—just multiply by 365. That estimate itself is an estimate of how many people one knows at a somewhat non-superficial level of familiarity.

I asked my Twitter followers that question, and this is what they answered: 

<blockquote class="twitter-tweet"><p lang="en" dir="ltr">How many people do you know that were born in the same day of the year as you?</p>&mdash; Nuño Sempere (@NunoSempere) <a href="https://twitter.com/NunoSempere/status/1760381265272885485?ref_src=twsrc%5Etfw">February 21, 2024</a></blockquote> <script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

Now, and here comes the nerd snipe: after seeing the results of that poll, what should my posterior estimate be for the distribution of how many people my pool of followers knows enough that they'd know their birthdays if they fell on the same day as one's own?

Here is a photo of two cats for those of my readers who don't want to be spoiled and want to sketch the solution before reading on.

![](https://gatitos.nunosempere.com/)

## The formal solution.

Consider distributions over how many people someone knows. Those distributions go from the natural numbers to a probability.

One example such distribution might be 

| Number of people known | Probability | 
| --- | --- |
| 10 | 5% |
| 21 | 94% |
| 1001 | 1% | 

Now, consider the likelihood of getting the Twitter poll results given a given distribution. Multiply that by the prior for that distribution, normalize, and then integrate over distributions to get your final result.

## The practical solution

The above is computationally intractable. For instance, for the number of cdfs, one could consider the number of monotonically increasing distributions from [0, 10k] to [0,1], and consider the cdf to be in increments of 0.0001. But iterating over those cdfs is basically uncomputable. So we turn to Monte Carlo approximations and other shortcuts. After tinkering for a bit, I ended considering distributions over logspace, and considering only number of people in the set: 16, 32, 64, 128, 128, 256, 512, 1024, 2046.

You can see the code [here](https://git.nunosempere.com/NunoSempere/peopleprobs/src/branch/master/probppl.go). It's written in go because I've been recently been learning its syntax, and it's only [moderately slower than C](https://github.com/NunoSempere/time-to-botec/) in exchange for a nicer developer experience.

With that code, the prior—integrated over the possible distributions—is:

| Num people known | % of people who know that many people| 
| --- | --- |
|   ~16 | 12.5%  |
|   ~32 | 12.5% |
|   ~64 | 12.5% |
|  ~128 | 12.5% |
|  ~256 | 12.5% |
|  ~512 | 12.5% |
| ~1024 | 12.5% |
| ~2048 | 12.5% |

and the posterior is:

| Num people known | % of people who know that many people |
| ---- | ----- |
|   ~16 |  7.2% |
|   ~32 |  8.6% |
|   ~64 | 11.0% |
|  ~128 | 15.7% |
|  ~256 | 21.9% |
|  ~512 | 22.6% |
| ~1024 |  9.5% |
| ~2048 |  3.4% |

<p>
  <section id='isso-thread'>
  <noscript>javascript needs to be activated to view comments.</noscript>
  </section>
</p>

