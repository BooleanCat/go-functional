package iter

import "iter"

// Iterator is a wrapper around [iter.Seq] that allows for method chaining of
// most iterators found in this package.
type Iterator[V any] iter.Seq[V]

func Lift[V any](slice []V) Iterator[V] {
	return Iterator[V](iter.Seq[V](func(yield func(V) bool) {
		for _, value := range slice {
			if !yield(value) {
				return
			}
		}
	}))
}

// Collect consumes an iterator and returns a slice of all values yielded.
func Collect[V any](iter Iterator[V]) []V {
	collection := make([]V, 0)

	for item := range iter {
		collection = append(collection, item)
	}

	return collection
}

// Collect is a convenience method for chaining [Collect] on [Iterator]s.
func (iter Iterator[V]) Collect() []V {
	return Collect[V](iter)
}
