package it

import "iter"

// Enumerate yields pairs of indices and values from an iterator.
func Enumerate[V any](delegate func(func(V) bool)) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		count := 0

		for value := range delegate {
			if !yield(count, value) {
				return
			}

			count++
		}
	}
}
