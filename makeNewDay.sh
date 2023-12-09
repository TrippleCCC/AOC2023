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

cat >"./day$1/solve.go" <<EOL
package day$1

import (
	"aoc2023/util"
	_ "embed"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

func init() {
	util.AOCDays.SetDay($1, Day$1{})
}

type Day$1 struct{}

func (Day$1) Solve1() any {
	return nil
}

func (Day$1) Solve2() any {
	return nil
}
EOL

# Make Go file
# printf "package day$1\n" >"./day$1/solve.go"
