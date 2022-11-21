Some data on the stock of EA™ funding 
=====================================

### Overall Open Philanthropy funding

Open Philanthropy's allocation of funding through time looks as follows:

![Bar graph of OpenPhil allocation by year. Global health leads for most years. Catastrophic risks are usually second since 2017. Overall spend increases over time.](https://i.imgur.com/RwD1pP9.png)

Dustin Moskovitz's wealth looks, per [Bloomberg](https://www.bloomberg.com/billionaires/profiles/dustin-a-moskovitz), like this:

![Line chart of Dustin Moskovitz's wealth over time, with a dip in 2019 and a peak in 2021.](https://i.imgur.com/cObIgOQ.png)

If we plot the two together, we don't see that much of a correlation:

![Combination of the previous two charts. Moskovitz's fortune does not match changes in total spend or category composition.](https://i.imgur.com/NhhiLqd.png)

Holden Karnofsky, head of Open Philanthropy, [writes](https://forum.effectivealtruism.org/posts/mCCutDxCavtnhxhBR/some-comments-on-recent-ftx-related-events) that the Blomberg estimates might not be all that accurate:

> Our available capital has fallen over the last year for these reasons. That said, as of now, public reports of Dustin Moskovitz and Cari Tuna’s net worth give a substantially understated picture of our available resources. That’s because, among other issues, they don’t include resources that are already in foundations. (I also note that META stock is not as large a part of their portfolio as some seem to assume)

Edited to add: Moskovitz replies:

<blockquote class="twitter-tweet tw-align-center"><p lang="en" dir="ltr">Actually the Bloomberg tracker looks pretty close, though missing 3B or so of foundation assets. The Forbes one is like half the Bloomberg estimate 🤷‍♂️</p>&mdash; Dustin Moskovitz (@moskov) <a href="https://twitter.com/moskov/status/1594337871355207680?ref_src=twsrc%5Etfw">November 20, 2022</a></blockquote> <script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

In mid 2022, Forbes put Sam Bankman-Fried's wealth at [$24B](https://www.forbes.com/profile/sam-bankman-fried/?sh=706b96804449). So in some sense, the amount of money allocated to or according to Effective Altruism™ peaked somewhere close to $50B.

### Funding flow restricted to longtermism & global catatrophic risks (GCRs)

The analysis becomes a bit more interesting if we look only at longtermism and GCRs:

![Bar graph of OpenPhil allocation to catastrophic risks by year. AI leads most years, followed by biosecurity.](https://i.imgur.com/OZwHMtV.png)

In contrast, per [Forbes](https://web.archive.org/web/20221116022228/https://fortune.com/2022/11/14/balkman-fried-ftx-collapse-threatens-effective-altruism-billions-charity-philanthropy/), the FTX Foundation had given out $160M by September 2022. My sense is that most (say, maybe 50% to 80%) of those grants went to "longtermist" cause areas, broadly defined. In addition, SBF and other FTX employees led a $580M funding round for [Anthropic](https://www.privateequitywire.co.uk/2022/05/05/314319/ftx-ceo-leads-580m-series-b-round-anthropic)

### Further analysis 

It's unclear what would have to happen for Open Philanthropy to pick up the slack here. In practical terms, I'm not sure whether their team has enough evaluation capacity for an additional $100M/year, or whether they will choose to expand that. 

Two somewhat informative posts from Open Philanthropy on this are [here](https://forum.effectivealtruism.org/posts/HPdWWetJbv4z8eJEe/open-phil-is-seeking-applications-from-grantees-impacted-by) and [here](https://forum.effectivealtruism.org/posts/mCCutDxCavtnhxhBR/some-comments-on-recent-ftx-related-events)

I'd be curious about both interpretative analysis and forecasting on these numbers. I am up for supporting the later by e.g., committing to rerunning this analysis in a year.

### Appendix: Code

The code to produce these plots can be found [here](./.source/analysis.R); lines 42 to 48 make the division into categories fairly apparent. To execute this code you will need a working R installation and a document named [grants.csv](./.source/grants.csv), which can be downloaded from [Open Philanthropy's website](https://www.openphilanthropy.org/grants/).

<p><section id="isso-thread">
  <noscript>Javascript needs to be activated to view comments.</noscript>
</section></p>
