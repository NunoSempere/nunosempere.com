# Alternative Visions of Effective Altruism

## Intro

EA was born out of a very specific cultural milieu, and holds a very specific yet historically contingent shape. So what if instead EA had been...

## Effective Altruism Distributed

![](.src/diagrams/coordination-1.svg)

**Day-trading the crypto to liberate the animals:**

> You are a group of friends who at some point realized that there was money to be made in the crypto world. You've done that for the last couple of years. And now that the the boom part of the cycle is over, you have accumulated a nice stash. You use that to donate to somewhat radical or unconventional groups aiming to reduce animal suffering. You've also begun to worry about AI safety, so a few of your projects are about doing independent research in the space.

**Alternating consulting and cultivating the forecasting community:**

> You work on consulting as a generally talented jack-of-all-tradesn, and you keep track of what's up in the forecasting community, which you then write up in a pretty well-read newsletter. You think that a few of the most impactful moves you've made in the last few years might have been about giving honest encouragement to new people in the space trying out new projects—though you don't really know whether that was counterfactually valuable. Otherwise you subsidize a bit of your own independent research with the money you earn from consulting. Over time, you develop a community of people working on stuff that's similar to yours.

**Jane Street trader to make bets on global health:**

> After studying Math at Oxbridge, you got a job at Jane Street. You fund Charity Entrepreneurship to start a few charities a year at $50k/charity. You're pretty sure that's counterfactual. Your expected impact is insane, but you're not really certain about what that translates to in QALYs. Part of what keeps you going is having a group of like-minded friends, each of which is earning good money and making their own neat philanthropic bets.

--- 

In these scenarios, you have small, unaccountable "bands of brothers" operating independently. They are much more robust, it's harder to unsettle them, and they have some localized knowledge about their own area. My sense is that an enlightened dictator would beat that approach. But capable autonomous cells would beat having an incompetent dictator. And CEA or Open Philanthropy aren't particularly close to being enlightened dictators.

At a slightly higher level of coordination, you could have an Effective Altruism Federal, something akin to the [Informal Anarchist Federation](https://en.wikipedia.org/wiki/Informal_Anarchist_Federation)[^informal] or a [clandestine cell system](https://en.wikipedia.org/wiki/Clandestine_cell_system). Here, you might have some common aims and identities and occasional collaboration, but different subgroups would still maintain operational independence.

[^informal]: what a great name.

## Earn to Build

Capitalism has been a powerful engine of prosperity, and companies like the [Grameen Bank](https://wikiless.nunosempere.com/wiki/Grameen_Bank?lang=en) or [Wave](https://wave.com) are pushing humanity further and making it more prosperous. Inspired by those successes, an "earn to build"[^earntogive] movement could aim to to build scalable companies that produce vast amounts of good, capture some of that profit, and scale up. Take the PayPal mafia playbook and attempt to replicate it.

[^earntogive]: See [Earning to Give](https://en.wikipedia.org/wiki/Earning_to_give). One could also imagine other variants beyond your regular "earn to survive", e.g., earn to empire-build, earn to attain financial independence, earn to learn, earn to explore, earn to art, etc.

The shape of an earn to build movement would depend on whether its startups try to [bootstrap](https://www.digitalocean.com/resources/article/how-to-successfully-bootstrap-your-startup)[^bootstrapping], try to get VC funding, or try to get funding some other way. Bootstrapping here means that founders use their own money to start their startup: at the cost of they starting out with less capital at the beginning, they'll retain more control if they succeed. Conversely, taking VC money might warp incentives: normally because VC's utility function is more linear, but also because maximizing profitability is not maximizing value. Altruists might have more linear utility functions, though. And it is not inconceivable that optimizing profitability and optimizing value could look similar: it wouldn't have been that implausible for Google to have retained its "don't be evil" motto. Though this looks less likely after recent OpenAI drama.

We could also imagine more niche forms of funding. We have the [Thiel Fellowship](https://en.wikipedia.org/wiki/Thiel_Fellowship), which is a one-off unconditional grant of $100k, some[^somethiel] of whose [recipients](https://en.wikipedia.org/wiki/Thiel_Fellowship#Notable_recipients) have gone on to do very interesting things. The unconditional grant format is interesting, since it gives grantees the opportunity to pursue their own goals fully. On the other extreme, we could imagine a [robinhansonesque](https://www.overcomingbias.com/) deal where a prospective founder sells a fraction of all of his [future income flow](https://www.overcomingbias.com/p/four-uses-for-personal-tax-assets), not just of any particular venture.

[^somethiel]: 20 fellows selected annually, times thirteen years of operation gives 286 total fellows. [19 of them](https://en.wikipedia.org/wiki/Thiel_Fellowship#Notable_recipients) are notable enough to have their own Wikipedia page, for a base rate of 6.6%. And they include Vitalik Buterin and Laura Demming.

If compared against the current EA movement, an "earn to build" movement would be more capitalistic, more right-wing. It would have different and probably better feedback loops. Funding mechanisms could create wacky incentives, but probably less wacky than EA as different VCs compete and their profit motive culls away some arbitrariness and incompetence. It would have more Peter Thiel-like qualities.

In the current EA milieu, you have some startups, like [Wave](https://www.wave.com/) and [Sparkwave](https://www.sparkwave.tech/) (unrelated). You also have [Founders Pledge](https://www.founderspledge.com/news/pledge-ventures-first-close) and, recently, their very cautious [^founderspledge] [associated venture capital branch](https://www.founderspledge.com/news/pledge-ventures-first-close).

[^founderspledge]: Their investment criteria are structured so that they only invest if: The founders or leadership of the company are Founders Pledge members, the company is raising at least $20 million in a series B or later, and a "top tier" VC is participating in the round. Seems like a good way to get some of the upside of their members, but also, because they are shadowing top-tier funds in series B or later, they aren't causing new companies to exist.

## Hierarchical Altruistic Networked Structure

Consider the Linux kernel, an example of a successful international collaboration[^linux]. It adopts the following structure:

[^linux]: Linux has both centralized and decentralized features. In terms of centralization, you have the fact that Linus Torvalds ultimately decides what goes in. In terms of decentralization, you have the fact that anyone can contribute, package it, fork it, improve it. I think that the centralization aspects would hurt and be more salient if Linus was incompetent and/or stringent, but as is, he seems competent and laid back enough that he can enable and channel a wide community around Linux, rather than holding it back. Note, though, that this is my impression from a distance, and the reader is welcome to take this as a pointer to how a social movement could work. Though e.g., tools like git *do* enable decentralized cooperation.

![](https://git-scm.com/book/en/v2/images/benevolent-dictator.png)

Now consider applying this structure to this [big list of cause candidates](https://forum.effectivealtruism.org/posts/SCqRu6shoa8ySvRAa/big-list-of-cause-candidates). You could have a structure that tried to implement some form of [Multi-armed bandit](https://en.wikipedia.org/wiki/Multi-armed_bandit) optimization to it:

- A dictator in charge of coordinating the different moving parts
  - A lieutenant in charge of value estimation
    - Sub-lieutenants in charge of value estimation for each cause area
      - Teams and people working on specific projects
  - A lieutenant in charge of coordinating exploration for valuable projects
    - Sub-lieutenants in charge of coordinating exploration in each cause area
- Heretics who disagree with some trade offs, and fork the git repository

In this story, you have people working coming to work together in one project, as represented by the artifact of a git repository, i.e., a bunch of files and abstractions for collaborating on them and attaining social consensus. Imagine that there was no EA Forum, only an EA Forum wiki and people making decisions on top of it. Imagine that Open Philanthropy was open, and you could see its employees working through the reasoning for their grants as they go, and review their assumptions. Imagine that instead of a hodgepodge of research institutes, a distributed research effort transparently determined the next actions taken by the distributed implementation arm, and that there was some more speculative version of Charity Entrepreneurship. Imagine that there was some well-understood mechanism for getting funding for starting a project, or for getting funding after the fact if you do some work that proves to be valuable.

Imagine that you are interested in contributing to the Hierarchical Altruistic Networked Structure, so you look at its giant to-do list, and you pick a few items that correspond to your interests and skills, do them, and get some livable wage as a result. Imagine that you had a behemoth like Google, but oriented towards altruistic enterprises. 

For less half-baked version of these ideas, consider the vision laid out in [What to do with people](https://forum.nunosempere.com/posts/oNY76m8DDWFiLo7nH/what-to-do-with-people), maybe one of the best and most underrated posts in the EA Forum over the last few years.

## Blue Rose Unfolded

Due to the tireless efforts of [kbog](https://forum.nunosempere.com/user/kbog) and his [Policy Platform for Maximum Happiness](https://happinesspolitics.org/platform.html), EA gets involved in party politics. It gets absorbed by some coalition and shapes it to be 5% more altruistic, 7% more effective. Personally, an effectiveness-oriented strand of millitant Bernie-Sanderism sounds pretty cool to me, as does a Tea party with EA characteristics—[Bjørn Lomborg](https://en.wikipedia.org/wiki/Bj%C3%B8rn_Lomborg) reads as right wing these days but he got started by noticing that climate change interventions didn't do well on [cost-effectiveness analysis](https://copenhagenconsensus.com/post-2015-consensus).

Or, alternatively, EA becomes defanged and party politics doesn't change. You start out with a movement committed to doing the most good, you let people in and you don't take care of teaching them the basics, and you end up with members justifying their petite bourgeoisie motions as though they had meaningful impact, or performatively clamoring about the issues of the day.

## Effective Troopers

In a parallel world, a charismatic [gay Mormon](https://tracingwoodgrains.substack.com/p/thank-you-for-my-service) created a high commitment civil organization dedicated to improving the world. Effective Troopers fills the empty hole left by religion in many Western people; sign up today!

You will be assessed, trained and deployed to tackle one of humanity's most pressing problems, in order of importance and neglectedness, weighted by your likelihood of success. We are firm believers in [Auftragstaktik](https://nunosempere.com/blog/2023/11/30/auftragstaktik/): you will be given clear objectives and the tools—and eventually manpower—needed to achieve those goals. 

Our high ambition means that about 30% of the time, you will fail: others will learn from your efforts, as they will be fed into our [multi-armed bandit](https://wikiless.nunosempere.com/wiki/Multi-armed_bandit?lang=en) estimations to choose missions at the right ambition/tractability threshold. You will need a disciplined and robust mind, we will help you develop it.

## Bonus

- Effective altruism veiled. What if EA was more akin a hellenistic mystery cult. What if it had initiation rites, clear progression hierarchies, expected commitments and rituals... You get introduced to this Pythagorean sect, and you see how deep the rabbit hole goes.
- Champagne socialism revisited. What if EA goes (¿continues to go?) in the direction of being a belief that is like an attire, without further consequences. People sip their wine and participate in the petty drama, but they don't act differently.
- From the creators of "if you haven't missed a flight, you are spending too much time in airports" comes "if you aren't producing an SBF or two, then your movement isn't being ambitious enough."
- Effective Altruism Hufflepuff. What if EA, but any implication of its core beliefs that was offensive or sounded ridiculous to someone was simply... slided over? Many people think that something like shrimp welfare is ridiculous[^singer]. What if instead you increased your mass appeal by not staying weird and sacrificing a few Robin Hansons?
- Effective Tyler-Durdenism? Never heard of it.
    
[^singer]: Or, more seriously, Peter Singer has a bunch of eyebrow-raising beliefs that follow pretty straightforwardly from utilitarianism.

## Conclusion

Perhaps these visions of EA show that it is much less constrained than it appears.
