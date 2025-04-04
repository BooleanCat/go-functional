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
	return MapUp(delegate, f)
}

// MapUp yields pairs of values from an iterator of single values that have had
// the provided function applied.
func MapUp[V, W, X any](delegate func(func(V) bool), f func(V) (W, X)) iter.Seq2[W, X] {
	return func(yield func(W, X) bool) {
		for value := range delegate {
			if !yield(f(value)) {
				return
			}
		}
	}
}

// MapDown yields single values from an iterator of pairs of values that have had
// the provided function applied.
func MapDown[V, W, X any](delegate func(func(V, W) bool), f func(V, W) X) iter.Seq[X] {
	return func(yield func(X) bool) {
		for v, w := range delegate {
			if !yield(f(v, w)) {
				return
			}
		}
	}
}
