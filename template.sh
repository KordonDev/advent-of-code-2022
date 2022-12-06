#!/bin/sh

echo "creating $1"

mkdir "day-$1"
cd "day-$1" || exit
touch "day-$1.go"
touch "$1-example.txt"
touch "$1.txt"
