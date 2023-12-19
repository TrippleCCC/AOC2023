package day14

import (
	"aoc2023/util"
	_ "embed"
	"fmt"
	"maps"
	"slices"
	"strings"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

func init() {
	util.AOCDays.SetDay(14, Day14{})
}

type Day14 struct{}

func (Day14) Solve1() any {
	layers := util.RemoveEmptyLines(strings.Split(input, "\n"))
	return calculateTotalLoad(layers)
}

func (Day14) Solve2() any {
	layers := util.RemoveEmptyLines(strings.Split(testInput, "\n"))
	balls, squares := layersToMap(layers)
	layerCount := len(layers)
	rowLen := len(layers[0])

	pastValues := map[int][]int{}
	type node struct {
		total int
		grid  string
	}
	v := node{}
	next := map[node]node{}
	lastCount := len(pastValues)
	sameCount := 5000
	i := 0
	for ; i < 1000000000; i++ {
		tiltNorth(balls, squares, rowLen)
		tiltWest(balls, squares, rowLen)
		tiltSouth(balls, squares, rowLen)
		tiltEast(balls, squares, rowLen)

		next[v] = node{total: calculateTotalLoadFromMap(balls, layerCount), grid: generateString(balls, squares, layerCount, rowLen)}
		v = next[v]
		pastValues[v.total] = append(pastValues[v.total], i)

		if lastCount == len(pastValues) {
			sameCount--
		}
		if sameCount == 0 {
			break
		}
		lastCount = len(pastValues)
	}

	// Find the length of pattern
	start := v
	length := 0
	cur := next[v]
	for ; cur != start; cur = next[cur] {
		fmt.Printf("cur: %v\n", cur)
		length++
	}
	fmt.Printf("cur: %v\n", cur)

	// Find the largest list
	fmt.Printf("length: %v\n", length)
	ret := v
	for i := (1000000000 % length); i >= 0; i-- {
		ret = next[ret]
	}

	return ret.total
}

func calculateTotalLoad(layers []string) int {
	numLayers := len(layers)
	lastPossibleLayer := make([]int, len(layers[0]))
	total := 0

	for i, l := range layers {
		for c, r := range l {
			switch r {
			case 'O':
				total += numLayers - lastPossibleLayer[c]
				lastPossibleLayer[c] += 1
			case '#':
				lastPossibleLayer[c] = i + 1
			}
		}
	}

	return total
}

func layersToMap(layers []string) (map[int][]int, map[int][]int) {
	ballLocations := map[int][]int{}
	squareLocations := map[int][]int{}
	for i := range layers {
		ballLocations[i] = []int{}
		squareLocations[i] = []int{}
	}

	for i, l := range layers {
		for c, r := range l {
			switch r {
			case 'O':
				ballLocations[i] = append(ballLocations[i], c)
			case '#':
				squareLocations[i] = append(squareLocations[i], c)
			}
		}
		slices.Sort(squareLocations[i])
	}

	return ballLocations, squareLocations
}

func tiltNorth(balls, squares map[int][]int, rowLen int) {
	lastPossibleLayer := make([]int, rowLen)

	for i := 0; i < len(balls); i++ {
		remaining := slices.Clone(balls[i])
		for _, c := range balls[i] {
			if lastPossibleLayer[c] != i {
				// Remove from current row
				remaining = slices.DeleteFunc(remaining, func(i int) bool { return i == c })

				// Add to new row
				balls[lastPossibleLayer[c]] = append(balls[lastPossibleLayer[c]], c)
				lastPossibleLayer[c] += 1
			} else if lastPossibleLayer[c] == i {
				lastPossibleLayer[c] += 1
			}
		}
		balls[i] = remaining
		for _, c := range squares[i] {
			lastPossibleLayer[c] = i + 1
		}
	}
}

func tiltSouth(balls, squares map[int][]int, rowLen int) {
	lastPossibleLayer := make([]int, rowLen)
	for i := range lastPossibleLayer {
		lastPossibleLayer[i] = len(balls) - 1
	}

	for i := len(balls) - 1; i >= 0; i-- {
		remaining := slices.Clone(balls[i])
		for _, c := range balls[i] {
			if lastPossibleLayer[c] != i {
				// Remove from current row
				remaining = slices.DeleteFunc(remaining, func(i int) bool { return i == c })

				// Add to new row
				balls[lastPossibleLayer[c]] = append(balls[lastPossibleLayer[c]], c)
				lastPossibleLayer[c] -= 1
			} else if lastPossibleLayer[c] == i {
				lastPossibleLayer[c] -= 1
			}
		}
		balls[i] = remaining
		for _, c := range squares[i] {
			lastPossibleLayer[c] = i - 1
		}
	}
}

func tiltWest(balls, squares map[int][]int, rowLen int) {
	for k, v := range maps.Clone(balls) {
		newBalls := []int{}
		squaresWithEdge := []int{-1}
		squaresWithEdge = append(squaresWithEdge, squares[k]...)
		for i := 0; i < len(squaresWithEdge); i++ {
			ss := len(util.FilterList(v, func(c int) bool {
				if i == len(squaresWithEdge)-1 {
					return c > squaresWithEdge[i]
				}
				return c > squaresWithEdge[i] && c < squaresWithEdge[i+1]
			}))
			for j := 1; j < ss+1; j++ {
				newBalls = append(newBalls, squaresWithEdge[i]+j)
			}
		}

		if len(newBalls) != 0 {
			balls[k] = newBalls
		}
	}
}

func tiltEast(balls, squares map[int][]int, rowLen int) {
	for k, v := range maps.Clone(balls) {
		newBalls := []int{}
		squaresWithEdge := []int{}
		squaresWithEdge = append(squaresWithEdge, squares[k]...)
		squaresWithEdge = append(squaresWithEdge, rowLen)
		for i := 0; i < len(squaresWithEdge); i++ {
			ss := len(util.FilterList(v, func(c int) bool {
				if i == 0 {
					return c < squaresWithEdge[i]
				}
				return c < squaresWithEdge[i] && c > squaresWithEdge[i-1]
			}))
			for j := 1; j < ss+1; j++ {
				newBalls = append(newBalls, squaresWithEdge[i]-j)
			}
		}

		if len(newBalls) != 0 {
			balls[k] = newBalls
		}
	}

}

func calculateTotalLoadFromMap(balls map[int][]int, layerCount int) int {
	total := 0
	for i, v := range balls {
		total += len(v) * (layerCount - i)
	}
	return total
}

func generateString(balls, squares map[int][]int, layerCount, rowLen int) string {
	b := strings.Builder{}
	for i := 0; i < layerCount; i++ {
		l := make([]byte, rowLen)
		for i := range l {
			l[i] = '.'
		}
		for _, c := range balls[i] {
			l[c] = 'O'
		}
		for _, c := range squares[i] {
			l[c] = '#'
		}
		b.Write(l)
		b.WriteByte('\n')
	}
	return b.String()
}
