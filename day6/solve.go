package day6

import (
	"aoc2023/util"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

func init() {
	util.AOCDays.SetDay(6, Day6{})
}

type Day6 struct{}

func (Day6) Solve1() any {
	lines := util.MapList(strings.Split(input, "\n"), func(s string) []int {
		if len(s) == 0 {
			return nil
		}
		l := strings.Fields(strings.Split(s, ":")[1])

		return util.MapList(l, func(n string) int {
			num, _ := strconv.Atoi(n)
			return num
		})
	})
	times := lines[0]
	distances := lines[1]

	ret := uint64(0)
	for i := 0; i < len(times); i++ {
		sols := GetNumSolutions(uint64(times[i]), uint64(distances[i]))
		if ret == 0 {
			ret = sols
		} else {
			ret *= sols
		}
	}

	return ret
}

func (Day6) Solve2() any {
	lines := util.MapList(strings.Split(input, "\n"), func(s string) uint64 {
		if len(s) == 0 {
			return 0
		}
		l := strings.Fields(strings.Split(s, ":")[1])

		var num uint64
		fmt.Sscanf(strings.Join(l, ""), "%d", &num)

		return num
	})
	time := lines[0]
	dist := lines[1]

	return GetNumSolutions(time, dist)
}

func GetNumSolutions(time, dist uint64) uint64 {
	sum := uint64(0)

	for i := uint64(1); i < time; i++ {
		remainingTime := time - i
		speed := i
		if remainingTime*speed > dist {
			sum++
		}
	}
	return sum
}
