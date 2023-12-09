package day8

import (
	"aoc2023/util"
	_ "embed"
	"fmt"
	"regexp"
	"slices"

	// "slices"
	// "strconv"
	"strings"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

func init() {
	util.AOCDays.SetDay(8, Day8{})
}

type Day8 struct{}

type LeftRight struct {
	left  string
	right string
}

var pat *regexp.Regexp = regexp.MustCompile(`(.+) = \((.+), (.+)\)`)

func createLeftRightMap(str string) map[string]LeftRight {
	ret := map[string]LeftRight{}
	for _, s := range strings.Split(str, "\n") {
		if len(s) > 0 {
			vals := pat.FindStringSubmatch(s)
			ret[vals[1]] = LeftRight{left: vals[2], right: vals[3]}
		}
	}

	return ret
}

func createLeftRightMapAndStartNodes(str string) (map[string]LeftRight, []string) {
	ret := map[string]LeftRight{}
	starts := []string{}
	for _, s := range strings.Split(str, "\n") {
		if len(s) > 0 {
			vals := pat.FindStringSubmatch(s)
			ret[vals[1]] = LeftRight{left: vals[2], right: vals[3]}
			if strings.HasSuffix(vals[1], "A") {
				starts = append(starts, vals[1])
			}
		}
	}

	return ret, starts
}

func (Day8) Solve1() any {
	in := strings.SplitN(input, "\n\n", 2)
	pattern := strings.TrimSpace(in[0])
	mapEntryLines := in[1]
	m := createLeftRightMap(mapEntryLines)

	index := 0
	steps := 0
	cur := "AAA"
	for cur != "ZZZ" {
		fmt.Println(cur, pattern[index], m[cur])
		switch pattern[index] {
		case 'L':
			cur = m[cur].left
		case 'R':
			cur = m[cur].right
		}

		index = (index + 1) % len(pattern)
		steps += 1
	}

	return steps
}

func allAtEnd(l []string) bool {
	for _, s := range l {
		if !strings.HasSuffix(s, "Z") {
			return false
		}
	}
	return true
}

func findLeastCommonMultiple(nums []uint64) uint64 {
	// primes to 1000
	primes := []uint64{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997}
	divs := slices.Clone(nums)
	eq := util.MapList(nums, func(uint64) uint64 { return 1 })
	pindex := 0
	ps := []uint64{}

	for !slices.Equal(eq, divs) {
		divided := false
		for i, v := range divs {
			if v%primes[pindex] == 0 {
				divided = true
				divs[i] = v / primes[pindex]
			}
		}

		if !divided {
			pindex += 1
			if pindex == len(primes) {
				fmt.Println(ps)
				panic("Not enough primes")
			}
		} else {
			ps = append(ps, primes[pindex])
		}
	}

	product := uint64(1)
	for _, v := range ps {
		product *= v
	}

	return uint64(product)
}

func (Day8) Solve2() any {
	in := strings.SplitN(input, "\n\n", 2)
	pattern := strings.TrimSpace(in[0])
	mapEntryLines := in[1]
	m, starts := createLeftRightMapAndStartNodes(mapEntryLines)

	finalSteps := []uint64{}
	for _, start := range starts {
		index := 0
		steps := uint64(0)
		cur := start
		for !strings.HasSuffix(cur, "Z") {
			switch pattern[index] {
			case 'L':
				cur = m[cur].left
			case 'R':
				cur = m[cur].right
			}

			index = (index + 1) % len(pattern)
			steps += 1
		}
		finalSteps = append(finalSteps, steps)
	}

	return findLeastCommonMultiple(finalSteps)
}
