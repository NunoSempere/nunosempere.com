## Thoughts on the "evolutionary anchor" of Aleja Cotra's report

![](https://upload.wikimedia.org/wikipedia/commons/a/ae/Darwin%27s_finches_by_Gould.jpg)

Within Effective Altruism, Aleja Cotra's report on artificial general intelligence (AGI) timelines has been influential in justifying or convincing members and organizations to work on AGI safety.

The report has a section on the "evolutionary anchor", i.e., an upper bound on how much compute it would take to reach artificial general intelligence. The section can be found in pages 24-28 of [this Google doc](https://docs.google.com/document/d/1k7qzzn14jgE-Gbf0CON7_Py6tQUp2QNodr_8VAoDGnY/edit#heading=h.gvc1xyxlemkd). As a summary, in the report's own words:

> This hypothesis states that we should assume on priors that training computation requirements will resemble the amount of computation performed in all animal brains over the course of evolution from the earliest animals with neurons to modern humans, because we should expect our architectures and optimization algorithms to be about as efficient as natural selection.

This anchor isn't all that important in the report's own terms: it only gets a 10% probability assigned to it in the final weighted average. But this bound is personally important to be because I do buy that if you literally reran evolution, or if you use as much computation as evolution, you would have a high chance of producing something as intelligent as humans, and so I think that it is particularly forceful as an "outside view".

However, I'm not sure I buy some of the details of how the author arrives at this estimate:

> The amount of computation done over evolutionary history can roughly be approximated by the following formula: (Length of time since earliest neurons emerged) \* (Total amount of computation occurring at a given point in time). My rough best guess for each of these factors is as follows:

> - **Length of evolutionary time:** Virtually all [animals](https://en.wikipedia.org/wiki/Animal) have neurons of some form, which means that the earliest nervous systems in human evolutionary history likely emerged around the time that the Kingdom _Animalia_ diverged from the rest of the [Eukaryotes](https://en.wikipedia.org/wiki/Eukaryote). According to [timetree.org](http://www.timetree.org/), an online resource for estimating when different taxa diverged from one another, this occurred around \~6e8 years ago. In seconds, this is \~1e16 seconds.

> - **Total amount of computation occurring at a given point in time:** [This blog post](https://reducing-suffering.org/how-many-wild-animals-are-there/) attempts to estimate how many individual creatures in various taxa are alive at any given point in time in the modern period. It implies that the total amount of brain computation occurring inside animals with very few neurons is roughly comparable to the amount of brain computation occurring inside the animals with the largest brains. For example, the population of [nematodes](https://en.wikipedia.org/wiki/Nematode) (a [phylum](https://en.wikipedia.org/wiki/Phylum) of small worms including _C. Elegans_) estimated to be \~1e20 to \~1e22 individuals. Assuming that each nematode performs \~10,000 FLOP/s,the number of FLOP contributed by the nematodes every second is \~1e21 \* 1e4 = \~1e25; this doesn't count non-nematode animals with similar or fewer numbers of neurons. On the other hand, the number of FLOP/s contributed by humans is (\~7e9 humans) \* (\~1e15 FLOP/s / person) = \~7e24. The human population is vastly larger now than it was during most of our evolutionary history, whereas it is likely that the population of animals with tiny nervous systems has stayed similar. This suggests to me that the _average_ ancestor across our entire evolutionary history was likely tiny and performed very few FLOP/s. I will assume that the "average ancestor" performed about as many FLOP/s as a nematode and the "average population size" was \~1e21 individuals alive at a given point in time. This implies that our ancestors were collectively performing \~1e25 FLOP every second on average over the \~1 billion years of evolutionary history.

> This implies that the total amount of computation done over the course of evolution from the first animals with neurons to humans was (\~1e16 seconds) \* (\~1e25 FLOP/s) = **\~1e41 FLOP**

Unlike what the reader might suspect, I don't particularly want to take issue with the assumption that "...the total amount of brain computation occurring inside animals with very few neurons is roughly comparable to the amount of brain computation occurring inside the animals with the largest brains." I haven't looked at the literature on this, and the assumption seems a prima facie plausible.

Instead, I don't buy the assumption that to simulate evolution, it would be enough to simulate the behaviour of all the neurons throughout history. Instead, I think that one would also have to simulate the stimuli to which these neurons are exposed. For instance, in the case of AlphaGo, one not only has to simulate the inner workings of the model, but also the state of the Go board. Likewise, to simulate evolution, one would not only have to simulate the neurons of the brains in it, but also the state of the world in which they are.

It's not clear to me how big a world would one need to spawn intelligence life. For instance, I think that Minecraft requires on the order of ~500 GFLOPs to run[^1]. Conversely, if we want a planet as large as Earth, with a complex atmosphere & so on, this would exceed the computing power of current supercomputers[^2].

Working through these details might be a neat research projects, though I'm not sure it would change many estimates other than my own. In any case, this would involve:

1. Coming up with estimates of what the least fine-grained world that we would expect might be able to produce intelligence life due to evolution
2. Calculating how much compute it would take to in fact simulate it.

[^1]: This page: <https://minecraft.fandom.com/wiki/Java_Edition_hardware_requirements> claims that recommended requirements are an Intel Core i5 for the CPU, and a Radeon 7 200 for the GPU. The Intel CPU has a processing speed of 37.73 GFLOPs (<https://setiathome.berkeley.edu/cpu_list.php>) and the Radeon of 467.2 GFLOPs (<https://en.wikipedia.org/wiki/Radeon_200_series#Radeon_R7_250>). I'm not sure what .
[^2]: E.g., see [this article](https://www.reuters.com/article/us-britain-weather-computer-idUSKBN20B00U) on a £1.2 billion, 60 petaFLOPS supercomputer built solely for better atmospheric prediction/simulation. Simulating the rest of the world would require more computing power on top of that. I haven't really
