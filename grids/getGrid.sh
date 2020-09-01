#!/bin/bash

# Grid from URL: https://qqwing.com/generate.html

GRID=$1
[ ! -f $GRID ] && echo "Error: file '$GRID' doesn't exist !" 2>&1 && exit 1

varName=$(basename $GRID)

echo "var $varName = [9][9]int{"
head -n9 $GRID |
    sed 's/\./0/g' | 
    sed 's/\([0-9]\)/\1,/g' |
    sed 's/^\([0-9]\)/{\1/' |
    sed 's/,$/},/'
echo "}"

exit 0