Simple electoral college simulation
===================================

[Here](https://git.nunosempere.com/NunoSempere/2024-election-modelling) is a simple model of the US electoral college. It aims to be conceptually simple and replicatable. It incorporates data from state-specific polls, and otherwise defaults to the state's electoral history baserate.

Other projects, like [538](https://en.wikipedia.org/wiki/FiveThirtyEight), [Nate Silver's substack](https://www.natesilver.net/) or [Gelman's model](https://github.com/TheEconomist/us-potus-model) are to this project as a sportscar is to a walking stick. They are much more sophisticated, and probably more accurate. However, they are also more difficult to understand and to maintain. 

Compare with: [Nuño's simple node version manager](https://github.com/NunoSempere/nsnvm), [squiggle.c](https://git.nunosempere.com/personal/squiggle.c), [Predict, Resolve & Tally](https://github.com/NunoSempere/PredictResolveTally)

## What stories does the model tell?

### The naïve baserate story

Consider Ohio. Bush won the state in 2000 and 2004, Obama in 2008 and 2012, and Trump again in 2016 and 2020. The base rate, the historical frequency for republicans in Ohio is therefore 4/6. 

A straightforward way of getting at a probability of an electoral college win is to just take the historical frequency for each state, and sample from it many times, and then build up the different electoral college results from those samples. 

If we do so, however, Republicans end up with only a 25% chance of winning the 2024 election.

Why is this? Well, consider the number of electoral college votes in the last few elections: 

| Year | Republican electoral college votes | Democrat electoral college votes | 
| ---- | --- | --- | 
| 2000 | 271 | 266 | 
| 2004 | 286 | 251 | 
| 2008 | 173 | 365 | 
| 2012 | 206 | 332 | 
| 2016 | 304 | 227 |
| 2020 | 232 | 232 | 

Essentially, Obama won by much more than Bush, Trump or Biden. But our naïve model doesn't see that those results were correlated. 

So the story here is that our model is not very sophisticated. But another might be that Obama was much more popular than Biden, and if Democrats can tap into that again, they will do better.

Still, *for states in which there is no polling*, the electoral history seems like a decent enough proxy: these are the states which are solid Republican or solid Democrat.

### The unadjusted polls story

If we only look at polls (and use baserates when there are no polls—which happens for states like Alabama, which lean strongly towards one party already), this time the Republicans win by a mile: with 95% probability. 

What's happening here is that:

- There aren't that many polls yet
- For the polls that do exist, Trump polling very well in Pennsylvania, Wisconsin, Arizona, Michigan, Florida, Nevada, Georgia, North Carolina
  - Trump is also polling decently in Minessota; Biden is polling well in Colorado
- In part, this is because Biden is just [unpopular](https://projects.fivethirtyeight.com/biden-approval-rating/), or at least more than [Trump](https://projects.fivethirtyeight.com/polls/favorability/donald-trump/)
- In part though, polls currently also ask about the third party vote: for Robert F. Kennedy, Cornel West and Jill Stein (Green party).
  - In a normal democracy, like in Spain, a protest party could amass some electors, and use them as bargaining chips to govern together with one of the other major parties. For instance, this is what happened with Ciudadanos in Spain. Perhaps third parties performing strongly could conceivably, create pressure to reform the US electoral system.
  - In the US, with the system as currently exists, these votes seem to favour Trump.

However, this 95% really doesn't feel right. It is only accounting, and very naively, for the sample size of the poll. It not only assumes that the poll is a representative sample, it also assumes that opinions will not drift between now and election time. This later assumption is fatal.

### The adjusted polls story 

If we look at how [Gallup presidential election polls](https://news.gallup.com/poll/110548/gallup-presidential-election-trial-heat-trends.aspx) did between 1936 and 2008, we get a sense that polls in mid April just aren't very informative as to the eventual result. Doing the tally, for republicans, polls have a standard error of 4-5 points: huge when races in battleground states tend to be close to 50/50 (49/51, 48/52, 47/53, etc.)

Moreover, these are national polls: polls in battleground states will have smaller samples and thus more uncertainty. And current pollsters are nor as good as gallup. And... there might be other sources of uncertainty that I'm missing. On the other hand, we have increased polarization, not all states are battleground states, and this variable seems like it requires a bit of finesse.

But incorporating reasonable estimates of uncertainty, the probability of a republican win the model gives is 50-60%. This does depend on how much uncertainty you inject. If you inject a lot of uncertainty, it moves closer to 50%. But on the other hand, one has to take care to not inject *too* much uncertainty, even for sure states, like, say, Alabama. This is now in line with [prediction markets](https://electionbettingodds.com/PresidentialParty2024.html).

### The "conviction didn't matter that much" story

Trump's recent conviction reminded me that I had this model lying around, so I reran it and wrote this blogpost. Last time I'd run it was on April 15th, at which time this model gave Trump a 59% chance of winning. Today on June 2nd, the model gives him a 55% chance instead. You can see more granular results by number of electoral college votes [here](https://git.nunosempere.com/NunoSempere/2024-election-modelling/src/branch/master/out/output-2024-06-02.txt).

There is still time for polls to catch up—I'm including polls from the last 30 days, and not that many state-specific polls have been published after the conviction. But so far, I think the conviction just isn't having that large an effect. You can sort of see this visually [here](https://projects.fivethirtyeight.com/polls/favorability/donald-trump/) and [here](https://projects.fivethirtyeight.com/biden-approval-rating/) as well. 

<p>
  <section id='isso-thread'>
  <noscript>javascript needs to be activated to view comments.</noscript>
  </section>
</p>
