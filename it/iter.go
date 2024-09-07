package it

import (
	"cmp"
	"iter"
)

// ForEach consumes an iterator and applies a function to each value yielded.
func ForEach[V any](iterator func(func(V) bool), fn func(V)) {
	for item := range iterator {
		fn(item)
	}
}

// ForEach2 consumes an iterator and applies a function to each pair of values.
func ForEach2[V, W any](iterator func(func(V, W) bool), fn func(V, W)) {
	for v, w := range iterator {
		fn(v, w)
	}
}

// Fold will fold every element into an accumulator by applying a function and
// passing an initial value.
func Fold[V, R any](iterator func(func(V) bool), fn func(R, V) R, initial R) R {
	result := initial

	for item := range iterator {
		result = fn(result, item)
	}

	return result
}

// Fold2 will fold every element into an accumulator by applying a function and
// passing an initial value.
func Fold2[V, W, R any](iterator func(func(V, W) bool), fn func(R, V, W) R, initial R) R {
	result := initial

	for v, w := range iterator {
		result = fn(result, v, w)
	}

	return result
}

// Max consumes an iterator and returns the maximum value yielded and true if
// there was at least one value, or the zero value and false if the iterator
// was empty.
func Max[V cmp.Ordered](iterator func(func(V) bool)) (V, bool) {
	var (
		max V
		set bool
	)

	next, _ := iter.Pull(iterator)
	first, ok := next()
	if !ok {
		return max, set
	}

	max = first
	set = true

	for item := range iterator {
		if item > max {
			max = item
		}
	}

	return max, set
}

// Min consumes an iterator and returns the minimum value yielded and true if
// there was at least one value, or the zero value and false if the iterator
// was empty.
func Min[V cmp.Ordered](iterator func(func(V) bool)) (V, bool) {
	var (
		min V
		set bool
	)

	next, _ := iter.Pull(iterator)
	first, ok := next()
	if !ok {
		return min, set
	}

	min = first
	set = true

	for item := range iterator {
		if item < min {
			min = item
		}
	}

	return min, set
}

// Find consumes an iterator until a value is found that satisfies a predicate.
// It returns the value and true if one was found, or the zero value and false
// if the iterator was exhausted.
func Find[V any](iterator func(func(V) bool), pred func(V) bool) (V, bool) {
	for item := range iterator {
		if pred(item) {
			return item, true
		}
	}

	var zero V
	return zero, false
}

// Find2 consumes an iterator until a pair of values is found that satisfies a
// predicate. It returns the pair and true if one was found, or the zero values
// and false if the iterator was exhausted.
func Find2[V, W any](iterator func(func(V, W) bool), pred func(V, W) bool) (V, W, bool) {
	for v, w := range iterator {
		if pred(v, w) {
			return v, w, true
		}
	}

	var zeroV V
	var zeroW W
	return zeroV, zeroW, false
}

// TryCollect consumes an [iter.Seq2] where the right side yields errors and
// returns a slice of values and the first error encountered. Iteration stops
// at the first error.
func TryCollect[V any](iterator func(func(V, error) bool)) ([]V, error) {
	var values []V

	for v, err := range iterator {
		if err != nil {
			return values, err
		}
		values = append(values, v)
	}

	return values, nil
}

// Collect2 consumes an [iter.Seq2] and returns two slices of values.
func Collect2[V, W any](iterator func(func(V, W) bool)) ([]V, []W) {
	var (
		lefts  []V
		rights []W
	)

	for v, w := range iterator {
		lefts = append(lefts, v)
		rights = append(rights, w)
	}

	return lefts, rights
}

// Len consumes an [iter.Seq] and returns the number of values yielded.
func Len[V any](iterator func(func(V) bool)) int {
	var length int

	for range iterator {
		length++
	}

	return length
}

// Len2 consumes an [iter.Seq2] and returns the number of pairs of values
// yielded.
func Len2[V, W any](iterator func(func(V, W) bool)) int {
	var length int

	for range iterator {
		length++
	}

	return length
}

// Contains consumes an [iter.Seq] until the provided value is found and
// returns true. If the value is not found, it returns false when the iterator
// is exhausted.
func Contains[V comparable](iterator func(func(V) bool), v V) bool {
	for value := range iterator {
		if value == v {
			return true
		}
	}

	return false
}
