#!/bin/bash

common_sed_filters='s/<.*//; s/.*=//; s/ //; s/,//'

function get_area_from_city_name_km2(){
  name="$(echo "$@" | sed 's/ /_/g')"
  api_response="$(curl -s "https://en.wikipedia.org/w/api.php?action=query&prop=revisions&rvprop=content&format=json&titles=$name&rvsection=0&rvslots=main")"
  area="$(echo "$api_response" | jq -r '.query.pages | to_entries[] | .value.revisions[0].slots.main."*"' | grep area_metro_km2 | sed "$common_sed_filters")"
  # km2
  if [ -z "$area" ]; then
    area=$(echo "$api_response" | jq -r '.query.pages | to_entries[] | .value.revisions[0].slots.main."*"' | grep "area_total_km2" | head -n 1 | sed "$common_sed_filters")
  fi
  if [ -z "$area" ]; then
    area=$(echo "$api_response" | jq -r '.query.pages | to_entries[] | .value.revisions[0].slots.main."*"' | grep "area_urban_km2" | head -n 1 | sed "$common_sed_filters" )
  fi
  if [ -z "$area" ]; then
    area=$(echo "$api_response" | jq -r '.query.pages | to_entries[] | .value.revisions[0].slots.main."*"' | grep "area.*_km2" | head -n 1 | sed "$common_sed_filters")
  fi
  # sq_miles
  if [ -z "$area" ]; then
    area=$(echo "$api_response" | jq -r '.query.pages | to_entries[] | .value.revisions[0].slots.main."*"' | grep "area_total_sq_mi" | head -n 1 | sed "$common_sed_filters")
    if [ -n"$area" ]; then
      area="$(echo "$area * 2.59" | bc)"
    fi
  fi
  if [ -z "$area" ]; then
    area=$(echo "$api_response" | jq -r '.query.pages | to_entries[] | .value.revisions[0].slots.main."*"' | grep "area_urban_sq_mi" | head -n 1 | sed "$common_sed_filters" )
    if [ -n"$area" ]; then
      area="$(echo "$area * 2.59" | bc)"
    fi
  fi
  if [ -z "$area" ]; then
    area=$(echo "$api_response" | jq -r '.query.pages | to_entries[] | .value.revisions[0].slots.main."*"' | grep "area.*_sq_mi" | head -n 1 | sed "$common_sed_filters")
    if [ -n"$area" ]; then
      area="$(echo "$area * 2.59" | bc)"
    fi
  fi
  echo $area
}


# Define the input and output CSV files
input_file="List_of_cities_with_over_one_million_inhabitants_1.csv"
output_file="List_of_cities_with_over_one_million_inhabitants_with_areas.csv"

rm $output_file

# Read the CSV file line by line, skipping the header
tail -n +2 "$input_file" | while IFS=, read -r city country inhabitants continent year definition
do
    # Here, you will call your function or operation to process the city name
    # Replace the following line with your actual function or operation
    processed_city=$(get_area_from_city_name_km2 "$city") 

    # Adding a new field - in this case, we're just appending "Processed" to indicate it's been processed
    new_field="${processed_city}"

    # Write the modified line to the new output file. Including the new field at the end.
    echo "$city,$country,$inhabitants,$continent,$year,$definition,$new_field" | tee -a "$output_file"
done

# Optionally, you can add the header with the new field to the output file
{
  echo "City,Country,Inhabitants,Continent,Year,City definition,Area (km2)"
    cat "$output_file"
} > temp_file && mv temp_file "$output_file"
