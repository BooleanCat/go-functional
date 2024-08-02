package itx

import "github.com/BooleanCat/go-functional/v2/it"

// Integers yields all integers in the range [start, stop) with the given step.
func Integers[V ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64](start, stop, step V) Iterator[V] {
	return Iterator[V](it.Integers(start, stop, step))
}

// NaturalNumbers yields all non-negative integers in ascending order.
func NaturalNumbers[V ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64]() Iterator[V] {
	return Iterator[V](it.NaturalNumbers[V]())
}
