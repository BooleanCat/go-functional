package iter

import "iter"

// Take limits the number of elements yielded by a delegate iterator to a
// maximum limit.
func Take[V any](delegate Iterator[V], limit int) Iterator[V] {
	return Iterator[V](iter.Seq[V](func(yield func(V) bool) {
		next, stop := iter.Pull(iter.Seq[V](delegate))
		defer stop()

		for {
			limit--

			if limit < 0 {
				return
			}

			v, ok := next()
			if !ok || !yield(v) {
				return
			}
		}
	}))
}

// Take is a convenience method for chaining [Take] after an iterator.
func (iter Iterator[V]) Take(limit int) Iterator[V] {
	return Take[V](iter, limit)
}
