package it

import "iter"

// Map yields values from an iterator that have been transformed by a function.
func Map[V, W any](delegate func(func(V) bool), transform func(V) W) iter.Seq[W] {
	return func(yield func(W) bool) {
		for value := range delegate {
			if !yield(transform(value)) {
				return
			}
		}
	}
}

// Map2 yields pairs of values from an iterator that have been transformed by a
// function.
func Map2[V, W, X, Y any](delegate func(func(V, W) bool), transform func(V, W) (X, Y)) iter.Seq2[X, Y] {
	return func(yield func(X, Y) bool) {
		for v, w := range delegate {
			if !yield(transform(v, w)) {
				return
			}
		}
	}
}
