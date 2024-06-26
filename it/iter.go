package it

import (
	"cmp"
	"iter"
)

// ForEach consumes an iterator and applies a function to each value yielded.
func ForEach[V any](iter iter.Seq[V], fn func(V)) {
	for item := range iter {
		fn(item)
	}
}

// ForEach2 consumes an iterator and applies a function to each pair of values.
func ForEach2[V, W any](iter iter.Seq2[V, W], fn func(V, W)) {
	for v, w := range iter {
		fn(v, w)
	}
}

// Reduce consumes an iterator and applies a function to each value yielded,
// accumulating a single result.
func Reduce[V any, R any](iter iter.Seq[V], fn func(R, V) R, initial R) R {
	result := initial

	for item := range iter {
		result = fn(result, item)
	}

	return result
}

// Reduce2 consumes an iterator and applies a function to each pair of values,
// accumulating a single result.
func Reduce2[V, W any, R any](iter iter.Seq2[V, W], fn func(R, V, W) R, initial R) R {
	result := initial

	for v, w := range iter {
		result = fn(result, v, w)
	}

	return result
}

// Max consumes an iterator and returns the maximum value yielded and true if
// there was at least one value, or the zero value and false if the iterator
// was empty.
func Max[V cmp.Ordered](iterator iter.Seq[V]) (V, bool) {
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
func Min[V cmp.Ordered](iterator iter.Seq[V]) (V, bool) {
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
