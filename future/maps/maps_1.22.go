//go:build go1.22 && goexperiment.rangefunc

// Package maps provides early implementations of map-related functions
// available in Go 1.23+.
//
// If using Go 1.22, the rangefunc experimental feature must be enabled.
//
// If using Go 1.23 or later, prefer the functions from the standard library.
//
// This package will be removed in Go 1.24.
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

// Collect collects key-value pairs from seq into a new map and returns it.
func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K]V {
	m := make(map[K]V)

	for key, value := range seq {
		m[key] = value
	}

	return m
}
