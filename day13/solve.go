package day13

import (
	"aoc2023/util"
	_ "embed"
	"slices"
	"strings"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

func init() {
	util.AOCDays.SetDay(13, Day13{})
}

type Day13 struct{}

func (Day13) Solve1() any {
	patterns := util.RemoveEmptyLines(strings.Split(input, "\n\n"))

	total := 0
OUTER:
	for _, p := range patterns {
		rows := util.RemoveEmptyLines(strings.Split(p, "\n"))
		cols := []string{}
		for i := 0; i < len(rows[0]); i++ {
			b := strings.Builder{}
			for _, r := range rows {
				b.WriteByte(r[i])
			}
			cols = append(cols, b.String())
			b.Reset()
		}

		indexes := map[int]int{}
		for _, row := range rows {
			for _, c := range findReflections(row) {
				if _, ok := indexes[c]; ok {
					indexes[c]++
				} else {
					indexes[c] = 1
				}
			}
		}
		for k, v := range indexes {
			if v == len(rows) {
				total += k
				continue OUTER
			}
		}

		indexes = map[int]int{}
		for _, col := range cols {
			for _, c := range findReflections(col) {
				if _, ok := indexes[c]; ok {
					indexes[c]++
				} else {
					indexes[c] = 1
				}
			}
		}
		for k, v := range indexes {
			if v == len(cols) {
				total += k * 100
			}
		}
	}

	return total
}

func (Day13) Solve2() any {
	patterns := util.RemoveEmptyLines(strings.Split(input, "\n\n"))

	total := 0
OUTER:
	for _, p := range patterns {
		rows := util.RemoveEmptyLines(strings.Split(p, "\n"))
		cols := []string{}
		for i := 0; i < len(rows[0]); i++ {
			b := strings.Builder{}
			for _, r := range rows {
				b.WriteByte(r[i])
			}
			cols = append(cols, b.String())
			b.Reset()
		}

		indexes := map[int]int{}
		for _, row := range rows {
			for _, c := range findReflections(row) {
				if _, ok := indexes[c]; ok {
					indexes[c]++
				} else {
					indexes[c] = 1
				}
			}
		}
		for k, v := range indexes {
			if v == len(rows)-1 {
				total += k
				continue OUTER
			}
		}

		indexes = map[int]int{}
		for _, col := range cols {
			for _, c := range findReflections(col) {
				if _, ok := indexes[c]; ok {
					indexes[c]++
				} else {
					indexes[c] = 1
				}
			}
		}
		for k, v := range indexes {
			if v == len(cols)-1 {
				total += k * 100
			}
		}
	}

	return total
}

func findReflections(str string) []int {
	left, right := 0, 1

	ret := []int{}
	for ; right < len(str); left, right = left+1, right+1 {
		length := min(left+1, len(str)-right)
		leftSide := str[left+1-length : left+1]
		rightSide := []rune(str[right : right+length])
		slices.Reverse(rightSide)
		if leftSide == string(rightSide) {
			ret = append(ret, right)
		}
	}

	return ret
}
