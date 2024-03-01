package iter

import "iter"

type (
	// Iterator is a wrapper around [iter.Seq] that allows for method chaining of
	// most iterators found in this package.
	Iterator[V any] iter.Seq[V]

	// Iterator is a wrapper around [iter.Seq2] that allows for method chaining
	// of most iterators found in this package.
	Iterator2[V, W any] iter.Seq2[V, W]
)

// Lift yields all values from a slice.
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

// LiftHashMap yields all key-value pairs from a map.
//
// The order of iteration is non-deterministic.
func LiftHashMap[K comparable, V any](m map[K]V) Iterator2[K, V] {
	return Iterator2[K, V](iter.Seq2[K, V](func(yield func(K, V) bool) {
		for key, value := range m {
			if !yield(key, value) {
				return
			}
		}
	}))
}

// ForEach consumes an iterator and applies a function to each value yielded.
func ForEach[V any](iter Iterator[V], fn func(V)) {
	for item := range iter {
		fn(item)
	}
}

// ForEach is a convenience method for chaining [ForEach] on [Iterator]s.
func (iter Iterator[V]) ForEach(fn func(V)) {
	ForEach[V](iter, fn)
}

// ForEach2 consumes an iterator and applies a function to each pair of values.
func ForEach2[V, W any](iter Iterator2[V, W], fn func(V, W)) {
	for v, w := range iter {
		fn(v, w)
	}
}

// ForEach2 is a convenience method for chaining [ForEach2] on [Iterator2]s.
func (iter Iterator2[V, W]) ForEach2(fn func(V, W)) {
	ForEach2[V, W](iter, fn)
}

// Reduce consumes an iterator and applies a function to each value yielded,
// accumulating a single result.
func Reduce[V any, R any](iter Iterator[V], fn func(R, V) R, initial R) R {
	result := initial

	for item := range iter {
		result = fn(result, item)
	}

	return result
}

// Reduce2 consumes an iterator and applies a function to each pair of values,
// accumulating a single result.
func Reduce2[V, W any, R any](iter Iterator2[V, W], fn func(R, V, W) R, initial R) R {
	result := initial

	for v, w := range iter {
		result = fn(result, v, w)
	}

	return result
}
