Legalize acetylcysteine: An open letter to the UK's MHRA
========================================================

## Part I: Demagoguery

This is the map of maximum Celtic expansion, in circa 270 BC, per [Wikipedia](https://upload.wikimedia.org/wikipedia/commons/0/08/Celtic_expansion_in_Europe.svg):

![](https://upload.wikimedia.org/wikipedia/commons/0/08/Celtic_expansion_in_Europe.svg)

Since then, the Spaniards have further developed into Gazpacho-drinking siesta-sleepers and the Britons have developed into tea-drinking weather-contemplators[^1]. Still, my understanding is that population differences are to a great degree cultural, and that the basic plumbing remains pretty much the same. My understanding is also that 

Imagine, then, my surprise, when in the middle of being sick in the UK, I find out that an extremely common medicine used to treat the cold in Spain throughout my childhood just wasn't commonly available in the UK. This medicine is [acetylcysteine](https://en.wikipedia.org/wiki/Acetylcysteine)—known in Spain under the brand name "Fluomicil". It's purpose is to decrease the thickness of the mucus so that it can be expelled, so that the patient can better breathe. In my experience, this is particularly crucial at night, because if the nose is blocked, you will breathe through the mouth and end up having a sore throat, and generally not sleep as well. 

Instead of using acetylcysteine, the UK uses other less efficaceous medicaments, such as nose sprays, which don't work as well through the night. They aren't as useful once the nose is already blocked. And they are more annoying to use, which means that people may forget or use them less.

## Part II: Cost-effectiveness analysis

I work as a forecaster, not as a doctor or as a medical researcher. So there are surely factors I'm missing. For instance, maybe living for two milenia under lousy weather has maybe made the population of Britain more immune to having blocked noses, and this could mean that nose sprays are a better tradeoff than acetylcysteine. I really wouldn't know, though it would surprise me. 	

Still, as a forecaster I can offer the following estimation:

Per the [NHS inform website](https://www.nhsinform.scot/illnesses-and-conditions/infections-and-poisoning/common-cold#colds-in-children): 

> Children get colds far more often than adults. While adults usually have two to four colds a year, children can catch as many as 8 to 12.

According to the [latest data from the Office of National Statistics](https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/populationestimates/bulletins/annualmidyearpopulationestimates/mid202), the population pyramid of the UK looks as follows:

![](./population.png)

meaning that there are 67,081,234 people, of which 20.1% (13,468,262) is under 16. I also estimate that acetylcysteine makes an illness somewhere between 1% and 10% better. 

Putting this together, I can [estimate the following](https://www.squiggle-language.com/playground#code=eNqFkV9rwjAUxb%2FKpbChnfPfnAPBR5%2FGXiZ7C4TYxBpIb7r0dlLE775E7aa1c2%2Fl9OR3TnJ2UbGx22WZZcJV0YxcqXoHaSE1WVcrGjVpYZafpU5To5bkNKbRLGI4GMCiIJ0JUrAqnVQIdg1SF0oUimFu89II0ha5XfOPV5jD9OXN687m1h30ZKONdP7cHIb98XAEHjke9kd3DMmSMFzI0hBPrJEFz5XjlRLOmztjIAuTLsRwlRJDZwSP0JLSram10AKeBPBo%2FAe5BfrDDChPuNH7AW7GM5Sl%2B8kL%2F8KLHfuEh5GiKuq08M23mja177xEDE3QeclD1m%2FTBmkAT9Nnhpfj6iwXCYVxRaKoMklVkNKowGLb8N7u7JfKFNKR3DgVxvZb%2B4v5qRmmQmPByfKV4htxeZlT2Rj%2BYZ4avysqHcJ96JIbUTHcMYQr3uxK6QVbo8isKXjTnmG0%2FwYbRy4m):

```
// Estimate burden of disease
population_of_UK = 67M
proportion_children = 0.201 // 20.1%
total_adult_colds_per_year = (2 to 4) * population_of_UK * (1 - proportion_children)
total_children_colds_per_year = (4 to 12) * population_of_UK * proportion_children
total_colds = total_adult_colds_per_year + total_children_colds_per_year
duration_of_cold = 6 to 12 // days
total_days_with_cold = total_colds * duration_of_cold
total_cold_years = total_days_with_cold / 365

// Estimate impact of acetylcysteine on burden of disease
improvement_with_acetylcysteine = 0.01 to 0.1
gains_to_be_had = total_cold_years * improvement_with_acetylcysteine

// Return & display
{
  total_cold_years: total_cold_years,
  gains_to_be_had: gains_to_be_had,
}
```

That is, I arrive at an estimate of 6M (1.7M to 9.2M) cummulative person-years spent having a cold in Britain:

![](./cold_years_per_year.png)

and a potential improvement from adopting acetylcysteine of 250,000 (53,000 to 640,000) "quality-adjusted-sickness-years"—an intutitive, ad-hoc unit that I just made up:

![](./gains-to-be-had.png)

The weakness of the method is that my subjective estimates of the 1% to 10% quality of life improvement might be off, or that my estimates of how often people are sick might be inaccurate—6M years of cold per year does seem a bit high. I'm also not really familiar with how potential alternatives, such as carbocisteine, are used in the UK. Still, I think that this rough calculation does show that having better medicaments is of great importance. And the Spanish doctors I've spoken expressed shock and disbelief that acetylcysteine was not available in the UK.

But while an abstract argument may have been made, the action and followup remains. And it falls on the brave and hardworking souls at the [MHRA](https://www.gov.uk/government/organisations/medicines-and-healthcare-products-regulatory-agency) to send [a Message to Garcia](https://courses.csail.mit.edu/6.803/pdf/hubbard1899.pdf): Legalize acetylcysteine. 

---

Sources

https://www.nhs.uk/medicines/carbocisteine/#:~:text=A%20mucolytic%20helps%20you%20cough,chronic%20obstructive%20pulmonary%20disease%20(COPD)
https://www.cochrane.org/CD003124/ARI_acetylcysteine-and-carbocysteine-to-treat-acute-upper-and-lower-respiratory-tract-infections-in-children-without-chronic-broncho-pulmonary-disease
https://www.medicines.org.uk/emc/product/2916/smpc
https://www.gov.uk/government/organisations/medicines-and-healthcare-products-regulatory-agency
https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/populationestimates/bulletins/annualmidyearpopulationestimates/mid202
https://www.boots.com/sitesearch?searchTerm=mucolytics
https://www.drugs.com/drug-interactions/acetylcysteine-with-fenesin-dm-97-0-846-8166.html
https://www.amazon.co.uk/N-Acetyl-Cysteine-Nutritional-Supplements/b?ie=UTF8&node=5977697031
https://www.dbth.nhs.uk/wp-content/uploads/2017/10/Patient-Information-Leaflet-Acetylcysteine.pdf
https://bnf.nice.org.uk/drugs/acetylcysteine/
https://www.waymade.co.uk/contact-us/
https://www.medicines.org.uk/emc/product/11366/smpc
https://www.medicines.org.uk/emc/product/11366/smpc
https://www.boots.com/boots-pharmaceuticals-effervescent-powder-10-sachets-10049868
https://www.waymade.co.uk/wp-content/uploads/2021/05/PIL-Acetylcysteine-200mg.pdf
https://mhraproducts4853.blob.core.windows.net/docs/1a9869bfeac300ad6cb531ff4434f1cf90243624
https://products.mhra.gov.uk/search/?search=acetylcysteine&page=1&ter=UK&rerouteType=0

[^1]: As in, "Isn't the weather nice today, darling?"
