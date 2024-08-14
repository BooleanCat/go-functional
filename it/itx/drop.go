package itx

import "github.com/BooleanCat/go-functional/v2/it"

// Drop is a convenience method for chaining [it.Drop] on [Iterator]s.
func (iterator Iterator[V]) Drop(count uint) Iterator[V] {
	return Iterator[V](it.Drop(iterator, count))
}

// Drop is a convenience method for chaining [it.Drop2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Drop(count uint) Iterator2[V, W] {
	return Iterator2[V, W](it.Drop2(iterator, count))
}

// DropWhile is a convenience method for chaining [it.DropWhile] on
// [Iterator]s.
func (iterator Iterator[V]) DropWhile(predicate func(V) bool) Iterator[V] {
	return Iterator[V](it.DropWhile(iterator, predicate))
}

// DropWhile is a convenience method for chaining [it.DropWhile2] on
// [Iterator2]s.
func (iterator Iterator2[V, W]) DropWhile(predicate func(V, W) bool) Iterator2[V, W] {
	return Iterator2[V, W](it.DropWhile2(iterator, predicate))
}
