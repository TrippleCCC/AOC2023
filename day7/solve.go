package day7

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
	util.AOCDays.SetDay(7, Day7{})
}

type Hand map[rune]int

type HandAndBet struct {
	tiebreaker []int
	handValue  int
	bet        int
}

func handToMap(hand string) Hand {
	ret := map[rune]int{}
	for _, r := range hand {
		if _, ok := ret[r]; ok {
			ret[r] += 1
		} else {
			ret[r] = 1
		}
	}
	return ret
}

func compareTiebreaker(l1, l2 []int) int {
	for i := range l1 {
		if l1[i] < l2[i] {
			return -1
		} else if l1[i] > l2[i] {
			return 1
		}
	}
	return 0
}

func createTiebreaker(hand, cardVals string) []int {
	ret := []int{}
	for _, c := range hand {
		ret = append(ret, strings.IndexRune(cardVals, c))
	}
	return ret
}

func (hand Hand) toValue() int {
	hv := []int{}
	for _, v := range hand {
		hv = append(hv, v)
	}
	slices.Sort(hv)

	if slices.Equal(hv, []int{5}) {
	} else if slices.Equal(hv, []int{5}) {
		return 7
	} else if slices.Equal(hv, []int{1, 4}) {
		return 6
	} else if slices.Equal(hv, []int{2, 3}) {
		return 5
	} else if slices.Equal(hv, []int{1, 1, 3}) {
		return 4
	} else if slices.Equal(hv, []int{1, 2, 2}) {
		return 3
	} else if slices.Equal(hv, []int{1, 1, 1, 2}) {
		return 2
	} else if slices.Equal(hv, []int{1, 1, 1, 1, 1}) {
		return 1
	} else {
		panic("unrecognized pattern")
	}

	return 0
}

func isFiveOfAKind(hand Hand) bool {
	return len(hand) == 1
}

func isFiveOfAKindJ(hand Hand) bool {
	if len(hand) == 1 {
		return true
	}

	ones := 0
	twos := 0
	threes := 0
	fours := 0
	jv := 0
	for k, v := range hand {
		if v == 1 {
			ones += 1
		} else if v == 2 {
			twos += 1
		} else if v == 3 {
			threes += 1
		} else if v == 4 {
			fours += 1
		}

		if k == 'J' {
			jv = v
		}
	}

	if jv == 0 {
		return false
	}

	return (jv == 1 && fours >= 1) ||
		(jv == 2 && threes >= 1) ||
		(jv == 3 && twos >= 1) ||
		(jv == 4 && ones >= 1)

}

func isFourOfAKind(hand Hand) bool {
	if len(hand) != 2 {
		return false
	}
	for _, v := range hand {
		if v == 1 || v == 4 {
			return true
		}
	}
	return false
}

func isFourOfAKindJ(hand Hand) bool {
	ones := 0
	twos := 0
	threes := 0
	fours := 0
	jv := 0
	for k, v := range hand {
		switch v {
		case 1:
			ones += 1
		case 2:
			twos += 1
		case 3:
			threes += 1
		case 4:
			fours += 1
		}

		if k == 'J' {
			jv = v
		}
	}

	// True case
	if fours >= 1 {
		return true
	}

	if jv == 0 {
		return false
	}

	switch jv {
	case 1:
		return threes >= 1
	case 2:
		return twos >= 2
	case 3:
		return ones >= 1
	default:
		return false
	}
}

func isFullHouse(hand Hand) bool {
	if len(hand) != 2 {
		return false
	}

	for _, v := range hand {
		if v == 2 || v == 3 {
			return true
		}
	}
	return false
}

func isFullHouseJ(hand Hand) bool {
	ones := 0
	twos := 0
	threes := 0
	jv := 0
	for k, v := range hand {
		switch v {
		case 1:
			ones += 1
		case 2:
			twos += 1
		case 3:
			threes += 1
		}
		if k == 'J' {
			jv = v
		}
	}

	if threes == 1 && twos == 1 {
		return true
	}

	if jv == 0 {
		return false
	}

	switch jv {
	case 1:
		return (ones >= 2 && threes >= 1) || (twos >= 2)
	case 2:
		return (twos >= 2 && ones >= 1)
	case 3:
		return ones >= 2
	default:
		return false
	}
}

func isThreeOfAKind(hand Hand) bool {
	if len(hand) != 3 {
		return false
	}
	for _, v := range hand {
		if !(v == 1 || v == 3) {
			return false
		}
	}

	return true
}

func isThreeOfAKindJ(hand Hand) bool {
	ones := 0
	twos := 0
	threes := 0
	jv := 0
	for k, v := range hand {
		switch v {
		case 1:
			ones += 1
		case 2:
			twos += 1
		case 3:
			threes += 1
		}
		if k == 'J' {
			jv = v
		}
	}

	if threes >= 1 {
		return true
	}
	if jv == 0 {
		return false
	}

	switch jv {
	case 1:
		return twos >= 1
	case 2:
		return true
	default:
		return false
	}
}

func isTwoPair(hand Hand) bool {
	if len(hand) != 3 {
		return false
	}

	for _, v := range hand {
		if !(v == 2 || v == 1) {
			return false
		}
	}

	return true
}

func isTwoPairJ(hand Hand) bool {
	ones := 0
	twos := 0
	jv := 0
	for k, v := range hand {
		switch v {
		case 1:
			ones += 1
		case 2:
			twos += 1
		}

		if k == 'J' {
			jv = v
		}
	}

	if twos == 2 && ones == 1 {
		return true
	}
	if jv == 0 {
		return false
	}

	switch jv {
	case 1:
		return (twos == 1 && ones >= 3)
	case 2:
		return true
	default:
		return false
	}
}

func isOnePair(hand Hand) bool {
	if len(hand) != 4 {
		return false
	}

	singles, pair := 0, 0
	for _, v := range hand {
		if v == 1 {
			singles += 1
		} else if v == 2 {
			pair += 1
		}
	}

	return singles == 3 && pair == 1
}

func isOnePairJ(hand Hand) bool {
	twos := 0
	jv := 0
	for k, v := range hand {
		switch v {
		case 2:
			twos += 1
		}

		if k == 'J' {
			jv = v
		}
	}

	if twos >= 1 {
		return true
	}
	if jv == 0 {
		return false
	}

	switch jv {
	case 1:
		return true
	default:
		return false
	}
}

func isHighCard(hand Hand) bool {
	return true
}

func calcuateStrength(hand Hand) int {
	funcs := []func(Hand) bool{isFiveOfAKind, isFourOfAKind, isFullHouse, isThreeOfAKind, isTwoPair, isOnePair, isHighCard}
	for i, f := range funcs {
		if f(hand) {
			return len(funcs) - i
		}
	}
	return -1
}

func calcuateStrengthJ(hand Hand) int {
	funcs := []func(Hand) bool{isFiveOfAKindJ, isFourOfAKindJ, isFullHouseJ, isThreeOfAKindJ, isTwoPairJ, isOnePairJ, isHighCard}
	for i, f := range funcs {
		if f(hand) {
			return len(funcs) - i
		}
	}
	return -1
}

func compareHands(hand1, hand2 Hand) int {
	s1, s2 := calcuateStrength(hand1), calcuateStrength(hand2)
	if s1 > s2 {
		return 1
	} else if s1 < s2 {
		return -1
	}
	return 0
}

func compareHandsJ(hand1, hand2 Hand) int {
	s1, s2 := calcuateStrengthJ(hand1), calcuateStrengthJ(hand2)
	if s1 > s2 {
		return 1
	} else if s1 < s2 {
		return -1
	}
	return 0
}

var cardStrenths map[rune]int = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

var cardStrenths2 map[rune]int = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
	'J': 0,
}

func compareHandAndBets(hb1, hb2 HandAndBet) int {
	if hb1.handValue > hb2.handValue {
		return 1
	} else if hb1.handValue < hb2.handValue {
		return -1
	}

	return compareTiebreaker(hb1.tiebreaker, hb2.tiebreaker)
}

type Day7 struct{}

func (Day7) Solve1() any {
	lines := strings.Split(input, "\n")
	lines = slices.DeleteFunc(lines, func(s string) bool { return s == "" })

	hbs := util.MapList(lines, func(s string) HandAndBet {
		fs := strings.Fields(s)
		handStr := fs[0]
		betStr := fs[1]

		bet, _ := strconv.Atoi(betStr)

		return HandAndBet{
			tiebreaker: createTiebreaker(handStr, "23456789TJQKA"),
			handValue:  handToMap(handStr).toValue(),
			bet:        bet,
		}
	})

	slices.SortFunc(hbs, compareHandAndBets)

	sum := 0
	for i, hb := range hbs {
		sum += hb.bet * (i + 1)
	}

	return sum
}

func (Day7) Solve2() any {
	lines := strings.Split(input, "\n")
	lines = slices.DeleteFunc(lines, func(s string) bool { return s == "" })

	hbs := util.MapList(lines, func(s string) HandAndBet {
		fs := strings.Fields(s)
		handStr := fs[0]
		betStr := fs[1]

		bet, _ := strconv.Atoi(betStr)

		return HandAndBet{
			tiebreaker: createTiebreaker(handStr, "J23456789TQKA"),
			handValue:  handToMap(handStr).toValue(),
			bet:        bet,
		}
	})

	slices.SortFunc(hbs, compareHandAndBets)

	sum := uint64(0)
	for i, hb := range hbs {
		sum += uint64(hb.bet * (i + 1))
	}

	return sum
}
