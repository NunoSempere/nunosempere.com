Use of &ldquo;I'd bet&rdquo; on the EA Forum is mostly metaphorical
==========================================================

Epistemic status: much ado about nothing.

**tl;dr**: I look at people saying "I'd bet" on the EA Forum. I find that they mostly mean this metaphorically. I suggest reserving the word "bet" for actual bets, offer to act as a judge for the next 10 bets that people ask me to judge, and mention that I'll be keeping an eye on people who offer bets on the EA Forum to consider taking them. Usage of the construction "I'd bet" is a strong signal of belief only if it is occasionally tested, and I suggest we make it so.

Inspired by [this manifold market created by Alex Lawsen](https://manifold.markets/AlexL/has-nuno-sempere-got-an-automated-s)—which hypothesized that I have an automated system to detect where people offer bets—and by [this exchange](https://forum.effectivealtruism.org/posts/xBfp3HGapGycuSa6H/i-m-less-approving-of-the-ea-community-now-than-before-the?commentId=Gafn8GSkTge6q6ezf#46wAwgBFBmHBmoLt4)—where someone said "I would bet all the money I have (literally not figuratively) that X" and then didn't accept one such bet—I wrote a small script[^1] to search for instances of the word "bet" on the EA Forum:

```
$ data="{
     \"query\": \"{
       comments( input: {
         terms: {
           limit: 5000
         }
       }){
         results{
           postId,
           htmlBody,
           postedAt
         }
       }
     }\"
   }"

$ response="$(curl 'https://forum.effectivealtruism.org/graphql/' \
    -X POST  -H 'content-type: application/json' --data "$(echo $data)")"

$ echo "$response" | jq .data.comments.results       |
    jq  'map(select(.htmlBody | contains(" bet ")))' |
    jq 'reverse'   | ack --passthru ' bet '

```

That script fetches the last 5k comments from the EA Forum, and searches for instances of the string " bet ". 

### 29 usages of the word "bet" in the last 5k comments.

In the last 5k comments, 29 mention the word "bet". I'll just paste them here, and then give some thoughts.

> I'd **bet** that the TEMS utility is lower than the aggregate TEMS costs, especially if you account for TEMS EOL costs in novel ways.  

> ---

> I'll **bet** those victims didn't do so well on an IQ test the day after their lives were destroyed.

> ---

> ...if you're feeling brave you can even decide to just look at the variants in some small subset of genes that you **bet** are more likely to have a role in determining your trait, but even after those extra assumptions, you still have many candidate variants.

> ---

> I appreciate pushback. I don't think this issue is being brigaded. If it were we'd see a load of chat on twitter. I think that explains too much. I just would **bet** it isn't the case. Some people are making new burners but I'd guess that's it.

> ---

>  We want to make bets on expertises that could come in handy in unpredictable ways. (CoI: I'm making one such **bet** myself as we speak!). 

> ---

> This isn't something we need to centrally decide, you can just... start doing it. Start a fund that makes decisions democratically! I **bet** you'd get funding from the EAIF even.

> ---

> I assume these "trained raters" were grad students who had thought about the problem for a couple days or something, and I **bet** that if you actually genuinely studied this you could get good at it, but probably very few people are in that reference class.

> ---

>  We don't really know where on subjective wellbeing scales people construe wellbeing to go from positive to negative. My best **bet** is around 2.5 on a 0 to 10 scale.

> ---

> Yeah I'd **bet** this is true. I think there are tradeoffs here though (and I have also talked to women who like the status quo and I assume men do). It's not clear to me that the obvious path forward is.

> ---

> I think I'm gonna make a **bet** on a skillset that isn't correlated with the rest of the movement

> ---

> I think it's probably disproportionately common for the times when your actions were followed by bad outcomes (even if that wasn't caused by your action, or was you making a good **bet** but getting unlucky) to become visible and salient.

> ---

> I don't necessarily endorse this but I think it's valuable for us to know what we think on stuff like this - I **bet** we disagree a surprising amount.

> ---

> Personally I am able to take healing from this post, and I **bet** I'm not the only one who is finding something positive to take out of it

> ---

> We don't know that any complainants would have been paying enough attention to notice announcements, but I **bet** at least some would have, and the Time journalist could have noticed at least.

> ---

> I would **bet** this is a big part of "being a good researcher". Writing this as someone who isn't doing an AI PhD.

> ---

> I wouldn't be surprised if the situation is quite a bit different in authoritarian regimes like China and Russia, where I **bet** there's much less animal liberation activity and, in China at least, a lot more state capacity to crack down on biohazardous practice

> ---

> I'd place a **bet** that the majority of the people who are concerned about this commitment know their content, and that the majority of the people who support it don't.

> ---

> Because IMO -isms are not going to erode this community - I would **bet** almost everyone here is against them, a lot of us actively. But losing EA values - good epistemic, avoiding virtue signalling and poseuring, avoiding group-think, etc. might.

> ---

> I would like for all involved to consider this, basically, a **bet** on "making and publishing this pledge" being an effective intervention on ... something. I'm not sure whether the something is "actual racism and sexism and other bigotry within EA," or "the median EA's discomfort at their uncertainty about whether racism and sexism are a part of EA," or what. But (in the spirit of the E in EA) I'd like that **bet** to be more clear, so since you were willing to leave a comment above: would you be willing to state with a little more detail which problem this was intended to solve, and how confident you (the group involved) are that it will be
 a good intervention?

> ---

> Although I can't say for sure, I would also **bet** that there's dozens of unofficial rationalist events (and a few unofficial EA events) that he attended in the last five years, given that he was literally hanging out in the miri/cfar reception area for hours per week, right until the time he was officially banned.

> ---

> Also, for Nuno I'll ask when OP gonna let us **bet** against them.

> ---

> There's a phenomenon where a gambler places their money on 32, and then the roulette wheel comes up 23, and they say "I'm such a fool; I should have **bet** 23". More useful would be to say "I'm such a fool; I should have noticed that the EV of this gamble is negative." Now at least you aren't asking for magic lottery powers. Even more useful would be to say "I'm such a fool; I had three chances to notice that this **bet** was bad: when my partner was trying to explain EV to me; when I snuck out of the house and ignored a sense of guilt; and when I suppressed a qualm right before placing the bet. I should have paid attention in at least one of those cases and internalized the arguments about negative EV, before gambling my money." Now at least you aren't asking for magic cognitive powers. My impression is that various EAs respond to crises in a manner that kinda rhymes with saying "I wish I had **bet** 23", or at best "I wish I had noticed this **bet** was negative EV"

> ---

>  Let's say you have to **bet** on whether a machine will turn on or off, however the machine is a 'frustrator', once you **bet** that it will turn off it will turn on and vise versa. You aren't predicting the outcome (you literally can't) you are causin the opposite outcome. Humans are the same, if I see that someone predicts I will do X, I might do Y just to assert my autonomy (or maybe not because of mindgames). So let's say a new prediction market opens on the machine frustrator, no one wants to be the first to make a bet, since you're guaranteed to lose money if no one else participates. Even if people are irrational enough to **bet** on it, we don't learn anything new and we just wasted a bunch of electricity and manhours for nothing. It's a waste of time and money to **bet** on frustrators, especially since you can get a 10% return on the stockmarket.

> ---

>  When the president promises not to manipulate the market about event X some people will trust him 90% and **bet** accordingly while others will trust him 10% and **bet** according to that. But that's not the same thing as people having 10% credence event X will happen, it can be that you think the president is trustworthy but the event itself is unlikely [...]

> ---

>  Compared to other interventions I would not **bet** a substantial amount on proofs for deep learning systems given its important, neglectedness, and tractability

> ---

>  The large underclass has almost no money to **bet** that it will, while the small upperclass bets a large chunk of their money that it won't. Predictably, more money is betted on it not increasing welfare and when the market closes, everyone gets their money back and the government decides not to implement it.

> ---

>  I'd **bet** you that multiple EA communities and orgs improved or are working on improving their processes after reading your post (and the comments as by EA NYC). Community builders talk (as you know) so if I had to bet, I'd **bet** changes would have reached the EA Forum (for grassroots implementation) or the CEA Groups Team (for top-down implementation) within a few months of now.

> ---

>  If you look at their notable recipients, and imagine a similar cohort producing a similar amount of impact in a less capitalistic direction, it seems at least conceivable that such a **bet** could be well worth it. That's my 2cts.

> ---

> However, maybe I'm blindspotted, but I can't find a better topic to **bet** on - would solve all problems solvable with resources I don't think I can find a non-emotional way to convince people to switch from we should not search to we should search (for infinite energy).

### Thoughts on this status quo

We notice that many people are using constructions like "I'd bet", "I would bet", "I'll bet" in a metaphorical sense. In some cases, no such bet is possible, because resolving the bet would not be possible. In other cases, a bet might be possible, but no operationalization is given, and it just seems very unlikely that a bet might arise. 

In a sense, this is ok and just corresponds to common usage. But in another sense, it just seems like there could be room for optimization.

### Proposed improvements, as suggestions rather than as hard rules

I think it would be better if there was some way of conveying enough belief in a topic that one is willing to make a monetary bet, which is a somewhat adversarial process which allows other people to extract money from you if you are wrong. That is, I think it would be better if people had a strong signal of belief available to them.

To that effect, here are some proposed improvements:

- If you would like to offer a bet for real, I would suggest...
  - That you propose propose an operationalization. This can be a clear, unambiguous indicator, e.g., "who will the president of the US in 2025 according to Wikipedia", "what will be US mortality according to Our World In Data in 2030". Or it can be an approximate or somewhat ambiguous resolution, such that there is some room for noise, and accept that sometimes the resolution will be wrong. For example, "we will solve this bet by mutual accord", "this bet is resolved by Nu&ntilde;o's best guess as to whether life in 2030 is better for the median American than life in 2025", "this bet resolves positively if 3 out of these five indicators are positive". The core thing here is that there can be some ambiguity, but also a method for arriving at a resolution despite that ambiguity.
  - That you propose the odds and the amount you'd be willing to bet. A $100:$10k bet isn't the same as a $10:$10 bet. This helps people consider whether the effort to engage is worth it. 
- On my side:
  - I am making the following offer: for the amount of $5 for each resolution, I will judge the next 10 bets of more than $100 that people ask me to judge[^2].
  - I will be listening for instances of the "bet" on the EA Forum, and will consider taking the bets that I think are wrong or uncalibrated.
- Additionally, I would find it convenient if people reserved the word "bet" for situations where they are willing to bet, and otherwise use like "I guess that", "I am confident in", "I imagine that", "it seems to me that", "I believe quite strongly that", "I am taking a chance on", and so on. This would make it easier for me programmatically search the EA forum for bet offers. 

One mildly subtle point here is that the more often people end up making bets when one bet is offered, the stronger the signal to offer a bet is. Conversely, if when people offer to make a bet it turns out that they meant so metaphorically, that they don't end up agreeing on a resolution, or that they initially meant it but then backed out, then the less point there is in trying to accept someone's bet, and the weaker the signal of offering a bet is.

### Example transformations

Together, these guidelines would change:

> I'd bet that Napoleon would agree with me on X

to either

> I'd be willing to bet my $500 against your $10 that if we ask a historian at the LSE, their best guess would be that Napoleon thought that X.

or to 

> I strongly believe that Napoleon would agree with me on X

or 

>  We don't really know where on subjective wellbeing scales people construe wellbeing to go from positive to negative. My best **bet** is around 2.5 on a 0 to 10 scale.

to

> ... My best guess is around 2.5 on a 0 to 10 scale.

or

> .... I'd bet my $60 against your $30 that if we did a survey on Positly of at least 50 people, we'd find that the average neutral point is between 2 and 4 on a 0 to 10 scale.

<br>

Anyways, that's all I have for today. I might post this in the EA forum in a few days. In the meantime, comments are welcome.

PS: You can subscribe to these posts [here](https://nunosempere.com/.subscribe/).

[^1]: Readers who think this is cool may want to learn more about Linux, and then maybe buy a Linux computer next time they are switching laptops. I'd recommend a [Lenovo Thinkpad](https://www.lenovo.com/us/en/d/linux-laptops-desktops/).

[^2]: To be clear, I think that, e.g., spending an hour researching the outcome of a bet costs me way more than $5, and I am providing this pretty much as a public service. The $5 are meant so that a) people get into the habit of paying for resolutions, which I think makes them more sustainable, and b) together with the $100 limit, so that I don't get summoned for trivial bets.


<p>
  <section id='isso-thread'>
  <noscript>Javascript needs to be activated to view comments.</noscript>
  </section>
</p>
