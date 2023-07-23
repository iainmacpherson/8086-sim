#!/bin/sh
# Format all go files in all directories below the current one
echo "Formatting all .go files below current directory:"
for FILE in $(find . -name "*.go")
do
    echo $FILE
    go fmt
done
echo "Done."
