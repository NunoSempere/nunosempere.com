Infinite Ethics 101: Stochastic and Statewise Dominance as a Backup Decision Theory when Expected Values Fail
====

In [Exceeding expectations: stochastic dominance as a general decision theory](https://globalprioritiesinstitute.org/wp-content/uploads/Christian-Tarsney_Exceeding-expectations-stochastic-dominance-as-a-general-decision-theory.pdf), Christian Tarsney presents stochastic dominance (to be defined) as a total replacement for expected value as a decision theory. He wants to argue that one decision is only rationally better as another one when it is stochastically dominant. For this, he needs to say that the choiceworthiness of a decision (how rational it is) is undefined in the case where one decision doesn't stochastically dominate another one.

I think this is absurd, and perhaps determined by academic incentives to produce more eye-popping claims rather than more restricted incremental improvements. Still, I thought that the paper made some good points about us still being able to make decisions even when expected values stop being informative. It was also my introduction to extending rational decision-making to infinite cases, and a great introduction at that. Below, I outline my rudimentary understanding of these topics.

![](https://i.imgur.com/9Uc0znM.png)

## Where expected values fail.

Consider a choice between:

- A: 1 utilon with probability 1/2, 2 utilons with probability 1/4th, 4 utilons with probability 1/8th, etc. The expected value of this choice is 1 × 1/2 + 2 × 1/4 + 4 × 1/8 + ... = 1/2 + 1/2 + 1/2 + ... = ∞
- B: 2 utilons with probability 1/2, 4 utilons with probability 1/4th, 8 utilons with probability 1/8th, etc. The expected value of this choice is 2 × 1/2 + 2 × 1/4 + 4 × 1/8 + ... = 1 + 1 + 1 + ... = ∞

So the expected value of choice A is ∞, as is the expected value of choice B. And yet, B is clearly preferable to A. What gives?

## Statewise dominance

Suppose that in the above case, there were different possible states, as if the payoffs for A and B were determined by the same coin throws:

- State i: A gets 1, B gets 2
- State ii: A gets 2, B gets 4
- State iii: A gets 4, B gets 8,
- State i^n: A gets 2^n, B gets 2 × 2^n.

Then in this case, B dominates A in every possible state. This is a reasonable decision principle that we can reach to ground our decision to choose B over A.

## Stochastic dominance

O stochastically dominates P if:

1. For any payoff x, the probability that O yields a payoff at least as good as x is equal
   to or greater than the probability that P yields a payoff at least as good as x, and
2. For some payoff x, the probability that O yields a payoff at least as good as x is
   strictly greater than the probability that P yields a payoff at least as good as x.

or, in math notation:

1. ∀x, Probability(Payoff(O) ≥ x) ≥ Probability(Payoff(P) ≥ x))
2. ∃x such that Probability(Payoff(O) ≥ x) > Probability(Payoff(P) ≥ x))

This captures a notion that O is, in a sense, strictly better than P, probabilistically.

In the case of A and B above, if their payoffs were determined by throwing independent coins:

- There is a 100% chance that B yields a payoff ≥ 1, and 100% that A yields a payoff ≥ 1
- There is a 50% chance that B yields a payoff ≥ 2, but only a 25% chance that A yields a payoff ≥ 2
- There is a 25% chance that B yields a payoff ≥ 4, but only a 12.5% chance that A yields a payoff ≥ 4
- There is a 12.5% chance that B yields a payoff ≥ 8, but only a 6.26% chance that A does so.
- There is a 1/2^n chance that B yields a payoff ≥ 2^n, but only a 1/2^(n+1) chance that A does so.

So the probability that B gets increasingly better outcomes is higher than the probability that A will do so. So in this case, B stochastically dominates A. Stochastic dominance is thus another decision principle that we could reach to compare choices with infinite expected values.

## Gaps left

The above notions of stochastic and statewise dominance could be expanded and improved. For instance, we could ignore a finite number of comparisons going the other way if the expected value of those options was finite but the expected value of the whole thing was infinite. For instance, in the following comparison:

- A: 100 utilons with probability 1/2, 2 utilons with probability 1/4th, 4 utilons with probability 1/8th, etc. The expected value of this choice is 1 × 1/2 + 2 × 1/4 + 4 × 1/8 + ... = 1/2 + 1/2 + 1/2 + ... = ∞
- B: 2 utilons with probability 1/2, 4 utilons with probability 1/4th, 8 utilons with probability 1/8th, etc. The expected value of this choice is 2 × 1/2 + 2 × 1/4 + 4 × 1/8 + ... = 1 + 1 + 1 + ... = ∞

I would still say that B is preferable to A in that case. And my impression is that there are many similar principles one could reach to, in order to resolve many but not all comparisons between infinite sequences.

---

Exercise for the reader: Come up with two infinite sequences which cannot be compared using statewise or stochastic dominance, or similar principles.

You can see some discussion of this post in the [EA Forum](https://forum.effectivealtruism.org/posts/GseREh8MEEuLCZayf/nunosempere-s-shortform?commentId=BMEaMkB8Fdrz7oanw).

<p><section id="isso-thread">
  <noscript>Javascript needs to be activated to view comments.</noscript>
</section></p>
