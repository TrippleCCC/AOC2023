package day12

import (
	"aoc2023/util"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

func init() {
	util.AOCDays.SetDay(12, Day12{})
}

type Day12 struct{}

func (Day12) Solve1() any {
	lines := util.RemoveEmptyLines(strings.Split(input, "\n"))
	total := uint64(0)
	for _, l := range lines {
		x := strings.Split(l, " ")
		arrg := x[0]
		nums := util.MapList(strings.Split(x[1], ","), func(s string) int {
			n, _ := strconv.Atoi(s)
			return n
		})

		total += memCount(arrg, nums)
	}
	return total
}

func (Day12) Solve2() any {
	lines := util.RemoveEmptyLines(strings.Split(input, "\n"))
	total := uint64(0)
	for _, l := range lines {
		x := strings.Split(l, " ")
		arrg := strings.Join([]string{x[0], x[0], x[0], x[0], x[0]}, "?")
		nums := util.MapList(strings.Split(x[1], ","), func(s string) int {
			n, _ := strconv.Atoi(s)
			return n
		})
		numsr := []int{}
		for i := 0; i < 5; i++ {
			numsr = append(numsr, nums...)
		}
		// fmt.Printf("%v %v\n", nums, numsr)

		c := memCount(arrg, numsr)

		total += c
	}
	return total
}

func count(cfg string, nums []int) int {
	if cfg == "" {
		return util.If(len(nums) == 0, 1, 0)
	}
	if len(nums) == 0 {
		return util.If(strings.Contains(cfg, "#"), 0, 1)
	}

	result := 0

	if slices.Contains([]byte{'.', '?'}, cfg[0]) {
		result += count(strings.Clone(cfg)[1:], slices.Clone(nums))
	}
	if slices.Contains([]byte{'#', '?'}, cfg[0]) {
		if nums[0] <= len(cfg) && !strings.Contains(cfg[:nums[0]], ".") && (nums[0] == len(cfg) || cfg[nums[0]] != '#') {
			var nextCfg string
			if len(cfg) == nums[0] {
				nextCfg = ""
			} else {
				nextCfg = strings.Clone(cfg[nums[0]+1:])
			}
			result += count(nextCfg, slices.Clone(nums)[1:])
		}
	}

	return result
}

func memCount(cfg string, nums []int) uint64 {
	results := map[string]uint64{}

	var memCount func(string, []int) uint64
	memCount = func(c string, n []int) uint64 {
		if c == "" {
			return util.If[uint64](len(n) == 0, 1, 0)
		}
		if len(n) == 0 {
			return util.If[uint64](strings.Contains(c, "#"), 0, 1)
		}

		numsKey := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(n)), ","), "[]")
		key := c + " " + numsKey
		if r, ok := results[key]; ok {
			// fmt.Println(key)
			return r
		}

		result := uint64(0)

		if slices.Contains([]byte{'.', '?'}, c[0]) {
			result += memCount(strings.Clone(c)[1:], slices.Clone(n))
		}

		if slices.Contains([]byte{'#', '?'}, c[0]) {
			if n[0] <= len(c) && !strings.Contains(c[:n[0]], ".") && (n[0] == len(c) || c[n[0]] != '#') {
				var nextCfg string
				if len(c) == n[0] {
					nextCfg = ""
				} else {
					nextCfg = strings.Clone(c)[n[0]+1:]
				}
				result += memCount(nextCfg, slices.Clone(n[1:]))
			}
		}

		results[key] = result

		return result
	}

	return memCount(cfg, nums)
}
