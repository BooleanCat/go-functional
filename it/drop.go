package it

import "iter"

// Drop yields all values from a delegate [Iterator] except the first `count`
// values.
func Drop[V any](delegate iter.Seq[V], count int) iter.Seq[V] {
	return func(yield func(V) bool) {
		for value := range delegate {
			if count > 0 {
				count--
				continue
			}

			if !yield(value) {
				return
			}
		}
	}
}

// Drop is a convenience method for chaining [Drop] on [Iterator]s.
func (iterator Iterator[V]) Drop(count int) Iterator[V] {
	return Iterator[V](Drop(iter.Seq[V](iterator), count))
}

// Drop2 yields all pairs of values from a delegate [Iterator2] except the
// first `count` pairs.
func Drop2[V, W any](delegate iter.Seq2[V, W], count int) iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {
		for v, w := range delegate {
			if count > 0 {
				count--
				continue
			}

			if !yield(v, w) {
				return
			}
		}
	}
}

// Drop is a convenience method for chaining [Drop] on [Iterator2]s.
func (iterator Iterator2[V, W]) Drop(count int) Iterator2[V, W] {
	return Iterator2[V, W](Drop2(iter.Seq2[V, W](iterator), count))
}
