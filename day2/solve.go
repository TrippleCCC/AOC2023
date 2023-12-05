package day2

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"

	"aoc2023/util"
)

//go:embed testinput.txt
var testInput string

//go:embed input.txt
var input string

type MaxCountPerGame struct {
	id, maxRed, maxGreen, maxBlue int
}

func (m MaxCountPerGame) power() int {
	return m.maxBlue * m.maxGreen * m.maxRed
}

func parseGameMax(game string) MaxCountPerGame {
	l := util.MapList(strings.Split(game, ":"), strings.TrimSpace)
	sets := l[1]
	setsList := strings.Split(sets, ";")
	maxCounts := util.MapList(setsList, parseSet)

	id, _ := strconv.Atoi(strings.Split(l[0], " ")[1])

	return MaxCountPerGame{
		id:       id,
		maxRed:   slices.Max(util.MapList(maxCounts, func(m MaxCountPerGame) int { return m.maxRed })),
		maxBlue:  slices.Max(util.MapList(maxCounts, func(m MaxCountPerGame) int { return m.maxBlue })),
		maxGreen: slices.Max(util.MapList(maxCounts, func(m MaxCountPerGame) int { return m.maxGreen })),
	}
}

func parseSet(set string) MaxCountPerGame {
	colors := util.MapList(strings.Split(set, ","), strings.TrimSpace)

	ret := MaxCountPerGame{}

	for _, c := range colors {
		l := strings.Split(c, " ")
		number, _ := strconv.Atoi(l[0])

		switch color := l[1]; color {
		case "red":
			ret.maxRed = number
		case "blue":
			ret.maxBlue = number
		case "green":
			ret.maxGreen = number
		}
	}

	return ret
}

func init() {
	util.AOCDays.SetDay(2, Day2{})
}

type Day2 struct{}

func (Day2) Solve1() any {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, l := range lines {
		if len(strings.TrimSpace(l)) == 0 {
			continue
		}
		m := parseGameMax(l)
		if m.maxRed <= 12 && m.maxBlue <= 14 && m.maxGreen <= 13 {
			sum += m.id
		}
	}

	return sum
}

func (Day2) Solve2() any {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, l := range lines {
		if len(strings.TrimSpace(l)) == 0 {
			continue
		}
		m := parseGameMax(l)
		sum += m.power()
	}

	return sum
}
