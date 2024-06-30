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
	return func(yield func(V) bool) {
		for value := range delegate {
			if !predicate(value) {
				if !yield(value) {
					return
				}
			}
		}
	}
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
	return func(yield func(V, W) bool) {
		for k, v := range delegate {
			if !predicate(k, v) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
