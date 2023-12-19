package day15

import (
	"aoc2023/util"
	_ "embed"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

func init() {
	util.AOCDays.SetDay(15, Day15{})
}

type Day15 struct{}

var stepRegex = regexp.MustCompile(`(^[a-z]+)([\-=])([0-9]*)`)

func (Day15) Solve1() any {
	steps := util.RemoveEmptyLines(strings.Split(strings.TrimSpace(input), ","))
	total := 0
	for _, s := range steps {
		total += hashString(s)
	}
	return total
}

func (Day15) Solve2() any {
	steps := util.RemoveEmptyLines(strings.Split(strings.TrimSpace(input), ","))

	type Lens struct {
		label string
		focal int
	}
	boxes := map[int][]Lens{}
	for _, s := range steps {
		matches := stepRegex.FindSubmatch([]byte(s))
		label := string(matches[1])
		op := matches[2][0]
		focal, _ := strconv.Atoi(string(matches[3]))

		box := hashString(label)
		switch op {
		case '=':
			if index := slices.IndexFunc(boxes[box], func(l Lens) bool { return l.label == label }); index != -1 {
				boxes[box][index] = Lens{label: label, focal: focal}
			} else {
				boxes[box] = append(boxes[box], Lens{label: label, focal: focal})
			}
		case '-':
			if lenses, ok := boxes[box]; ok {
				if index := slices.IndexFunc(lenses, func(l Lens) bool { return l.label == label }); index != -1 {
					boxes[box] = slices.Delete(boxes[box], index, index+1)
				}
			}
		}
	}

	total := 0
	for box, lenses := range boxes {
		for slot, lens := range lenses {
			total += (box + 1) * (slot + 1) * lens.focal
		}
	}

	return total
}

func hashString(str string) int {
	ret := 0
	for i := 0; i < len(str); i++ {
		ret = ((ret + int(str[i])) * 17) % 256
	}
	return ret
}
