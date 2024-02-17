package iter

import "iter"

// Enumerate yields pairs of indices and values from an iterator.
func Enumerate[V any](delegate Iterator[V]) Iterator2[int, V] {
	return Iterator2[int, V](iter.Seq2[int, V](func(yield func(int, V) bool) {
		delegate, stop := iter.Pull(iter.Seq[V](delegate))
		defer stop()

		for i := 0; ; i++ {
			value, ok := delegate()

			if !ok || !yield(i, value) {
				return
			}
		}
	}))
}

// Enumerate is a convenience method for chaining [Enumerate] on [Iterator]s.
func (iter Iterator[V]) Enumerate() Iterator2[int, V] {
	return Enumerate[V](iter)
}
