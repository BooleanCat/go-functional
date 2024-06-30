package it

import "iter"

// Enumerate yields pairs of indices and values from an iterator.
func Enumerate[V any](delegate iter.Seq[V]) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		delegate, stop := iter.Pull(delegate)
		defer stop()

		for i := 0; ; i++ {
			value, ok := delegate()

			if !ok || !yield(i, value) {
				return
			}
		}
	}
}
