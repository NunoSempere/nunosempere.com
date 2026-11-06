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
  
  curl -s "https://en.wikipedia.org/w/api.php?action=query&prop=revisions&rvprop=content&format=json&titles=$name&rvsection=0&rvslots=main" | jq -r '.query.pages | to_entries[] | .value.revisions[0].slots.main."*"' 
  echo $area
}
get_area_from_city_name_km2 "$@"
