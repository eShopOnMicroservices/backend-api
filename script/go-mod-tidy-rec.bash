#!/usr/bin/bash

for f in $(find gen -name go.mod)
do (cd $(dirname $f); go mod tidy)
done
