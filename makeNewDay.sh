#!/bin/bash

if [ -d "./day$1" ]; then
	echo "Directory ./day$1 already exists"
	exit 0
fi

# Make day directory
mkdir "./day$1"

# Make input files
touch "./day$1/input.txt"
touch "./day$1/testInput.txt"

# Make Go file
printf "package day$1\n" >"./day$1/solve.go"
