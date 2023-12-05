package util

import (
	"fmt"
)

var AOCDays = &AOC{}

type Day interface {
	Solve1() any
	Solve2() any
}

type AOC struct {
	days [25]Day
}

func (a *AOC) SetDay(dayNumber int, day Day) {
	a.days[dayNumber] = day
}

func (a *AOC) RunSolveFunc(dayNumber, part int) {
	day := a.days[dayNumber]

	if day != nil {
		switch part {
		case 1:
			fmt.Printf("Day %d Part 1: %v\n", dayNumber, day.Solve1())
		case 2:
			fmt.Printf("Day %d Part 2: %v\n", dayNumber, day.Solve2())
		default:
			goto END
		}
		return
	}

END:
	fmt.Printf("There is no solution for day %d, part %d\n", dayNumber, part)
}
