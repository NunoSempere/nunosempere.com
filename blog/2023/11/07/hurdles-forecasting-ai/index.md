Hurdles of using forecasting as a tool for making sense of AI progress
======================================================================

<h3>Introduction</h3>

<p>In recent years there have been various attempts at using forecasting to discern the shape of the future development of artificial intelligence, like the <a href="https://www.metaculus.com/tournament/ai-progress/">AI progress Metaculus tournament</a>, the Forecasting Research Institute&rsquo;s <a href="https://forum.effectivealtruism.org/posts/un42vaZgyX7ch2kaj/announcing-forecasting-existential-risks-evidence-from-a">existential risk forecasting tournament/experiment</a>, <a href="https://forum.effectivealtruism.org/posts/EG9xDM8YRz4JN4wMN/samotsvety-s-ai-risk-forecasts">Samotsvety forecasts</a> on the topic of AI progress and dangers, or various questions osn <a href="https://www.infer-pub.com">INFER</a> on short-term technological progress.</p>

<p>Here is a list of reasons, written with early input from Misha Yagudin, on why using forecasting to make sense of AI developments can be tricky, as well some casual suggestions of ways forward.</p>

<h3>Excellent forecasters and Superforecasters™ have an imperfect fit for long-term questions</h3>

<p>Here are some reasons why we might expect longer-term predictions to be more difficult:</p>

<ol>
<li>No fast feedback loops for long-term questions. You can&rsquo;t get that many predict/check/improve cycles, because questions many years into the future, tautologically, take many years to resolve. There are shortcuts, like this <a href="https://www.quantifiedintuitions.org/pastcasting">past-casting</a> app, but they are imperfect.</li>
<li>It&rsquo;s possible that short-term forecasters might acquire habits and intuitions that are good for forecasting short-term events, but bad for forecasting longer-term outcomes. For example, &ldquo;things will change more slowly than you think&rdquo; is a good heuristic to acquire for short-term predictions, but might be a bad heuristic for longer-term predictions, in the same sense that &ldquo;people overestimate what they can do in a week, but underestimate what they can do in ten years&rdquo;. This might be particularly insidious to the extent that forecasters acquire intuitions which they can see are useful, but can&rsquo;t tell where they come from. In general, it seems unclear to what extent short-term forecasting skills would generalize to skill at longer-term predictions.</li>
<li>&ldquo;Predict no change&rdquo; in particular might do well, until it doesn&rsquo;t. Consider a world which has a 2% probability of seeing a worldwide pandemic, or some other large catastrophe. Then on average it will take 50 years for one to occur. But at that point, those predicting a 2% will have a poorer track record compared to those who are predicting a ~0%.</li>
<li>In general, we have been in a period of comparative technological stagnation, and forecasters might be adapted to that, in the same way that e.g., startups adapted to low interest rates.</li>
<li>Sub-sampling artifacts within good short-term forecasters are tricky. For example, my forecasting group Samotsvety is relatively bullish on transformative technological change from AI, whereas the Forecasting Research Institute&rsquo;s pick of forecasters for their existential risk survey was more bearish.</li>
</ol>


<h3>Forecasting loses value when decontextualized, and current forecasting seems pretty decontextualized</h3>

<p>Forecasting seems more valuable when it is commissioned to inform a specific decision. For instance, suppose that you were thinking of starting a new startup. Then it would be interesting to look at:</p>

<ul>
<li>The base rate of success for startups</li>
<li>The base rate of success for all new businesses</li>
<li>The base rate of success for startups that your friends and wider social circle have started</li>
<li>Your personal rate of success at things in life</li>
<li>The inside view: decomposing the space between now and potential success into steps and giving explicit probabilities to each step</li>
<li>etc.</li>
</ul>


<p>With this in mind, you could estimate the distribution of monetary returns to starting a startup, vs e.g., remaining an employee somewhere, and make the decision about what to do next with that estimate as an important factor.</p>

<p>But our impression is that AI forecasting hasn&rsquo;t been tied to specific decisions like that. Instead, it has tended to ask questions that might contribute to an &ldquo;holistic understanding&rdquo; of the field. For example, look at <a href="https://www.metaculus.com/tournament/ai-progress/">Metaculus' AI progress tournament</a>. The first few questions are:</p>

<ul>
<li><a href="https://www.metaculus.com/questions/6299/nlo-e-prints-2021-01-14-to-2030-01-14/">How many Natural Language Processing e-prints will be published on arXiv over the 2021-01-14 to 2030-01-14 period?</a></li>
<li><a href="https://www.metaculus.com/questions/5958/it-as--of-gdp-in-q4-2030/">What percent will software and information services contribute to US GDP in Q4 of 2030?</a></li>
<li><a href="https://www.metaculus.com/questions/11241/top-price-performance-of-gpus/">What will be the average top price performance (in G3D Mark /$) of the best available GPU on the following dates?</a></li>
</ul>


<p>My impression is that these questions don&rsquo;t have the immediacy of the previous example about startups failing; they aren&rsquo;t incredibly connected to impending decisions. You could draft questions which are more connected to impending decisions, like asking about whether specific AI safety research agendas would succeed, whether AI safety organizations that were previously funded would be funded again, or about how Open Philanthropy would evaluate its own AI safety grant-making in the future. However, these might be worse qua forecasting questions, or at least less Metaculus-like.</p>

<p>Overall, my impression is that forecasting questions about AI haven&rsquo;t been tied to specific decisions in a way that would make them incredibly valuable. This is curious, because if we look at the recent intellectual history of forecasting, its original raison d'être was to make US intelligence reports more useful, and those reports were directly tied to decisions. But now forecasts are presented separately. In our experience, it has often been more meaningful for forecasters to look in depth at a topic, and then produce a report which contains predictions, rather than producing predictions alone. But this doesn&rsquo;t happen often.</p>

<h3>The phenomena of interest are really imprecise</h3>

<p>Misha Yagudin recalls that he knows of at least five different operationalizations of &ldquo;human-level AGI&rdquo;. &ldquo;Existential risk&rdquo; is also ambiguous: does it refer to human extinction? or to losing a large fraction of possible human potential? if so, how is &ldquo;human potential&rdquo; specified?</p>

<p>To deal with this problem, one can:</p>

<ul>
<li>Not spend much time on operationalization, and accept that different forecasters will be talking about slightly different concepts.</li>
<li>Try to specify concepts as precisely as possible, which involves a large amount of effort.</li>
</ul>


<p>Neither of those options is great. Although some platforms like Manifold Markets and Polymarket are experimenting with under-specified questions, forecasting seems to work best when working with clear definitions. And the fact that this is expensive to do makes the topic of AI a bit of a bad fit for forecasting.</p>

<p>CSET had a great report trying to address this difficulty: <a href="https://search.nunosempere.com/search?q=Future%20Indices">Future Indices</a>. By having a few somewhat overlapping questions on a topic, e.g., a few distinct operationalizations of AGI, or a few proxies that capture different aspects of a domain of interest, we can have a summary index that better captures the fuzzy concept that we are trying to reason about than any one imperfect question.</p>

<p>That approach does make dealing with imprecise phenomena easier. But it increases costs, and a bundle of very similar questions can sometimes be dull to forecast on. It also doesn&rsquo;t solve this problem completely—some concepts, like &ldquo;disempowering humanity&rdquo;, still remain very ambiguous.</p>

<p>Here are some high-level examples for which operationalization might still be a concern:</p>

<ul>
<li>You might want to ask about whether &ldquo;AI will go well&rdquo;. The answer depends whether you compare this against &ldquo;humanity&rsquo;s maximum potential&rdquo; or with human extinction.</li>
<li>You might want to ask whether any AI startup will &ldquo;have powers akin to that of a world government&rdquo;.</li>
<li>You might want to ask about whether measures taken by AI labs are &ldquo;competent&rdquo;.</li>
<li>You might want to ask about whether some AI system is &ldquo;human-level&rdquo;, and find that there are wildly different operationalizations available for this</li>
</ul>


<p>Here are some lower-level but more specific examples:</p>

<ul>
<li>Asking about FLOPs/$ seems like a tempting abstraction at first, because then you can estimate the FLOPs if the largest experiment is willing to spend $100M, $1B, $10B, etc. However, the abstraction ends up breaking down a bit when you look at specifics.

<ul>
<li>Dollars are unspecified: For example, consider a group like <a href="https://www.reuters.com/technology/inflection-ai-raises-13-bln-funding-microsoft-others-2023-06-29/">Inflection</a>, which raises $1B from NVIDIA and Microsoft, and pays NVIDIA and Microsoft $1B to buy the chips and build the datacenters. Then the FLOPs/$ is very under-defined. OpenAI&rsquo;s deal with Microsoft also makes their FLOPS/$ ambiguous. If China becomes involved, their ability to restrict emigration and the pre-eminent role of their government in the economy also makes FLOPs/$ ambiguous.</li>
<li>FLOPS are under-specified. Do you mean 64-bit precision bits? 16-bit precision? 8-bit precision? Do you count a <a href="https://wikiless.nunosempere.com/wiki/Multiply%E2%80%93accumulate_operation?lang=en">multiply-accumulate</a> operation as one FLOP or two FLOPs?</li>
</ul>
</li>
<li>Asking about what percentage of labor is automated gets tricky when, instead of automating exactly past labor, you automatize a complement. For example, instead of automatizing a restaurant as is, you design the menu and experience that is most amenable to being automated. Portable music devices don&rsquo;t automate concert halls, they provide a different experience. These differences matter when asking short-term resolvable questions about automation.</li>
<li>You might have some notion of a &ldquo;leading lab&rdquo;. But operationalizing this is tricky, and simply enumerating current &ldquo;leading labs&rdquo; risks them being sidelined by an upstart, or that list not including important Chinese labs, etc. In our case, we have operationalized &ldquo;leading lab&rdquo; as &ldquo;a lab that has performed a training run within 2 OOM of the largest ever at the time of the training run, within the last 2 years&rdquo;, which leans on the inclusive side, but requires keeping good data of what the largest training data is at each point in time, like <a href="https://epochai.org/research/ml-trends">here</a>, which might not be available in the future.</li>
</ul>


<h3>Many questions don&rsquo;t resolve until it&rsquo;s already too late</h3>

<p>Some of the questions we are most interested in, like &ldquo;will AI permanently disempower humanity&rdquo;, &ldquo;will there be a catastrophe caused by an AI system that kills >5%, or >95% of the human population&rdquo;, or &ldquo;over the long-term, will humanity manage to harness AI to bring forth a flourishing future &amp; achieve humanity&rsquo;s potential?&rdquo; don&rsquo;t resolve until it&rsquo;s already too late.</p>

<p>This adds complications, because:</p>

<ul>
<li>Using short-term proxies rather than long-term outcomes brings its own problems</li>
<li>Question resolution after transformative AI poses incentive problems. E.g., the answer incentivized by &ldquo;will we get unimaginable wealth?&rdquo; is &ldquo;no&rdquo;, because if we do get unimaginable wealth, the reward is worth less.</li>
<li>You may have <a href="https://en.wikipedia.org/wiki/Prevention_paradox">&ldquo;prevention paradox&rdquo;</a> and fixed-point problems, where asking a probability reveals that some risk is high, after which you take measures to reduce that risk. You could have asked about the probability conditional on taking no measures, but then you can&rsquo;t resolve the forecasting question.</li>
<li>You can chain forecasts, e.g., ask &ldquo;what will [another group] predict that the probability of [some future outcome] is, in one year&rdquo;. But this adds layers of indirection and increases operational burdens.</li>
</ul>


<p>Another way to frame this is that some stances about how the future of AI will go are unfalsifiable until a hypothesized treacherous turn in which humanity dies, but otherwise don&rsquo;t have strong enough views on short-term developments that they are willing to bet on short-term events. That seems to be the takeaway from the <a href="https://www.lesswrong.com/s/n945eovrA3oDueqtq">late 2021 MIRI conversations</a>, which didn&rsquo;t result in a string of $100k bets. While this is a disappointing position to be in, not sure that forecasting can do much here beyond pointing it out.</p>

<h3>More dataset gathering is needed</h3>

<p>A pillar of Tetlock-style forecasting is looking at historical frequencies and extrapolating trends. For the topic of AI, it might be interesting to do some systematic data gathering, in the style of Our World In Data-type work, on measures like:</p>

<ul>
<li>Algorithmic improvement for [chess/image classification/weather prediction/&hellip;]: how much compute do you need for equivalent performance? what performance can you get for equivalent compute?</li>
<li>Price of FLOPs</li>
<li>Size of models</li>
<li>Valuation of AI companies, number of AI companies through time</li>
<li>Number of organizations which have trained a model within 1, 2 OOM of the largest model</li>
<li>Performance on various capability benchmarks</li>
<li>Very noisy proxies: Machine learning papers uploaded to arXiv, mentions in political speeches, mentions in American legislation, Google n-gram frequency, mentions in major newspaper headlines, patents, number of PhD students, number of Sino-American collaborations, etc.</li>
<li>Answers to AI Impacts' survey of ML researchers through time</li>
<li>Funding directed to AI safety through time</li>
</ul>


<p>Note that datasets for some of these exist, but systematic data collection and presentation in the style of <a href="https://ourworldindata.org/">Our World In Data</a> would greatly simplify creating forecasting pipelines about these questions, and also produce an additional tool for figuring out &ldquo;what is going on&rdquo; at a high level with AI. As an example, there is a difference between &ldquo;Katja Grace polls ML researchers every few years&rdquo;, and &ldquo;there are pipelines in place to make sure that that survey happens regularly, and forecasting questions are automatically created five years in advance and included in forecasting tournaments with well-known rewards&rdquo;. <a href="https://epochai.org/">Epoch</a> is doing some good work in this domain.</p>

<h3>Forecasting AI hits the limits of Bayesianism in general</h3>

<p>One could answer worries about Tetlock-style forecasting by saying: sure, that particular brand of forecasting isn&rsquo;t known to work on long-term predictions. But we have good theoretical reasons to think that Bayesianism is a good model of a perfect reasoner: see for example the review of <a href="https://en.wikipedia.org/wiki/Cox%27s_theorem">Cox&rsquo;s theorem</a> in the first few chapters of <a href="https://annas-archive.org/md5/ddec0cf1982afa288d61db3e1f7d9323">Probability Theory. The Logic of Science</a>. So the thing that we should be doing is some version of subjective Bayesianism: keeping track of evidence and expressing and sharpening our beliefs with further evidence. See <a href="https://nunosempere.com/blog/2022/08/31/on-cox-s-theorem-and-probabilistic-induction/">here</a> for a blog post making this argument in more length, though still informally.</p>

<p>But Bayesianism is a good model of a perfect reasoner with <em>infinite compute</em> and <em>infinite memory</em>, and in particular access to a bag of hypotheses which contains the true hypothesis. However, humans don&rsquo;t have infinite compute, and sometimes don&rsquo;t have the correct hypothesis in mind. <a href="https://en.wikipedia.org/wiki/Knightian_uncertainty">Knightian uncertainty</a> and <a href="https://en.wikipedia.org/wiki/The_Structure_of_Scientific_Revolutions">Kuhnian revolutions</a><sup id="fnref:1"><a href="#fn:1" rel="footnote">1</a></sup>, <a href="https://en.wikipedia.org/wiki/Black_swan_theory">Black swans</a> or <a href="https://en.wikipedia.org/wiki/Ambiguity_aversion">ambiguity aversion</a> can be understood as consequences of normally being able to get around being approximately Bayesian, but sometimes getting bitten by that approximation being bounded and limited.</p>

<p>So there are some situations where we can get along by being approximately Bayesian, like coin flips and blackjack tables, domains where we pull our hairs and accept that we don&rsquo;t have infinite compute, like maybe some turbulent and chaotic physical systems or trying to predict dreams. Then we have some domains in which our ability to predict is meaningfully improving with time, like for example weather forecasts, where we can throw supercomputers and PhD students at it, because we care.</p>

<p>Now the question is where AI in particular falls within that spectrum. Personally, I suspect that it is a domain in which we are likely to not have the correct hypothesis in our prior set of hypotheses. For example, observers in general, but also the <a href="https://intelligence.org/">Machine Intelligence Research Institute</a> in particular, failed to predict the rise of LLMs and to orient their efforts into making such systems safer, or into preventing such systems from coming into existence. I think this tweet, though maybe meant to be hurtful, is also informative about how tricky of a domain predicting AI progress is:</p>

<blockquote class="twitter-tweet"><p lang="en" dir="ltr">eliezer has IMO done more to accelerate AGI than anyone else.<br><br>certainly he got many of us interested in AGI, helped deepmind get funded at a time when AGI was extremely outside the overton window, was critical in the decision to start openai, etc.</p>&mdash; Sam Altman (@sama) <a href="https://twitter.com/sama/status/1621621724507938816?ref_src=twsrc%5Etfw">February 3, 2023</a></blockquote>


<p> <script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script></p>

<p>However, consider the following caveat: imagine that instead of being interested in AI progress, we were interested in social science, and concerned that they couldn&rsquo;t arrive at the correct conclusion in cases where it was Republican-flavored. Then, one could notice that moving from p-values to likelihood ratios and Bayesian calculations wouldn&rsquo;t particularly help, since Bayesianism doesn&rsquo;t work unless your prior assigns a sufficiently high prior probability to the correct hypothesis. In this case, I think one easy mistake to make might be to just shrug and keep using p-values.</p>

<p>Similarly, for AI progress, one could notice that there is this subtle critique of forecasting and Bayesianism, and move to using, I don&rsquo;t know, scenario planning, which arguendo could be even worse, assume even more strongly that you know the shape of events to come, or not provide mechanisms for noticing that none of your hypotheses are worth much. I think that would be a mistake.</p>

<h3>Forecasting also has a bunch of other limitations as a genre</h3>

<p>You can see forecasting as a type of genre. In it, someone writes a forecasting question, that question is deemed sufficiently robust, and then forecasters produce probabilities on it. As a genre, it has some limitations. For instance, when curious about a topic, not all roads lead to forecasting questions, and working in a project such that you <em>have</em> to produce forecasting questions could be oddly limited.</p>

<p>The conventions of the forecasting genre also dictate that forecasters will spend a fairly short amount of time researching before making a prediction. Partly this is a result of, for example, the scoring rule in Metaculus, which incentivized forecasting on many questions. Partly this is because forecasting platforms don&rsquo;t generally pay their forecasters, and even those that are <a href="https://www.openphilanthropy.org/grants/applied-research-laboratory-for-intelligence-and-security-forecasting-platforms/">well funded</a> pay their forecasters badly, which leads to forecating being a hobby, rather than a full-time occupation. If one thinks that some questions require one to dig deep, and that one will otherwise easily produce shitty forecasts, this might be a particularly worrying feature of the genre.</p>

<p>Perhaps also as a result of its unprofitability, the forecasting community has also tended to see a large amount of churn, as hobbyist forecasters rise up in their regular careers and it becomes more expensive for them in terms of income lost to forecast on online platforms. You also see this churn in terms of employees of these forecasting platforms, where maybe someone creates some new project—e.g., Replication Markets, Metaculus' AI Progress Tournament, Ought&rsquo;s Elicit, etc.—but then that project dies as its principal person moves on to other topics.</p>

<p>Forecasting also makes use of scoring rules, which aim to reward forecasters such that they will be incentivized to input their true probabilities. Sadly, these often have the effect of incentivizing people to not collaborate and share information. This can be fixed by using more capital-intensive scoring rules that incentivize collaboration, like <a href="https://github.com/SamotsvetyForecasting/optimal-scoring">these ones</a> or by grouping forecasters into teams such that they will be incentivized to share information within a team.</p>

<h3>As an aside, here is a casual review of the track record of long-term predictions</h3>

<p>If we review the track record of superforecasters on longer term questions, we find that&hellip; there isn&rsquo;t that much evidence here—remember that the <a href="https://wikiless.nunosempere.com/wiki/Aggregative_Contingent_Estimation_Program?lang=en">ACE program</a> started in 2010. In <em>Superforecasting</em> (2015), Tetlock wrote:</p>

<blockquote><p>Taleb, Kahneman, and I agree there is no evidence that geopolitical or economic forecasters can predict anything ten years out beyond the excruciatingly obvious—“there will be conflicts”—and the odd lucky hits that are inevitable whenever lots of forecasters make lots of forecasts. These limits on predictability are the predictable results of the butterfly dynamics of nonlinear systems. In my EPJ research, the accuracy of expert predictions declined toward chance five years out. And yet, this sort of forecasting is common, even within institutions that should know better.</p></blockquote>

<p>However, in p. 33 of <a href="https://papers.ssrn.com/sol3/papers.cfm?abstract_id=4377599">Long-Range Subjective-Probability Forecasts of Slow-Motion Variables in World Politics: Exploring Limits on Expert Judgment</a> (2023), we see that the experts predicting &ldquo;slow-motion variables&rdquo; 25 years into the future attain a Brier score of 0.07, which isn&rsquo;t terrible.</p>

<p>Karnofsky, the erstwhile head-honcho of Open Philanthropy, <a href="https://www.cold-takes.com/the-track-record-of-futurists-seems-fine/">spins</a> some research by Arb and others as saying that the track record of futurists is &ldquo;fine&rdquo;. <a href="https://danluu.com/futurist-predictions/">Here</a> is a more thorough post by Dan Luu which concludes that:</p>

<blockquote><p>&hellip;people who were into &ldquo;big ideas&rdquo; who use a few big hammers on every prediction combined with a cocktail party idea level of understanding of the particular subject to explain why a prediction about the subject would fall to the big hammer generally fared poorly, whether or not their favored big ideas were correct. Some examples of &ldquo;big ideas&rdquo; would be &ldquo;environmental doomsday is coming and hyperconservation will pervade everything&rdquo;, &ldquo;economic growth will create near-infinite wealth (soon)&rdquo;, &ldquo;Moore&rsquo;s law is supremely important&rdquo;, &ldquo;quantum mechanics is supremely important&rdquo;, etc. Another common trait of poor predictors is lack of anything resembling serious evaluation of past predictive errors, making improving their intuition or methods impossible (unless they do so in secret). Instead, poor predictors often pick a few predictions that were accurate or at least vaguely sounded similar to an accurate prediction and use those to sell their next generation of predictions to others.</p>

<p>By contrast, people who had (relatively) accurate predictions had a deep understanding of the problem and also tended to have a record of learning lessons from past predictive errors. Due to the differences in the data sets between this post and Tetlock&rsquo;s work, the details are quite different here. The predictors that I found to be relatively accurate had deep domain knowledge and, implicitly, had access to a huge amount of information that they filtered effectively in order to make good predictions. Tetlock was studying people who made predictions about a wide variety of areas that were, in general, outside of their areas of expertise, so what Tetlock found was that people really dug into the data and deeply understood the limitations of the data, which allowed them to make relatively accurate predictions. But, although the details of how people operated are different, at a high-level, the approach of really digging into specific knowledge was the same.</p></blockquote>

<h3>In comparison with other mechanisms for making sense of future AI developments, forecasting does OK.</h3>

<p>Here are some mechanisms that the Effective Altruism community has historically used to try to make sense of possible dangers stemming from future AI developments:</p>

<ul>
<li>Books, like Bostrom&rsquo;s <em>Superintelligence</em>, which focused on the abstract properties of highly intelligent and capable agents in the limit.</li>
<li><a href="https://www.openphilanthropy.org/research/?q=&amp;focus-area%5B%5D=potential-risks-advanced-ai&amp;content-type%5B%5D=research-reports">Reports</a> by Open Philanthropy. They either try to model AI progress in some detail, like <a href="https://www.lesswrong.com/posts/KrJfoZzpSDpnrv9va/draft-report-on-ai-timelines">example 1</a>, or look at priors on technological development, like <a href="https://www.openphilanthropy.org/research/semi-informative-priors-over-ai-timelines/">example 2</a>.</li>
<li>Mini think tanks, like Rethink Priorities, Epoch or AI impacts, which produce their own research and reports.</li>
<li>Larger think tanks, like CSET, which produce reports like <a href="https://cset.georgetown.edu/publication/future-indices/">this one</a> on Future Indices.</li>
<li>Online discussion on lesswrong.com, that typically assumes things like: intelligence gains would be fast and explosive, that we should aim to design a mathematical construction that guarantees safety, that iteration would not be advisable in the face of fast intelligence gains, etc.</li>
<li>Occasionally, theoretical or mathematical arguments or models of risk.</li>
<li>One-off projects, like Drexler&rsquo;s <a href="https://www.fhi.ox.ac.uk/reframing/">Comprehensive AI systems</a></li>
<li>Questions on forecasting platforms, like Metaculus, that try to solidly operationalize possible AI developments and dangers, and ask their forecasters to anticipate when and whether they will happen.</li>
<li>Writeups from forecasting groups, like <a href="https://forum.effectivealtruism.org/posts/EG9xDM8YRz4JN4wMN/samotsvety-s-ai-risk-forecasts">Samotsvety</a></li>
<li>More recently, the Forecasting Research Institute&rsquo;s <a href="https://forecastingresearch.org/xpt">existential risk tournament/experiment writeup</a>, which has tried to translate geopolitical forecasting mechanisms to predicting AI progress, with mixed success.</li>
<li>Deferring to intellectuals, ideologues, and cheerleaders, like Toby Ord, Yudkowsky or MacAskill.</li>
</ul>


<p>None of these options, as they currently exist, seem great. Forecasting has the hurdles discussed above, but maybe other mechanisms have even worse downsides, particularly the more pundit-like ones. Conversely, forecasting will be worse than deferring to a brilliant theoretical mind that is able to grasp the dynamics and subtleties of future AI development, like perhaps Drexler&rsquo;s on a good day.</p>

<p>Anyways, you might think that this forecasting thing shows potential. Were you a billionnaire, money would not be a limitation for you, so&hellip;</p>

<h3>In this situation, here are some strategies of which you might avail yourself</h3>

<h4>A. Accept the Faustian bargain</h4>

<ol>
<li>Make a bunch of short-term and long-term forecasting questions on AI progress</li>
<li>Wait for the short-term forecasting questions to resolve</li>
<li>Weight the forecasts for the long-term questions according to accuracy in the short term questions</li>
</ol>


<p>This is a Faustian bargain because of the reasons reviewed above, chiefly that short-term forecasting performance is not a guarantee of longer term forecasting performance. A cheap version of this would be to look at the best short-term forecasters on the AI categories on Metaculus, and report their probabilities on a few AI and existential risk questions, which would be more interpretable than the current opaque &ldquo;Metaculus prediction&rdquo;.</p>

<p>If you think that your other methods of making sense of what it&rsquo;s going on are sufficiently bad, you could choose this and hope for the best? Or, conversely, you could anchor your beliefs on a weighted aggregate of the best short-term forecasters and the most convincing theoretical views. Maybe things will be fine?</p>

<h4>B. Attempt to do a Bayesianism</h4>

<p>Go to the effort of rigorously formulating hypotheses, then keep track of incoming evidence for each hypothesis. If a new hypothesis comes in, try to do some version of <a href="https://nunosempere.com/blog/2023/02/04/just-in-time-bayesianism/">just-in-time Bayesianism</a>, i.e., monkey-patch it after the fact. Once you are specifying your beliefs numerically, you can deploy some cute incentive mechanisms and <a href="https://github.com/SamotsvetyForecasting/optimal-scoring/blob/master/3-amplify-bayesian/amplify-bayesian.pdf">reward people who change your mind</a>.</p>

<p>Hope that keeping track of hypotheses about the development of AI at least gives you some discipline, and enables you to shed untrue hypotheses or frames a bit earlier than you otherwise would have. Have the discipline to translate the worldviews of various pundits into specific probabilities<sup id="fnref:2"><a href="#fn:2" rel="footnote">2</a></sup>, and listen to them less when their predictions fail to come. And hope that going to the trouble of doing things that way allows you to anticipate stuff 6 months to 2 years sooner than you would have otherwise, and that it is worth the cost.</p>

<h4>C. Invest in better prediction pipelines as a whole</h4>

<p>Try to build up some more speculative and <a href="https://nunosempere.com/blog/2023/07/19/better-harder-faster-stronger/">formidable</a> type of forecasting that can deal with the hurdles above. Be more explicit about the types of decisions that you want better foresight for, realize that you don&rsquo;t have the tools you need, and build someone up to be that for you.</p>
<div class="footnotes">
<hr/>
<ol>
<li id="fn:1">
To spell this out more clearly, Kuhn was looking at the structure of scientific revolutions, and he notices that you have these &ldquo;paradigm changes&rdquo; every now in a while. As a naïve Bayesian, those paradigm changes are kinda confusing, and shouldn&rsquo;t have any special status. You should just have hypotheses, and they should just rise and fall in likelihood according to Bayes rule. But as a Bayesian who knows he has finite compute/memory, you can think of Kuhnian revolutions as encountering a true hypothesis which was outside your previous hypothesis space, and having to recalculate. On this topic, see <a href="https://nunosempere.com/blog/2023/02/04/just-in-time-bayesianism/">Just-in-time Bayesianism</a> or <a href="https://nunosempere.com/blog/2023/03/01/computable-solomonoff/">A computable version of Solomonoff induction</a>.<a href="#fnref:1" rev="footnote">&#8617;</a></li>
<li id="fn:2">
Back in the day, Tetlock received a <a href="https://www.openphilanthropy.org/grants/university-of-pennsylvania-philip-tetlock-on-forecasting/#2-about-the-grant">grant</a> to &ldquo;systematically convert vague predictions made by prominent pundits into explicit numerical forecasts&rdquo;, but I haven&rsquo;t been able to track what happened to it, and I suspect it never happened.<a href="#fnref:2" rev="footnote">&#8617;</a></li>
</ol>
</div>

<p>
  <section id='isso-thread'>
  <noscript>javascript needs to be activated to view comments.</noscript>
  </section>
</p>
