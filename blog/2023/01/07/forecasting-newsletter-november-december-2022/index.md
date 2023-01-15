Forecasting Newsletter for November and December 2022
================

## Highlights

- PredictIt nears its [probable demise](https://insightprediction.com/m/18505/will-predictit-survive)
- American Civics Exchange offers [political betting](https://otc.amciv.com/) to Americans using weird legal loophole
- Forecasting community member Avraham Eisenberg [arrested for $100M+ theft](https://www.coindesk.com/policy/2022/12/27/mango-markets-exploiter-eisenberg-arrested-in-puerto-rico/) ([a](https://web.archive.org/web/20230106173959/https://www.coindesk.com/policy/2022/12/27/mango-markets-exploiter-eisenberg-arrested-in-puerto-rico/)).
- Forecasting Research Institute [launches publicly](https://ea.greaterwrong.com/posts/kEd5qWwg8pZjWAeFS/announcing-the-forecasting-research-institute-we-re-hiring)
- Blogpost [suggests that GiveWell use uncertainty](https://www.metacausal.com/givewells-uncertainty-problem/), wins $20k
- Contrarian [offers $500k bet](https://stevekirsch.substack.com/p/1m-bet-rules), then chickens out
- Walter Frick writes a [resource to introduce journalists to prediction markets](https://www.niemanlab.org/2022/12/journalists-wake-up-to-the-power-of-prediction-markets/)
- Yearly feedback form

## Index

- Prediction Markets and Forecasting Platforms
  - PredictIt
  - American Civics Exchange
  - Polymarket
  - Insight Prediction
  - Kalshi
  - INFER
  - Metaculus
  - Manifold
  - Odds and Ends
- Opportunities
- Research
  - Shortform
  - Longform
- In the News
- Extra: The tax-evaded Nazi money which influences German politics and US/NATO relationships has opened a new forecasting platform.
- Yearly housekeeping

Browse past newsletters [here](https://forecasting.substack.com/) ([a](http://web.archive.org/web/20230105001738/https://forecasting.substack.com/)), or view this newsletter on the EA forum here. If you have a content suggestion or want to reach out, you can leave a comment or find me on [Twitter](https://twitter.com/NunoSempere). 
## Prediction Markets and Forecasting Platforms

### PredictIt

Racing against the clock, PredictIt is trying to get permission from the CFTC to continue existing. From PredictIt’s newsletter:

> As you may know, the Commodity Futures Trading Commission (CFTC) decided in August to revoke the authorization that allowed PredictIt to operate [effective Feb. 15, 2023](https://www.predictit.org/platform-announcements) ([a](http://web.archive.org/web/20221223063826/https://www.predictit.org/platform-announcements)) ([original CFTC source](https://www.cftc.gov/PressRoom/PressReleases/8567-22) ([a](http://web.archive.org/web/20221217222553/https://www.cftc.gov/PressRoom/PressReleases/8567-22))). More than a dozen of your fellow PredictIt traders, as well as academic data users and PredictIt’s parent company Aristotle, have challenged this decision in federal court with the goal of keeping election markets open and operating past the CFTC’s arbitrary Feb. 15, 2023, deadline.
> 
> Given the time-sensitive nature of this fight, the [plaintiffs filed a motion](https://www.casino.org/news/predictit-to-seek-cftc-injunction-in-appeals-court-after-lower-court-fails-to-act/) ([a](http://web.archive.org/web/20230103080014/https://www.casino.org/news/predictit-to-seek-cftc-injunction-in-appeals-court-after-lower-court-fails-to-act/)) last week to move the case to the United States Court of Appeals for the Fifth Circuit, which is the quickest path towards attaining justice for our traders and academic partners. We understand that legal battles do not generally make for exhilarating newsletter content, but as valued members of our community, we want you to know that we’re fighting on every available battlefield for the future of PredictIt and political prediction markets.
> 
> To that end, Aristotle, has also applied to the CFTC for permission to operate a Designated Contract Market (DCM), which will offer event contracts that include certain political outcomes. We’re well along in that process and hope to have approval in the coming months.
> 
> If approved, PredictIt Exchange will list many of the same contracts that you love: polling averages, nominations, confirmations and legislative outcomes, among other topics.
> 
> This next step toward becoming a regulated entity has been a long-time goal of the PredictIt team. We’ll be in touch soon with more information about how to sign up for PredictIt Exchange and other exciting developments in the works for next year. Stay tuned\!
 
### American Civics Exchange

The newly created American Civics Exchange is offering [over-the-counter betting](https://otc.amciv.com/dashboard) ([a](http://web.archive.org/web/20221122165636/https://otc.amciv.com/dashboard)) on political events.

I'm not clear on how [this is legal](https://en.wikipedia.org/wiki/Over-the-counter_%28finance%29) ([a](https://web.archive.org/web/20230106180304/https://en.wikipedia.org/wiki/Over-the-counter_%28finance%29)), but readers might consult their "What is ACE otc's regulatory framework?" section in their [FAQ](https://otc.amciv.com/page/faq) ([a](http://web.archive.org/web/20230105001904/https://otc.amciv.com/page/faq)). 

The man behind the thing is [Filip Pidot](https://twitter.com/FlipPidot) ([a](https://web.archive.org/web/20230106192455/https://twitter.com/FlipPidot)), who previously ran the [Old Bull TV](https://www.youtube.com/@OldBullTV/videos) ([a](http://web.archive.org/web/20230105002037/https://www.youtube.com/@OldBullTV/videos)) youtube channel, and is generally considered a respectable community member.

### Polymarket

Polymarket deploys an [adapter](https://github.com/Polymarket/uma-ctf-adapter) ([a](http://web.archive.org/web/20221013234314/https://github.com/Polymarket/uma-ctf-adapter)) between its Keynesian beauty contest resolution mechanism, [UMA](https://umaproject.org/) ([a](http://web.archive.org/web/20230104180400/https://umaproject.org/)), and actual resolutions. There is an "emergencyResolve" feature, which judging by the name could be used to override or bypass the UMA resolution. This is marginally concerning, because it could be used to siphon money away from markets, by betting on the less likely outcome and then resolving to it.

![Image of the Polymarket-UMA adapter](https://i.imgur.com/vsep53y.jpg)

Ultimately, I don't think this matters very much, because if Polymarket was compromised, users' wallets could probably be drained through some change in the webpage. Still, this downgrades UMA from its previous status as a source of truth to merely a convenient resolution mechanism,

### Insight Prediction

Insight Prediction [introduced](https://twitter.com/Insight_Predict/status/1587428761456193536) ([a](http://web.archive.org/web/20230105001841/https://twitter.com/Insight_Predict/status/1587428761456193536)) some innovations around "sudden death markets". It's difficult to do market-making for markets about discrete events, because if the event happens, traders will quickly make bets and rake in profit before the market-maker notices. And so:

> what we've implemented is the ability to have markets where we return all bets after a particular event has transpired. Bets placed before the event will be honored, while bets placed after will be returned. This of course requires bets in the market to be locked (72 hours)

You can see this functionality in this market: [Will Elon Musk Father Another Child Before August 2023?](https://insightprediction.com/m/67960/will-elon-musk-father-another-child-by-august-2023) ([a](http://web.archive.org/web/20230105001846/https://insightprediction.com/m/67960/will-elon-musk-father-another-child-by-august-2023)) market.

### Kalshi

Kalshi [updates their API](https://kalshi.com/blog/apiv2-is-here/) ([a](https://web.archive.org/web/20230106192534/https://kalshi.com/blog/apiv2-is-here/)) and, separately, gets [lampooned in Hacker News](https://news.ycombinator.com/item?id=33696486) ([a](https://web.archive.org/web/20230106174123/https://news.ycombinator.com/item?id=33696486)):

> That's neat but how does one hedge a $800k loan with a max position of $25k? Idea seems spurious. The difference is substantially higher. Compare a pre-hike 2.8% with a present 6.8%. Your $25k hedge is going to be pointless.
> 
> > Before Kalshi, markets that allowed you to trade on economically relevant events were illegal or unregulated.
> 
> Totally untrue. PredictIt was legal and regulated and well-loved for many years, operating with the permission of the CFTC under a no-action letter. Then Kalshi hired former CFTC commissioner Brian Quintenz and soon, PredictIt was suddenly deemed by the CFTC to have committed still-unenumerated violations of the no-action letter.
> 
> This is a rotten way to do business and shame on YC for being involved in it.

### INFER

INFER has a few [quarterly reports](https://www.infer-pub.com/archive) ([a](http://web.archive.org/web/20230105002053/https://www.infer-pub.com/archive)) on their forecasts on Synthetic Biology, AI and Microelectronics. It also has a [$5.25k](https://www.infer-pub.com/future-bowl-tournament) ([a](http://web.archive.org/web/20230105002109/https://www.infer-pub.com/future-bowl-tournament)) tournament.

When Metaculus got a [$5.5M grant](https://www.openphilanthropy.org/grants/metaculus-platform-development/) ([a](http://web.archive.org/web/20230105002123/https://www.openphilanthropy.org/grants/metaculus-platform-development/)) earlier this year, I could mostly see this reverberating, them hiring more people, a few teams being created, people I was keeping tabs on getting hired, more forecasting contests happening, and so on. You know, $5M is a small fortune, and I try to listen to the pulse of these things. Manifold Markets got built on top of $1M.

So for a while now, it has been very unclear to me where the [$8M that INFER's parent organization got from Open Philanthropy](https://www.openphilanthropy.org/grants/applied-research-laboratory-for-intelligence-and-security-forecasting-platforms/) ([a](http://web.archive.org/web/20230105002137/https://www.openphilanthropy.org/grants/applied-research-laboratory-for-intelligence-and-security-forecasting-platforms/)) was going. Things just weren't very visibly different from what I would expect to see if that grant hadn't really happened. INFER's forecasting rewards continued to be pretty small.

Taking a page from [Andrew Gelman](https://statmodeling.stat.columbia.edu/) ([a](http://web.archive.org/web/20230101023816/https://statmodeling.stat.columbia.edu/)), who is extremely meek when pointing these kinds of issues out, there could be many explanations. For example, it's possible that it does add up, that infrastructure and putting these kinds of reports out does take that much money. Or that they are spending the money slowly and sustainably for the long term. Or that the University of Maryland administration takes a really large chunk of funds. I am also assuming that the $8M in two years for two platforms doesn't correspond to something in the very rough neighbourhood of $2M per year per platform, but it could be that that's not the case. It's also possible that the INFER forecasting sphere and the sphere that I pay the most attention to don't overlap as much, which could be a positive. Still, the difference between the funds they receive and the impact that I perceive as for a while a small mystery to me.

Now, in its yearly review, INFER [mentions](https://www.infer-pub.com/the-pub/season-review-2022) ([a](http://web.archive.org/web/20221222200606/https://www.infer-pub.com/the-pub/season-review-2022)) working on "an internal U.S. Government-wide site" for 2023. That does seem like it could absorb large amounts of funds. Still, personally, I'd love to see a bill of how much was spent where.

### Metaculus

Metaculus held a [Million Predictions Hackathon](https://scribe.froth.zone/announcing-metaculuss-million-predictions-hackathon-91c2dfa3f39) ([a](https://web.archive.org/web/20230106174232/https://scribe.froth.zone/announcing-metaculuss-million-predictions-hackathon-91c2dfa3f39)), and posted their [year in review](https://scribe.froth.zone/metaculus-year-in-review-2022-c92ff9cd87a0) ([a](https://web.archive.org/web/20230106174302/https://scribe.froth.zone/metaculus-year-in-review-2022-c92ff9cd87a0)). They are also hosting a [beginner tournament](https://www.metaculus.com/tournament/beginner-tournament/) ([a](http://web.archive.org/web/20221120092344/https://www.metaculus.com/tournament/beginner-tournament/)), which seems like a neat way of introduction to forecasting.

### Manifold

Vincent Luczkow [analyzes](https://github.com/vluzko/manifoldpy) ([a](http://web.archive.org/web/20221228201434/https://github.com/vluzko/manifoldpy)) whether Manifold Markets are well calibrated, and finds that they are:

![manifold calibration chart](https://i.imgur.com/06mU9fG.png)

A Manifold Community member analyzed the performance of various prediction platforms on the [US elections](https://firstsigma.substack.com/p/midterm-elections-forecast-comparison) ([a](http://web.archive.org/web/20221215030200/https://firstsigma.substack.com/p/midterm-elections-forecast-comparison)). I find that these analyses don't really take into account that political outcomes are massively correlated. And so in years in which Democrats do well, platforms with a Democratic bias will do well, and vice-versa. Still, this isn't totally uninformative.

Manifold adds swipe functionality to bet. I think this is pure evil genius, in the sense of being a very addictive skinner box, and I'm very impressed. You can try it [here](https://manifold.markets/swipe) ([a](http://web.archive.org/web/20221031121657/https://manifold.markets/swipe)).

Manifold now has [Android](https://play.google.com/store/apps/details?id=com.markets.manifoldhttps://play.google.com/store/apps/details?id=com.markets.manifold&pli=1) ([a](https://web.archive.org/web/20230106174337/https://play.google.com/store/apps/details?id=com.markets.manifoldhttps://play.google.com/store/apps/details?id=com.markets.manifold&pli=1)) and [iOS](https://apps.apple.com/us/app/manifold-markets/id6444136749) ([a](http://web.archive.org/web/20221208084400/https://apps.apple.com/us/app/manifold-markets/id6444136749)) apps.

### Odds and ends

The FTX Foundation, a major founder of innovative forecasting projects, exists no longer after the downfall of FTX. More on this in the next edition of this newsletter.

The Effective Altruism Forum has started a [subforum for forecasting and estimation](https://forum.effectivealtruism.org/posts/azL6uPcbCqBfJ37TJ/announcing-a-subforum-for-forecasting-and-estimation) ([a](https://web.archive.org/web/20230104122631/https://forum.effectivealtruism.org/posts/azL6uPcbCqBfJ37TJ/announcing-a-subforum-for-forecasting-and-estimation)). Right now it mostly has some seed discussions started by employees of the Center for Effective Altruism. 

I don't recommend that subforum for general forecasting discussion because I think that the EA Forum has too onerous "be nice" rules, whereas I tend to most like spaces that are maximally uncensorious and are willing to sacrifice other values in the pursuit of accuracy. For instance, Discord spaces from Metaculus, Polymarket & others have tended to be extremely unmoderated. 

Some comments from Good Judgment Open from the last long while: 

- [julick](https://www.gjopen.com/comments/1545611) ([a](https://web.archive.org/web/20230106174359/https://www.gjopen.com/comments/1545611)) summarizes some [Cambridge research on regime change](https://www.cambridge.org/core/journals/american-political-science-review/article/abs/war-and-the-fate-of-regimes-a-comparative-analysis/88056B8C9BE0413DB349086653B58520) ([a](http://web.archive.org/web/20220914032710/https://www.cambridge.org/core/journals/american-political-science-review/article/abs/war-and-the-fate-of-regimes-a-comparative-analysis/88056B8C9BE0413DB349086653B58520)) and points out that "For regimes that initiated wars and lost them the regime change by violence within 3 years is 60-80%. For an initiator that won the probability is 20-35%." 
- [Anneinak](https://www.gjopen.com/comments/1554036) ([a](https://web.archive.org/web/20230106174419/https://www.gjopen.com/comments/1554036)) gives a 15% to Twitter crashing by December 2023, and points to [this article](https://www.gjopen.com/comments/1554036) ([a](https://web.archive.org/web/20230106174419/https://www.gjopen.com/comments/1554036)): "If the \[EU\] Commission elects to designate Twitter a VLOP \[Very Large Online Platform\] under the DSA \[Digital Services Act\] the business will face an accelerated compliance timetable with oversight kicking in in February next year — rather than in February 2024 — and with a tougher set of requirements to assess and mitigate risks on its platform." 
- Some [interesting discussion](https://www.gjopen.com/questions/2659-between-9-december-2022-and-9-june-2023-will-north-korea-detonate-a-nuclear-device) ([a](https://web.archive.org/web/20230106174450/https://www.gjopen.com/questions/2659-between-9-december-2022-and-9-june-2023-will-north-korea-detonate-a-nuclear-device)) on whether North Korea will detonate a nuclear device, which includes and would most likely be a nuclear test. 
- Forecaster [belikewater](https://www.gjopen.com/comments/1568206) ([a](https://web.archive.org/web/20230106174513/https://www.gjopen.com/comments/1568206)) does a nice, textbook base rate calculation to conclude that the current Italian Prime Minister, Meloni, will probably not (10%) stop being PM by October 2023. 

A decentralized sports prediction market has gotten [circa $5M](https://www.theblock.co/post/194496/susquehanna-bets-on-decentralized-sports-prediction-market-frontrunner) ([a](https://web.archive.org/web/20230106174536/https://www.theblock.co/post/194496/susquehanna-bets-on-decentralized-sports-prediction-market-frontrunner)) in [funding](https://techfundingnews.com/frontrunner-secures-4-75m-to-create-decentralised-sports-prediction-market/) ([a](http://web.archive.org/web/20221224043943/https://techfundingnews.com/frontrunner-secures-4-75m-to-create-decentralised-sports-prediction-market/)).

Hypermind continues to host its [Russia Punished](https://prod.hypermind.com/ngdp/en/showcase2/showcase.html?sc=RUSSANCT1) ([a](https://web.archive.org/web/20230106174607/https://prod.hypermind.com/ngdp/en/showcase2/showcase.html?sc=RUSSANCT1)) series of questions.

Scott Alexander wrote a [Prediction Market FAQ](https://astralcodexten.substack.com/p/prediction-market-faq) ([a](https://web.archive.org/web/20230106174747/https://astralcodexten.substack.com/p/prediction-market-faq)) and is hosting a [prediction contest](https://astralcodexten.substack.com/p/2023-prediction-contest) ([a](http://web.archive.org/web/20230103085602/https://astralcodexten.substack.com/p/2023-prediction-contest)) with a few days left to enter. If someone calculates expected reward over expected time, please post it as a comment.

Web3 betting platform [loses $0.4M in hack](https://newsbtc.com/news/company/dexsport-loses-400000-in-hack-user-funds-remain-unaffected/) ([a](http://web.archive.org/web/20221211201410/https://newsbtc.com/news/company/dexsport-loses-400000-in-hack-user-funds-remain-unaffected/)).

[Against the Odds](https://againsttheodds.es/) ([a](http://web.archive.org/web/20230105125805/https://againsttheodds.es/)) is a platform available to Spaniards which offers the opportunity to earn around $500/month through "crossed bets", i.e., betting one thing in one prediction market and the opposite in another, in order to make a guaranteed profit. The service itself seems legit and comes recommended by a friend. But interested Spaniards should calculate whether this is worth it, e.g,. earning 7€ for 5 mins of work is 84€/hour, and in practice something like 50€/hour including time to set up things.

Forecasting community member Avraham Eisenberg was both despised and admired after pulling off a [$114M heist](https://www.coindesk.com/business/2022/10/15/114m-mango-markets-exploiter-outs-himself-returns-most-of-the-money/) ([a](http://web.archive.org/web/20221223162406/https://www.coindesk.com/business/2022/10/15/114m-mango-markets-exploiter-outs-himself-returns-most-of-the-money/)) on a prediction platform. Now, he has been [arrested in Puerto Rico](https://www.coindesk.com/policy/2022/12/27/mango-markets-exploiter-eisenberg-arrested-in-puerto-rico/) ([a](https://web.archive.org/web/20230106173959/https://www.coindesk.com/policy/2022/12/27/mango-markets-exploiter-eisenberg-arrested-in-puerto-rico/)), and will be facing trial. He remains 47th on the Metaculus leaderboard, as can be seen on [his profile](https://www.metaculus.com/accounts/profile/114222/) ([a](http://web.archive.org/web/20221225212754/https://www.metaculus.com/accounts/profile/114222/)).

## Opportunities

The [Forecasting Research Institute](https://forecastingresearch.org/) ([a](http://web.archive.org/web/20221222190557/https://forecastingresearch.org/)) is a new think tank which has some senior research talent. Their [research page](https://forecastingresearch.org/research) ([a](http://web.archive.org/web/20221214224323/https://forecastingresearch.org/research)) should give readers a good taste of what they are about; my impression is that they have a strong academic component and will be aiming to be a bit more institutional than other efforts in the space. I hesitate to mention this, but Phil Tetlock takes the perhaps somewhat ceremonial or advisory role of President.

[Here](https://ea.greaterwrong.com/posts/kEd5qWwg8pZjWAeFS/announcing-the-forecasting-research-institute-we-re-hiring) ([a](https://web.archive.org/web/20230106174029/https://ea.greaterwrong.com/posts/kEd5qWwg8pZjWAeFS/announcing-the-forecasting-research-institute-we-re-hiring)) is a post on the EA forum introducing them. They are currently [hiring](https://forecastingresearch.org/apply) ([a](http://web.archive.org/web/20221222190529/https://forecastingresearch.org/apply)) for research analyst ($70k/year), data analyst ($70k), content editor ($60k) and research assistant ($15/hour) positions.

Other people who are hiring are 

- [Metaculus](https://apply.workable.com/metaculus/) ([a](http://web.archive.org/web/20221209055644/https://apply.workable.com/metaculus/)) 
- [Blue Rose research](https://statmodeling.stat.columbia.edu/2023/01/06/blue-rose-research-is-hiring/) ([a](http://web.archive.org/web/20230106074310/https://statmodeling.stat.columbia.edu/2023/01/06/blue-rose-research-is-hiring/)) 
- Insight Prediction [was hiring](https://www.econjobrumors.com/topic/insight-prediction-is-hiring) ([a](https://web.archive.org/web/20230106174830/https://www.econjobrumors.com/topic/insight-prediction-is-hiring)), and maybe still is 

## Writeups

### Shortform

The US 2022 congressional elections happened. Some relevant pieces on
that:


- Arpit Gupta looks at the [relative performance](https://arpitrage.substack.com/p/loss-for-prediction-markets-in-the) ([a](http://web.archive.org/web/20230105002406/https://arpitrage.substack.com/p/loss-for-prediction-markets-in-the)) of PredictIt, 538, Manifold Markets and Metaculus in the 2022 US election. He concludes that Metaculus does best.  
- This substack also [compares PredictIt, 538 and Manifold Markets](https://mikesaintantoine.substack.com/p/scoring-midterm-election-forecasts) ([a](http://web.archive.org/web/20221126214755/https://mikesaintantoine.substack.com/p/scoring-midterm-election-forecasts)). 
- In postmortem analysis, FiveThrirtyEight [finds that a data aggregation error meaningfully changed their forecasts](https://fivethirtyeight.com/features/how-a-data-processing-error-changed-our-deluxe-forecast/) ([a](http://web.archive.org/web/20221201012643/https://fivethirtyeight.com/features/how-a-data-processing-error-changed-our-deluxe-forecast/)). 
- Tanae Rao collects [all trades on Polymarket for the 2022 US midterm elections](https://github.com/tanaerao/polymarket-midterms) ([a](http://web.archive.org/web/20221120230038/https://github.com/tanaerao/polymarket-midterms)). 

With regard to comparisons between platforms, I don’t think that those analyses are completely uninformative. But I'd be hesitant to draw conclusions from one electoral season, because election outcomes are correlated. So for instance it could be that PredictIt is biased in favour of Republicans, Metaculus is biased in favour of Democrats, and Democrats did better this time.  

In [Tracking the money flows in forecasting](https://forum.effectivealtruism.org/posts/ytHCpLbT6A4gxqH8s/tracking-the-money-flows-in-forecasting) ([a](http://web.archive.org/web/20230105002212/https://forum.effectivealtruism.org/posts/ytHCpLbT6A4gxqH8s/tracking-the-money-flows-in-forecasting)), I post some rough monetary valuation estimates for a large number of forecasting or prediction market organizations. The most out-of-place platform there is INFER, as reviewed above.

[This online presentation](https://csustan.csustan.edu/~tom/Lecture-Notes/Non-Ergodic/Non-Ergodic-Short/Non-Ergodic-Short.pdf) ([a](http://web.archive.org/web/20221115165242/https://csustan.csustan.edu/~tom/Lecture-Notes/Non-Ergodic/Non-Ergodic-Short/Non-Ergodic-Short.pdf)) gives a few examples of systems whose expected value is far apart from the most likely value. h/t [Adam Marblestone](https://twitter.com/AdamMarblestone) ([a](https://web.archive.org/web/20221206194732/https://twitter.com/AdamMarblestone)).

Gelman [points to a history of forecasting competitions](https://statmodeling.stat.columbia.edu/2022/11/12/history-of-time-series-forecasting-competitions/) ([a](https://web.archive.org/web/20221127082244/https://statmodeling.stat.columbia.edu/2022/11/12/history-of-time-series-forecasting-competitions/)).

Jaime Sevilla posts [Some research ideas in forecasting](https://forum.effectivealtruism.org/posts/ozPL3mLGShqvjhiaG/some-research-ideas-in-forecasting) ([a](https://web.archive.org/web/20230105002420/https://forum.effectivealtruism.org/posts/ozPL3mLGShqvjhiaG/some-research-ideas-in-forecasting)).

The great Robin Hanson proposes [clean politician bonds](https://www.overcomingbias.com/2022/11/clean-politician-bonds.html) ([a](https://web.archive.org/web/20221224124910/https://www.overcomingbias.com/2022/11/clean-politician-bonds.html)):

> Assume you will vote for 1 of 2 candidates for US Mayor, Governor, or House of Repr. Only 1 has posted $1M bond payable to anyone who proves their misbehavior re lies, bribes, or other cheating. On average, what is chance you vote for this candidate?

[This brief blogpost](https://breakingthemarket.com/the-most-misunderstood-force-in-the-universe/) ([a](http://web.archive.org/web/20221117115030/https://breakingthemarket.com/the-most-misunderstood-force-in-the-universe/)) looks at when the geometric average of returns is a misleading indicator.

Alvaro de Menard writes [Forecasting Forecasting](https://fantasticanachronism.com/2022/11/21/forecasting-forecasting/) ([a](http://web.archive.org/web/20221215042936/https://fantasticanachronism.com/2022/11/21/forecasting-forecasting/)), a short article looking at his perspective on the future of forecasting. Curiously, the article is fairly obsequious towards my forecasting group, Samotsvety.

There is an AI risk tournament whose results should be coming up soon that had superforecasters look at the chance of existential risk tournament due to AI. Participant Ege Erdil [writes](https://twitter.com/EgeErdil2/status/1594723484860665858) ([a](https://web.archive.org/web/20230106175129/https://twitter.com/EgeErdil2/status/1594723484860665858)):

> i was one of the people who participated in the AI forecasting thing you cite here, and the experience was unfortunately very bad due to various reasons
> 
> i would strongly push back on the idea that the people selected were in any sense "top forecasters"

Lao Mein writes [Human-level Diplomacy was my fire alarm](https://www.lesswrong.com/posts/AZHHEPYWvTovvtikz/human-level-diplomacy-was-my-fire-alarm) ([a](http://web.archive.org/web/20221202162217/https://www.lesswrong.com/posts/AZHHEPYWvTovvtikz/human-level-diplomacy-was-my-fire-alarm)).

Nikos Bosse is working on a [database of baserates](https://forum.effectivealtruism.org/posts/H7xWzvwvkyywDAEkL/creating-a-database-for-base-rates) ([a](http://web.archive.org/web/20221223122549/https://forum.effectivealtruism.org/posts/H7xWzvwvkyywDAEkL/creating-a-database-for-base-rates)).

A newsletter from Cultivate Labs highlights an article on [Crowd Forecasting and Effective Altruism](https://www.cultivatelabs.com/posts/crowd-forecasting-and-effective-altruism) ([a](http://web.archive.org/web/20221223123956/https://www.cultivatelabs.com/posts/crowd-forecasting-and-effective-altruism)).

Ben West looks at the [rate of fraud in YCombinator startups](https://ea.greaterwrong.com/posts/5mghcxCabxuaK4WTs/ycombinator-fraud-rates) ([a](https://web.archive.org/web/20230106175234/https://ea.greaterwrong.com/posts/5mghcxCabxuaK4WTs/ycombinator-fraud-rates)), and finds that "1-2% of $100M+ YCombinator-backed companies have faced serious allegations of fraud".

This [old Washington Post article](https://www.washingtonpost.com/news/wonk/wp/2016/03/18/how-the-likelihood-of-breaking-up-changes-as-time-goes-by/) ([a](http://web.archive.org/web/20211005190327/https://www.washingtonpost.com/news/wonk/wp/2016/03/18/how-the-likelihood-of-breaking-up-changes-as-time-goes-by/)) shows the annual risk of breakup for married and unmarried couples:

![Breakup rates through time for different groups](https://i.imgur.com/tvg9CAt.png)

Cultivate Labs' CEO [reviews](https://www.cultivatelabs.com/posts/2022-in-the-rearview) ([a](http://web.archive.org/web/20221223123956/https://www.cultivatelabs.com/posts/2022-in-the-rearview)) his company's activities in 2022. In particular:

> The INFER project run by UMD’s ARLIS will be announcing a co-sponsored forecasting tournament with a significant player in the AI space next month, and will also be formally launching a U.S. Government-wide forecasting platform in the spring.

The European Commission, an organ of the European Union, has [a page with economic forecasts and trends](https://commission.europa.eu/statistics/economic-forecasts-and-trends_en) ([a](http://web.archive.org/web/20221225144109/https://commission.europa.eu/statistics/economic-forecasts-and-trends_en)). Its economic forecasts are [here](https://economy-finance.ec.europa.eu/economic-forecast-and-surveys/economic-forecasts_en) ([a](http://web.archive.org/web/20221227042949/https://economy-finance.ec.europa.eu/economic-forecast-and-surveys/economic-forecasts_en)) and [here](https://ec.europa.eu/economy_finance/publications/european_economy/forecasts/index_en.htm) ([a](http://web.archive.org/web/20221101065157/https://ec.europa.eu/economy_finance/publications/european_economy/forecasts/index_en.htm)). It would be a neat research project to track results and check whether they are calibrated. It would also be easy to propose possible improvements: use range forecasts instead of point estimates, and use very liquid markets if you really care.

Here are [some forecasts on British tourism](https://www.visitbritain.org/2023-tourism-forecast) ([a](http://web.archive.org/web/20230104165249/https://www.visitbritain.org/2023-tourism-forecast)).

Pinterest has a flashy [Pinterest Predicts](https://business.pinterest.com/en-us/pinterest-predicts/) ([a](http://web.archive.org/web/20230104173909/https://business.pinterest.com/en-us/pinterest-predicts/)) webpage (methodology explained [here](https://business.pinterest.com/en-us/blog/how-we-forecast-trends-pinterest-predicts/) ([a](http://web.archive.org/web/20221228013359/https://business.pinterest.com/en-us/blog/how-we-forecast-trends-pinterest-predicts/))). Pinterest is used by people to plan stuff, and Pinterest can rely on the law of straight lines to predict that trends that are growing will continue to do so.

[Does (a measure of) CEO overconfidence affect management forecasting and subsequent earnings management?](https://sci-hub.se/https://onlinelibrary.wiley.com/doi/epdf/10.1111/1911-3846.12144) ([a](https://web.archive.org/web/20230104231627/https://sci-hub.se/https://onlinelibrary.wiley.com/doi/epdf/10.1111/1911-3846.12144)). A paper finds that it does.

The International Journal of Forecasting has a [special issue](https://www.sciencedirect.com/journal/international-journal-of-forecasting/vol/38/issue/4) ([a](https://web.archive.org/web/20230000000000*/https://www.sciencedirect.com/journal/international-journal-of-forecasting/vol/38/issue/4)) on the M5 competition.

The Wall Street Journal reports on how companies use [uncertainty and rising energy prices as the justification for increasing price](https://www.wsj.com/articles/unilever-cfo-struggles-with-forecasting-accuracy-amid-high-inflation-11666904322) ([a](http://web.archive.org/web/20230105211045/https://www.wsj.com/articles/unilever-cfo-struggles-with-forecasting-accuracy-amid-high-inflation-11666904322)). You could take them at their word. On the other hand, in Spain, where food prices have risen a whole lot, some people view these as illegitimate cash grabs.

[This paper](https://www.thelancet.com/journals/lanpub/article/PIIS2468-2667%2822%2900203-1/fulltext) ([a](http://web.archive.org/web/20230105002641/https://www.thelancet.com/journals/lanpub/article/PIIS2468-2667%2822%2900203-1/fulltext)) tries to forecast the burden on the Chinese health system by tracking a large sample of older people.

Amazon offers some [automated ML forecasting](https://techcrunch.com/2022/11/29/aws-brings-natural-language-forecasting-queries-to-quicksight-q) ([a](https://web.archive.org/web/20230000000000*/https://techcrunch.com/2022/11/29/aws-brings-natural-language-forecasting-queries-to-quicksight-q)), as does [Oracle](https://www.oracle.com/sk/artificial-intelligence/forecasting/) ([a](http://web.archive.org/web/20230105002702/https://www.oracle.com/sk/artificial-intelligence/forecasting/)).

### Longform

Jared Leibowich looks at [Modeling the End of Monkeypox](https://asteriskmag.com/issues/1/modeling-the-end-of-monkeypox) ([a](http://web.archive.org/web/20221122015655/https://asteriskmag.com/issues/1/modeling-the-end-of-monkeypox)) in a new EA magazine, [Asterisk](https://asteriskmag.com/) ([a](http://web.archive.org/web/20221212215202/https://asteriskmag.com/)).

[This blogpost](https://www.metacausal.com/givewells-uncertainty-problem/) ([a](http://web.archive.org/web/20221222172413/https://www.metacausal.com/givewells-uncertainty-problem/)) won $20k from GiveWell’s "Change our Mind Contest." It looks at how their analyses could be improved by considering the uncertainty in their estimates, and does this at some depth, providing examples and recommendations.

Prominent anti-vaccine personality Steve Kirsh offered [$500k bets](https://stevekirsch.substack.com/p/1m-bet-rules) ([a](http://web.archive.org/web/20221104130431/https://stevekirsch.substack.com/p/1m-bet-rules)) on the result of a public debate on whether vaccines save lives or not. [Rootclaim](https://www.rootclaim.com/) ([a](https://web.archive.org/web/20230104080248/https://www.rootclaim.com/)), a hardcore Bayesian analysis organization, put significant efforts into doing their homework, and [accepted the bet](https://blog.rootclaim.com/rootclaim-accepts-500000-challenge-on-covid-vaccine-safety-efficacy/) ([a](http://web.archive.org/web/20221129133112/https://blog.rootclaim.com/rootclaim-accepts-500000-challenge-on-covid-vaccine-safety-efficacy/)). Then Kirsh [dishonourably chickened out](https://stevekirsch.substack.com/p/1m-bet-rules) ([a](http://web.archive.org/web/20221104130431/https://stevekirsch.substack.com/p/1m-bet-rules)). A transcript of the negotiations can be found [here](https://www.skirsch.com/covid/SaarWilf.pdf) ([a](http://web.archive.org/web/20221224061743/https://www.skirsch.com/covid/SaarWilf.pdf)).

The Swift Centre has a report on [what would happen if Russia uses a nuclear weapon](https://www.swiftcentre.org/what-would-be-the-consequences-of-a-nuclear-weapon-being-used-in-the-russia-ukraine-war/) ([a](http://web.archive.org/web/20230104231909/https://www.swiftcentre.org/what-would-be-the-consequences-of-a-nuclear-weapon-being-used-in-the-russia-ukraine-war/)), and is accepting commissions for topics to forecasting. They can be contacted at info@swiftcentre.org.

Andrew Gelman [points](https://statmodeling.stat.columbia.edu/2022/11/04/new-version-of-posteriordb-a-database-of-posterior-distributions-for-evaluating-bayesian-computation-algorithms/) ([a](http://web.archive.org/web/20230104231944/https://statmodeling.stat.columbia.edu/2022/11/04/new-version-of-posteriordb-a-database-of-posterior-distributions-for-evaluating-bayesian-computation-algorithms/)) to [PosteriorDB](https://github.com/stan-dev/posteriordb) ([a](http://web.archive.org/web/20230105002803/https://github.com/stan-dev/posteriordb)) a database of posteriors.

Friend of the newsletter Pratik Chougule recently published [American Universities in the Middle East and U.S. Foreign Policy](https://brill.com/display/title/63192) ([a](http://web.archive.org/web/20230105002822/https://brill.com/display/title/63192)). It seems oddly specific, so I'd instead recommend his older & cheaper [How to Make Money from Political Predictions](https://www.amazon.com/How-Make-Money-Political-Predictions-ebook/dp/B01KEGLRJ2) ([a](https://web.archive.org/web/20230000000000*/https://www.amazon.com/How-Make-Money-Political-Predictions-ebook/dp/B01KEGLRJ2)).

Michael Story writes about [why he generally doesn't recommend internal prediction markets or forecasting tournaments to organisations](https://mwstory.substack.com/p/why-i-generally-dont-recommend-internal) ([a](https://web.archive.org/web/20230106175846/https://mwstory.substack.com/p/why-i-generally-dont-recommend-internal)).

## In the News

Walter Frick of [Nonrival](https://nonrival.pub/) ([a](http://web.archive.org/web/20221222000715/https://nonrival.pub/)) writes [Journalists wake up to the power of prediction markets](https://www.niemanlab.org/2022/12/journalists-wake-up-to-the-power-of-prediction-markets/) ([a](http://web.archive.org/web/20230105002710/https://www.niemanlab.org/2022/12/journalists-wake-up-to-the-power-of-prediction-markets/)), a short piece introducing the concept and resources to journalists.

The New York Times has an article ([a](https://archive.is/FLB4p)) on "Forecasting the Future of Election Prediction Markets". It has this nice tidbit:

> PredictIt is something of a hybrid. It is a joint venture. One partner is Victoria University of Wellington, a New Zealand institution. The other is Aristotle, a for-profit American political consulting, compliance, data and software company, **whose founder**, John Aristotle Phillips, first gained national attention in 1976 as “the A-bomb kid” — a Princeton undergraduate who **successfully designed an atomic bomb** in a physics class project.

I also find it interesting how the NYT article mostly cites academics rather than, e.g., [Sam Altman's comments to the CFTC](https://comments.cftc.gov/PublicComments/ViewComment.aspx?id=69699) ([a](https://web.archive.org/web/20230106194025/https://comments.cftc.gov/PublicComments/ViewComment.aspx?id=69699)) or other [interesting figures](https://comments.cftc.gov/PublicComments/CommentList.aspx?id=7311) ([a](https://web.archive.org/web/20230000000000*/https://comments.cftc.gov/PublicComments/CommentList.aspx?id=7311)).

Vegas Slots Online has a short article [CFTC Ready to Deny Political Betting Request From Kalshi as PredictIt Levies a Lawsuit](https://www.vegasslotsonline.com/news/2022/10/31/cftc-ready-to-deny-political-betting-request-from-kalshi/) ([a](http://web.archive.org/web/20221101221514/https://www.vegasslotsonline.com/news/2022/10/31/cftc-ready-to-deny-political-betting-request-from-kalshi/)).

El País, a major Spanish newspaper, [directly embedded prediction markets](https://english.elpais.com/usa/2022-11-01/us-election-forecast-republicans-will-control-the-house-the-senate-is-too-close-to-call.html) ([a](http://web.archive.org/web/20221208220725/https://english.elpais.com/usa/2022-11-01/us-election-forecast-republicans-will-control-the-house-the-senate-is-too-close-to-call.html)) in its reporting on the US election.

Newsweek, on [Marjorie Taylor Greene's Chances of Losing Georgia Re-election](https://www.newsweek.com/marjorie-taylor-greene-chances-losing-georgia-re-election-1756671) ([a](http://web.archive.org/web/20230106180010/https://www.newsweek.com/marjorie-taylor-greene-chances-losing-georgia-re-election-1756671)).

Fast Company has a lengthy profile on ["the wild subculture of traders who bet on elections"](https://www.fastcompany.com/90796054/inside-the-wild-subculture-of-traders-who-bet-on-elections) ([a](http://web.archive.org/web/20221109203752/https://www.fastcompany.com/90796054/inside-the-wild-subculture-of-traders-who-bet-on-elections)).

[The Atlantic](https://www.theatlantic.com/technology/archive/2022/11/political-betting-elections-forecasting-future/672113/) ([a](http://web.archive.org/web/20221202193500/https://www.theatlantic.com/technology/archive/2022/11/political-betting-elections-forecasting-future/672113/)): "In the eight years since PredictIt began, online sports gambling went from being outlawed nearly everywhere in the country to a booming industry... In this context, the crackdown on political betting seems somewhat silly... As it turns out, on another site, Insight Prediction, you can bet on PredictIt’s survival. The price of Yes, as of this writing, is 10 cents. Even if the CFTC follows through on its initial decision to shut the site down, it’s hard not to feel that PredictIt has, in some more meaningful way, already won."

The Financial Times reviews [a few of the wackier $1M investments](https://www.ft.com/content/9a3b4b78-b18b-4d1f-b28b-05a73525aa61) ([a](http://web.archive.org/web/20221227114544/https://www.ft.com/content/9a3b4b78-b18b-4d1f-b28b-05a73525aa61)) FTX made, including Manifold Markets.

Vox [reviews](https://www.vox.com/future-perfect/2022/12/31/23509593/2022-predictions-macron-bolsanaro-covid-inflation-unemployment-climate-change-olympics) ([a](http://web.archive.org/web/20230105091649/https://www.vox.com/future-perfect/2022/12/31/23509593/2022-predictions-macron-bolsanaro-covid-inflation-unemployment-climate-change-olympics)) its yearly numerical predictions and makes new ones.

### Extra: The tax-evaded Nazi money which influences German politics and US/NATO relationships has opened a new forecasting platform.

A smallish company [rises](https://www.campaignlive.co.uk/article/truth-comes-bertelsmann-admits-nazi-past/160317) ([a](http://web.archive.org/web/20230105002941/https://www.campaignlive.co.uk/article/truth-comes-bertelsmann-admits-nazi-past/160317)) to become a major German publisher during the [Nazi times](https://en.wikipedia.org/wiki/Bertelsmann#1933_to_1945). After World War II, the now disgraced father gives control of the company to one of his surviving sons. The son manages the business competently, and greatly expands it. In an attempted tax evasion scheme, he divides the company's stock into a small percentage of voting stock and a large percentage of non-voting stock which he gives to a private foundation—I think this is the same trick which failed for the [Ford Foundation](https://en.wikipedia.org/wiki/Ford_Foundation) ([a](http://web.archive.org/web/20230105003036/https://en.wikipedia.org/wiki/Ford_Foundation)).

Except that this time, the [trick](https://taz.de/!5137745/)[succeeds](http://www.bertelsmannkritik.de/index.htm) ([a](http://web.archive.org/web/20230105003044/http://www.bertelsmannkritik.de/index.htm)), and after the death of the founder, it allows the family to keep control of assets while minimizing tax exposure. But the logic of the scheme requires that the company give out as few dividends as possible—because they would then be trapped inside the foundation and be difficult to get out. Which is why Bertelsmann has greatly expanded and is now a $30B behemoth which owns [Penguin Random House](https://en.wikipedia.org/wiki/Random_House) ([a](http://web.archive.org/web/20221224150408/https://en.wikipedia.org/wiki/Random_House)) as one of many items in an absurdly [long list of assets](https://en.wikipedia.org/wiki/List_of_assets_owned_by_Bertelsmann) ([a](http://web.archive.org/web/20221031225743/https://en.wikipedia.org/wiki/List_of_assets_owned_by_Bertelsmann)). Amusingly, they also [have the same person](https://en.wikipedia.org/wiki/Brigitte_Mohn#Career) ([a](https://web.archive.org/web/20221114210641/https://en.wikipedia.org/wiki/Brigitte_Mohn#Career)) represent the interest of the family and of the foundation.

![Heimat wider Heimat cover](https://i.imgur.com/zTrM1SN.jpg)
*<br>Cover of the "Heimat, wider Heimat" book, which translates to something like "Motherland, oh Motherland again", published by Bertelsmann. Note dog whistle connotations of [Heimat](https://en.wikipedia.org/wiki/Heimat#Modern_history). The author, Gustav Schröer, is one of four authors that I could quickly find that were both published by Bertelsmann and signed the "[Declaration of Allegiance of German Writers](https://de.wikipedia.org/wiki/Gel%C3%B6bnis_treuester_Gefolgschaft) ([a](http://web.archive.org/web/20230105003103/https://de.wikipedia.org/wiki/Gel%C3%B6bnis_treuester_Gefolgschaft))" to Adolf Hitler. The others are Rudolf G. Binding, Otto Flake and Fritz Müller-Partenkirchen. The case of Otto Flake is particularly interesting, since Bertelsmann rescued him from poverty in 1958 by reprinting old books of his.*

Now, the Bertelsmann foundation has made a small contribution to the forecasting world by starting the [Range Forecasting](https://www.rangeforecasting.org/) ([a](http://web.archive.org/web/20230105003141/https://www.rangeforecasting.org/)) platform, which runs on Cultivate Labs software. Per a recent Cultivate Labs newsletter:

> Earlier this year, the Bertelsmann Stiftung and Bertelsmann Foundation, in partnership with Cultivate Labs, launched RANGE (Rethinking Assumptions in a New Geostrategic Environment) to forecast on topics critical to the transatlantic relationship. RANGE asks questions related to systemic rivalry with China, climate change, democracy, geopolitics, trade, and economics, all to gain a greater perspective of the transatlantic relationship and to provide analysis and recommendations for U.S. and European policymakers. 
> 
> At a critical time for strengthening global relations, the Stiftung and Foundation aim to engage an international audience in understanding how the transatlantic relationship affects us all, and how it can and should be shaped for the future. Read our recent announcement about the partnership and learn more at rangeforecasting.org.    RANGE was featured this week at the 2022 Futures Forum in Münster, Germany, a side event of the G7 minister’s meeting, which brought together American and German experts and rising leaders from civil society, academia, the private sector, and government to discuss transatlantic relations and the intersection of democracy and technology.
> 
> Earlier today at the Forum, Cultivate Labs CEO Adam Siegel spoke alongside Bertelsmann Foundation Executive Director Irene Braam about RANGE and the ways that crowd forecasting can help us better understand and strengthen transatlantic relations.
> 
I do think that this new forecasting platform might be of some use. And I want to express my continued high regard for Cultivate Labs as one of the few platforms that have the proven ability to bring forth a new platform into existence, and maintain it through time. But I guess that reputational laundering of billionaires' money is easier to notice when it is billionaires I don't care much about. And in this particular case, the provenance of the money seems so illegitimate, so rotten, as to be almost comical.

## Yearly housekeeping

[Here](https://forms.gle/L6kv9GKh5eSf2Kmb6) ([a](https://web.archive.org/web/20230106180045/https://docs.google.com/forms/d/e/1FAIpQLSd3eFDhdDReceLrSly4gWmGliHdCX7LNEJygGUEttkLi57M_Q/viewform?usp=send_form#short_url=https://forms.gle/L6kv9GKh5eSf2Kmb6)) is my yearly feedback form for this newsletter. I care most about readers' thoughts on whether this newsletter should be posted once a month or once every two months, and about freeform thoughts.

Also note that as the price I assign to my time rises, I'm becoming more ambivalent about whether to continue publishing this newsletter. If you have strong feelings about whether this newsletter should continue existing, consider becoming a paid subscriber, because a) I consider that a strong signal that I should continue writing this newsletter, b) I like the money, and c) the rate of current paid subscriptions is a pitiful [1.6%](https://substack.com/going-paid-guide#grow-paid-list) ([a](https://web.archive.org/web/20230000000000*/https://substack.com/going-paid-guide#grow-paid-list)).

---

Note to the future: All links are added automatically to the Internet Archive, using this [tool](https://github.com/NunoSempere/longNowForMd) ([a](http://web.archive.org/web/20230105003212/https://github.com/NunoSempere/longNowForMd)). "(a)" for archived links was inspired by [Milan Griffes](https://www.flightfromperfection.com/) ([a](http://web.archive.org/web/20230105003223/https://www.flightfromperfection.com/)), [Andrew Zuckerman](https://www.andzuck.com/) ([a](http://web.archive.org/web/20220316214638/https://www.andzuck.com/)), and [Alexey Guzey](https://guzey.com/) ([a](http://web.archive.org/web/20230101090040/https://guzey.com/)).

---

> If I had to put my money where my mouth is, my mouth wouldn’t be very busy. I’m just sayin’… — [Anonymous commenter](https://statmodeling.stat.columbia.edu/2022/10/27/you-can-read-for-free-but-comments-cost-money-or-is-it-the-other-way-around/) ([a](https://web.archive.org/web/20230000000000*/https://statmodeling.stat.columbia.edu/2022/10/27/you-can-read-for-free-but-comments-cost-money-or-is-it-the-other-way-around/)) 
