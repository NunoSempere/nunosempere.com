Five slightly more hardcore Squiggle models.
==============

Following up on [Simple estimation examples in Squiggle](https://forum.effectivealtruism.org/posts/vh3YvCKnCBp6jDDFd/simple-estimation-examples-in-squiggle), this post goes through some more complicated models in Squiggle.

## Initial setup

As well as in the [playground](https://www.squiggle-language.com/playground), Squiggle can also be used inside [VS Code](https://code.visualstudio.com/), after one installs [this extension](https://github.com/quantified-uncertainty/squiggle/tree/develop/packages/vscode-ext), following the instructions [here](https://github.com/quantified-uncertainty/squiggle/blob/develop/packages/vscode-ext/README.md). This is more convenient when working with more advanced models because models can be more quickly saved, and the overall experience is nicer.

<img src='https://i.imgur.com/ldmrmmX.png' class='.img-medium-center'>

## Models

### AI timelines at every point in time

Recently, when talking [about AI timelines](https://forum.effectivealtruism.org/posts/W7C5hwq7sjdpTdrQF/announcing-the-future-fund-s-ai-worldview-prize), people tend to give probabilities of AGI by different points in time, and about slightly different operationalizations. This makes different numbers more difficult to compare. 

But the problem with people giving probabilities about different years could be solved by asking or producing probabilities for all years. For example, we could write something like this:

```
// Own probability
_sigma(slope, top, start, t) = {
    f(t) = exp(slope*(t - start))/(1 + exp(slope*(t-start)))
    result = top * (f(t) - f(start))/f(start)
    result
}

advancedPowerSeekingAIBy(t) = {
    sigma_slope = 0.02
    max_prob = 0.6
    first_year_possible = 3

    // sigma(t) = exp(sigma_slope*(t - first_year_possible))/(1 + exp(sigma_slope*(t-first_year_possible)))
    // t < first_year_possible ? 0 :  (sigma(t) - sigma(first_year_possible))/sigma(first_year_possible)*max_prob
    sigma(t) = _sigma(sigma_slope, max_prob, first_year_possible, t)
    t < first_year_possible ? 0 : sigma(t) 
}
instantaneousAPSrisk(t) = {
    epsilon = 0.01
    (advancedPowerSeekingAIBy(t) - advancedPowerSeekingAIBy(t-epsilon))/epsilon
}

xriskIfAPS(t) = {
    0.5
}

xriskThroughAps(t) = advancedPowerSeekingAIBy(t) * xriskIfAPS(t)
```

This produces the cumulative and instantaneous probability of “advanced power-seeking AI” by/at each point in time:

<img src='https://i.imgur.com/7wlIKtp.png' class='.img-medium-center'>

And then, assuming a constant 95% probability of x-risk given advanced power-seeking AGI, we can get the probability of such risk by every point in time:

<img src='https://i.imgur.com/hPuGhUE.png' class='.img-medium-center'>

Now, the fun is that the x-risk is in fact not constant. If AGI happened tomorrow we’d be much less prepared than if it happens in 70 years, and a better model would incorporate that. 

For individual forecasts, rather than for models which combine different forecasts, <[forecast.elicit.org](https://forecast.elicit.org/)\> had a more intuitive interface. Some forecasts produced using that interface can be seen [here](https://www.lesswrong.com/posts/hQysqfSEzciRazx8k/forecasting-thread-ai-timelines). However, that interface is currently unmaintained. Open Philanthropy has also produced a number of models, generally written in Python.

### More detailed expected value estimates for potential career pathways

In the [preceding post](https://forum.effectivealtruism.org/posts/vh3YvCKnCBp6jDDFd/simple-estimation-examples-in-squiggle#Expected_value_for_a_list_of_things__complexity___2_10_), I presented some quick relative estimates for possible career pathways. Shortly after that, Benjamin Todd reached out about estimating the value of various career pathways he was considering. As a result, I created [this more complicated spreadsheet](https://docs.google.com/spreadsheets/d/1QATMTzLUdmxBqD2snhiAkH-_KvwbhGdlYaU8Ho7kjDY/edit?usp=sharing):  
 

<img src='https://i.imgur.com/MT1aVtk.png' class='.img-medium-center'>

You can see a higher quality version of this image here: <[https://i.imgur.com/hvq0SeM.png](https://www.google.com/url?q=https://i.imgur.com/hvq0SeM.png&sa=D&source=docs&ust=1665398336469760&usg=AOvVaw1SMBMKFOfTRKgaOAsxR9np)\>

Instead of using relative values, each row estimates the value of a broad career option in dollars, i.e., in relation to how much EA should be prepared to pay for particular outcomes (e.g., the creation of 80,000 hours, or CSET). One interesting feature of dollars is that they are a pretty intuitive measure, but it breaks down a bit under interrogation (dollars in which year? adjusted for inflation? are we implicitly assuming that twice as many dollars are twice as good?). But as long as the ratios between estimates are meaningful, they are still useful for prioritization.

For a model in the above style which is more hardcore and more complex, see [here](https://github.com/quantified-uncertainty/squiggle-models/blob/master/bahamas/hierarchical-impact-estimate-ftx-bahamas-joel2.squiggle) (or [here](https://www.squiggle-language.com/playground#code=eNrNWW2P47YR%2Fiu8%2FRLL1sqyHV8uToOiQdPkQ4MGveZTNxfQEmUTK5E6kvLGPe9%2FzwwpyXqhdo0CQWLgvLI4MxzOPPNwyPt0p4%2Fy6X1VFFSd73ZGVSy0r75NuZGqecMFN5zm7z9W%2FHDI2XujuDjc7e6WS%2FI9y0um9IP4H1OSfE2EVAXNZ3FI4iiO41XwIHRVwMCnS861uRDF0iphM%2FwRElQKYYgmSUhONL8QeCILfHwOnh%2FEg4Ap%2FlGJxHApYJLkKKVmMxE%2BBmAyowk4CY7NRECWZNZ9Qe4JyMw7r%2BAnOLPnQhb2ZyjCEq3wjDz%2B5WtBzJEJ0k7gtMsPj%2FA9W4G1MvjQWGW5ZiSuvfueM0VVcuQJzUnO9w%2FiwAz4b5dslwRfUTaD74A8N6OKJZXS%2FMSsGGjnqWLi8iBI%2FUmkMEwYDeMgXjE9a4SCqxCMoEBBy1kjHxI3wUDqQTTR%2FMnwnJvzg6jcAzrgZBNZlFKgkV3zCj%2BlkpAxw08g3BvAj6iKkskyZzsCUBEJNRA8B4F1HL4LwlUcbrZBCBNTQ6g4E5iBGF6wvh2j6InlOzJbQyK3McZ%2BaG%2BFkFoFYRzCV187rRRFhOzIOiQwVQHBOOqRp0cJMd%2BzTKquuxtiJPkSjIdk9XZoudGimWGqqxStrFq0mVBkWcYwaEwwrYsqN7zMASlPkETGRcbyXD7pIy93sKh3aGoVba3vsBYoLkQjuUr1bVtAyOxJqscdKX6drR%2BtgTh%2BDIl73trn%2F8bRlzZoPw%2B9wxj19baxU%2Fyhr4bVgRl7oohFEACcaZ4yRXKJXlKe62hkmwuSyjynSi9tKojMCHrbF8wATRdE16X%2F3sYdYIlDUQuwsUzayDTpH4uYRsThayzwsYKFuSpoprvme9l75YDjNZF3LNya%2BLEhFAYrrUvz1vRY1hGMnbCDBl8cgb5SoC2DJWVnmKPyWBLS9qOtdKaBNLnKz%2BTID0fCNFQrQB6YRUhDdKWYBcWJqXoI4aoLc%2BzbfL7%2BfO6gL%2BcZw%2FLfA8%2Ba341OIDClVIgJmTmDkLunoyzstFyXUmu%2B900DsN8i%2BNeBLew4%2BnxYOy%2FbPtKyZEInsgI%2BVrj7QA7zc2ei2FFHjPWO1HEjc3ABWyPLYbMxLD0zqC1b%2B6u6hENyv5q7X5s4DnpVDOlfDSZpbcEyuKjNNS7%2BG%2FJuwM81WgNG%2FSIe%2BogK%2BdkizzHQ1lEJsEf9OCSSPyv%2F2KRNEJAde4mBylUrdSviPEbWNxqZgNbYImsN3owlP3U6Ix6seAipFe9gw89HELU5rnoOjsI2j4T0h9FRgjDZS7c4%2FbvxkaLQ93WM%2FRO636gq%2FyNnYICghSGGURiWDS61am%2Bwx12%2FQUZv30Fpr4M3Q3VYVbsowNQeOnQ0Bn9Set65znztSCiON7btgJ4f%2BkhYgJ0VC8ZN8Ndpy30g%2Bqh0E8bR28BPc11LVyJZdSiheR5SQq%2FAX164NvTMUiggVq961ax6%2FZo7XdWXHbP7Rbx9xbVBj%2BprUS0hWUcuvoJASmrOJ03%2BgW8DP3mt49jDNWgCJ4hehMhYMcm8mn0ITLUpY70Jdigx3JMuXjPinWhSd5BOP9U5tSZJUGObGLMJ3xCPehvxshSwEFT8EcrfHLmG3DMtPoMf9BGISQBOgD4xSrajp4V9bKvrK%2BK1ieY%2B04QXwNYJN4gp1MYg0L07tYEJ2mcvMKnQ2chrEs6Asj7KAaOWihkwAig9NE55OF3MZ0k2L%2BeYxXuXnHkd6GBubuJXLROO3PBp1EPtTaVgOzmwE89dcTa1uRo3XHtwGDYx05O%2FFvNQw2b8IGXqpueIUZq46osHspn5FTZZUwmo81X8DZr8PP5mIOTsWLN4F6JZnqElDKySsiBggaRc04NirAA%2FyRG2o8hT3c7Qxb9nu8FoEBx%2FFXdk%2B7Hxn5Nq8etix2L6KuZZrs9j2MXhX0oWRE%2BCAT9%2BaOz5IeP6OMKGxT9UkoZTV56SPeBecq2HxzugwNbCaG%2FewsZsN%2BdtfNNhYT82hUB0%2BFr5NzAben6ShuZe7a3TXvu1bXBbPf8GM9pf%2Fr8dpZ5lqsl1o9E1nlN9bi3YCddUM9uR7Edoom%2Bs5bsx8cO4kWx5%2BnWK7necJwTsTcwFZ3VFNSQUtja8L%2FIwmGvIoe7ogXLha4De2aZli5dWCAF0txIcSrCYufNf9M4HK2uPinTS7mZk92rW3pONOspO81%2FAFCdLUzyrZ2jPpJvQAyAbiguZwIUdjcbxmADHVby3Ss91SCv8ku9jPTgEL9CzBVnPy3UAh%2BJFEdyU8pQW9DBOdXMh88u3f1OsrExddNea%2FaGp2G3dEr61Fwc%2Fe1Pr9u%2BBLQ9vTNwO2IQ4R30Z%2BfiLgrC58cjruCctXZ0pH%2F1TzVF5OrbusQmxdb2%2Bde7ddDfXzuBFPRxdX14F%2FVfSoORuvGdXnagrEEyd9ryafZFg2Ml4ddxQMOAOr2ifVYLRbuhVqseCiVs0r05PIhhi3KvihjqywFZRFPkSMOtlYDEM7KKJ1qJd12Lg833tSTC3IamR4v7U%2F1lx6oa7BkZQ%2Fy%2FG37kuc4o3o8tlXyJ6KZnL5atoWy4%2F2RksyHavYSt0sm6i3et4quWbtngSS7Vc2yBMg6KRdPeprwGhlm6YbhIEtZzNze4WHC%2BXTc5chlpGgzchuYeN7jtgs%2B%2BsJPlA%2FiXyc9M1YyK%2FIooVwOzaHnTwWu%2FN3fNvDguDhg%3D%3D) in the Squiggle playground.)

### A sketch of a more parsimonious estimate of AMF’s impact

The estimates in [this post](https://forum.effectivealtruism.org/posts/4Qdjkf8PatGBsBExK/adding-quantified-uncertainty-to-givewell-s-cost), and overall GiveWell’s [estimates](https://docs.google.com/spreadsheets/d/1tytvmV_32H8XGGRJlUzRDTKTHrdevPIYmb_uc6aLeas/edit#gid=1377543212) of the value of AMF had been bothering me because they divide the population into very coarse chunks. This is somewhat suboptimal because the first chunk ranges from 0 to 4 years, but malaria mortality [differs a fair amount between a newborn and a four-year-old](https://nunosempere.com/blog/2022/09/28/granular-AMF/).

Instead, we could express impact estimates in a more elegant functional form. I’ll sketch how this would look like, but I’ll stop halfway through because at some point the functional form would require more research about mortality at each age.

The core of the impact estimate is a function that takes the number of beneficiaries, the age distribution of a population, and the benefit of that intervention for someone of a given age, and outputs an estimate of impact.

In Squiggle, this would look as follows:

```
valueOfInterventionInPopulation(num_beneficiaries, population_age_distribution, benefitForPersonOfAge) = {
  age_of_beneficiaries = sampleN(population_age_distribution, num_beneficiaries)
  benefits_array = List.map(age_of_beneficiaries, {|a| benefitForPersonOfAge(a)})
  total_benefits = List.reduce(benefits_array, 0, {|acc, value| acc + value})
  total_benefits
}
```

So for example, if we [feed](https://www.squiggle-language.com/playground/#code=eNq9kk1rhDAQhv9K8KQ0LRaWHhZ66KGFhdJd8BqQWR1tIE5sTLZd1P%2FeqLWwXXePvU3m43nnI23QvOvPxFUVmGOwtsYhH13PubTazB5J0kpQyYeTZakwsUZSGayDAyiH22JDFs0ByUpNG9rp2ikY7JBcle6RsJCZBCOx4az%2BjaZQYprLxsP2bnBwNuXaF212aBpN2%2BKpxIg9slYQY0O%2BLk6BPtZAVSt8C6%2BSzzqJBuKPXpOCMXD0rFdfdFdBHS5pcdZ20C03GULUj0irLah0Bs9Ig7nLMDzV4ywekVnG2bjJjnmb3UyPJZygXpCgs2G8zH0cx4Ku7GDM8Ty28nlKFpjiV42ZBcqGyVdj7MHHLsw3HeJv4e1wFUH%2F8xGC%2Fhu08wBc) the following variables to our function:

```
num_beneficiaries = 1000
population_age_distribution = 10 to 40
life_expectancy = 40 to 60
benefitForPersonOfAge(age) = life_expectancy - age
valueOfInterventionInPopulation(num_beneficiaries, population_age_distribution, benefitForPersonOfAge)
```

Then we are saying that we are reaching 1000 people, whose age distribution looks like this:

<img src='https://i.imgur.com/GvyMyqW.png' class='.img-medium-center'>

This could use a bit more work to resemble an actual population pyramid.

and that the benefit is just the remaining life expectancy. This produces the following estimate, in person-years:

<img src='https://i.imgur.com/BSbneRi.png' class='.img-medium-center'>

But the assumptions we have used aren’t very realistic. We are essentially assuming that we are creating clones of people at different ages, and that they wouldn’t die until the end of their natural 40 to 50-year lifespan.

To shed these unrealistic assumptions, and produce something we can use to estimate the value of the AMF, we have to:

1.  Add uncertainty about the shape of the population, i.e., uncertainty about how the population pyramid looks
2.  Add uncertainty about how many people a distribution reaches
3.  Change the shape of the uncertainty about the benefit to more closely resemble the effects of bednet distribution

The first two are relatively easy to do.

For uncertainty about the number of beneficiaries, we could naïvely write:

```
valueWithUncertaintyAboutNumBeneficiaries(num_beneficiaries_dist, population_age_distribution, benefitForPersonOfAge) = {
  numSamples = 1000
  num_beneficiaries_samples_list = sampleN(num_beneficiaries_dist, numSamples)
  benefits_list = List.map(num_beneficiaries_samples_list, 
    {|n| 
      valueOfInterventionInPopulation(n, population_age_distribution, benefitForPersonOfAge)
    })
  result = mixture(benefits_list)
  result
}
```

  
However, that would be very slow, because we would be repeating an expensive calculation unnecessarily. Instead, we can do [this](https://www.squiggle-language.com/playground/#code=eNrFk9tKw0AQhl9lyVWqUVPwAIIXFRQKUoWg3gTCNp2kC5vZuAe1tH13dzc9m1bQC%2B82M7PfzM7%2FZxqosfhITFVROQmutTQQ%2BdDdiGkhlxGGTDPKkzfDypJDoiXDMrgO3ik38Fj0UYN8B9RMYB%2BfRG04decQTZUNAaFgOaOSgYpIvcpmtIRsxJSFDY0LRKSp1fdCPoFUAh%2BLXgkdckOmKRLi6kWxDbQ5RauawyA8SP42SccRF%2F1URqWkE8t6sJdOK1qHbb0iMp3RWfuQIe3MPVILTXm2BC%2BREkYmh3C7X0Rij8zziPhNzog9k%2BPmow2X4jzFFH3%2BlenxM%2BYgNWWoJ72hMHpgqtvNib8L4NfyJxUkFCDBNh7YSDeO463gixvNJn5yxpryq2HcFogTNfHiq2aU7iK48%2BTGICrjFrzhl327WVO3PbK4vrLI4UZOWZztLOYIz9YPbwSWoAx33Ip9aiM3POIoGyUL7VNsn9sSruLYGYZ0L5woB7bqt%2BUqz20dZwVk8FlDrinm7ic497lLm9tj9cYNuxdP3A%2F6T%2BYM5l83mcoE):

```
valueWithUncertaintyAboutNumBeneficiaries(num_beneficiaries_dist, population_age_distribution, benefitForPersonOfAge) = {
  referenceN = 1000
  referenceValue = valueOfInterventionInPopulation(referenceN, population_age_distribution, benefitForPersonOfAge)

  numSamples = 1001
  num_beneficiaries_samples_list = sampleN(num_beneficiaries_dist, numSamples)
  benefits_list = List.map(num_beneficiaries_samples_list, {|n| referenceValue*n/referenceN})
  result = mixture(benefits_list)
  result
}
```

  
That is, we are calculating the value for a beneficiary population of 1000, and then we are scaling this up. This takes about 6 seconds to compute in Squiggle.

Now, when adding uncertainty about the shape of the population, we are not going to be able to use that trick, and computation will become more expensive. In the limit, maybe I would want to have a distribution of distributions. But in the meantime, I’ll just have a list of possible population shapes, and [compute the shape of uncertainty over those](https://www.squiggle-language.com/playground/#code=eNrFVN9r2zAQ%2FleEn5zOa5VRd1DoQwctFEZaMFsf5mEU5%2BwI5JMnS11DnP99ku3UTpofkD307Xz36btP95219Kq5%2FBuZomBq4V1rZSBoUnczrqVaZzhyzZmI%2Fhie5wIirTjm3rX3woSBx%2BwBNagXQM0lPuCTLI1gLvbRFMkUEDKecqY4VAEp36oJyyGZ8cqSTY1LBKTF6nupnkBVEh%2Bz2xxG5IYsYyTE4WW2SWhrFStKARP%2FIPM7JSPH2PWrEqYUW1iu7%2FbQecFKf1evgCxrVu8W6bPRqqHUUjORrInXlApmJgV%2Fs19AaEOZpgFpJlkTG5NP7ccuuhhXMcbY1J%2B5nv%2FAFJRmHPXidiqNnpji21DxewOasfyXCwoyUGAbT2xmTCndSP500mzh2Gb0LCeJcVMgztSoMb9qpYy75NaV2wWpEmGJB%2FuybzY96%2BaOdMffVuRwI%2Bcs1luDOcOL%2FuKtwQoqIxxvwV%2B1UYMdcSwDyDHv%2B%2BFGc1bCKdZ3wg%2F7v3cex4mX9QFMTT5krU9z4dhdLdMvLf0vAbmko4C0YdiHV10Y9oCwB4Q9YEx7hIvDQWwxv2PcPQsr4Cul7v0g49D9o4JnkMBrCalmmLq37pK66pWt7XnRWtO3D3527%2FAH7mGM3uofoJ96rQ%3D%3D):

```
valueWithUncertaintyAboutPopulationShape(
  num_beneficiaries_dist, 
  population_age_distribution_list, 
  benefitForPersonOfAge
  ) = {
  benefits_list = List.map(population_age_distribution_list,
   {|population_age_distribution| 
     valueWithUncertaintyAboutNumBeneficiaries(
       num_beneficiaries_dist, 
       population_age_distribution, 
       benefitForPersonOfAge
   )})
  result = mixture(benefits_list)
  result
}
population_age_distribution_list = [
  to(2, 40), to(2, 50), to(2, 60), 
  to(5, 40), to(5, 50), to(5, 60), 
  to(10, 40), to(10, 50), to(10, 60)
]
```

  
We still have to tweak the benefits to better capture the benefits of distributing of malaria nets. One first attempt might look as follows:

```
benefitForPersonOfAge(age) = {
  result = age > 5 ? mixture(0) : {
    counterfactual_child_mortality = SampleSet.fromDist(0.01 to 0.07)
    // https://apps.who.int/gho/data/view.searo.61200?lang=en
    child_mortality_after_intervention = counterfactual_child_mortality/2
    chance_live_before = (1-(counterfactual_child_mortality))^(5-age)
    chance_live_after =   (1-(child_mortality_after_intervention))^(5-age)
    value = (chance_live_after - chance_live_before) * (life_expectancy - age)
    value
  }
  result
}
```

  
That is, we are modelling this example intervention of halving child mortality, and for child mortality to be pretty high. The final result looks like [this](https://www.squiggle-language.com/playground/#code=eNrFVFFvmzAQ%2FiunPEFHwDRJK0Xqqk7bpEpTWyna9rBsyCUHWAKbGTttleS%2FzwZSkjRNpO6hTxx35%2B8%2B33e%2BRa%2FKxMNEFwWVT72xkhq92vVlxpSQaw%2FjTDGaT%2F5qlqY5TpRkPO2Ne3Oaa7xNrrlCOUeumODX%2FE6UOqfWdrguonvkmLCYUcmw8qB8jkY0xWjGKgN2r63DgyZXfRXyDmUl%2BG1ylaILF7CYcgCbL5JtQBOraFHmeOMcRH7BxLWIbb0qolLSJ4P1zRzyC1o6%2B2p5sFjS5X6SDnVXNaQSiubRGngNKXGmY3S263lAasg49qDu5BKMDR%2Ban31wU76a8imv4z%2BZyr7zGKWijKunq3uh1Y0uPm0yfilA3Zb%2FUkFighJN4RvjCQkhW84flpoJHJuMDuVNZGwXwIo6qcWvGiph69y5cjMgVZQb4I15ea03Her2jLTHn0fkcCGrLF%2FuNOaEB93FG4ElVjq3uAV7VFpuzIhF2Ug5pn3X3ElGS3yL9C3xw%2Fq%2F2o%2FjwIvlgZwlvMtYv02FY3c1SL%2BUcEIPBsT1Gmv4bJ0Zy7xs57QJt%2BawM9cJgy5h0CUMuoShByPi%2Fp7y%2FR0xNM4JsVsEwpF9qTlLMMLHEmNFeWw33pDY6JmJvbLXtp5%2B2yXjhI8wgksoHh3iwriJA8RC2xef0Fhps7fijOWzqBBGzJwpW655WRNUfiJF8dlwdIhPQsvBfM%2FdBiYIIFOqrMZBQMuy8h8y4ZtpCNJMBDOqaDBn%2BOBXSKXwz8JTQi5zytML5C2L7bIRTQyniG3sIsPkMNXgdA1l%2BoRG0zma%2FiZC2uXmhH3n8HHX%2FeOM%2BrTeVLswNRuDAg3OUa67WPN2xTovQft7%2BLpwAs6u7n3YwbPmamfM323VTHlv9Q%2B7JiAl) (archived [here](https://gist.github.com/NunoSempere/715fd697ff3ebbb704e4a239e559d148)). The output unit is years of life saved. As is, it doesn’t particularly correspond to the impact of any actual intervention, but hopefully, it could be a template that GiveWell could use, after some research. 

But for reference, the distribution’s impact looks as follows:

<img src="https://i.imgur.com/bejosHk.png" class='.img-medium-center'>

### Calculate optimal allocation given diminishing marginal values

Suppose that we have some diminishing marginal return functions. Then, we may want to estimate the optimal allocation to each opportunity.

We can express diminishing marginal returns functions using two possible syntaxes:

```
diminishingMarginalReturns1(funds) = 1/funds
diminishingMarginalReturns2 = {|funds| 1/(funds^2)}
```

The first syntax is more readable, but the second one can be used without a function definition, which is useful for manipulating functions as objects and defining them programmatically, as explained in this footnote ⤻ [\[1\]](#fnmxi7a8ll6t).

Once we have a few diminishing marginal return curves, we can put them in a list/array:

```
diminishingMarginalReturns1(funds) = 1/(100 + funds)
diminishingMarginalReturns2 = {|funds| 1/(funds^1.1)}
diminishingMarginalReturns3 = {|funds| 100/(1k + funds^1.5)}
diminishingMarginalReturns4 = {|funds| 200/(funds^2.2)}
diminishingMarginalReturns5 = {|funds| 2/(100*funds + 1)}
uselessDistribution(funds) = 0
negativeOpportunity(funds) = 0

listOfDiminishingMarginalReturns = [
  diminishingMarginalReturns1, diminishingMarginalReturns2,
  diminishingMarginalReturns3, diminishingMarginalReturns4,
  diminishingMarginalReturns5, uselessDistribution,
  negativeOpportunity, {|funds| {1/(1 + funds + funds^2)}}
]
```

And then we can specify our amount of funds;

```
availableFunds = 1M // dollars
calculationIncrement  = 1 // calculate dollar by dollar
Danger.optimalAllocationGivenDiminishingMarginalReturnsForManyFunctions(
  listOfDiminishingMarginalReturns,
  availableFunds, 
  calculationIncrement
) 
```

So in this case, the difficulty comes not from applying a function, but from adding that function to Squiggle. This can be seen [here](https://github.com/quantified-uncertainty/squiggle/blob/develop/packages/squiggle-lang/src/rescript/FR/FR_Danger.res#L278).

Other software (e.g,. Python, R) could also do this, but the usefulness of the above comes from integrating that into Squiggle. For example, we could have an uncertain function produced by some other program, then take its mean (representing its expected value), and feed it to that calculator.

The [Survival and Flourishing Fund](https://survivalandflourishing.fund/) has some [software](https://youtu.be/jWivz6KidkI?t=487) to do something like this. It has a graphical interface which people can tweak, at the expense of being a bit more simple—their diminishing marginal returns are only determined by three points

### Defining a toy world

Lastly, we will define a simple toy world, which has some population growth and some economic growth, as well as some chance of extinction each year. And its value is defined as a function of the consumption of each person, times the chance that the world is still standing. 

For practical purposes, after some set point we stop calculations, and we calculate the remaining value as some function of the current value. We can understand this as either a) the heat death of the universe, or b) an arbitrary limit such that we are interested in the behaviour of the system as that limit goes to infinity, but we can only extend that limit with more computation.

This setup allows us to coarsely compare an increase in consumption vs an increase in economic growth vs a reduction in existential risk. In particular, given this setup, existential risk and economic growth would be valued less than in the infinite horizon case, so if their value is greater than some increase in consumption in this toy world, we will have reason to think that this would also be the case in the real world.

The code is a bit too large to simply paste into an EA Forum post, but it can be seen [here](https://github.com/quantified-uncertainty/squiggle-models/blob/master/toy-world/toy-world.squiggle). For a further tweak, you can see leaner code [here](https://github.com/quantified-uncertainty/squiggle-models/blob/master/toy-world/toy-world.squiggleU) which relies on the import functionality of the [squiggle-cli-experimental package](https://github.com/quantified-uncertainty/squiggle/tree/develop/packages/cli).

We can also look at the impact that various interventions have on our toy world, with further details [here](https://docs.google.com/spreadsheets/d/1WnplTYJJMeh0zXVUTPBaihE7n1kneW5LDidLvJcGcv4/edit?usp=sharing):

<img src='https://i.imgur.com/HCg2g5r.png' class='.img-medium-center'>

We see that of the sample interventions, increasing population growth by 0.5% has the highest impact. But 0.5%/year is a pretty large amount, and it would be pretty difficult to engineer. So further work could look at the relative difficulty of each of those interventions. Still, that table may serve to make a qualitative argument that interventions such as increasing population growth, economic growth, or reducing existential risk, are probably more valuable than directly increasing consumption.

## Conclusion

I presented a few more advanced Squiggle models. 

A running theme was that expressing estimates as functions—e.g., the chance of AGI at every point in time, the impact of an intervention for all possible ages, a list of diminishing marginal return functions for a list of interventions, a toy world with a population assigned some value at every point in time—might allow us to come up with better and more accurate estimates. Squiggle is not the only software that can do this, but hopefully it will make such estimation easier.

1.  **[^](#fnrefmxi7a8ll6t)**
    
    We can write:
    
    ```
    listOfFunctions = [ {|funds| 1/(funds^2)},  {|funds| 1/(funds^3)}]
    ```
    
    or even 
    
    ```
    multiplyByI(i) = {|x| x*i}
    listOfFunctions = List.map(List.upTo(0,10), {|i| multiplyByI(i)})
    ```
    
    ```
    or without the need for a helper:
    listOfFunctions2 = List.map(List.upTo(0,10), {|i| {|x| x*i}})
    listOfFunctions2[4](2) // 4 * 2 = 8
    ```
    
    This is standard functional programming stuff, and some functionality is missing from Squiggle, such as _List.length_ function. But still.
