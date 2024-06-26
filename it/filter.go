package it

import "iter"

// Filter yields values from an iterator that satisfy a predicate.
func Filter[V any](delegate iter.Seq[V], predicate func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for value := range delegate {
			if predicate(value) {
				if !yield(value) {
					return
				}
			}
		}
	}
}

// Exclude yields values from an iterator that do not satisfy a predicate.
func Exclude[V any](delegate iter.Seq[V], predicate func(V) bool) iter.Seq[V] {
	return Filter(delegate, not(predicate))
}

// Filter2 yields values from an iterator that satisfy a predicate.
func Filter2[V, W any](delegate iter.Seq2[V, W], predicate func(V, W) bool) iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {
		for k, v := range delegate {
			if predicate(k, v) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// Exclude2 yields values from an iterator that do not satisfy a predicate.
func Exclude2[V, W any](delegate iter.Seq2[V, W], predicate func(V, W) bool) iter.Seq2[V, W] {
	return Filter2(delegate, not2(predicate))
}

func not[V any](predicate func(V) bool) func(V) bool {
	return func(value V) bool {
		return !predicate(value)
	}
}

func not2[V, W any](predicate func(V, W) bool) func(V, W) bool {
	return func(k V, v W) bool {
		return !predicate(k, v)
	}
}
