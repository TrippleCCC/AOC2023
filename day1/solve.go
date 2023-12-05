package day1

import (
	"aoc2023/util"
	_ "embed"
	"slices"
	"strconv"
	"strings"
)

//go:embed testinput.txt
var testInput string

//go:embed part1.txt
var input string

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func getNumber(str string) int {
	digits := []rune{}
	for _, c := range str {
		if isDigit(c) {
			digits = append(digits, c)
		}
	}

	if len(digits) == 0 {
		return 0
	}

	numberStr := "" + string(digits[0]) + string(digits[len(digits)-1])
	num, _ := strconv.Atoi(numberStr)

	return num
}

func getNumberWithWords(str string) int {

	type numberRep struct {
		str string
		val string
	}

	type numberPos struct {
		pos int
		val string
	}

	numsStrs := []numberRep{{"one", "1"}, {"two", "2"}, {"three", "3"}, {"four", "4"}, {"five", "5"}, {"six", "6"}, {"seven", "7"}, {"eight", "8"}, {"nine", "9"}}
	pos := map[int]string{}

	for _, ns := range numsStrs {
		for i := 0; i < len(str)-len(ns.str)+1; i++ {
			if str[i:i+len(ns.str)] == ns.str {
				pos[i] = ns.val
			}
		}
	}

	for i, c := range str {
		if isDigit(c) {
			pos[i] = string(c)
		}
	}

	if len(pos) == 0 {
		return 0
	}

	keys := make([]int, 0, len(pos))
	for k := range pos {
		keys = append(keys, k)
	}

	maxIndex := slices.Max(keys)
	minIndex := slices.Min(keys)
	num, _ := strconv.Atoi(pos[minIndex] + pos[maxIndex])

	return num
}

func init() {
	util.AOCDays.SetDay(1, Day1{})
}

type Day1 struct{}

func (Day1) Solve1() any {
	values := strings.Split(input, "\n")

	sum := 0
	for _, v := range values {
		sum += getNumber(v)
	}

	return sum
}

func (Day1) Solve2() any {
	values := strings.Split(input, "\n")

	sum := 0
	for _, v := range values {
		sum += getNumberWithWords(v)
	}

	return sum
}
