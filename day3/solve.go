package day3

import (
	"aoc2023/util"

	_ "embed"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

//go:embed testInput.txt
var testInput string

func checkIfAdjacent(index, rowLen int, grid string) bool {
	topLeftIndex := index - (rowLen + 2)
	topIndex := index - (rowLen + 1)
	topRightIndex := index - rowLen
	leftIndex := index + 1
	rightIndex := index - 1
	bottomLeftIndex := index + rowLen
	bottomIndex := index + (rowLen + 1)
	bottomRightIndex := index + (rowLen + 2)
	indexes := []int{topLeftIndex, topIndex, topRightIndex, leftIndex, rightIndex, bottomLeftIndex, bottomIndex, bottomRightIndex}

	gridLength := len(grid)
	for _, i := range indexes {
		if i >= 0 && i < gridLength && !unicode.IsDigit(rune(grid[i])) && grid[i] != '\n' && grid[i] != '.' {
			return true
		}
	}

	return false
}

func checkIfAdjacentToStar(index, rowLen int, grid string) (bool, int) {
	topLeftIndex := index - (rowLen + 2)
	topIndex := index - (rowLen + 1)
	topRightIndex := index - rowLen
	leftIndex := index + 1
	rightIndex := index - 1
	bottomLeftIndex := index + rowLen
	bottomIndex := index + (rowLen + 1)
	bottomRightIndex := index + (rowLen + 2)
	indexes := []int{topLeftIndex, topIndex, topRightIndex, leftIndex, rightIndex, bottomLeftIndex, bottomIndex, bottomRightIndex}

	gridLength := len(grid)
	for _, i := range indexes {
		if i >= 0 && i < gridLength && !unicode.IsDigit(rune(grid[i])) && grid[i] != '\n' && grid[i] == '*' {
			return true, i
		}
	}

	return false, -1
}

func calculateSum(grid string) int {
	sum := 0
	rowLen := strings.Index(grid, "\n")

	i := 0
	for i < len(grid) {
		if unicode.IsDigit(rune(grid[i])) {
			startIndex := i
			adjacentFlag := false

			for unicode.IsDigit(rune(grid[i])) {
				if !adjacentFlag && checkIfAdjacent(i, rowLen, grid) {
					adjacentFlag = true
				}
				i++
			}

			// index here should be the end of the number
			if adjacentFlag {
				num, _ := strconv.Atoi(grid[startIndex:i])
				sum += num
			}
		}

		i++
	}

	return sum
}

func calculateTotalGearRatio(grid string) int {
	rowLen := strings.Index(grid, "\n")

	starToNumbers := map[int][]int{}

	i := 0
	for i < len(grid) {
		if unicode.IsDigit(rune(grid[i])) {
			startIndex := i
			starIndexes := []int{}

			for unicode.IsDigit(rune(grid[i])) {
				isAdj, index := checkIfAdjacentToStar(i, rowLen, grid)
				if isAdj && !slices.Contains(starIndexes, index) {
					starIndexes = append(starIndexes, index)
				}
				i++
			}

			// index here should be the end of the number
			if len(starIndexes) > 0 {
				num, _ := strconv.Atoi(grid[startIndex:i])
				for _, index := range starIndexes {
					nums, ok := starToNumbers[index]
					if !ok {
						nums = []int{}
					}
					starToNumbers[index] = append(nums, num)
				}
			}
		}

		i++
	}

	sum := 0
	for _, nums := range starToNumbers {
		if len(nums) == 2 {
			sum += nums[0] * nums[1]
		}
	}

	return sum
}

func init() {
	util.AOCDays.SetDay(3, Day3{})
}

type Day3 struct{}

func (Day3) Solve1() any {
	sum := calculateSum(input)

	return sum
}

func (Day3) Solve2() any {
	sum := calculateTotalGearRatio(input)

	return sum
}
