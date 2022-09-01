A comment on Cox's theorem and probabilistic inductivism.
=====================================================================================

I'm currently reading _What is this thing called science_, an introduction to philosophy of science by Alan Chalmers. Around page 48, Chalmers presents the following chain of reasoning:

> One attempt to avoid the problem of induction involves weakening the demand that scientific knowledge be proven true, and resting content with the claim that scientific claims can be shown to be probably true in the light of the evidence. So the vast number of observations that can be invoked to support the claim that materials denser than air fall downwards on earth, although it does not permit us to prove the truth of the claim, does warrant the assertion that the claim is probably true.

> In line with this suggestion we can reformulate the principle of induction to read, ‘if a large number of As have been observed under a wide variety of conditions, and if all these observed As have the property B, then all As probably have the property B’. This reformulation does not overcome the problem of induction. The reformulated principle is still a universal statement. It implies, on the basis of a finite number of successes, that all applications of the principle will lead to general conclusions that are probably true. 

> Consequently, attempts to justify the probabilistic version of the principle of induction by an appeal to experience involve an appeal to inductive arguments of the kind that are being justified just as the principle in its original form did.

I think this is not a good representation of what the probabilistic view would say. Here is what I think is a more convincing argument:

1. Per [Cox's theorem](https://en.wikipedia.org/wiki/Cox%27s_theorem), any "reasonable" view of probability is Bayesian.
2. Per Bayesian probability, hypotheses which predict future events which happen become more plausible, whereas hypotheses which don't predict the future loose plausibility.
3. Hypotheses such as "induction works", or "induction works in physics" perform well. Hypotheses such as "induction does't work", or "induction works well in the [soft sciences](https://en.wikipedia.org/wiki/Replication_crisis)" perform poorly.
4. As a result, we should view "induction works" as a highly probable statement.

To elaborate on this a bit, Cox's theorem starts with a few reasonable assumptions, like:

- "plausibility can be represented as a real number"
- "ways to derive plausibility have to be consistent"
- The plausibility of A & B is given by some function of "the plausibility of" and "plausibility of B given that A is the case", and this function is monotonic[^1].
- In the case of certainty, probability reverts to the logic of predicates.
- etc.

Using these assumptions, we arrive at:

<p><img src="https://i.imgur.com/gkymKm9.png"></p>

and by substiting P = w^m, we arrive at:

<p><img src="https://i.imgur.com/xw3pvyx.png"  class="img-medium-center"></p>

For a good intro to Cox's theorem, see the first chapters in E.T. Jaynes' _Probability Theory: The Logic of Science_.

In any case, if C is our background knowledge, A is "induction works ok in physics", and B is "induction worked ok in this particular instance", then the probability of A rises as we observe more Bs. Conversely, the probability that induction doesn't work, or works badly, goes down every time that we observe that induction steps do work. A could also be "the sun will continue rising", and B could be "the sun has risen on day i".

To say this in other words: Every time we observe one induction step working ("if x has been the case for 6 steps, it will also be the case for 6+1 steps"), we increase our probability that "induction works for x generally", and we increase our probability that "induction works ok when we try to apply it practice".

So we get around the infinite regress in the problem of induction by bootstrapping the probability of "induction works" and "induction works in this particular case".

Now, in the spirit of philosophical questioning, we can point out a few problems and complications with this.

1. We can question the assumptions of Cox's theorem. One popular assumption to question is that the degree of plausibility ought to be represented by a real number.
2. We can point out that humans can only imperfectly apply Bayesian probability. In particular, it is very much non-trivial to deduce the implications of a theory about future events.
3. This scheme doesn't work if the initial, prior, degree of plausiblity that we assign to the principle of induction is zero or very (extremely) low. But because we have observed induction working ok so many times, starting with any reasonable probability (e.g,. 0.0001%, 0.1%, 1%, 10%) should lead us to increase our plausibility of induction to very high levels.
4. The principle of induction does in fact not work some times (e.g., when we encounter literal or metaphorical black swans, when we observe the replication crisis, or when we predict that the sun will go out in the far future). So the probabilistic view would not assign probability 100% to the principle of induction. And it would in fact assign different plausibilities to it depending on the domain. Teasing this out seems like it could be a bit tricky.
5. You sometimes have situations which are anti-inductive! E.g., as an authority becomes more reliable people trust it more, which means that it is more profitable or advantageous to subvert it, e.g., for political means.
6. Induction may not be a good guide in the presence of [anthropic effects](<https://en.wikipedia.org/wiki/Anthropic_Bias_(book)>).
7. There might be the suspicion that we have reduced the problem of induction to something which is ultimately equivalent.
8. We can think about . I get the impression that past philosophers were mostly considering whether induction works "as a whole", but if we are doing Bayesian probability, we have to assign probability mass to many other hypothesis and classes of hypotheses, like "induction works reliably in the hard sciences and not in the soft sciences".

Personally I am most worried about 2. and 5, and I view 4. as as a feature, not as a bug. But I also think that the scheme works _in practice_, and I'm very much attracted to Bayesian epistemology as a foundation for knowledge. 

I don't think that 7. is the case: we could imagine a world in which past behavior is not correlated with future behaviour, and in which our probability machinery picks that up. For instance, in very liquid markets, it could be that the moment useful patterns are identified, they are exploited until they disappear, and we could easily update on that[^2].

I'm also excited about moving these problems away from philosophy and into [engineering](https://en.wikipedia.org/wiki/Information_theory) and practice, maybe using mechanisms such as [replication markets](https://replicationmarkets.com/) or [fraud markets](https://insightprediction.com/m/183/will-ryan-enos-have-a-paper-retracted-in-2022).

PS: I see that the [Stanford Encyclopedia of Philosophy](https://plato.stanford.edu/entries/induction-problem/#BayeSolu) has a page on this. I think it gives a good overview of the rigoristic camp, but it misses some of the points that a subjectivist would raise—e.g., that in cases with a wealth of data in favor of induction, like in the case of physical laws, the initial prior doesn't matter as long as it isn't "really crazy" (e.g., zero or extremely close to it), because one will reach pretty much the same conclusion of high certainty with any prior that isn't crazy.

[^1]: The reasonableness of requiring monotonicity might not be immediately apparent, so here is an explanation: The terms of that function are "plausibility of A" and "plausibility of B given that you know A is the case". By "monotonic", we mean that if we increase (resp. decrease) either of the inputs while keeping the other constant, or if we increase (resp. decrease) both at once, then the output—the plausibility of A and B—also has to increase (resp. decrease). 

[^2]: We could also imagine living in a "maelstrom of chaos" universe in which past behavior is not predictive of future behavior. This could be the case at different levels of meta, i.e., it could also be difficult to predict whether future events are difficult to predict or understand. A real life example which bears some resemblance to this might be [Terry A. Davis](https://en.wikipedia.org/wiki/Terry_A._Davis), whose speech alternated between the utterings of a raving lunatic and gifted technical insights about programming.
