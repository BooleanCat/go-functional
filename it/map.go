package it

import "iter"

// Map yields values from an iterator that have had the provided function
// applied to each value.
func Map[V, W any](delegate func(func(V) bool), f func(V) W) iter.Seq[W] {
	return func(yield func(W) bool) {
		for value := range delegate {
			if !yield(f(value)) {
				return
			}
		}
	}
}

// Map2 yields pairs of values from an iterator that have had the provided
// function applied to each pair.
func Map2[V, W, X, Y any](delegate func(func(V, W) bool), f func(V, W) (X, Y)) iter.Seq2[X, Y] {
	return func(yield func(X, Y) bool) {
		for v, w := range delegate {
			if !yield(f(v, w)) {
				return
			}
		}
	}
}

// MapError yields values from an iterator that have had the provided function
// applied to each value where the function can return an error.
func MapError[V, W any](delegate func(func(V) bool), f func(V) (W, error)) iter.Seq2[W, error] {
	return func(yield func(W, error) bool) {
		for value := range delegate {
			if result, err := f(value); err != nil {
				if !yield(result, err) {
					return
				}
			} else {
				if !yield(result, nil) {
					return
				}
			}
		}
	}
}
