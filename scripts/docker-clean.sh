#!/bin/bash

filter=${1:-"until=2h"}

echo "[Clean up images] '<none>'"
docker image rm $(docker image list -f 'dangling=true' -q --no-trunc)
echo ""

echo "[Clean up builder] filter $filter"
docker builder prune --filter $filter
echo ""