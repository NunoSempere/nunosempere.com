Prediction Markets in The Corporate Setting
==============

What follows is a report that Misha Yagudin, Nuño Sempere, and Eli Lifland wrote back in October 2021 for [Upstart](https://wikiless.org/wiki/Upstart_(company)?lang=en), an AI lending platform that was interesting in exploring forecasting methods in general and prediction markets in particular. 

We believe that the report is of interest to EA as it relates to the [institutional decision-making](https://forum.effectivealtruism.org/tag/institutional-decision-making) cause area and because it might inform EA organizations about which forecasting methods, if any, to use. In addition, the report covers a large number of connected facts about prediction markets and forecasting systems which might be of interest to people interested in the topic.

Note that since this report was written, Google has started a new [internal prediction market](https://cloud.google.com/blog/topics/solutions-how-tos/design-patterns-in-googles-prediction-market-on-google-cloud). Note also that this report mostly concerns company-internal prediction markets, rather than external prediction markets or forecasting platforms, such as Hypermind or Metaculus. However, one might think that the concerns we raise still apply to these. 

## Executive Summary

*  We reviewed the academic consensus on and corporate track record of prediction markets.
*  We are much more sure about the fact that prediction markets fail to gain adoption than about any particular explanation of why this is.
*  The academic consensus seems to overstate their benefits and promisingness. Lack of good tech, the difficulty of writing good and informative questions, and social disruptiveness are likely to be among the reasons contributing to their failure.
*  We don't recommend adopting company-internal prediction markets for these reasons. We see room for exceptions: using them in limited contexts or delegating external macroeconomic questions to them.
*  We survey some alternatives to prediction markets. Generally, we prefer these alternatives' pros and cons.

## Introduction

This section:

*  Defines prediction markets
*  Outlines their value proposition

### What are prediction markets

Prediction markets are markets in which contracts are traded that have some value if an event happens, and no value if an event doesn't happen. For example, a share of "Democrat" in a prediction market on the winner of the 2024 US presidential election will pay $1 if the winner of the 2024 election is a Democrat, and $0 if the winner is not.

Prices in prediction markets can be interpreted as probabilities. For example, the expected value of a "Democrat" contract in the previous market is $ 1 ⋅ p + $ 0 ⋅ ( 1 − p ) , where p is the chance that a Democrat will win. To the extent that the market is efficient, one expects the expected value of a contract to be equal to its current value. So if one observes a contract price of $0.54, one can deduce the expected probability by setting $ 0.54 = $ 1 ∗ p + $ 0 ∗ ( 1 − p ) , and thus p = 0.54 = 54 % . It is also in this sense that one says that "the market as a whole expects" Democrats to win with 54% probability.

One might expect markets _not_ to be efficient, for instance after remembering Keynes' adage that "markets can remain irrational longer than you can remain solvent." And we do see inefficiencies in modern prediction markets, sometimes glaring. However, note that, unlike the stock market, prediction markets have hard deadlines after which the market comes in contact with reality and gets resolved.

Besides binary prediction markets, there are also markets with multiple options—e.g., "Who will win the 2024 election?", with multiple contracts only one of which will pay out in the end,— or markets that pay out proportionally to some yet unknown number—e.g., "How many Senate seats will Republicans control after the 2022 elections?", which pays out proportionally to the number of seats.

Predictions markets thrive in some niches:

*  Bookmakers are offering odds and taking bets for notable sports events.
*  Markets like PredictIt and BetFair attract quite a bit of money for major political events \[cf. predict-it-growth\].
*  Some financial markets can be thought of as prediction markets.

 \[cf. predict-it-growth\]: US Presidential elections volumes on Betfair [have been growing](https://www.lesswrong.com/posts/4XXnMXfTrYXqpugwB/growth-of-prediction-markets-over-time?commentId=t2jmTnhMB9x2GzhrP) at an implied ~35%/year.

### Value proposition

The core value proposition of prediction markets is that they may produce accurate, calibrated and useful probabilities. They create an incentive for participants to seek information that would give them an edge. And when those participants trade on a prediction market they reveal some information about their degree of conviction. Finally, prediction markets provide an aggregate of the differing perspectives, namely the current price. At their best, the market mechanism aggregates more information than what could fit in the working memory of any one individual.

Ideally, prediction markets would elicit knowledge that wouldn't have otherwise been shared—e.g., the knowledge that a given deadline is unrealistic—and that knowledge would be used to make better decisions. This is the core pathway to impact, and if prediction markets don't end up changing decisions, their impact will normally be negligible, no matter how accurate they otherwise are.

In addition to producing accurate forecasts which drive better decisions, prediction markets could also be used to manage risk. For instance, a company fearing rising taxes or an increased regulatory burden might make a bet to hedge against that risk, i.e., make a bet on the side of the undesirable event so that if it happens, the company mitigates part of its downside. As of today, few prediction markets would have high enough volume and liquidity to allow for meaningful hedging, with the possible exception of Nadex and FTX, the latter of which has recently seen behaviour consistent with a large actor hedging against the chance of the Tokyo Olympics being cancelled. Note that this point, while perhaps of interest to Upstart more generally, doesn't feature in the rest of the report, which focuses on internal prediction markets.

Besides better decisions and risk hedging, one can also speculate that prediction markets may have a range of social benefits. For example, Robin Hanson speculates that prediction markets, if widely used, might create a visible expert or social consensus with clear incentives for honest contribution for arbitrary questions. For example, input from prediction markets can help people operate under shared and trustworthy assumptions like "we will not return to the office in the next half a year." Widespread prediction markets might also decrease the influence of especially persuasive individuals and groups: it would be harder for them to manufacture consensus just through good oratory without having a good track record or putting money on the line.

Further, Robin Hanson developed a proposal for governance called Futarchy \[cf. hanson-2013\]. In a futarchy, conditional prediction markets are used for estimating welfare (e.g. GDP or market capitalization) conditional on taking decisions under consideration, then the decision leading to the highest welfare is chosen. \[cf. crypto-futarchy\]

But overall, better information on its own has very little value if doesn't eventually end up changing any decisions. One can thus calculate the value of information as the difference in value between the decisions before and after gaining that information. In particular, the value of information may be lower than the cost to acquire it, and this could be a reason for the lack of adoption of prediction markets.

 \[cf. hanson-2013\]: Hanson, Robin. “Shall We Vote on Values, But Bet on Beliefs?: Shall We Vote on Values, But Bet on Beliefs?” _Journal of Political Philosophy_ 21, no. 2 (June 2013): 151–78. [https://doi.org/10.1111/jopp.12008](https://doi.org/10.1111/jopp.12008).

 \[cf. crypto-futarchy\]: In 2016, one of the most ambitious futarchy-based projects, the [DAO](https://en.wikipedia.org/wiki/The_DAO_(organization)), was launched and failed. But more crypto organizations have tried out futarchy-inspired governance mechanisms [before and since](https://en.wikipedia.org/wiki/Decentralized_autonomous_organization#List_of_DAOs). It seems plausible to us that in the coming years a larger number of organizations might try similar models.

## Track record

This section

*  Gives examples of high-profile companies using prediction markets
*  Outlines the academic consensus on prediction markets, and the limitations of this consensus

### High-profile companies that have used prediction markets.

In our literature review, we have found that a large number of high-profile for-profit organizations have used prediction markets. These include Arcelor Mittal, Best Buy, Boeing, CNBC, Chevron, Chrysler, Deutsche Bank, Electronic Arts, Ford, General Electric, Goldman Sachs, Google, Hewlett Packard, Intel, J&J, Koch Industries, Lockheed Martin, MITRE Corp, Microsoft, Motorola, Nokia, PayPal, Proctor and Gamble, Qualcomm, Siemens, Yahoo and Yandex, among others.

Sources for some of the above are given in the following table. The rest are taken from Table 1 of [Cowgill and Zitzewitz, 2015](https://academic.oup.com/restud/article-abstract/82/4/1309/2607345).

| Company                      | Source                           | Notes                                                                                                 |
|------------------------------|----------------------------------|-------------------------------------------------------------------------------------------------------|
| Eli Lilly                    | The End Of Management, The Times | Large American pharmaceutical corporation                                                             |
| Ford & others                | Cowgill and Zitzewitz, 2015      |                                                                                                       |
| Goldman Sachs, Deutsche Bank | Wolfers and Zitzewitz, 2004      | High volume markets, website only remains on the Internet Archive                                     |
| Google                       | Cowgill, Wolfers, et al., 2009   |                                                                                                       |
| Hewlett Packard              | Chen and Plott, 2002             | Markets were thinly traded, but still performed better than HP's own predictions                      |
| Koch Industries              | Cowgill and Zitzewitz, 2013      | Early draft of the 2015 paper mentions Koch Industries, though this was removed in the final version. |
| Microsoft                    | Prediction Markets at Microsoft  | Very clear pdf; very much worth downloading and reading.                                              |
| Nokia                        | Hankins and Lee, 2011            |                                                                                                       |
| Siemens                      | Ortner, 1998                     | Prediction markets predicted deadlines better than other processes.                                   |
| Yahoo                        | Bloomberg                        |                                                                                                       |
| Yandex                       | Interview with Yandex employee.  |                                                                                                       |

### Academic consensus

In practice, prediction markets have done well at predicting elections, sports outcomes, or at anticipating future data releases such as COVID infection numbers. Multiple academic studies—as well as analogies with the more liquid stock markets—give reason to believe that prediction markets should produce accurate probabilities. To sample some representative quotes from the academic literature:

> Ortner (1998) described an experiment at Siemens in which an internal market predicted that the firm would definitely fail to deliver on a software project on time, even when traditional planning tools suggested that the deadline could be met.
> 
> An internal market at Hewlett-Packard produced more accurate forecasts of printer sales than the firm’s internal processes (Chen and Plott, 2002).
> 
> In each case, the firms ran real money exchanges, with only a relatively small trading population (20 – 60 people), and subsidized participation in the market, by either endowing traders with a portfolio or matching initial deposits. The predictive performance of even these very thin markets was quite striking. (Wolfers and Zitzewitz, 2004)
> 
> Despite theoretically adverse conditions, we find these markets are relatively efficient, and improve upon the forecasts of experts at all three firms by as much as a 25% reduction in mean squared error (Cowgill and Zitzewitz, 2015)

Already in 2008, a group of prestigious economists, including four Nobel Prize recipients published _The Promise of Prediction Markets_ \[cf. arrow-2008\], which outlined their common optimism about prediction markets as a tool for producing forecasts with lower prediction error than conventional forecasting methods, and urged the Commodity Futures Trading Commission as well as US state and federal legislatures to establish safe-harbour rules to encourage the use and research of prediction markets.

In addition, there are elegant theoretical reasons to think that prediction markets may perform optimally. For instance, per (Begelzimmer et al., 2012) \[cf. beygelzimer-2012\], if bettors are following a Kelly betting strategy the market learns at the optimal rate, i.e., the market price reacts as if updating according to Bayes' Law.

 \[cf. arrow-2008\]: Arrow, Kenneth J., Robert Forsythe, Michael Gorham, Robert Hahn, Robin Hanson, John O. Ledyard, Saul Levmore, et al. “The Promise of Prediction Markets.” _Science_ 320, no. 5878 (May 16, 2008): 877. [https://doi.org/10.1126/science.1157679](https://doi.org/10.1126/science.1157679).

 \[cf. beygelzimer-2012\]: Beygelzimer, Alina, John Langford, and David Pennock. “Learning Performance of Prediction Markets with Kelly Bettors.” ArXiv:1201.6655 \[Cs, q-Fin\], January 31, 2012. [http://arxiv.org/abs/1201.6655](http://arxiv.org/abs/1201.6655).

### What is left unsaid in the academic literature

However, in most cases, prediction markets are only used for a short time, and early adopters chose not to incorporate markets into their business after initial trials. For instance, Hewlett Packard's prediction market was in collaboration with CalTech academics, and once the collaboration ended, it seems that so did the use of prediction markets within HP. Similarly, we know that the use of prediction markets at Google ceased after their main advocates left the organization.

In contrast with the literature, which is generally very optimistic, an internal Microsoft document provides a more honest outlook over the advantages and disadvantages of prediction markets. Further, it lists observed pros and cons of using prediction markets in the various contexts: project schedules, external reviews, sales, etc. It can be found [here](https://users.wfu.edu/strumpks/PMConf_2007/HenryBerg(PredictionPoint%20KC%20071101).pdf) ([archive link](https://web.archive.org/web/2020*/https://users.wfu.edu/strumpks/PMConf_2007/HenryBerg(PredictionPoint%20KC%20071101).pdf)). Although outdated and terse, it also has healthier epistemic incentives. We have also gathered excerpts related to prediction markets from interviews by Tyler Cowen of industry and academic experts at \[private\]. These are also worth reading.

To express this point explicitly, the academic literature on prediction markets gives a distorted perspective on their efficacy \[cf. page-2020\] and suitability for companies. This is probably because of perverse academic incentives which cause academics to overstate the impact and promisingness of their research. In particular, the academic literature makes emphasis on how prediction markets are interesting and promising, but makes less emphasis on how they are underdeveloped as a technology. Academics that study prediction markets for a time and then move on to other more fruitful areas also do not generally write up the reasons for their change in focus—academic bloggers like Hanson are an exception.

Further, we would like to note that while prediction markets performed better than sales and scheduling techniques used 20 years ago, they will face much stronger competition now because forecasting tools have improved a lot (e.g., through data science methods, but also through e.g., tools like Fibonacci point estimation \[cf. fibonacci\].) We dive into specialized SaaS systems later in the report.

Proponents like Robin Hanson argue that prediction markets are not popular because they are too politically disruptive: they expose hypocrisy and there is a lot of hypocrisy; they remove excuses. They can be thought of as a very direct socially awkward person who doesn’t navigate interpersonal relationships well and misses social conventions around excuses, failure, or the need to convince stakeholders. This is unpleasant for people interacting with them. While this person might be right, they are bad at coalition politics (and hence are ostracized) and/or they are harmful because their interactions—while candid—are not productive in a more socially competent world.

To give a recent example, forecasters at platforms like [pandemic.metaculus.com](https://pandemic.metaculus.com) were communicating fairly accurate predictions about the coronavirus pandemic, while the US CDC was, to our taste, fairly misleading. This can be partially explained by the different roles forecasters and bureaucrats play. Forecasters are just seeking truth, whereas bureaucrats are coordinating a large number of people who may have very different beliefs and decisions to make. Just being candid by saying that "we are unsure but this and that seems likely" works for the Metaculus crowd advising friends and relatives. But it might not work for a government agency communicating to millions of indifferent and millions of anxious people looking for guidance or needing some coaxing to do anything at all. Of course, less charitable interpretations of the CDC's behavior are also possible.

Another guess as to why prediction markets fail to gain adoption is that it's hard to incorporate them into the making of decisions. Partly, social forces outlined in previous paragraphs are to blame. But also, it's just hard to operationalize questions such that they are useful for making decisions. Even if one did, it might be difficult to create a forecasting system around those questions which has low enough overhead for them to be worth it. This suggests trying to use some of prediction markets' advantages while minimizing overhead. In a later section, we outline how this might be possible.

However, there are many other hypotheses as to why prediction markets fail to be useful or to get adopted; some are outlined in a later section. Some of these hypotheses may apply to Upstart; they seem difficult or expensive to rigorously falsify, and would likely require high overhead and coordination costs to anticipate, diagnose and overcome.

The rest of this report

*  provides details about how prediction markets ought to be structured in order to increase their usefulness,
*  presents hypotheses about the failure of prediction markets in the corporate setting,
*  outlines alternative information aggregation mechanisms,
*  and concludes with recommendations.

 \[cf. page-2020\]: Cf. [this review](https://twitter.com/page_eco/status/1308049931278512128) of experimental markets which [finds](https://twitter.com/page_eco/status/1308054640911556609) that markets are much better at incorporating public information than private information. And one of the selling points of the internal prediction market is to elicit information from everyone across the organization.

 \[cf. fibonacci\]: cf. [Fibonacci Agile Estimation](https://www.productplan.com/glossary/fibonacci-agile-estimation/)

## Requirements and challenges for a well-functioning prediction market

We believe the most important challenges are:

1. Maintaining a prediction market given current technologies
  *  Our research finds that current technologies add substantial overhead to maintaining prediction markets.
2. Writing questions that are both informative for decision-making and attractive to traders
  *  In particular, questions must a) target areas where employees have useful knowledge to share, b) concern topics where management has significant uncertainties about how to proceed, and c) be suitable for clear operationalization.
  *  There are many failure modes questions can fall into, as detailed below.
3. Attracting enough predictors to get accurate forecasts while also preventing the market from taking up too much of the employees' time and attention.
4. Managing the social effects and potential backlash of prediction markets.

However, we are fairly uncertain because it is difficult to observe the reasons for the lack of adoption by corporations. In any case, we note that the practical evidence for the usefulness of prediction markets is not so overwhelming that they have penetrated industry. In particular, if prediction markets were known to be highly useful, it would seem difficult to keep this knowledge limited to any one company, given imperfect employee retention.

### Categorization scheme

The requirements and challenges for well-functioning prediction markets can be broken down into the following categories:

1. The markets must have a low enough cost to create and maintain.
2. The markets must provide more value to decision-makers than the cost to create them and to subsidize predictions on them.
3. The markets must be attractive enough to traders to elicit accurate predictions.
4. The markets must not have large negative side-effects, such as costs to the company's dynamics and morale.

What follows is a summary of the most important challenges, followed by details on each requirement and challenges to achieving it. We quote in full several requirements and challenges from Zvi Mowshowitz's [Prediction Markets: When Do They Work?](https://thezvi.wordpress.com/2018/07/26/prediction-markets-when-do-they-work/) and Tyler Cowen's [Why don’t more businesses use prediction markets?](https://marginalrevolution.com/marginalrevolution/2006/03/why_dont_busine.html).

### The market must have a low enough cost to create and maintain

The market must be low enough cost to create and maintain. As documented above, the creation step has happened at many companies so the main plausible bottleneck seems to be the maintenance step. Challenges for maintaining prediction markets include:

1. Chicken-and-egg problems of an immature technology: 
  *  On the engineering side: Prediction markets and forecasting platforms are not mature as a technology. They are not streamlined, they are clunky \[cf. clunky-ux\] to use, and there are not many people familiar with the practice. This is unlike, e.g., Scrum or other development frameworks. In particular, it's hard work to incorporate prediction markets into decision-making or into the "rhythm-of-the-business", and this hard work hasn't been done yet. If "prediction markets as a service" was an established product, they might be worth buying, but they aren't.
  *  On the social side: Practically nobody is experienced at judgmental forecasting or at prediction markets. On the one hand, this makes it harder and more expensive to acquire the related skills \[cf. superforecaster-selection\]. But on the other hand, introducing a new metric to rank current managers or to decide whom to promote to managerial positions might be perceived as disruptive, unfair, and might otherwise be subject to the costs of innovation.
  *  On the rationality-waterline side: On well-functioning financial and prediction markets, participants self-select from the general population by being good at beating the market, and come from a huge pool of candidates. Corporate prediction markets don't have the luxury of a large pool to select from and e.g. lack of necessary skills like probabilistic thinking shrinks an already thin market.
2. Invisible improvements: Prediction markets might lead to a significant improvement in decision-making, say, to decisions that are 20% better. But that improvement might be too small to be perceivable, so the methodology gets abandoned after an initial experimental period. Peter Thiel's _Zero to One_ suggests that startups need a 10x improvement in their product for that difference to be so noticeable as to produce a "wow" factor which is conducive to a monopoly. Prediction markets as they are right now don't seem to reach that threshold. Note also that in the case of a startup offering prediction markets as a service, it might only capture a fraction of the improvements in decision quality that their clients can detect, so it might never reach profitability \[cf. personal-communication\].

 \[cf. clunky-ux\]: It seems plausible that the lack of software with a great UX (user experience) might be one particularly significant bottleneck. For example, two of the authors couldn't stick with forecasting on Metaculus because of the messy UX around forecasting distributions and keeping track of forecasts. Current software (both prediction markets and forecasting platforms) isn't shiny and pleasant to use, nor it is very collaborative or socially rewarding. Perhaps because competent people have many demands on their time, not finding platforms attractive could likely prevent them from engaging.

 \[cf. superforecaster-selection\]: In contrast, superforecasters are selected in two stages: they predict on at least 100 questions on GJ Open and pass a 3-month long trial period. And from [research into superforecasting](https://aiimpacts.org/evidence-on-good-forecasting-practices-from-the-good-judgment-project-an-accompanying-blog-post/#easy-footnote-10-1260) we know that talent-spotting is responsible for the most notable jump of accuracy.

One can also speculate that part of Robinhood's success in introducing millions to day trading can be attributed to its superior UX. 

 \[cf. personal-communication\]: In personal communications, a celebrated data scientist consultant/executive said that introducing "big data" to already competitive traditional companies leads to limited gains (~1-2%).

### The questions provided by the market must provide more value to decision-makers than the cost to create and predict on them

For a market to be useful, good predictions must provide substantial value to decision-makers. The time to create and predict on questions is non-negligible and the benefits need to be large enough to outweigh the time costs \[cf. caveat-misha\]. Additionally, perhaps the bottleneck for improving the company isn't necessarily better decisions after all.

To give a few examples:

*  Sales are important but forecasting sales might not be as important. In particular, forecasting sales might not meaningfully help increase them. Salespeople already have well-developed incentive structures to encourage them to sell more. Sales forecast might be very important for just-in-time manufacturing, but this is only crucial for a few industries.
*  Forecasting project scheduling is important. But it might not restore confidence in likely-to-be-missed deadlines, and measures taken to catch up to deadlines would or would not be implemented regardless of the forecast.

The following are some specific hypotheses as to why it may be hard to create questions that provide substantial value to decision-makers:

1. No matter what they pretend, businesses are not much interested in forecasting many future variables. Successful businesses find product markets they can control for a long time. They do a few things well, and let a surprisingly large number of tasks slide. \[cf. company-model\]
2. For many—but not all—variables of interest, we already implicit betting markets in the form of resource prices. These may be explicit—like aluminium prices, or slightly fuzzy—like the market salary for software engineers. Firms can look at those prices in outside markets for the information they need, rather than to potential internal prediction markets. 
  *  Note that this applies to some questions of interest, but not to others. For instance, most of the information about whether "we should move to Columbus/Texas/.../outside the Bay area" might be contained in the market salary and property prices at those locations, which are external markets. But for markets that more intimately depend on the firm's specific details—e.g, whether a project will be done on time, or how much of a product will be sold in a given timeframe—an external market might not be desirable or feasible. In this latter case, one might then use an internal information aggregation method, or use a SaaS to generate an estimate.
3. Large corporate companies are far more constrained than most outsiders imagine. Interest groups must be courted, coordinated, and sometimes fought every step of the way. When it comes to choice, there are fewer degrees of freedom than one might think. The real question is not what to do, but rather having the will and effectiveness to do it. Prediction markets don’t help much in this regard.
4. There is a challenge about which markets to choose. Markets about schedules might lead to self-fulfilling prophecies or undermine management. Markets about intimately decision-relevant factors might not see much participation, or participants might not have as much information. Markets about external events are likely to attract more interest, but also be less useful.
5. The time taken for specialized employees to familiarize themselves with prediction markets is also costly. For example, if 20 employees who are paid $50 to $500 per hour spend one to four hours on a given prediction market, costs can quickly balloon. Also note that most people are unfamiliar with probabilities, so prediction market benefits may not accrue until participants become more familiar with probabilistic thinking. Note that under this objection, prediction markets could still add new information, just not enough to justify their cost vis-à-vis a product manager. 
  *  Note that we need to compare the costs of prediction markets with the costs of alternatives. Some alternatives are maybe pretty costly as well. The default direct substitute for eliciting information and building consensus is meetings. It is well-known that they can be very wasteful (as they have quadratic costs in the number of participants). We are unsure about how much prediction markets will help with reducing the costs of meetings. It's plausible that some disagreements which arise during meetings could be resolved faster by deferring to whatever prediction markets say. But at least initially, prediction markets and regular procedures will co-exist, so initial costs might be pretty high.
6. Prediction markets and similar forecasting tournaments can have an addictive quality that could hurt further hurt productivity. The addictive qualities of gambling are well-documented and prediction markets share some of these addictive features. Eli has experience with being sucked into spending more time predicting on forecasting platforms than seems best, due to trying to increase his score. A similar thing may happen with trying to increase one's profits in the market. Even if employees aren't spending up much raw time on prediction markets, they may use up too large a portion of their attention \[cf. attention\].
7. The regulatory landscape is hostile towards prediction markets, and this increases their costs for risk-averse actors. In particular, real-money prediction markets have very attractive incentive structures, but fall prey to onerous regulations pertaining to gambling or the trading of securities. 
   
  *  Even if internal prediction markets are not outlawed right now, they might become outlawed in the future. Delegating a fraction of one's decision-making capacity to a procedure that might get outlawed might be unnecessarily risky. It also doesn't feel outside the realm of possibility that a ruthless competitor will lobby legislators to be hard on gambling, or frame prediction markets in a negative light to the public.

 \[cf. company-model\]: Perhaps when companies are small, there isn't much need for prediction markets as they add overhead and there isn't much need for more information aggregation. When a company is large, it has often already found a niche and should mostly keep doing the same thing. It may be tricky to find the middle ground where prediction markets can provide real value.

That said, it's hard to know whether such a model is true in practice, and e.g., perhaps small companies could use prediction markets or forecasting tournaments to better make early strategic decisions.

 \[cf. caveat-misha\]: Though note that at Google, employees participated in a number of prediction markets as their 20% projects. This might have been fairly low cost if one thinks that the time would otherwise have been wasted. On the other hand, we know the number of traders, and a [quick guesstimate](https://www.getguesstimate.com/models/19293) arrives at an implied cost of wages of between $50k and $15M, with a point estimate of $1.5M.

 \[cf. attention\]: For more on the importance of managing attention, see this [blog post by a CTO on managing attention](https://www.benkuhn.net/attention/) or this [book on maintaining enough focused work](https://www.goodreads.com/book/show/25744928-deep-work). Further, prediction markets are uniquely demanding: their nature requires constant vigilance as profits and losses are immediate. In contrast, on other forecasting platforms, one can update predictions semi-regularly without noticeable loss in performance.

### The market must be attractive enough to traders to elicit accurate predictions

A potential failure mode of prediction markets is not generating enough interest from traders to get accurate predictions. In particular, some requirements to be attractive to traders—from [an experienced trader](https://thezvi.wordpress.com/)—are:

**1\. Questions must be well-defined**

*  Resolving questions when unforeseen edge cases arise will lead to needless controversy and will put participants off from participating in future markets.
*  It's very unpleasant for traders not to get paid when they think they ought to per their reasonable interpretation of resolution criteria.
*  A bit more abstractly, the risk of not getting paid because of a misinterpretation or ambiguity in the rules decreases traders' expected reward and therefore discourages participation.

Here is an example of how it can be unexpectedly hard to write a well-defined question for a seemingly straightforward matter: "Who will control US Senate after the elections?". "Control" is a colloquial term so some kind of operationalization is needed. "Would Democrats have >50% of seats in the US Senate?" sounds simple enough. However, it suffers from a few problems, namely:

*  In the case of a 50/50 split, the VP breaks ties, so one might want to redefine control as ">50% seats or 50% seats + the VP"
*  How should one count independent politicians? Today Wikipedia states that Democrats are in majority despite having only 48 senators. This is because two independent senators, Angus King of Maine and Bernie Sanders of Vermont, caucus with the Democrats.
*  How should one count "rebellious" senators? In particular, Joe Manchin of West Virginia can't or doesn't want to vote on measures that would jeopardize the Democratic Party's standing on West Virginia, which is fairly conservative, and Kyrsten Sinema of Arizona likewise feels attached to bipartisanship. As a result, the Democratic Party cannot implement some of its priorities.

So a seemingly straightforward question faces many possible technicalities. This could be circumvented by taking a reasonable yet ultimately arbitrary stance on confusing details, e.g., by defining "control" as ">50% of seats". But then one might find oneself talking about and trying to forecast a strange slice of reality.

In any case, unforeseen corner cases lead to heated online discussions with no clear right way to resolve a question. And such a scenario is not particularly unlikely \[cf. metaculus-2021\] unless one drafts questions very carefully.

 \[cf. metaculus-2021\]: And corner cases are not particularly unlikely, see [this question](https://www.metaculus.com/questions/7240/what-is-a-counterparty-risk-of-polymarket/) about Polymarket hitting an edge case. As mentioned above, corner cases are bad as they make traders less certain about their interpretation of resolution and hence push for higher edges for trading.

**2\. Questions resolve soon** \[cf. hedgehog-markets\]

*  When traders make bets, they lock in funds that could have been used for other purposes (e.g., for other bets, for investing in an index fund, for getting crypto yields, etc.).
*  If markets are sufficiently deep and liquid, this issue can be mitigated, because participants can exit their positions after profiting from discovering new information earlier than others.

 \[cf. hedgehog-markets\]: Recent developments in crypto prediction markets make this point somewhat moot, because the funds at stake can be invested to generate yield using a crypto contract, and then given out to the winner. However, this mechanism seems far away from being incorporated into prediction market platforms usable by corporate entities.)

**3\. Questions are likely to be resolved**

*  Some markets resolve ambiguously or don't resolve at all in certain circumstances. For example,
  *  if markets are conditional on another event and the conditional doesn't happen
    *  e.g. [_Conditional on President Trump being convicted of "incitement of insurrection," what will the Senate's average Bipartisan Index score be from 2021-2022?_](https://www.cset-foretell.com/questions/108); or
  *  if resolution criteria are no longer evaluable e.g. a resolution referencing to a rating that stops existing or changes dramatically
    *  e.g. [_What will U.S. holiday season retail sales be for 2020 relative to the 2019 holiday season?_](https://www.gjopen.com/questions/1672-what-will-u-s-holiday-season-retail-sales-be-for-2020-relative-to-the-2019-holiday-season) was voided because the data provider changed methodology and mentioned that methodology may be adjusted again.
*  Such questions lock in money and offer a much reduced expected return, and hence disincentivize participation. More precisely, participants win (or lose) nothing if the question is voided. If this happens with non-negligible probability, expected returns are reduced proportionally to that chance

**4\. Sources of disagreements and profits (also known as "suckers at the table")**

*  Prediction markets require sources of disagreements. These can either be:
  *  Direct subsidies
    *  VC money willing to subsidize a market to capture market share, like on Polymarket.
    *  An external actor willing to lose money to elicit more accurate probabilities.
  *  Traders not looking to make money, but who still seek to extract value
    *  Actors hedging against risk, like airlines or countries hedging against movement in oil prices.
    *  Actors who want to manipulate prices, like parties who want to inflate their perceived chances of winning.
  *  Unsophisticated traders
    *  Inexperienced or naïve traders,
    *  Gamblers who bet based on emotional reasons (e.g., on their favourite team, or on principle),
    *  Traders who think they know something, but are mistaken.

**5\. Limited hidden information**

Limited hidden information, or lack of insider trading, is also a factor that influences whether traders are attracted to questions. However, we will delegate discussion to a footnote. \[cf. insider-trading\]

 \[cf. insider-trading\]: Zvi considers:

Limited Hidden Information

*  Insider trading does add more information to the market.
*  However, insider trading—or the suspicion that it could happen—drives other traders away
  *  "The first season of Survivor, there was a market on who would win. The production crew found out. Then there was no market."
*  If some group of actors can control the outcome, it might also not be worth it to bet against them.
*  Access to information might be seen as unfair or, in more extreme cases, as favouritism. (We elaborate on that in the section about social costs.)

This last point interacts strangely with well-subsidized markets in the corporate setting because:

*  All employees are "inside traders", and this is well-known and expected from the start
*  It is not a problem if only insider traders participate, since they would be doing so against the liquidity pool set-up by the employer. And the employer would get the information regardless, which fulfills the purpose of the prediction market.
*  Insider-trading can be noticed, and other participants could update on the moves of others. However, note that this leads to an arms-race dynamic in terms of sneakiness and counter-sneakiness.

A 1998 article by Bainbridge, [Insider Trading: An Overview](https://papers.ssrn.com/sol3/papers.cfm?abstract_id=132529) provides a list of 261 papers that have discussed insider trading, and then succinctly summarizes the arguments for and against allowing insider trading as follows:

*  In favour of deregulating insider trading: Insider trading causes the market price of the affected security to move toward the price that the security would command if the inside information were publicly available. If so, both society and the firm benefit through increased price accuracy. Secondly, insider trading might be an efficient way of compensating entrepreneurially inclined managers for having produced information and value. If so, the firm benefits directly (and society indirectly) because managers have a greater incentive to produce additional value to the firm.
*  In favour of regulating insider trading: Traditionally the argument was based on fairness issues, which had little traction in the law and economics communities. Instead, the economic argument in favour of mandatory insider trading prohibitions has typically rested on features of the economics of property rights in information: a ban would reduce adverse selection costs and perverse incentives, increase liquidity, improve confidence in the market, reduce interference in corporate plans, and motivate large shareholders to monitor management instead of seeking to profit from inside information.

But even if a market attracts as many dedicated traders as one can inside a company, this can still not be enough to produce accurate probabilities:

1. Most employees may have no rational basis on which to bet. If someone knows the truth, but is otherwise locked out from credibly signalling that knowledge to management, something is wrong with the organization of the company \[cf. anonymous\]. The small prizes from corporate prediction markets won’t be enough to elicit that knowledge from them in any case.
2. Too small a scale: Prediction markets at companies are too small. For example, maybe they are not liquid enough, and thus inefficient, or they have such a small number of questions that scoring is not really possible. In particular, we might expect overconfident forecasters to make some gains at the beginning, and for the market to not be able to realize this until overconfident predictors have gone bust for the first time. This might result in the initial estimates from the prediction markets being biased by overconfident players.

 \[cf. anonymous\]: In that case, an anonymous market might indeed reveal said information because betting on an anonymous market is low threat vs. openly speaking up. On the other hand, making markets credibly anonymous would add additional overhead.

### The market must not have too large negative side-effects, such as costs to the company's dynamics and morale

Prediction markets, being an infrequently used tool, may also have unintentional negative side effects:

1. As mentioned in the above section, prediction markets may be generally too politically disruptive, lacking the social intelligence to present their conclusions in the right way. To add a few specific examples: 
  *  It interferes with coalition politics. To quote Robin Hanson: "You have a couple of job candidates and you want to hire the best one for the company supposedly right? Well, I think actually when a person volunteers to be in charge of a hiring committee they don’t intend to pick the best person for the company. They intend to pick the best person for their coalition in the company. Forcing these metrics of who is best for the company would interfere with their plan to pick someone decent for the company but even better for their coalition. \[...\] You would be uncomfortable setting up this process that didn’t give you the flexibility to pretend to do A while really doing B."
  *  There is also general cultural pressure. Per [The Elephant in the Brain](https://en.wikipedia.org/wiki/The_Elephant_in_the_Brain) our behaviours and beliefs are optimized for living in a social group and very often we are self-deceptive and unaware of our motives. Prediction markets might reveal these hypocrisies: the disagreement between the likely outcome and our stated goals. This has the potential to be very unpleasant and awkward for everyone involved.
    *  For example, sometimes companies hire consultants to deal with workplace problems. These consultants sometimes serve more like a face-saving device than an actual solution (which might be fine). Having a prediction market around their effect will be awkward. \[cf. consultants\]
2. Prediction markets threaten the hierarchical control of top managers. It would become too obvious that most managers are less capable of planning/predicting the future than their confident behaviour suggests. 
  *  And this behaviour might not be irrational, bad or harmful (if not revealed) as leadership is not about making accurate quantitative forecasts.
3. Prediction markets make a big chunk of the bettors into "losers." Yet within a company, morale is very important. Businesses proceed by soliciting feedback, and by reshaping their plans to pretend that everyone is on board and has an ego stake in the final outcome. Prediction markets make this coordination more difficult. Once people make bets, they start rooting for their bet to win and for the other bet to lose. They move away from maximizing the value of the firm and develop an oppositional mentality vis-a-vis other employees. Furthermore, it is disruptive to have a running tally on who are the winners and losers each day.
4. When reward systems are created, employees view them as a means to distribute further privileges to insiders and favourites. Prediction markets would be viewed the same way. Who else is going to win all those bets? Do corporations really need more insider favouritism? 
  *  For example, companies might have both quantitatively oriented employees and non-quantitatively oriented ones, and adding predictions markets might be perceived or construed as a way to dole out rewards to the former.

 \[cf. consultants\]: See also: [Too Much Consulting?](https://www.overcomingbias.com/2012/01/why-so-much-consulting.html)

## Other Information Aggregation Mechanisms

### External platforms

Forecasting tournaments such as Metaculus, forecasting services such as Good Judgment or Maby forecasting, or public prediction markets such as Polymarket seem like a superior alternative to internal prediction markets for questions that don't require large amounts of context.

Metaculus is cheaper; a tournament or many questions costs on the order of $1000 to $5000, and may provide access to the inner reasoning of forecasters if incentives are structured correctly. However, quality and engagement may depend on the ability to present interesting questions to forecasters. Metaculus can produce somewhat sensible long-term forecasts.

Polymarket is much more expensive, on the order of $500 to $3000 on "sacrificed" liquidity needed to entice trades per question. So far, it is also limited to short term questions. However, the concept of winning money attracts very sharp sharks to even otherwise extremely dry questions. They are only just experimenting with including external questions.

Hypermind is a French forecasting platform/prediction market whose chief advantage is that its legal status is clear. Kalshi is based in the US, and its legal status is also clear, but is newer and thus might see lower volume. FTX might be particularly suitable for hedging for relatively large amounts, but it's unclear how one would go about coordinating this.

Good Judgment offers a degree of legibility, forecaster professionalism and experience which is not really available anywhere else. In particular, variance in forecast quality is likely to be lower—Good Judgment forecasts are likely to be of consistently high quality. However, it is much more expensive and inflexible. They may also be slightly overconfident. Forecasts can be confidential.

Other forecasting consultancies such as Maby forecasting, Azul Foresight or our own Samotsvety Forecasting are likely to be cheaper, but also more capacity constrained and less formidable than Good Judgment (but still fairly formidable). For instance, the two Maby forecasters were top Good Judgment superforecasters, but are still likely to be somewhat less effective in the absence of a broader forecasting team. They are also likely to have their own strengths; I view Azul Foresight as being strong on red-team aspects of forecasting.

### Specialized machine learning/data-analysis systems

Unlike prediction markets, startups that aim to predict a very narrow slice of reality have seen some success and high valuations. By restricting predictions to only one specific metric, current ML/data analysis pipelines can be thrown at that metric and the marginal cost of an additional prediction becomes very low. This is in contrast to prediction markets run by humans, in which the marginal cost of new predictions generally remains high.

Some instances of this kind of service in the management space:

*  Evidence Based Scheduling: [FogBugz](https://fogbugz.com/evidence-based-scheduling/), [Liquid Planner](https://www.liquidplanner.com/). See also: [Evidence Based Scheduling](https://www.joelonsoftware.com/2007/10/26/evidence-based-scheduling/) from Joel on Software.
*  Employee retention/turnover: [Peakon](https://peakon.com/solutions/experience-and-retention/) specializes in this. [This guide](https://bonus.ly/employee-retention-guide/employee-retention-tools) contains a guide to employee retention tools by a company that offers an adjacent service.
*  Sales/revenue forecasting. Searching for those keywords in Google uncovers plenty of alternatives. As a highlight, [People.ai](https://people.ai/) recently raised $100M on a $1.1 billion valuation, and was acquired by SalesForce. [HSBC](https://www.business.hsbc.uk/en-gb/tomorrow-ready-programme/optimise-your-cash-flow-and-finance) also offers cash flow forecasting.

These are generally SaaS startups. Once they get started, they can improve their predictive power by learning from their different clients, and give each client better predictions than if each only had access to their own data. On the negative side, they often come bundled with other products, like general HR software, and they might make clients dependent on the startup.

On the more elaborate end, the [Makridakis 5](https://en.wikipedia.org/wiki/Makridakis_Competitions#Fifth_competition,_started_on_March_3,_2020,_ended_on_July_1,_2020) competition can also be viewed as a [sanity check](https://www.sciencedirect.com/science/article/abs/pii/S0169207021001023?via%3Dihub) on Wallmart's own sales forecasting methodology, and as a more expensive and elaborate forecasting product.

### Internal forecasting competitions

Forecasting competitions are the main alternative to prediction markets. Unlike prediction markets, they don't have to be zero-sum competitions, but on the other hand, incentives are more loosely aligned. Nonetheless, most of the analysis above about the lack of adoption inside companies still applies.

Right now, the platforms which can support ongoing forecasting would be Cultivate Labs, Foretold, and maybe private instances of Metaculus or of Maby's platform. If money is not much of a concern, Cultivate Labs would be convenient, and does have the capacity to provide "forecasting competitions as a service". Otherwise, depending on in-house forecasting capabilities, creating one's own forecasting platform might be possible. Based on our experience with Foretold, I'd recommend against it because it is easy to underestimate the amount of hassle.

Note that—unlike prediction markets—[forecasting platforms](https://arxiv.org/abs/2106.11248) are very difficult to incentivize correctly. In particular, rewarding only the top forecasters incentivizes high-variance strategies. And although the hope for forecasting competitions might be that one also gets access to the comments, this is normally strictly disincentivized by scoring rules.

Further, if one tries to have a proper scoring rule \[cf. proper-scoring\] for a forecasting competition, this in our experience ends up producing dynamics that are not too different from prediction markets. One can also take care to design cooperative scoring rules, but these are generally underdeveloped and there isn't much literature studying their effects. There is also a tension between incentivizing effort and incentivizing collaboration.

Lastly, consider that forecasting platforms like Metaculus or Good Judgment Open are paying significantly below market rate, or getting essentially free labour. In the case of Metaculus, this is done by providing a product in some ways better than PredictionBook for keeping track of one's predictions, by incentivizing users with reputational benefits, and by appealing to users' altruistic motives. In the case of Good Judgment Open, users are enticed by the chance to attain the title of "superforecaster". It's unclear whether forecasting competitions within a company would be able to replicate a similarly strong appeal without recourse to monetary incentives.

**Judgemental forecasting might be better than prediction markets in some cases**

In Superforecasting Tetlock writes:

> The results were clear-cut each year. Teams of ordinary forecasters beat the wisdom of the crowd by about 10%. Prediction markets beat ordinary teams by about 20%. And superteams beat prediction markets by 15% to 30%. I can already hear the protests from my colleagues in finance that the only reason the superteams beat the prediction markets was that our markets lacked liquidity… they may be right. It is a testable idea, and one worth testing.

We guess that this might be due to prediction markets having much higher overhead than judgemental forecasting. On prediction markets, winning/losing is instantaneous, so markets give a premium on constant vigilance. Vigilance is stressful and takes a lot of attention, which is especially detrimental to the productivity of cognitive workers.

A forecaster, on the other hand, can update his forecast every other week, and doesn't lose many points due to not responding to breaking news fast enough. Or, if a fraction of forecasters catches any particular new development on time, this can be incorporated into the aggregate earlier by weighing more recent forecasts more.

 \[cf. proper-scoring\]: I.e., one in which inputting one's true belief is incentivized.

### Delphi Method

The Delphi Method is a well-known method for collaborative forecasting. Its chief characteristics are:

*  Participants are anonymous to each other
*  There are a series of rounds in which:
  *  Participants write down their quantified perspectives and/or the reasons for these.
  *  Participants read other's perspectives and/or a statistical summary thereof.
*  and then make their all-things-considered forecast

This is thought to help counteract various biases of normal face-to-face meetings—e.g., more extroverted or self-confident participants dominating the conversation. For a brief introduction, see pages 7 to 13 of "Analysis of the Future: The Delphi Method", an early paper by a RAND corporation researcher on the topic.

### Automatic Prediction Markets, Pseudo Prediction Markets

An automatic prediction market is one in which participants merely enter their honest probabilities, and their bets are calculated based on that probability. This allows one to track the accuracy of the different participants, by looking at their increasing or decreasing budgets. It might also avoid the zero-sum strategic trading aspect of current prediction markets if all participants bet against the house (more technically, if each participant bets against an automated market-maker whose liquidity is given by the house.)

For instance, in current prediction markets, a common trading pattern is to flip shares close to $0, e.g, to buy at $0.02 and sell at $0.04, even if one ultimately strongly suspects that the underlying event will not happen. This might be profitable, but may not contribute to increased market accuracy, and, crucially, would not happen in an automatic prediction market design.

The Kelly criterion can be automatically used as the rule for automatically determining betting amounts. Per [Learning Performance of Prediction Markets with Kelly Bettors](https://arxiv.org/abs/1201.6655), the market price in a market with participants willing to bet according to the Kelly criterion corresponds to the average of participants' private forecasts, weighted by their budget.

As some participants do better, their opinion is counted more within the aggregate. This automatic prediction market design could be combined with the Delphi method to produce a setup like the following:

*  Participants reveal their initial predictions (and these are recorded)
*  Participants outline the reasons for their predictions, with each participant being allocated an equal amount of time.
*  Participants update their predictions (and these are recorded)
*  Perhaps secretly: A prediction market is simulated using the methods from [Learning Performance of Prediction Markets with Kelly Bettors](https://arxiv.org/pdf/1201.6655.pdf). This produces a principled probability aggregate that takes into account the past accuracy of participants.

For a slower version more faithful to the original Delphi method, participants can anonymously write down both their predictions and their reasoning, and then anonymously update after reading everyone else's reasoning. One would then run two prediction markets for all questions, the one before reading others' probabilities and reasoning, and the one after.

### Low-tech options: Surveys and Interviews

Besides prediction markets, forecasting tournaments or Delphi methods, the option of doing very low-tech stuff—like sending out surveys or asking people what their informed opinion is on some—topic, still remains. One can also just hang a large post-it sheet in a public place and use a simple system like [BitBets](https://forum.effectivealtruism.org/posts/B8aWmfnWESxSQCxEL/bitbets-a-simple-scoring-system-for-forecaster-training) to give out rewards.

With regards to surveys specifically, [this](https://forum.effectivealtruism.org/posts/DCcciuLxRveSkBng2/a-review-of-two-books-on-survey-making) might be a good introduction to biases and pitfalls to take into account when designing surveys. We would also recommend abandoning Google Surveys and using [TypeForm](https://www.typeform.com/), which has beautiful typography and more functionalities.

## Conclusion

We:

*  Considered the academic consensus on prediction markets
*  Covered the corporate track record of prediction markets
*  Discussed hypotheses for the lack of adoption of prediction markets
*  Discussed alternative information aggregation mechanisms

At this point, because we do not get to observe the reasons for the lack of adoption of prediction markets, and because we are analyzing a complex question, we have slightly different overall views:

### Nuño Sempere

Nuño thinks that Upstart should not jump to using prediction markets, because the reasons for their failure are poorly understood. That is, we are much more sure about the fact that prediction markets fail to gain adoption than about any particular explanation of why this is.

To some extent, this conclusion depends on to what extent one thinks that something like the [efficient market hypothesis](https://en.wikipedia.org/wiki/Efficient-market_hypothesis) holds in management techniques. And I tend to think that it holds enough to not make prediction markets an attractive intervention \[cf. misha-disagrees\]. In particular, I think that if prediction markets were as good as advertised, they would have been adopted after initial trials, and then quickly spread. More specifically, I think that if prediction markets/forecasting competitions among employees were worth it, companies such as Google or Microsoft—which we know have experimented with prediction markets—would already be using them. I'm also very hesitant about suggesting interventions that may disrupt (in a negative way) current mechanisms.

Moreover, I observe that machine-learning or model-based or data-analysis solutions on forecasting weather, pandemics, supply chain, sales, etc. are happily adopted, and the startups that produce them reach quite high valuations. When trying to explain why prediction markets are not adopted, this makes me favor explanations based on high overhead, low performance and low applicability over Robin Hanson-style explanations based on covert and self-serving status moves.

Although prediction markets seem like an interesting technology, developing them to maturity seems too costly for a company centered on something else. Going on on a tangent, I'm also not sure that a startup that developed prediction markets would be able to capture enough of a percentage of the value of decision improvements to become profitable.

Many of the reasons for the lack of adoption of prediction markets also seem like they would apply to forecasting tournaments with proper incentives, unless one takes particular care to make them collaborative.

The above is more strongly worded than Misha and Eli's more nuanced perspectives below, with which I also partially agree. This might be a reaction of being in shock at the strong contrast between the optimism in the academic literature and the realities of prediction markets being hard to usefully implement in practice.

Like Eli below, I am also in favour of starting with small interventions and titrating one's way towards more significant ones. In particular, I feel attracted to:

*  Delphi-like automatic prediction markets built on top of dead-simple polls, and/or
*  forecasting setups using collaborative scoring rules—e.g., in which all employees (perhaps including \[private\]) play against \[private\]'s initial guess.

 \[cf. misha-disagrees\]: Misha holds the opinion that the efficient market hypothesis obviously doesn't hold (a lot of good techniques are not used or widely known; management was and is improving a lot; there are a bunch of startups delivering value through improving management). And he feels that is not really relevant for head-to-head comparison between different management techniques.

### Misha Yagudin

Misha thinks that Upstart shouldn't adopt company-wide prediction markets because prediction markets are disruptive to the social fabric and their effects on morale and culture are unclear. Further, Upstart is fairly small and markets shine when there are a lot of self-selected skilled participants. It seems better to select good managers who are excellent at conventional tools.

At the same time, I think Upstart might want to consider niche markets or even better forecasting tournaments (because these, as far as I can tell, take less attention and might get fairly accurate). The following condition should be met for these forecasting systems to feel reasonable to me:

*  should not mess up the social fabric too much;
*  should target people who might be a good cultural fit
  *  once again not to disturb up social fabric too much and to get people who might be better at probabilistic reasoning
  *  e.g. engineers and especially data scientists
*  should focus participants attention on something worthy

First, from Microsoft report and other sources, we know that prediction markets are quite accurate at predicting how long something will take.

If there are teams that are already using flexible approaches to development (like Scrum), it might be a fun idea to try a market or forecasting competition on top of it. This market is likely to be fairly active as new information comes regularly. And it is pretty fair because everyone syncs during stand-ups. Further, it's unlikely to affect managers in any way because they are already exposed by effort estimates produced via Scrum.

Second, another Microsoft success was with predicting something fuzzy like external product reviews. This is an understandable success because employees have different perspectives and contributed to different aspects of the product.

Something similar might be found in Upstart's workflow, for example, markets about how valuable new data sources will be (in terms of improved metrics). Further, this is what developers should care about; and developing a sense for better data seems beneficial.

### Eli Lifland

Eli feels fairly uncertain about what Upstart should do since this is a complicated issue. He would be more certain if he had a better model backed by more evidence on why existing corporate prediction markets have usually disappeared.

His best guess for what Upstart should do is:

1. To the extent that Upstart is interested in public predictions like macroeconomic trends and thinks more perspectives could benefit, solicit forecasts from platforms like Metaculus.
2. For internal predictions, start with interventions that take the least amount of employee time and are the least likely to damage morale then work your way up from there as it continues to be helpful.

*  For example, maybe start with using the Delphi method for one-time predictions on \[private\], or the value of projects for trying to improve Upstart's ML model.
*  If these go well, consider expanding to a small prediction market with similarly benign questions, before adding questions about e.g. project timelines.

1. A lot of the value that prediction markets could provide may be achieved through cultivating a culture of making quantified predictions to justify decisions and encouraging feedback/input from all members of the team on these types of predictions. You may already be doing this. This could be an 80/20 solution that brings a lot of the benefits without as much downside risk.

One uncertainty he has is about how strongly to take the weight of the evidence that many companies have started using prediction markets but then stopped; it's unclear to what extent this was because prediction markets actually don't provide more costs than benefits, versus other factors such as managers not prioritizing maintaining them due to undermining the managers' value or the improvements being hard to measure.

One's belief about this is likely correlated to one's belief about how close to maximally efficient for-profit businesses are given the efficient market hypothesis. Eli feels they are fairly inefficient but is not sure whether they are efficient to such an extent that the lack of adoption of prediction markets is enough evidence to recommend not trying them out.
