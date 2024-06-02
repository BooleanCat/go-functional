// Package maps provides early implementations of map-related functions
// available in Go 1.23+.
package maps

import "iter"

// All returns an iterator over key-value pairs from m. The iteration order is
// not specified and is not guaranteed to be the same from one call to the
// next.
func All[Map ~map[K]V, K comparable, V any](m Map) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for key, value := range m {
			if !yield(key, value) {
				return
			}
		}
	}
}
