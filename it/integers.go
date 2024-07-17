package it

import "iter"

// Integers yields all integers in the range [start, stop) with the given step.
func Integers[V ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64](start, stop, step V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := start; i < stop; i += step {
			if !yield(i) {
				return
			}
		}
	}
}

// NaturalNumbers yields all non-negative integers in ascending order.
func NaturalNumbers[V ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64]() iter.Seq[V] {
	return func(yield func(V) bool) {
		var i V

		for {
			if !yield(i) {
				return
			}
			i++
		}
	}
}
