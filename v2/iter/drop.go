package iter

import "iter"

// Drop discards the first `count` elements of a delegate iterator before
// yielding the rest.
func Drop[V any](delegate Iterator[V], count int) Iterator[V] {
	return Iterator[V](iter.Seq[V](func(yield func(V) bool) {
		next, stop := iter.Pull(iter.Seq[V](delegate))
		defer stop()

		for ; count > 0; count-- {
			if _, ok := next(); !ok {
				return
			}
		}

		for {
			if v, ok := next(); !ok || !yield(v) {
				return
			}
		}
	}))
}

// Drop is a convenience method for chaining [Drop] after an iterator.
func (iter Iterator[V]) Drop(count int) Iterator[V] {
	return Drop[V](iter, count)
}
