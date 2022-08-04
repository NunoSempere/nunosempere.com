Some thoughts on Turing.jl
==========================

[Turing](https://turing.ml/stable/) is a cool probabilistic programming new language written on top of [Julia](https://julialang.org/). Mostly I just wanted to play around with a different probabilistic programming language, and discard the low-probability hypothesis that things that I am currently doing in [Squiggle](https://www.squiggle-language.com/) could be better implemented in it.

My thoughts after downloading it and playing with it a tiny bit are as follows:

1\. **Installation is annoying**: The program is pretty heavy, and it requires several steps (you have to install Julia and then Turing as a package, which is annoying (e.g., I had to figure out where in the filesystem to put the Julia binaries).

- Node.js installations can also be pretty gnarly (though there is [nvm](https://github.com/nvm-sh/nvm)), but Turing doesn't have an equivalent online playground. My sense is that running Julia online would also be pretty annoying (?).

2\. **Compilation and running the thing is _slow_**; 9 seconds until I get an error (I hadn't installed a necessary package), and then 1 min 26 seconds to run their [_simplest example_](https://turing.ml/dev/docs/using-turing/get-started) (!!)

```
using Turing
using StatsPlots

# Define a simple Normal model with unknown mean and variance.
@model function gdemo(x, y)
  s² ~ InverseGamma(2, 3)
  m ~ Normal(0, sqrt(s²))
  x ~ Normal(m, sqrt(s²))
  y ~ Normal(m, sqrt(s²))
end

#  Run sampler, collect results
chn = sample(gdemo(1.5, 2), HMC(0.1, 5), 1000)

# Summarise results
describe(chn)

# Plot and save results
p = plot(chn)
savefig("gdemo-plot.png")
```

This seems like this is a problem with [Julia more generally](<https://www.reddit.com/r/Julia/comments/lmznx7/running_scripts_is_terribly_slow_im_new_to_julia/>). Btw, the [Julia webpage](https://julialang.org/) mentions that Julia "feels like a scripting language", which seems like a bold-faced lie.

A similar but not equivalent [^1] model in Squiggle would run in seconds, and allow for the fast iteration that I know and love:

```
s = (0.1 to 1)^(1/2) // squiggle doesn't have the inverse gamma function yet
m = normal(0, s)
x = normal(m, s)
y = normal(m, s)
```

3\. Turing is able to do **Bayesian inference** over parameters, which seems cool & intend to [learn more about](https://github.com/rmcelreath/stat_rethinking_2022).

It's probably kind of weird that Squiggle, as a programming language that manipulates distributions, doesn't allow for Bayesian inference.

4\. Turing seems **pretty integrated with Julia**, and the documentation seems to assume familiarity with Julia. This can have pros and cons, but made it difficult to just grasp what they are doing.

- The pros are that it can use all the Julia libraries, and this looks like it is _very_ powerful
- The cons are that it requires familiarity with Julia.

It's possible that there could be some workflows with Squiggle where we go back and forth between Squiggle and javascript in node; Turing seems like it has that kind of integration down-pat.

5\. Turing seems like it could drive some **hardcore setups**. E.g., [here](https://withdata.io/election/media/) is a project using it to generate election forecasts.

---

Overall, I dislike the slowness and, as an outsider, the integration with Julia, but I respect the effort. It's possible but not particularly likely that we may want to first script models in Squiggle and then translate them to a more powerful languages like Turing when speed is not a concern and we need capabilities not natively present in Squiggle (like Baysian inference).

---

See also:

- [Why We Use Julia, 10 Years Later](https://julialang.org/blog/2022/02/10years/)
- [What's bad about Julia?](https://viralinstruction.com/posts/badjulia/#whats_bad_about_julia)
- [Julia in HackerNews](https://hn.algolia.com/?dateRange=all&page=0&prefix=false&query=Julia%20&sort=byPopularity&type=story)


[^1]: Even if Squiggle had the inverse gamma function, it's not clear to me that the two programs are doing the same thing, because Turing could be doing something trickier even in that simple example (?). E.g., Squiggle is drawing samples whereas Turing is (?) representing the space of distributions with those pararmeters. This is something I didn't understand from the documentation.
