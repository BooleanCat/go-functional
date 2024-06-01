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
