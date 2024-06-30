package it

import "iter"

// Cycle yields values from an iterator repeatedly.
//
// Note: this is an infinite iterator.
//
// Note: memory usage will grow until all values from the underlying iterator
// are stored in memory.
func Cycle[V any](delegate iter.Seq[V]) iter.Seq[V] {
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

// Cycle is a convenience method for chaining [Cycle] on [Iterator]s.
func (iterator Iterator[V]) Cycle() Iterator[V] {
	return Iterator[V](Cycle(iter.Seq[V](iterator)))
}

// Cycle2 yields pairs of values from an iterator repeatedly.
//
// Note: this is an infinite iterator.
//
// Note: memory usage will grow until all values from the underlying iterator
// are stored in memory.
func Cycle2[V, W any](delegate iter.Seq2[V, W]) iter.Seq2[V, W] {
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

// Cycle is a convenience method for chaining [Cycle2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Cycle() Iterator2[V, W] {
	return Iterator2[V, W](Cycle2(iter.Seq2[V, W](iterator)))
}
