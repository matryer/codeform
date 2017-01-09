#!/bin/bash

# build.sh
# Packages all files in static/lib into static/dist/all.* files

echo "Remove old files..."
rm -rf static/dist
mkdir static/dist

echo "Building all.js..."
cat static/lib/*.js > static/dist/all.js
cat static/js/*.js >> static/dist/all.js
echo "Building all.css..."
cat static/lib/*.css > static/dist/all.css
cat static/css/*.css >> static/dist/all.css
echo "Copying themes..."
cp -R static/lib/themes static/dist/themes/
echo "Done."
