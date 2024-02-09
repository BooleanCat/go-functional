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

// Zip yields pairs of values from two iterators.
func Zip[V, W any](left Iterator[V], right Iterator[W]) Iterator2[V, W] {
	return Iterator2[V, W](iter.Seq2[V, W](func(yield func(V, W) bool) {
		left, stop := iter.Pull(iter.Seq[V](left))
		defer stop()

		right, stop := iter.Pull(iter.Seq[W](right))
		defer stop()

		for {
			leftValue, leftOk := left()
			rightValue, rightOk := right()

			if !leftOk || !rightOk {
				return
			}

			if !yield(leftValue, rightValue) {
				return
			}
		}
	}))
}
