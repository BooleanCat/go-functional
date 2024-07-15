package itx

import (
	"slices"

	"github.com/BooleanCat/go-functional/v2/it"
)

// Chain is a convenience method for chaining [it.Chain] on [Iterator]s.
func (iterator Iterator[V]) Chain(iterators ...func(func(V) bool)) Iterator[V] {
	iter := (func(func(V) bool))(iterator)
	return Iterator[V](it.Chain(slices.Insert(iterators, 0, iter)...))
}

// Chain is a convenience method for chaining [it.Chain2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Chain(iterators ...func(func(V, W) bool)) Iterator2[V, W] {
	iter := (func(func(V, W) bool))(iterator)
	return Iterator2[V, W](it.Chain2(slices.Insert(iterators, 0, iter)...))
}
