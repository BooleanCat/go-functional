package itx

import "github.com/BooleanCat/go-functional/v2/it"

// Cycle is a convenience method for chaining [it.Cycle] on [Iterator]s.
func (iterator Iterator[V]) Cycle() Iterator[V] {
	return Iterator[V](it.Cycle(iterator))
}

// Cycle is a convenience method for chaining [it.Cycle2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Cycle() Iterator2[V, W] {
	return Iterator2[V, W](it.Cycle2(iterator))
}
