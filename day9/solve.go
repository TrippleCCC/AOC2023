package day9

import (
	"aoc2023/util"
	_ "embed"
	"slices"
	"strconv"
	"strings"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

func init() {
	util.AOCDays.SetDay(9, Day9{})
}

func allZero(l []int) bool {
	for _, v := range l {
		if v != 0 {
			return false
		}
	}
	return true
}

func computeNextValue(l []int) int {
	layers := [][]int{l}
	currLayer := l
	for !allZero(currLayer) {
		nextLayer := []int{}
		for i := 0; i < len(currLayer)-1; i++ {
			nextLayer = append(nextLayer, currLayer[i+1]-currLayer[i])
		}
		layers = append(layers, nextLayer)
		currLayer = nextLayer
	}

	nextValue := 0
	slices.Reverse(layers)
	for _, h := range layers[1:] {
		nextValue = h[len(h)-1] + nextValue
	}
	return nextValue
}

func computePrevValue(l []int) int {
	layers := [][]int{l}
	currLayer := l
	for !allZero(currLayer) {
		nextLayer := []int{}
		for i := 0; i < len(currLayer)-1; i++ {
			nextLayer = append(nextLayer, currLayer[i+1]-currLayer[i])
		}
		layers = append(layers, nextLayer)
		currLayer = nextLayer
	}

	nextValue := 0
	slices.Reverse(layers)
	for _, h := range layers[1:] {
		nextValue = h[0] - nextValue
	}
	return nextValue
}

type Day9 struct{}

func (Day9) Solve1() any {
	histories := util.MapList(slices.DeleteFunc(strings.Split(input, "\n"), func(s string) bool { return s == "" }), func(s string) []int {
		return util.MapList(strings.Fields(s), func(s string) int {
			n, _ := strconv.Atoi(s)
			return n
		})
	})

	sum := 0
	for _, h := range histories {
		sum += computeNextValue(h)
	}
	return sum
}

func (Day9) Solve2() any {
	histories := util.MapList(slices.DeleteFunc(strings.Split(input, "\n"), func(s string) bool { return s == "" }), func(s string) []int {
		return util.MapList(strings.Fields(s), func(s string) int {
			n, _ := strconv.Atoi(s)
			return n
		})
	})

	sum := 0
	for _, h := range histories {
		sum += computePrevValue(h)
	}
	return sum
}
