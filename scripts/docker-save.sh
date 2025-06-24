#!/bin/bash

BASEDIR=$(dirname "$0")

tag=$1

imageName="friday.go"
imagePath="$BASEDIR/../.docker"

mkdir -p $imagePath

echo "[docker save] tag=$tag"
echo "imageName=$imageName"

docker save -o $imagePath/$imageName-$tag.tar $imageName:$tag