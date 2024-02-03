package iter

import "iter"

// Lift yields all items in the provided slice.
func Lift[V any](slice []V) Iterator[V] {
	return Iterator[V](iter.Seq[V](func(yield func(V) bool) {
		for _, item := range slice {
			if !yield(item) {
				return
			}
		}
	}))
}

// LiftHashMap yields all items in the provided map as key-value pairs.
func LiftHashMap[K comparable, V any](hashmap map[K]V) Iterator[Pair[K, V]] {
	return Iterator[Pair[K, V]](iter.Seq[Pair[K, V]](func(yield func(Pair[K, V]) bool) {
		for k, v := range hashmap {
			if !yield(Pair[K, V]{k, v}) {
				return
			}
		}
	}))
}
