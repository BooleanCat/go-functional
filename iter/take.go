package iter

import "iter"

// Take yields the first `limit` values from a delegate [Iterator].
func Take[V any](delegate Iterator[V], limit int) Iterator[V] {
	return Iterator[V](iter.Seq[V](func(yield func(V) bool) {
		if limit <= 0 {
			return
		}

		for value := range delegate {
			if !yield(value) {
				return
			}

			limit--
			if limit <= 0 {
				return
			}
		}
	}))
}

// Take is a convenience method for chaining [Take] on [Iterator]s.
func (iter Iterator[V]) Take(limit int) Iterator[V] {
	return Take[V](iter, limit)
}
