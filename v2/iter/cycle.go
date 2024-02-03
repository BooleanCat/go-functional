package iter

import "iter"

// Cycle yields all items in the provided slice, repeating the sequence
// indefinitely.
//
// Since Cycle stores the members from the underlying iterator, it will grow in
// size as items are yielded until the underlying iterator is exhausted.
//
// In most cases this iterator is infinite, except when the underlying iterator
// is immediately exhausted in which case this iterator will exhaust.
func Cycle[V any](delegate Iterator[V]) Iterator[V] {
	return Iterator[V](iter.Seq[V](func(yield func(V) bool) {
		next, stop := iter.Pull(iter.Seq[V](delegate))
		defer stop()

		items := make([]V, 0, 8)

		for {
			v, ok := next()
			if !ok || !yield(v) {
				break
			}

			items = append(items, v)
		}

		if len(items) == 0 {
			return
		}

		for {
			for _, item := range items {
				if !yield(item) {
					return
				}
			}
		}
	}))
}

// Cycle is a convenience method for chaining [Cycle] after an iterator.
func (iter Iterator[V]) Cycle() Iterator[V] {
	return Cycle[V](iter)
}
