package iter

import "iter"

// Repeat yields an infinite sequence of the same value.
func Repeat[V any](value V) Iterator[V] {
	return Iterator[V](iter.Seq[V](func(yield func(V) bool) {
		for {
			if !yield(value) {
				return
			}
		}
	}))
}
