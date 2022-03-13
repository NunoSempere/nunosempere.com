#!/usr/bin/bash
configdir="/home/uriel/workspace/werc-1.5.0/sites/nunosempere.com/_werc"
configfile="$configdir/config"
configfiletemp="$configdir/config_temp"
line="$(sort -Ru "$configdir/titles.txt" | head -n 1)"
# echo "$line"
title="$(echo "$line" | sed 's/: .*//g')"
subtitle="$(echo "$line" | sed 's/.*: //g')"

echo "$title"
echo "$subtitle"

rm -f "$configfiletemp"
echo "masterSite=nunosempere.com" > "$configfiletemp"
echo "siteTitle='$title'" >> "$configfiletemp"
echo "siteSubTitle='$subtitle'" >> "$configfiletemp"

mv "$configfiletemp" "$configfile"
## masterSite=nunosempere.com
## siteTitle='Learning'
## siteSubTitle='is about to occur'

