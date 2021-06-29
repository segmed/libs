#!/bin/bash

contents='// Copyright (c) 2021 Segmed Inc.
package namedetect
var AllFrenchNames = []string{
'

contents+=$(curl -s https://raw.githubusercontent.com/ThinkR-open/prenoms/master/data-raw/nat2019.csv \
  | awk 'BEGIN{FS=";";OFS="\t"} {printf "    \"%s\",\n",tolower($2)}' \
  | awk '!a[$0]++' | grep -v -e preusuel -e _prenoms_rares -e '"a"')
contents+=$'\n}'

echo "$contents" > frenchnames.go
