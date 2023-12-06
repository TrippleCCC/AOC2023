package day5

import (
	"aoc2023/util"
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"sync"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

type Day5 struct{}

func init() {
	util.AOCDays.SetDay(5, Day5{})
}

func (Day5) Solve1() any {
	input := strings.Split(input, "\n\n")

	seeds := util.MapList(strings.Fields(strings.Split(input[0], ":")[1]), func(s string) uint64 {
		i, _ := strconv.Atoi(s)
		return uint64(i)
	})
	fmt.Println(seeds)

	var mapFunction func(uint64) uint64

	for _, m := range input[1:] {
		lines := strings.Split(m, "\n")[1:]

		buckets := []MapBucket{}
		for _, l := range lines {
			mb := MapBucket{}
			fmt.Sscanf(l, "%d %d %d", &mb.destinationStart, &mb.sourceStart, &mb.length)
			buckets = append(buckets, mb)
		}

		mf := func(in uint64) uint64 {
			for _, b := range buckets {
				if next, ok := b.GetDest(in); ok {
					return next
				}
			}

			return in
		}

		if mapFunction == nil {
			mapFunction = mf
		} else {
			mapFunction = Compose(mapFunction, mf)
		}
	}

	return slices.Min(util.MapList(seeds, mapFunction))
}

func (Day5) Solve2() any {
	input := strings.Split(input, "\n\n")

	seeds := util.MapList(strings.Fields(strings.Split(input[0], ":")[1]), func(s string) uint64 {
		i, _ := strconv.Atoi(s)
		return uint64(i)
	})

	seedRanges := []struct{ start, length uint64 }{}
	for i := 0; i < len(seeds); i += 2 {
		seedRanges = append(seedRanges, struct {
			start  uint64
			length uint64
		}{start: seeds[i], length: seeds[i+1]})
	}
	fmt.Printf("%v\n", seedRanges)

	var mapFunction func(uint64) uint64

	for _, m := range input[1:] {
		lines := strings.Split(m, "\n")[1:]

		buckets := []MapBucket{}
		for _, l := range lines {
			mb := MapBucket{}
			fmt.Sscanf(l, "%d %d %d", &mb.destinationStart, &mb.sourceStart, &mb.length)
			buckets = append(buckets, mb)
		}

		mf := func(in uint64) uint64 {
			for _, b := range buckets {
				if next, ok := b.GetDest(in); ok {
					return next
				}
			}

			return in
		}

		if mapFunction == nil {
			mapFunction = mf
		} else {
			mapFunction = Compose(mapFunction, mf)
		}
	}

	wg := sync.WaitGroup{}
	rc := make(chan uint64, len(seedRanges))
	wg.Add(len(seedRanges))
	for _, r := range seedRanges {
		go func(sr struct{ start, length uint64 }) {
			var smallest uint64 = math.MaxUint64
			for i := sr.start; i < sr.start+sr.length; i++ {
				if l := mapFunction(i); l < smallest {
					smallest = l
				}
			}
			rc <- smallest
			wg.Done()
		}(r)
	}

	wg.Wait()
	close(rc)

	var smallest uint64 = math.MaxUint64
	for i := range rc {
		if i < smallest {
			smallest = i
		}
	}

	return smallest
}

type MapBucket struct {
	destinationStart, sourceStart, length uint64
}

func Compose[A, B, C any](f1 func(A) B, f2 func(B) C) func(A) C {
	return func(a A) C {
		return f2(f1(a))
	}
}

func (mb MapBucket) GetDest(source uint64) (uint64, bool) {
	if source >= mb.sourceStart && source <= mb.sourceStart+mb.length-1 {
		return mb.destinationStart + (source - mb.sourceStart), true
	}
	return 0, false
}
