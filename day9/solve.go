package day9

import (
	"aoc2023/util"
	_ "embed"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

func init() {
	util.AOCDays.SetDay(9, Day9{})
}

type Day9 struct{}

func (Day9) Solve1() any {
	return nil
}

func (Day9) Solve2() any {
	return nil
}
