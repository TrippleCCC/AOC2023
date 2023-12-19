package day16

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
	util.AOCDays.SetDay(16, Day16{})
}

type Day16 struct{}

func (Day16) Solve1() any {
	return calculateEnergizedTiles(input, 0, 'R')
}

func (Day16) Solve2() any {
	biggest := 0
	in := input
	for _, n := range CalculateEdgeDirections(in) {
		// fmt.Printf("n.index: %v\n", n.index)
		// fmt.Printf("n.direction: %v\n", string(n.direction))
		e := calculateEnergizedTiles(in, n.index, n.direction)
		// fmt.Printf("e: %v\n\n", e)
		if e > biggest {
			biggest = e
		}
	}
	return biggest
}

type N struct {
	index     int
	direction byte // 'U' Up, 'D' Down, 'L' Left, 'R' Right
}

func calculateEnergizedTiles(tiles string, startingIndex int, startingDir byte) int {
	queue := []N{{index: startingIndex, direction: startingDir}}
	seen := map[int][]byte{}
	rowLen := len(strings.Split(tiles, "\n")[0])

	for len(queue) != 0 {
		n := queue[0]
		queue = queue[1:]

		// Check if index is valid
		if n.index < 0 || n.index >= len(tiles) || tiles[n.index] == '\n' {
			continue
		}

		// Check if we have already been here in the same direction
		if dirs, ok := seen[n.index]; ok {
			if slices.Contains(dirs, n.direction) {
				continue
			}
		}
		seen[n.index] = append(seen[n.index], n.direction)

		// Add next queues
		queue = append(queue, generateNextNs(n.index, rowLen, tiles[n.index], n.direction)...)
	}

	// println(strings.TrimSpace(buildGraph(seen, tiles)))

	// calculate energy
	total := 0
	for range seen {
		// if slices.Contains([]byte("\\/|-"), tiles[i]) {
		// 	total += 1
		// } else {
		// 	total += len(dirs)
		// }
		total += 1
	}

	return total
}

func buildGraph(seen map[int][]byte, tiles string) string {
	b := strings.Builder{}
	for i := 0; i < len(tiles); i++ {
		if tiles[i] == '\n' {
			b.WriteByte('\n')
		} else if dirs, ok := seen[i]; ok {
			count := len(dirs)
			if slices.Contains([]byte("\\/|-"), tiles[i]) {
				b.WriteRune('#')
			} else if count == 1 {
				b.WriteByte('#')
			} else {
				b.WriteString(strconv.Itoa(count))
			}
		} else {
			b.WriteRune('.')
		}
	}

	return b.String()
}

func generateNextNs(index, rowLen int, modifier, direction byte) []N {
	var nextDirections []byte
	switch direction {
	case 'U':
		switch modifier {
		case '\\':
			nextDirections = append(nextDirections, 'L')
		case '/':
			nextDirections = append(nextDirections, 'R')
		case '-':
			nextDirections = append(nextDirections, 'L', 'R')
		default:
			nextDirections = append(nextDirections, direction)
		}
	case 'D':
		switch modifier {
		case '\\':
			nextDirections = append(nextDirections, 'R')
		case '/':
			nextDirections = append(nextDirections, 'L')
		case '-':
			nextDirections = append(nextDirections, 'L', 'R')
		default:
			nextDirections = append(nextDirections, direction)
		}
	case 'L':
		switch modifier {
		case '\\':
			nextDirections = append(nextDirections, 'U')
		case '/':
			nextDirections = append(nextDirections, 'D')
		case '|':
			nextDirections = append(nextDirections, 'U', 'D')
		default:
			nextDirections = append(nextDirections, direction)
		}
	case 'R':
		switch modifier {
		case '\\':
			nextDirections = append(nextDirections, 'D')
		case '/':
			nextDirections = append(nextDirections, 'U')
		case '|':
			nextDirections = append(nextDirections, 'U', 'D')
		default:
			nextDirections = append(nextDirections, direction)
		}
	}

	return util.MapList(nextDirections, func(b byte) N {
		return N{index: GenerateIndex(index, rowLen, b), direction: b}
	})
}

func GenerateIndex(currIndex, rowLen int, direction byte) int {
	switch direction {
	case 'U':
		return currIndex - rowLen - 1
	case 'D':
		return currIndex + rowLen + 1
	case 'L':
		return currIndex - 1
	case 'R':
		return currIndex + 1
	default:
		panic("Invalid Direction: " + string(direction))
	}
}

func CalculateEdgeDirections(tiles string) []N {
	rowLen := len(strings.Split(tiles, "\n")[0])
	ret := []N{
		{
			index:     0,
			direction: 'R',
		},
		{
			index:     0,
			direction: 'D',
		},
		{
			index:     rowLen - 1,
			direction: 'L',
		},
		{
			index:     rowLen - 1,
			direction: 'D',
		},
		{
			index:     len(tiles) - 1,
			direction: 'U',
		},
		{
			index:     len(tiles) - 1,
			direction: 'L',
		},
		{
			index:     len(tiles) - rowLen,
			direction: 'R',
		},
		{
			index:     len(tiles) - rowLen,
			direction: 'U',
		},
	}

	// top
	for i := 1; i < rowLen-1; i++ {
		ret = append(ret, N{index: i, direction: 'D'})
	}

	// sides
	for i := rowLen + 1; i < len(tiles)-rowLen; i += rowLen + 1 {
		ret = append(ret, N{index: i, direction: 'R'}, N{index: i + rowLen - 1, direction: 'L'})
	}

	// bottom
	for i := len(tiles) - rowLen; i < len(tiles)-1; i++ {
		ret = append(ret, N{index: i, direction: 'U'})
	}

	return ret
}
