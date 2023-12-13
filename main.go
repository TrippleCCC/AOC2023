package main

import (
	_ "aoc2023/day1"
	_ "aoc2023/day11"
	_ "aoc2023/day12"
	_ "aoc2023/day13"
	_ "aoc2023/day2"
	_ "aoc2023/day3"
	_ "aoc2023/day4"
	_ "aoc2023/day5"
	_ "aoc2023/day6"
	_ "aoc2023/day7"
	_ "aoc2023/day8"
	_ "aoc2023/day9"
	"aoc2023/util"

	"flag"
	"fmt"
	"os"
	"slices"
)

var (
	dayFlag  = flag.Int("day", -1, "Advent of code Day")
	partFlag = flag.Int("part", -1, "1st or second part")
)

func main() {
	flag.Parse()

	if !slices.Contains([]int{1, 2}, *partFlag) {
		fmt.Printf("Invalid day part (%d), only 1 and 2 are allowed.\n", *partFlag)
		os.Exit(1)
	}

	if *dayFlag > 25 || *dayFlag < 1 {
		fmt.Printf("Invalid day (%d), must be inbetween 1 and 25.\n", *dayFlag)
		os.Exit(1)
	}

	util.AOCDays.RunSolveFunc(*dayFlag, *partFlag)
}
