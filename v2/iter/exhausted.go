package iter

import "iter"

func Exhausted[V any]() Iterator[V] {
	return Iterator[V](iter.Seq[V](func(yield func(V) bool) {}))
}
