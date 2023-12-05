package day5

import (
	"aoc2023/util"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

func init() {
	util.AOCDays.SetDay(4, Day4{})
}

type Day4 struct{}

func (Day4) Solve1() any {
	cardStrs := strings.Split(input, "\n")
	sum := 0

	for _, s := range cardStrs {
		if len(s) > 0 {
			card := ParseToCard(s)
			score := card.calculateScore()
			sum += score
		}
	}

	return sum
}

func (Day4) Solve2() any {
	cardStrs := strings.Split(input, "\n")
	duped := map[int]int{}
	sum := 0

	for _, s := range cardStrs {
		if len(s) > 0 {
			card := ParseToCard(s)
			if _, ok := duped[card.id]; !ok {
				duped[card.id] = 1
			} else {
				duped[card.id] += 1
			}

			sum += duped[card.id]

			for _, id := range card.dupedCards() {
				if _, ok := duped[id]; !ok {
					duped[id] = duped[card.id]
				} else {
					duped[id] += duped[card.id]
				}
			}

			delete(duped, card.id)
		}
	}

	return sum
}

type Card struct {
	id             int
	numbers        []int
	winningNumbers map[int]struct{}
}

func (c Card) calculateScore() int {
	total := 0

	for _, num := range c.numbers {
		if _, ok := c.winningNumbers[num]; ok {
			if total == 0 {
				total = 1
			} else {
				total *= 2
			}
		}
	}

	return total
}

func (c Card) dupedCards() []int {
	total := 0

	for _, num := range c.numbers {
		if _, ok := c.winningNumbers[num]; ok {
			total += 1
		}
	}

	if total == 0 {
		return nil
	}

	ret := []int{}
	for i := 1; i <= total; i++ {
		ret = append(ret, c.id+i)
	}

	return ret
}

func SplitNumbers(nums string) []int {
	ret := util.MapList(strings.Fields(nums), func(s string) int {
		num, _ := strconv.Atoi(s)
		return num
	})
	return ret
}

func ParseToCard(str string) Card {
	cardSplit := strings.Split(str, ":")
	gotAndWin := strings.Split(cardSplit[1], "|")
	cardNumber := 0
	fmt.Sscanf(cardSplit[0], "Card %d", &cardNumber)

	wmap := map[int]struct{}{}
	for _, a := range SplitNumbers(gotAndWin[0]) {
		wmap[a] = struct{}{}
	}

	return Card{
		id:             cardNumber,
		numbers:        SplitNumbers(gotAndWin[1]),
		winningNumbers: wmap,
	}
}
