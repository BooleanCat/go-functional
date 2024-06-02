package iter

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
