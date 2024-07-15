package it

import "iter"

// Chain yields values from multiple iterators in sequence.
func Chain[V any](iterators ...func(func(V) bool)) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, iterator := range iterators {
			iterator(yield)
		}
	}
}

// Chain2 yields values from multiple iterators in sequence.
func Chain2[V, W any](iterators ...func(func(V, W) bool)) iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {
		for _, iterator := range iterators {
			iterator(yield)
		}
	}
}
