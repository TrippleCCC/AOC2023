package day10

import (
	"aoc2023/util"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

func init() {
	util.AOCDays.SetDay(10, Day10{})
}

type Day10 struct{}

func (Day10) Solve1() any {
	maze := input

	start, nodes, _ := traverseMaze(maze)
	f, _ := findFurthest(maze, start, nodes)
	return f
}

func (Day10) Solve2() any {
	maze := testInput

	start, nodes, tiles := traverseMaze(maze)
	_, path := findFurthest(maze, start, nodes)
	area := 0
	fmt.Println(path)
	for _, t := range tiles {
		if withinMaze(maze, t, path) {
			fmt.Println(t)
			area++
		}
	}

	return area
}

type Pipe struct {
	index, end1, end2 int
}

func traverseMaze(maze string) (int, map[int]Pipe, []int) {
	m := map[int]Pipe{}
	start := -1
	rowLen := len(strings.Split(maze, "\n")[0]) + 1
	t := []int{}

	for i, p := range maze {
		switch p {
		case 'S':
			start = i
		case '|':
			m[i] = Pipe{index: i, end1: i - rowLen, end2: i + rowLen}
		case '-':
			m[i] = Pipe{index: i, end1: i - 1, end2: i + 1}
		case 'L':
			m[i] = Pipe{index: i, end1: i - rowLen, end2: i + 1}
		case 'J':
			m[i] = Pipe{index: i, end1: i - 1, end2: i - rowLen}
		case '7':
			m[i] = Pipe{index: i, end1: i - 1, end2: i + rowLen}
		case 'F':
			m[i] = Pipe{index: i, end1: i + rowLen, end2: i + 1}
		case '.':
			t = append(t, i)
		default:
			continue
		}
	}

	return start, m, t
}

func findFurthest(maze string, start int, nodes map[int]Pipe) (int, map[int]struct{}) {
	rowLen := len(strings.Split(maze, "\n")[0]) + 1
	starts := []int{start - 1, start - rowLen, start + 1, start + rowLen}
	validIndex := func(i int) bool {
		return i >= 0 && i < len(maze) && maze[i] != '\n' && maze[i] != '.'
	}

OUTER:
	for _, s := range starts {
		if !validIndex(s) {
			continue
		}

		pathLen := 1
		prev := start
		currIndex := s
		path := map[int]struct{}{prev: struct{}{}, currIndex: struct{}{}}

		for validIndex(currIndex) && currIndex != start {
			if p, ok := nodes[currIndex]; ok {
				if p.end1 == prev {
					prev = currIndex
					currIndex = p.end2
				} else if p.end2 == prev {
					prev = currIndex
					currIndex = p.end1
				} else {
					continue OUTER
				}
				path[currIndex] = struct{}{}
				pathLen++
			} else {
				panic("?????")
			}
		}

		if currIndex == start {
			if pathLen%2 == 0 {
				return pathLen / 2, path
			} else {
				return pathLen/2 + 1, path
			}
		}
	}

	return 0, nil
}

func withinMaze(maze string, tile int, path map[int]struct{}) bool {
	// left
	curr := tile
	collisions := 0
	for {
		if curr < 0 {
			break
		}

		switch maze[curr] {
		case '\n':
			break
		default:
			if _, ok := path[curr]; ok {
				switch maze[curr] {
				case '|':
				case 'J':
				case 'L':
					collisions += 1
				}
			}
		}
		curr--
	}
	fmt.Println(tile, collisions)
	if collisions%2 == 0 {
		return false
	}

	return true
}
