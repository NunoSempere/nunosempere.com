sum:
	cat List_of_cities_with_over_one_million_inhabitants_with_areas_manual.csv | cut -d, -f9 | paste -sd+ | sed 's|^+||' | bc

sort:
	cat List_of_cities_with_over_one_million_inhabitants_with_areas_manual.csv | cut -d, -f9  | sort -n -r
