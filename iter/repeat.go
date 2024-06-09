package iter

import "iter"

// Repeat yields the same value indefinitely.
func Repeat[V any](value V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for {
			if !yield(value) {
				return
			}
		}
	}
}
