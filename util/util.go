package util

import (
	"slices"
)

type Set[T comparable] map[T]struct{}

type Ordered[T any] interface {
	// if positive value is returned then caller is larger than parameter
	// if zero then values are equal
	compare(T any) int
}

func Max[T Ordered[T]](values ...T) T {
	if len(values) == 0 {
		panic("values is empty")
	}

	least := values[0]
	for _, v := range values[1:] {
		if least.compare(v) < 0 {
			least = v
		}
	}
	return least
}

func Min[T Ordered[T]](values ...T) T {
	if len(values) == 0 {
		panic("values is empty")
	}

	least := values[0]
	for _, v := range values[1:] {
		if least.compare(v) > 0 {
			least = v
		}
	}
	return least
}

func MapList[A, B any](l []A, f func(A) B) []B {
	ret := make([]B, 0, len(l))
	for _, a := range l {
		ret = append(ret, f(a))
	}
	return ret
}

func FilterList[A comparable](l []A, f func(A) bool) []A {
	ret := make([]A, 0, len(l))
	for _, a := range l {
		if f(a) {
			ret = append(ret, a)
		}
	}
	return ret
}

func RemoveEmptyLines(l []string) []string {
	return FilterList(l, func(s string) bool {
		return s != ""
	})
}

func If[T any](cond bool, trueValue, falseValue T) T {
	if cond {
		return trueValue
	}
	return falseValue
}

func GroupBy[A any](l []A, gs int) [][]A {
	ret := [][]A{}
	currIndex := 0

	for currIndex < len(l) {
		endIndex := slices.Min([]int{currIndex + gs, len(l)})
		ret = append(ret, slices.Clone(l[currIndex:endIndex]))
		currIndex = endIndex
	}

	return ret
}

type Tuple2[A, B any] struct {
	first  A
	second B
}

func Zip[A, B any](a []A, b []B) []Tuple2[A, B] {
	if len(a) != len(b) {
		panic("lists are not the same size")
	}

	ret := make([]Tuple2[A, B], 0, len(a))
	index := 0
	for index < len(a) {
		ret = append(ret, Tuple2[A, B]{a[index], b[index]})
	}
	return ret
}
