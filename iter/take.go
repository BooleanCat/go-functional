package iter

import "iter"

// Take yields the first `limit` values from a delegate [Iterator].
func Take[V any](delegate iter.Seq[V], limit int) iter.Seq[V] {
	return func(yield func(V) bool) {
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
	}
}

// Take is a convenience method for chaining [Take] on [Iterator]s.
func (iterator Iterator[V]) Take(limit int) Iterator[V] {
	return Iterator[V](Take[V](iter.Seq[V](iterator), limit))
}
