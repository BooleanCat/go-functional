package itx

import "github.com/BooleanCat/go-functional/v2/it"

// Chain is a convenience method for chaining [it.Chain] on [Iterator]s.
func (iterator Iterator[V]) Chain(iterators ...func(func(V) bool)) Iterator[V] {
	return Iterator[V](it.Chain(append([]func(func(V) bool){iterator}, iterators...)...))
}

// Chain is a convenience method for chaining [it.Chain2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Chain(iterators ...func(func(V, W) bool)) Iterator2[V, W] {
	return Iterator2[V, W](it.Chain2(append([]func(func(V, W) bool){iterator}, iterators...)...))
}
