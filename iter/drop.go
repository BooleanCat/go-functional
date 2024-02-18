package iter

import "iter"

// Drop yields all values from a delegate [Iterator] except the first `count`
// values.
func Drop[V any](delegate Iterator[V], count int) Iterator[V] {
	return Iterator[V](iter.Seq[V](func(yield func(V) bool) {
		for value := range delegate {
			if count > 0 {
				count--
				continue
			}

			if !yield(value) {
				return
			}
		}
	}))
}

// Drop is a convenience method for chaining [Drop] on [Iterator]s.
func (iter Iterator[V]) Drop(count int) Iterator[V] {
	return Drop[V](iter, count)
}
