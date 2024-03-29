package iter

import "iter"

// Iterator is a wrapper around the iter.Seq type that allowed for method
// chaining of iterators found in this package.
type Iterator[V any] iter.Seq[V]

// Collect consumes an iterator and returns a slice of all items yielded.
func Collect[V any](iter Iterator[V]) []V {
	collection := make([]V, 0)

	for item := range iter {
		collection = append(collection, item)
	}

	return collection
}

// Collect is a convenience method for chaining [Collect] after an iterator.
func (iter Iterator[V]) Collect() []V {
	return Collect[V](iter)
}
