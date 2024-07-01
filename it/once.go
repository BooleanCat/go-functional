package it

import "iter"

// Once yields the provided value once.
func Once[V any](value V) iter.Seq[V] {
	return func(yield func(V) bool) {
		if !yield(value) {
			return
		}
	}
}

// Once2 yields the provided value pair once.
func Once2[V, W any](v V, w W) iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {
		if !yield(v, w) {
			return
		}
	}
}
