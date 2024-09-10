What's the chance your startup will succeed?
=================================================

### The baserate for all companies

The U.S. Bureau of Labor Statistics [shares data](https://www.bls.gov/bdm/us_age_naics_00_table7.txt) on the survivorship rates of a all companies. Around 30-35% die before the third year, and around 50% by the fifth year, with surprising regularity. Of all 4.66M companies founded in the year 2000, 2.7M remain, or 18.2%. This seems like a good starting point, or baserate.

The BLS also shares data [disaggregated by industry](https://www.bls.gov/bdm/bdmage.htm). For instance, for the [information](https://www.bls.gov/bdm/us_age_naics_51_table7.txt) or [technical services](https://www.bls.gov/bdm/us_age_naics_54_table7.txt) fields. But the numbers look pretty similar for all industry (with agriculture dying substantially less often and mining dying substantially more often)

![](./by-industry.png)

### The data is unreliable for Venture Capital

We can also restrict our analysis to startups that receive venture capital. But the data there is much more spotty. Still, a Wall Street Journal [article from 2012](https://www.wsj.com/articles/SB10000872396390443720204578004980476429190) ([archive link](https://archive.is/s0I9j)) mentions that:

> Of the 6,613 U.S.-based companies initially funded by venture capital between 2006 and 2011, 84% now are closely held and operating independently, 11% were acquired or made initial public offerings of stock and 4% went out of business, according to Dow Jones VentureSource. Less than 1% are currently in IPO registration.

Various other articles mention that ["9 out of 10 startups fail"](https://blog.privateequitylist.com/untitled-4/#:~:text=their%20innovative%20ideas.-,Failure%20Rates%20of%20Venture%2DBacked%20Startups,out%20of%20every%2010%20companies.), but these seem to be ultimately sourced from a [random 2014 Forbes article](https://fortune.com/2014/09/25/why-startups-fail-according-to-their-founders/) that doesn't cite sources. I wouldn't consider this number particularly reliable.

Paul Graham, who founded YC combinator, initially expected [a third of startups funded by YC combinator to succeed](https://paulgraham.com/die.html), though later mentions that [way more](https://www.paulgraham.com/swan.html) than 30% of the startups in their batches raise funding by "Demo day" (via [Benjamin Hoffmann](http://benjaminrosshoffman.com/approval-extraction-advertised-as-production/#Treatment_Effects)).

### Combine this with the inside view

You could also looks at your personal success rate for projects you've tried out throughout your life, and try to figure out if you're the kind of person for whom projects tend to succeed, or to fail. You could get a larger population to estimate that by looking at the success of projects or companies in your highschool circle, your family, your university group, friends, or even every young organization you've been involved with.

You could further decompose success into steps, giving explicit probabilities to each step, and then multiplying the probability of each step to arrive at the probability of success. Some generic broad initial steps could be getting to an MVP, raising enough funding or early revenue, and reaching sustainable revenue. 

You could also look at initial proxies of success, like the founding team, product market fit, difficulty or easiness raising funding, and so on. Maybe if these are all good, you might have a 66% chance of surviving by the fifth year, rather than the 50% for all companies. 

But a hard-earned lesson from forecasting is that the baserate, the historical frequency, is king. and one shouldn't move too much away from it.

### Then there is also the degree of success

Beyond that, you'll also want to know the magnitude of success; being alive making $10K a year is not the same as being alive making $10M a year. Data here is much more spotty, though maybe [number of employees](https://www.bls.gov/charts/county-employment-and-wages/establishments-by-size.htm) is a good proxy for profit. For instance, for the information industry, of the 280K companies, 76% have less than 5 employees, and 96% have less than 50 employees, though this includes all businesses throughout their lifetime. Again, this would be the approach top down bottom up size of industry.

### Initially successful nonprofits fail more slowly

The data sources for nonprofits is different. The [National Center for Charitable Statistics](https://urbaninstitute.github.io/nccs/) provides [data](https://urbaninstitute.github.io/nccs/datasets/core/) for US charities. However, its [best dataset](https://urbaninstitute.github.io/nccs/datasets/core/) excludes very small nonprofits:

> The Core Series does not include 990N ePostcard filers, even though they are the most common type of nonprofit. These organizations have annual revenues below $50,000, and the 990N form they file includes minimal information, including updates to the organization address and key contact.

In addition, even sometimes quite large projects don't sometimes have their own 501c3 status, but rather are fiscally sponsored by some other larger organization. This wouldn't show up in the data below.

So, excluding those, for organizations which at some point have over $50K in donations, and have registered as 501c3, the survival rates for organizations created in 2005 (chosen as a convenient year), are as follows:

| Year | Surviving establishments | Survival rate since birth |
| --- | --- | --- |
| 2005 | 25502 | 100% |
| 2006 | 25242 | 98% |
| 2007 | 23518 | 92% |
| 2008 | 20646 | 80% |
| 2009 | 19396 | 76% |
| 2010 | 17957 | 70% |
| 2011 | 17029 | 66% |
| 2012 | 15706 | 61% |
| 2013 | 15081 | 59% |
| 2014 | 15667 | 61% |
| 2015 | 15084 | 59% |
| 2016 | 14513 | 56% |
| 2017 | 14013 | 54% |
| 2018 | 13510 | 52% |
| 2019 | 12997 | 50% |

So they fall much slower than the dataset of all businesses, but it's not an apples to apples comparison, since we're not comparing all charities against all businesses. And here we also fall into seeing-like-a-state problems, where we aren't asking the question we want to ask, but rather the question which the data allows us to ask. 

Some alternatives might have been to look at the survivorship rate of organizations incubated by Charity Entrepreneurship, of organizations in Open Philanthropy's database. These could provide alternative reference classes, but would be less useful if one is not in fact incubated by Charity Entrepreneurship or funded by Open Philanthropy. We might also think that fiscally sponsored projects are drawn from a distribution that's closer to all businesses than to all 501c3-registered nonprofits, since 501c3 registration and compliance seems like it has more operations costs than setting up a business. 

### Conclusion

So we've looked at some reference classes, and found out that actuarial tables for new projects are steep, but not extremely steep. When applying these numbers to individual cases, there is probably a temptation to adjust them based on the inside view, on personal factors that make you confident that you're better than the masses. But forecasting practice suggests that deviating too much from the baserate is a mistake: 50% of new businesses die by the fifth year.

---

<p>
  <section id='isso-thread'>
  <noscript>javascript needs to be activated to view comments.</noscript>
  </section>
</p>
