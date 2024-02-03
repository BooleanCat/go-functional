package iter

import "iter"

// Filter yields the values from delegate that satisfy predicate.
func Filter[V any](delegate Iterator[V], predicate func(V) bool) Iterator[V] {
	return Iterator[V](iter.Seq[V](func(yield func(V) bool) {
		next, stop := iter.Pull(iter.Seq[V](delegate))
		defer stop()

		for {
			v, ok := next()
			if !ok {
				return
			}

			if predicate(v) && !yield(v) {
				return
			}
		}
	}))
}

// Filter is a convenience method for chaining [Filter] after an iterator.
func (iter Iterator[V]) Filter(predicate func(V) bool) Iterator[V] {
	return Filter[V](iter, predicate)
}

// Exclude drops the values from delegate that satisfy the predicate.
func Exclude[V any](delegate Iterator[V], predicate func(V) bool) Iterator[V] {
	return Filter[V](delegate, func(v V) bool {
		return !predicate(v)
	})
}

// Exclude is a convenience method for chaining [Exclude] after an iterator.
func (iter Iterator[V]) Exclude(predicate func(V) bool) Iterator[V] {
	return Exclude[V](iter, predicate)
}
