//go:build go1.23

// Package maps provides early implementations of map-related functions
// available in Go 1.23+.
package maps

import (
	"iter"
	"maps"
)

// All returns an iterator over key-value pairs from m. The iteration order is
// not specified and is not guaranteed to be the same from one call to the
// next.
func All[Map ~map[K]V, K comparable, V any](m Map) iter.Seq2[K, V] {
	return maps.All(m)
}

// Collect collects key-value pairs from seq into a new map and returns it.
func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K]V {
	return maps.Collect(seq)
}
