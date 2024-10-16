The share of the surface of the earth very urban areas is roughly ~1.7%
=======================================================

The share of the surface of the Earth which is very urban turns out to be a recurring question when estimating various catastrophic risks, like asteroids or volcanoes. Once we estimate the frequency of such events, we want to estimate the chance that they'd hit an urban area. 

There is some pre-existing literature on the topic, but I thought that making my own estimate would be quicker and easier to fit with my estimation problem.

My methodology was to download this [List of cities with over one million inhabitants](https://en.wikipedia.org/wiki/List_of_cities_with_over_one_million_inhabitants) from Wikipedia, then use a script to fetch the area in kilometers from the Wikipedia infobox. I chose to take the *metropolitan* area when this existed, i.e., including urban sprawl. I did not include the giant "prefecture-level" Chinese cities of [Chongqing](https://en.wikipedia.org/wiki/Chongqing) and [Jilin](https://en.wikipedia.org/wiki/Jilin), though I did include the "macrometropolis" around São Paulo.

My initial script had to be improved a fair bit, since not all cities had infoboxes, or in the same format. The final version can be seen [here](./src/areas.sh). In addition, I complemented it with some manual data gathering, either still from the Wikipedia page, or roughly estimating city size from Google maps. Finally, I manually checked for transcription errors (I had a few, particularly mixing up prefectures and cities, or taking the population number instead of surface area).

The resulting spreadsheet can be found [here](./src/List_of_cities_with_over_one_million_inhabitants_with_areas_manual.csv). 

All in all, metropolitan areas add up to 1.57M km2. The total surface of the Earth is 510M, so "very urban" share per our operationalization above corresponds to 0.31% of it.



