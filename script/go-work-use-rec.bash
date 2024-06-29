#!/usr/bin/bash

for f in $(find gen -name go.mod)
do (go work use $(dirname $f))
done
