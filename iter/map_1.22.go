//go:build go1.22 && goexperiment.rangefunc

package iter

import "iter"

// Map yields values from an iterator that have been transformed by a function.
func Map[V, W any](delegate iter.Seq[V], transform func(V) W) iter.Seq[W] {
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
func Map2[V, W, X, Y any](delegate iter.Seq2[V, W], transform func(V, W) (X, Y)) iter.Seq2[X, Y] {
	return func(yield func(X, Y) bool) {
		for v, w := range delegate {
			if !yield(transform(v, w)) {
				return
			}
		}
	}
}
