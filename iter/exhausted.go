//go:build go1.23

package iter

import "iter"

// Exhausted is an iterator that yields no values.
func Exhausted[V any]() iter.Seq[V] {
	return func(yield func(V) bool) {}
}

// Exhausted2 is an iterator that yields no values.
func Exhausted2[V, W any]() iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {}
}
