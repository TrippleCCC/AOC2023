package day11

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
	util.AOCDays.SetDay(11, Day11{})
}

type Day11 struct{}

func (Day11) Solve1() any {
	return solve(input, 2)
}

func (Day11) Solve2() any {
	return solve(input, 1000000)
}

type universeInfo struct {
	rows, cols int
}

type galaxy struct {
	row, col int
}

func solve(universe string, expansionFactor uint64) uint64 {
	lines := util.RemoveEmptyLines(strings.Split(universe, "\n"))
	cols := len(lines[0])

	// Calculate empty rows
	emptyRow := strings.Repeat(".", cols)
	emptyRows := []int{}
	for i, l := range lines {
		if l == emptyRow {
			emptyRows = append(emptyRows, i)
		}
	}

	emptyCols := []int{}
	// Calculate empty rows
OUTER:
	for i := 0; i < cols; i++ {
		for _, l := range lines {
			if l[i] != '.' {
				continue OUTER
			}
		}
		emptyCols = append(emptyCols, i)
	}

	galaxys := []galaxy{}
	for i, l := range lines {
		for k, g := range l {
			if g == '#' {
				galaxys = append(galaxys, galaxy{row: i, col: k})
			}
		}
	}

	// Generate all combos
	combos := [][]galaxy{}
	for i, g := range galaxys {
		for k := i + 1; k < len(galaxys); k++ {
			combos = append(combos, []galaxy{g, galaxys[k]})
		}
	}

	sum := uint64(0)
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	for _, c := range combos {
		rdist := uint64(abs(c[0].col-c[1].col) + abs(c[0].row-c[1].row))
		minRow := min(c[0].row, c[1].row)
		maxRow := max(c[0].row, c[1].row)
		minCol := min(c[0].col, c[1].col)
		maxCol := max(c[0].col, c[1].col)

		for i := minCol; i < maxCol; i++ {
			if slices.Contains(emptyCols, i) {
				rdist += expansionFactor - 1
			}
		}
		for i := minRow; i < maxRow; i++ {
			if slices.Contains(emptyRows, i) {
				rdist += expansionFactor - 1
			}
		}

		sum += rdist
	}

	return sum
}
