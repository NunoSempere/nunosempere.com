Shapley values: Better than counterfactuals
==============

\[Epistemic status: Pretty confident. But also, enthusiasm on the verge of partisanship\]

One intuitive function which assigns impact to agents is the counterfactual, which has the form:

> CounterfactualImpact(Agent) = Value(World) - Value(World/Agent)

which reads "The impact of an agent is the difference between the value of the world with the agent and the value of the world without the agent".

It has been discussed in the effective altruism community that this function leads to pitfalls, paradoxes, or to unintuitive results when considering scenarios with multiple stakeholders. See:

*   [Triple counting impact in EA](https://forum.effectivealtruism.org/posts/fnBnEiwged7y5vQFf/triple-counting-impact-in-ea)
*   [The counterfactual impact of agents acting in concert](https://forum.effectivealtruism.org/posts/EP8x3vHRQJP57TjFL/the-counterfactual-impact-of-agents-acting-in-concert)

In this post I'll present some new and old examples in which the counterfactual function seems to fail, and how, in each of them, I think that a less known function does better: the Shapley value, a concept from cooperative game theory which has also been brought up before in such discussions. In the first three examples, I'll just present what the Shapley value outputs, and halfway through this post, I'll use these examples to arrive at a definition.

I think that one of the main hindrances in the adoption of Shapley values is the difficulty in its calculation. To solve this, I have written a Shapley value calculator and made it available online: [shapleyvalue.com](http://shapleyvalue.com/). I encourage you to play around with it.

## Example 1 & recap: Sometimes, the counterfactual impact exceeds the total value.

> Suppose there are three possible outcomes: P has cost $2000 and gives 15 utility to the world Q has cost $1000 and gives 10 utility to the world R has cost $1000 and gives 10 utility to the world
> 
> Suppose Alice and Bob each have $1000 to donate. Consider two scenarios:
> 
> Scenario 1: Both Alice and Bob give $1000 to P. The world gets 15 more utility. Both Alice and Bob are counterfactually responsible for giving 15 utility to the world.
> 
> Scenario 2: Alice gives $1000 to Q and Bob gives $1000 to R. The world gets 20 more utility. Both Alice and Bob are counterfactually responsible for giving 10 utility to the world.
> 
> From the world's perspective, scenario 2 is better. However, from Alice and Bob's individual perspective (if they are maximizing their own counterfactual impact), scenario 1 is better. This seems wrong, we'd want to somehow coordinate so that we achieve scenario 2 instead of scenario 1.  
> [Source](https://forum.effectivealtruism.org/posts/EP8x3vHRQJP57TjFL/the-counterfactual-impact-of-agents-acting-in-concert#9KJM54ydQiucy22Gy)  
> Attribution: rohinmshah

In Scenario 1:  
Counterfactual impact of Alice: 15 utility.  
Counterfactual impact of Bob: 15 utility.  
Sum of the counterfactual impacts: 30 utility. Total impact: 15 utility.

The Shapley value of Alice would be: 7.5 utility.  
The Shapley value of Bob would be: 7.5 utility.  
The sum of the Shapley values always adds up to the total impact, which is 15 utility.

In Scenario 2:  
Counterfactual impact of Alice: 10 utility.  
Counterfactual impact of Bob: 10 utility.  
Sum of the counterfactual impacts: 20 utility. Total impact: 20 utility.

The Shapley value of Alice would be: 10 utility.  
The Shapley value of Bob would be: 10 utility.  
The sum of the Shapley values always adds up to the total impact, which is 10+10 utility = 20 utility.

In this case, if Alice and Bob were each individually optimizing for counterfactual impact, they'd end up with a total impact of 15. If they were, each of them, individually, optimizing for the Shapley value, they'd end up with a total impact of 20, which is higher.

It would seem that we could use a function such as

> CounterfactualImpactModified = CounterfactualImpact / NumberOfStakeholders

to solve this particular problem. However, as the next example shows, that sometimes doesn't work. The Shapley value, on the other hand, has the property that it always adds up to total value.

Property 1: The Shapley value always adds up to the total value.

## Example 2: Sometimes, the sum of the counterfactuals is less than total value. Sometimes it's 0.

Consider the invention of Calculus, by Newton and Leibniz at roughly the same time. If Newton hadn't existed, Leibniz would still have invented it, and vice-versa, so the counterfactual impact of each of them is 0. Thus, you can't normalize like above.

The Shapley value doesn't have that problem. It has the property that equal people have equal impact, which together with the requirement that it adds up to total value is enough to assign 1/2 of the total impact to each of Newton and Leibniz.

Interestingly, GiveWell has Iodine Global Network as a standout charity, but not as a recommended charity, because of considerations related to the above. If it were the case that, had IGN not existed, another organization would have taken its place, its counterfactual value would be 0, but its Shapley value would be 1/2 (of the impact of iodizing salt in developing countries).

Property 2: The Shapley assigns equal value to equivalent agents.

## Example 3: Order indifference.

Consider Scenario 1 from Example 1 again.

> P has cost $2000 and gives 15 utility to the world.
> 
> Suppose Alice and Bob each have $1000 to donate. Both Alice and Bob give $1000 to P. The world gets 15 more utility. Both Alice and Bob are counterfactually responsible for giving 15 utility to the world.

Alice is now a pure counterfactual-impact maximizer, but something has gone wrong. She now views Bob adversarially. She thinks he's a sucker, and she waits until Bob has donated to make her own donation. There are no worlds in which he doesn't donate before her, and Alice assigns all 15 utility to herself, and 0 to Bob. Note that she isn't exactly calculating the counterfactual impact, but something slightly different.

The Shapley value doesn't consider any agent to be a sucker, doesn't consider any variables to be in the background, and doesn't care whether people try to donate strategically before or after someone else. Here is a perhaps more familiar example:

Scenario 1:  
Suppose that the Indian government creates some big and expensive infrastructure to vaccinate people, but people don't use it. Suppose an NGO then comes in, and sends reminders to people to vaccinate their people, and some end up going.

Scenario 2:  
Suppose that an NGO could be sending reminders to people to vaccinate their children, but it doesn't, because the vaccination infrastructure is nonexistent, so there would be no point. Then, the government steps in, and creates the needed infrastructure, and vaccination reminders are sent.

Again, it's tempting to say that in the first scenario, the NGO gets all the impact, and in the second scenario the government gets all the impact, perhaps because we take either the NGO or the Indian government to be in the background. To repeat, the Shapley value doesn't differentiate between the two scenarios, and doesn't leave variables in the background. For how this works numerically, see the examples below.

Property 3: The Shapley value doesn't care about who comes first.

## The Shapley value is uniquely determined by simple properties.

These properties:

*   Property 1: Sum of the values adds up to the total value (Efficiency)
*   Property 2: Equal agents have equal value (Symmetry)
*   Property 3: Order indifference: it doesn't matter which order you go in (Linearity). Or, in other words, if there are two steps, Value(Step1 + Step2) = Value(Step1) + Value(Step2).

And an extra property:

*   Property 4: Null-player (if in every world, adding a person to the world has no impact, the person has no impact). You can either take this as an axiom, or derive it from the first three properties.

are enough to force the Shapley value function to take the form it takes:

![](https://images.nunosempere.com/blog/2019/10/10/shapley-formula.svg)

At this point, the reader may want to consult [Wikipedia](https://en.wikipedia.org/wiki/Shapley_value) to familiarize themselves with the mathematical formalism, or, for a book-length treatment, [_The Shapley value: Essays in honor of Lloyd S. Shapley_](http://www.library.fa.ru/files/Roth2.pdf). Ultimately, a quick way to understand it is as "the function uniquely determined by the properties above".

I suspect that order indifference will be the most controversial option. Intuitively, it prevents stakeholders from adversarially choosing to collaborate earlier or later in order to assign themselves more impact.

Note that in the case of only one agent the Shapley value reduces to the counterfactual function, and that the Shapley value uses many counterfactual comparisons in its formula. It sometimes just reduces to CounterfactualValue/ NumberOfStakeholders (though it sometimes doesn't). Thus, the Shapley value might be best understood as an extension of counterfactuals, rather than as something completely alien.

## Example 4: The Shapley value can also deal with leveraging

> Organisations can leverage funds from other actors into a particular project. Suppose that AMF will spend $1m on a net distribution. As a result of AMF’s commitment, the Gates Foundation contributes $400,000. If AMF had not acted, Gates would have spent the $400,000 on something else. Therefore, the counterfactual impact of AMF’s work is:  
> AMF’s own $1m on bednets plus Gates’ $400,000 on bednets minus the benefits of what Gates would otherwise have spent their $400,000 on.  
> If Gates would otherwise have spent the money on something worse than bednets, then the leveraging is beneficial; if they would otherwise have spent it on something better than bednets, the leveraging reduces the benefit produced by AMF.  
> Source: [The counterfactual impact of agents acting in concert](https://forum.effectivealtruism.org/posts/EP8x3vHRQJP57TjFL/the-counterfactual-impact-of-agents-acting-in-concert).

Let's consider the case in which the Gates Foundation would otherwise have spent their $400,000 on something half as valuable.

Then the counterfactual impact of the AMF is 1,000,000+400,000-(400,000)\*0.5 = $1,2m.  
The counterfactual impact of the Gates Foundation is $400,000.  
And the sum of the counterfactual impacts is $1,6m, which exceeds total impact, which is $1,4m.

The Shapley value of the AMF is $1,1m.  
The Shapley value of the Gates Foundation is $300,000.

Thus, the Shapley value assigns to the AMF part, but not all, of the impact of the Gates Foundation donation. It takes into account their outside options when doing so: if the Gates Foundation would have invested on something equally valuable, the AMF wouldn't get anything from that.

## Example 5: The Shapley value can also deal with funging

> Suppose again that AMF commits $1m to a net distribution. But if AMF had put nothing in, DFID would instead have committed $500,000 to the net distribution. In this case, AMF funges with DFID. AMF’s counterfactual impact is therefore:  
> AMF’s own $1m on bednets minus the $500,000 that DFID would have put in plus the benefits of what DFID in fact spent their $500,000 on.  
> [Source](https://forum.effectivealtruism.org/posts/EP8x3vHRQJP57TjFL/the-counterfactual-impact-of-agents-acting-in-concert)

Suppose that the DFID spends their money on something half as valuable.

The counterfactual impact of the AMF is $1m - $500,000 + ($500,000)\*0.5 = $750,000.  
The counterfactual impact of DFID is $250,000.  
The sum of their counterfactual impacts is $1m; lower than the total impact, which is $1,250,000.

The Shapley value of the AMF is, in this case, $875,000.  
The Shapley value of the DFID is $375,000.  
The AMF is penalized: even though it paid $1,000,000, its Shapley value is less than that. The DFID's Shapley-impact is increased, because it could have invested its money in something more valuable, if the AMF hadn't intervened.

For a perhaps cleaner example, consider the case in which the DFID's counterfactual impact is $0: It can't use the money except to distribute nets, and the AMF got there first. In that scenario:

The counterfactual impact of the AMF is $500,000.  
The counterfactual impact of DFID is $0.  
The sum of their counterfactual impacts is $500,000. This is lower than the total impact, which is $1,000,000.

The Shapley value of the AMF is $750,000.  
The Shapley value of the DFID is $250,000.  
The AMF is penalized: even though it paid $1,000,000, its Shapley value is less than that. The DFID shares some of the impact,

## Example 6: The counterfactual value doesn't deal correctly tragedy of the commons scenarios.

Imagine a scenario in which many people could replicate the GPT-2 model and make it freely available, but the damage is already done once the first person does it. Imagine that 10 people end up doing it, and that the damage done is something big, like -10 million utility.

Then the counterfactual damage done by each person would be 0, because the other nine would have done it regardless.

The Shapley value deals with this by assigning an impact of -1 million utility to each person.

### Example 7: Hiring in EA

Suppose that there was a position in an EA org, for which there were 7 qualified applicants which are otherwise "idle". In arbitrary units, the person in that position in that organization can produce an impact of 100 utility.

The counterfactual impact of the organization is 100.  
The counterfactual impact of any one applicant is 0.

The Shapley value of the the organization is 85.71.  
The Shapley value of any one applicant is 2.38.

As there are more applicant, the value skews more in favor of the organization, and the opposite happens with less applicants. If there were instead only 3 applicants, the values would be 75 and 8.33, respectively. If there were only 2 applicants, the Shapley value of the organization is 66.66, and that of the applicants is 16.66. With one applicant and one organization, the impact is split 50/50.

In general, I suspect, but I haven't proved it, that if there are n otherwise iddle applicants, the Shapley value assigned to the organization is (n-1)/n. This suggests that a lot of the impact of the position goes to whomever created the position.

## Example 8: The Shapley value makes the price of a life rise with the number of stakeholders.

Key:

*   Shapley value - counterfactual value / counterfactual impact
*   Shapley price - counterfactual price. The amount of money needed to be counterfactually responsible for 1 unit of X / The amount of money needed for your Shapley value to be 1 unit of X.
*   Shapley cost-effectiveness - counterfactual cost-effectiveness.

Suppose that, in order to save a life, 4 agents have to be there: AMF to save a life, GiveWell to research them, Peter Singer to popularize them and a person to donate $5000. Then the counterfactual impact of the donation would be 1 life, but its Shapley value would be 1/4th. Or, in other words, the Shapley cost of saving a life though a donationis four times higher than the counterfactual cost.

Why is this? Well, suppose that, to save a life, each of the organizations spent $5000. Because all of them are necessary, the counterfactual cost of a life is $5000 for any of the stakeholders. But if you wanted to save an additional life, the amount of money which would be spend must be $5000\*4 = $20,000, because someone would have to go through the four necessary steps.

If, instead of 4 agents there were 100 agents involved, then the counterfactual price stays the same, but the Shapley price rises to 100x the counterfactual price. In general, I've said "AMF", or "GiveWell", as if they each were only one agent, but that isn't necessarily the case, so the Shapley price (of saving a life) might potentially be even higher.

This is a problem because if agents are reporting their cost-effectiveness in terms of counterfactuals, and one agent switches to consider their cost-effectiveness in terms of Shapley values, their cost effectiveness will look worse.

This is also a problem if organizations are reporting their cost-effectiveness in terms of counterfactuals, but in some areas there are 100 necessary stakeholders, and in other areas there are four.

## Shapley value and cost effectiveness.

So we not only care about impact, but also about cost-effectiveness. Let us continue with the example in which an NGO sends reminders to undergo vaccination, and let us give us some numbers.

Lets say that a small Indian state with 10 million inhabitants spends $60 million to vaccinate 30% of their population. An NGO which would otherwise be doing something really ineffective (we'll come back to this), comes in, and by sending reminders, increases the vaccination rate to 35%. They do this very cheaply, for $100,000.

The Shapley value of the Indian government would be 32.5%, or 3.25 million people vaccinated.  
The Shapley value of the small NGO would be 2.5%, or 0.25 million people vaccinated.

Dividing this by the amount of money spent:  
The cost-effectiveness in terms of the Shapley value of the Indian government would be $60 million / 3.25 million vaccinations = $18.46/vaccination.  
The cost-effectiveness in terms of the Shapley value of the NGO would be $100,000 / 250,000 vaccinations = $0.4/vaccination.

So even though the NGO's Shapley value is smaller, it's cost-effectiveness is higher, as one might expect.

If the outside option of the NGO were something which has a similar impact to vaccinating 250,000 people, we're back at the funging/leveraging scenario: because the NGO's outside option is better, its Shapley value rises.

## Cost effectiveness in terms of Shapley value changes when considering different groupings of agents.

Continuing with the same example, consider that, instead of the abstract "Indian government" as a homogeneous whole, there are different subagents which are all necessary to vaccinate people. Consider: The Central Indian Government, the Ministry of Finance, the Ministry of Health and Family Welfare, and within any one particular state: the State's Council of Ministers, the Finance Department, the Department of Medical Health and Family Welfare, etc. And within each of them there are sub-agencies, and sub-subagencies.

In the end, suppose that there are 10 organizations which are needed for the vaccine to be delivered, for a nurse to be there, for a hospital or a similar building to be available, and for there to be money to pay for all of it. For simplicity, suppose that the budget of each of those organizations is the same: $60 million / 10 = $6 million. Then the Shapley-cost effectiveness is different:

The Shapley value of each governmental organization would be 1/10 \* (30 million + 10/11 \* 0.5 million) = 345,454 people vaccinated.  
The Shapley value of the NGO would be 1/11 \* 500,000 = 45,454 people vaccinated.

The cost effectiveness of each governmental organization would be ($6 million)/(345,454 vaccinations) = $17 / vaccination.  
The cost effectiveness of the NGO would be $100,000 / 45,454 vaccinations = $2.2 / vaccination.

That's interesting. These concrete numbers are all made up, but they're inspired by reality and "plausible", and I was expecting the result to be that the NGO would be less cost-effective than a government agency. It's curious to see that, in this concrete example, the NGO seems to be robustly more cost-efficient than the government under different groupings. I suspect that something similar is going on with 80,000h.

## Better optimize Shapley.

If each agent individually maximizes their counterfactual impact per dollar, we get suboptimal results, as we have seen above. In particular, consider a toy world in which twenty people can either:

*   Each be an indispensable part of a project which has a value of 100 utility, for a total impact of 100 utility
*   Each can by themselves undertake a project which has 10 utility, for a total impact of 200 utility.

Then if each person was optimizing for the counterfactual impact, they would all choose the first option, for a lower total impact. If they were optimizing for their Shapley value, they'd choose the second option.

Can we make a more general statement? Yes. Agents individually optimizing for cost-effectiveness in terms of Shapley value globally optimize for total cost-effectiveness.

Informal proof: Consider the case in which agents have constant budgets and can divide them between different projects as they like. Then, consider the case in which each $1 is an agent: projects with higher Shapley value per dollar get funded first, then those with less impact per dollar, etc. Total cost-effectiveness is maximized. Because of order indifference, both cases produce the same distribution of resources. Thus, agents individually optimizing for cost effectiveness in terms of Shapley-value globally optimize for total cost-effectiveness.

Note: Thinking in terms of marginal cost-effectiveness doesn't change this conclusion. Thinking in terms of time/units other than money probably doesn't change the conclusion.

## Am I bean counting?

I don't have a good answer to that question.

## Conclusion

The counterfactual impact function is well defined, but it fails to meet my expectations of what an impact function ought to do when considering scenarios with multiple stakeholders.

On the other hand, the Shapley value function flows from some very general and simple properties, and can deal with the examples in which the counterfactual function fails. Thus, instead of optimizing for counterfactual impact, it seems to me that optimizing for Shapley value is less wrong.

Finally, because the Shapley value is not pretty to calculate by hand, [here is a calculator](http://shapleyvalue.com/).

Question: Is there a scenario in which the Shapley value assigns impacts which are clearly nonsensical, but with which the counterfactual value, or a third function, deals correctly?

---

## Addendum: The Shapley value is not easily computable.

For large values the Shapley value will not be computationally tractable (but approximations might be pretty good), and work on the topic has been done in the area of interpreting machine learning results. See, for example:

> This was a very simple example that we’ve been able to compute analytically, but these won’t be possible in real applications, in which we will need the approximated solution by the algorithm. Source: [https://towardsdatascience.com/understanding-how-ime-shapley-values-explains-predictions-d75c0fceca5a](https://towardsdatascience.com/understanding-how-ime-shapley-values-explains-predictions-d75c0fceca5a)

Or

> The Shapley value requires a lot of computing time. In 99.9% of real-world problems, only the approximate solution is feasible. An exact computation of the Shapley value is computationally expensive because there are 2^k possible coalitions of the feature values and the “absence” of a feature has to be simulated by drawing random instances, which increases the variance for the estimate of the Shapley values estimation. The exponential number of the coalitions is dealt with by sampling coalitions and limiting the number of iterations M. Decreasing M reduces computation time, but increases the variance of the Shapley value. There is no good rule of thumb for the number of iterations M. M should be large enough to accurately estimate the Shapley values, but small enough to complete the computation in a reasonable time. It should be possible to choose M based on Chernoff bounds, but I have not seen any paper on doing this for Shapley values for machine learning predictions. Source: [https://christophm.github.io/interpretable-ml-book/shapley.html#disadvantages-13](https://christophm.github.io/interpretable-ml-book/shapley.html#disadvantages-13)

That being said, here is a nontrivial example:

### Foundations and projects.

Suppose that within the EA community, OpenPhilantropy, a foundation whose existence I appreciate, has the opportunity to fund 250 out of 500 projects every year. Say that you also have 10 smaller foundations: Foundation1,..., Foundation10, each of which can afford to fund 20 projects, that there aren't any more sources of funding, and that each project costs the same.

On the other hand, we will also consider the situation in which OpenPhil is a monopoly. In the end, perhaps all these other foundations and centers might be founded by OpenPhilantropy themselves. Consider the assumption that OpenPhil has the opportunity to fund 450 projects out of 500, and that there are no other sources in the EA community.

Additionally, we could model the distribution of projects with respect to how much good they do in the world by ordering all projects from 1 to 500, and saying that:

*   Impact1 of the k-th project = I1(k) = 0.99^k.
*   Impact2 of the k-th project = I2(k) = 2/k^2 (a power law).

With that in mind, here are our results for the different assumptions. Power Index= Shapley(OP) / Total Impact


| Monopoly? | Impact measure | Total Impact | Shapley(OP) | Power index |
|--------------------------|----------------|--------------|-------------------|-------------------------------------------------|
| 0                        | I(k) = 0.99^k | 97.92        | 7.72              | 7.89%                                           |
| 0                        | I(k) = 2/k^2   | 3.29         | 0.028             | 0.86%                                           |
| 1                        | I(k) = 0.99^k | 97.92        | 48.96             | 50%                                             |
| 1                        | I(k) = 2/k^2   | 3.29         | 1.64              | 50%                                             |

For a version of this table which has counterfactual impact as well, see [here](https://raw.githubusercontent.com/NunoSempere/nunosempere.github.io/master/ea/OpenPhil.jpg).

The above took some time, and required me to beat the formula for the Shapley value into being computationally tractable for this particular case (see [here](https://nunosempere.github.io/ea/ShapleyComputation.jpg) for some maths, which as far as I'm aware, are original, and [here](https://github.com/NunoSempere/nunosempere.github.io/blob/master/ea/ShapleyValueCode.R) for some code).
