MARKDOWN=/usr/bin/markdown -f fencedcode -f ext -f footnote -f latex
build:
	$(MARKDOWN) index.md > temp
	cat title.md temp  isso-snippet.txt > ../index.md
	rm temp
