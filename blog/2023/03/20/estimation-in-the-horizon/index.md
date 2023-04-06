Some estimation work in the horizon
===================================

This post outlines some work in altruistic estimation that seems currently doable. Some of it might be pursued by, for example, my team at the Quantified Uncertainty Research Institute. But together this work adds up to more than what our small team can achieve.

Two downsides of this post are that a) it looks at things that are more salient to me, and doesn't comprehensively review all estimation work being done, and b) it could use more examples.

<img src="https://i.imgur.com/ybVJNqo.png" alt="Saruman in Isengard looking at an army of orcs" class="img-medium-center" />
*<br>Saruman in Isengard looking at an army of orcs*

## Directions 

### Produce more specific high-value estimates

The idea here is to produce more estimates in a way that results in better decisions.

Some examples might be:

- Produce estimates of the value of 80,000 hours' top career paths
- Produce estimates of the value of EA organizations
- Produce estimates of the value of past EA Funds and Open Philanthropy grants.
- Produce estimates of the value of different research directions
- Produce estimates of the value of different new cause areas

The value of these estimates generally has two moving parts:

1. Their immediate payoff in terms of affecting near-term decisions
2. Making it easier to create future estimates of a similar kind 

For example, if I help someone estimate the value of various career options they are considering, that has the benefit of improving their decision about which career moves they make. But if those estimates are then posted on the EA Forum, or in some repository of models, they might make it easier for people to reuse and tweak these estimates, which would be cheaper than coming up with their own estimates anew.

At QURI in particular, we might work on estimating the value of some of 80,000 hours' top career paths, building upon [the results from this contest](https://forum.effectivealtruism.org/posts/hGdsgaRiF2zH3vX5M/winners-of-the-squiggle-experimentation-and-80-000-hours), because this seems like a somewhat ambitious type of estimation that would produce useful templates and allow us to explore new strategies. And I personally expect to produce a few such estimates this year.

### Explore different types of estimates

At QURI, we have been experimenting mainly with two broad classes of estimation:

- End-to-end estimation, where the final estimate is a quantity like QALYs, existential risk reduced, that we care more or less directly about.
- Relative value estimation, where we don't estimate and endline quality, but rather how much you value some option relative to some other related option.

As we do more work there, we learn and begin to implement some stuff, like:

- End-to-end estimation is generally too effortful, and you can reduce that effort by making estimates more modular and reusing their parts
- Relative value estimation is still too unwieldy, so you can do relative value estimates within a cluster, and then elicit the relative estimate between different clusters, or between representatives of each clusters. 

But in general, are there any ingenious types of estimates which would make estimation easier, more scalable, or doable for previously inaccessible topics? 

### Build estimation pipelines

In [August 2020](https://forum.effectivealtruism.org/posts/Ze2Je5GCLBDj3nDzK/how-many-ea-billionaires-five-years-from-now), @Erich_Grunewald estimated that there was a 5-10% chance that EA would lose a billionaire in the next few years. In hindsight, this estimate seems decent. But it wasn't vetted, distributed, and incorporated into people's decisions. And although it was part of some QURI contests, it seems more like a one-off than a continuous series of estimations of importance. 

Instead, we could have pipelines that continuously produce estimates in a way which produces value. Some pipelines which it might be valuable to build might be:

- Evaluations of the results of projects which receive funds from EA organizations. This would require figuring out what kinds of evaluations of past funding decisions would influence future funding decisions. It would also be neat to add a consent form to funding applications which would allow for more inquisitive public evaluations in the future. I think there would be value in this being a pipeline rather than a one-off evaluation of some grants, because it would allow for more lessons learnt to accumulate.
- Adapting GiveWell's estimates to incorporate uncertainty. Although there is already some early work on this, it hasn't been incorporated into GiveWell's own estimates. It could also be updated year on year.
- Pipelines to produce forecasts. Right now, there are various groups producing forecasts, from forecasting platforms like Metaculus and Manifold Markets, to prediction markets like Polymarket, to groups of forecasters like those at Samotsvety and the Swift Center. But these don't cover all that could be forecasted, and could produce more forecasts, sooner.
- One very general pipeline would involve more widespread use of prediction markets. This has two challenges. One is that it's currently very legally effortful to do things legally in the US. Another challenge is that it's not clear how to structure prediction markets so that they produce value. That is, there the clear between putting a market up in Polymarket and good things happening could be made much stronger, but this requires figuring out.

Here, by pipeline, I mean something distinct from a one-off, and that starts to produce value and can continue to do so. My sense is that those types of work are more valuable.

### Work on the theory behind estimation

Some topics which could benefit from a theoretical treatment might be:

- How can we address Goodhart's law and the optimizer's curse? You can make a Bayesian adjustment, but how would it look like in practice?
- What are the different types of estimations, and how best can we tackle them? What is the best way to divide and organize the space?
- What are some good proper scoring rules that one can use in forecasting tournaments? This is kind of an interesting intellectual topic, but it can also be a huge time-waster.
- Although Superforecasters do something similar to Bayesian updating, they can't be literally doing Bayesian updating (the world is too big, and their heads too small to represent all possible hypotheses). How then is the best way to approximate Bayesian updates given finite resources? 
- What can we learn by looking at the history of forecasting?

The hope here would be that taking a step back and theorizing could add some clarity which could improve your day to day estimation.

### Produce better estimation tooling 

Better tooling seems like it could lead to the production of more and better estimates, faster. For an extreme example, [here](https://github.com/NunoSempere/time-to-botec/blob/master/C/samples/samples.c) is a simple BOTEC in C (a low level estimation language), and [here](https://www.squiggle-language.com/playground#code=eNptUEFuwjAQ%2FMrKp6QykFCQUKSeSl%2BQY42QIRuIhG1kO6II8ffuxuaGDzvj8ezOyg8Rzu7WjsZofxdN9CPKSfrphuj8S%2Bmw1%2BMlfrsORSOUve41fEE13zA9THTN9EiU3z6oHpRVthtC3FdsyLwmXmcenEG6ttpcL9hinPfemS29FDVEB59l9hlt7%2B99S%2FbVVfmKCmT7TZkSUl5GzsqUx%2B2UveFwOqcOypvRxsdScl0sE6wywI7Hewz0A%2BQ2w18cPRZTnoQ8hlZ4gLJAJzmbjDKJBrVtploknRqe4vkPXHOBOQ%3D%3D) is that same estimate in Squiggle, a language for rapid estimation.

Some progress in the tooling domain might involve:

- Better dashboards, and ways to display estimates
- Adding more useful functionality, like a [metalog distribution](https://github.com/quantified-uncertainty/squiggle/pull/1444), which allows for easier construction of distributions based on arbitrary confidence intervals.
- Making tooling which is usable by more people, that could be used on mobile, or that has a more gentle learning curve

One consideration for this kind of enterprise is that it can be fairly time intensive, and fairly long until you learn whether it has been worth it or not.

## Who is working on this type of thing? 

Within or adjacent to EA:

- The Quantified Uncertainty Research Institute
- Rethink Priorities
- The Swift Center
- Samotsvety Forecasting
- The Forecasting Research Institute
- Various funders and funds
- Manifold Markets
- Metaculus
- etc.

Outside EA, probably too many to mention at once, but some which stand out to me are:

- Startups working in the prediction market space: Polymarket, Kalshi, etc.
- Startups or early companies working on the estimation space: Causal, Rootclaim, FiveThirtyEight, etc.
- Large academic fields around statistics, etc.
- Most corporations have some revenue and demand forecasting, or other [statistical software](https://en.wikipedia.org/wiki/SAS_%28software%29).
- etc.

<p>
  <section id='isso-thread'>
  <noscript>Javascript needs to be activated to view comments.</noscript>
  </section>
</p>


