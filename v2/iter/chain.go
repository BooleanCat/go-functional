package iter

import "iter"

// Chain yields the elements of each iterator in turn.
func Chain[V any](iterators ...Iterator[V]) Iterator[V] {
	return Iterator[V](iter.Seq[V](func(yield func(V) bool) {
		for _, iterator := range iterators {
			next, stop := iter.Pull(iter.Seq[V](iterator))
			defer stop()

			for {
				if v, ok := next(); !ok || !yield(v) {
					stop()
					break
				}
			}
		}
	}))
}

// Chain is a convenience method for chaining [Chain] after an iterator.
func (iter Iterator[V]) Chain(iterators ...Iterator[V]) Iterator[V] {
	return Chain[V](append([]Iterator[V]{iter}, iterators...)...)
}
