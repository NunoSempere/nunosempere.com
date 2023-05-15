year="$1" ## bash traverse.sh 2019

echo "## In $year..."
echo

for dir in */*/*
do
	index_path="$(pwd)/$dir/index.md"
	title="$(cat $index_path | head -n 1)"
	url="https://nunosempere.com/$year/$dir"
	# echo $dir
	# echo $index_path
	# echo $title
	echo "- [$title]($url)"
done
