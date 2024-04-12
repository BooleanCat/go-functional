package iter

import "iter"

// Map yields values from an iterator that have been transformed by a function.
func Map[V, W any](delegate Iterator[V], transform func(V) W) Iterator[W] {
	return Iterator[W](iter.Seq[W](func(yield func(W) bool) {
		for value := range delegate {
			if !yield(transform(value)) {
				return
			}
		}
	}))
}
