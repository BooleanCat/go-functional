package iter

import "iter"

// Filter yields values from an iterator that satisfy a predicate.
func Filter[V any](delegate Iterator[V], predicate func(V) bool) Iterator[V] {
	return Iterator[V](iter.Seq[V](func(yield func(V) bool) {
		delegate, stop := iter.Pull(iter.Seq[V](delegate))
		defer stop()

		for {
			value, ok := delegate()

			if !ok || predicate(value) && !yield(value) {
				return
			}
		}
	}))
}

// Filter is a convenience method for chaining [Filter] on [Iterator]s.
func (iter Iterator[V]) Filter(predicate func(V) bool) Iterator[V] {
	return Filter[V](iter, predicate)
}

// Filter2 yields values from an iterator that satisfy a predicate.
func Filter2[V, W any](delegate Iterator2[V, W], predicate func(V, W) bool) Iterator2[V, W] {
	return Iterator2[V, W](iter.Seq2[V, W](func(yield func(V, W) bool) {
		delegate, stop := iter.Pull2(iter.Seq2[V, W](delegate))
		defer stop()

		for {
			left, right, ok := delegate()

			if !ok || predicate(left, right) && !yield(left, right) {
				return
			}
		}
	}))
}

// Filter2 is a convenience method for chaining [Filter2] on [Iterator2]s.
func (iter Iterator2[V, W]) Filter2(predicate func(V, W) bool) Iterator2[V, W] {
	return Filter2[V, W](iter, predicate)
}
