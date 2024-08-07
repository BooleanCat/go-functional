package it

import "iter"

// Cycle yields values from an iterator repeatedly.
//
// Note: this is an infinite iterator.
//
// Note: memory usage will grow until all values from the underlying iterator
// are stored in memory.
func Cycle[V any](delegate func(func(V) bool)) iter.Seq[V] {
	return func(yield func(V) bool) {
		var items []V

		for item := range delegate {
			items = append(items, item)
			if !yield(item) {
				break
			}
		}

		for {
			for _, item := range items {
				if !yield(item) {
					return
				}
			}
		}
	}
}

// Cycle2 yields pairs of values from an iterator repeatedly.
//
// Note: this is an infinite iterator.
//
// Note: memory usage will grow until all values from the underlying iterator
// are stored in memory.
func Cycle2[V, W any](delegate func(func(V, W) bool)) iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {
		var items []struct {
			v V
			w W
		}

		for v, w := range delegate {
			items = append(items, struct {
				v V
				w W
			}{v, w})
			if !yield(v, w) {
				break
			}
		}

		for {
			for _, item := range items {
				if !yield(item.v, item.w) {
					return
				}
			}
		}
	}
}
