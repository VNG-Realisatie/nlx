#!/bin/bash

# Make sure you have installed drawio-desktop
# https://github.com/jgraph/drawio-desktop/releases

set -euxo pipefail

for drawioFile in ./*.drawio
do
	filebase=$(basename ${drawioFile} .drawio)

	pngFile=${filebase}.png
	rm ${pngFile} || true
	draw.io --export --border 16 --output ${pngFile} ${drawioFile}

	svgFile=${filebase}.svg
	rm ${svgFile} || true
	draw.io --export --border 16 --output ${svgFile} ${drawioFile}
done
