package itx

import "github.com/BooleanCat/go-functional/v2/it"

// Take is a convenience method for chaining [it.Take] on [Iterator]s.
func (iterator Iterator[V]) Take(limit uint) Iterator[V] {
	return Iterator[V](it.Take(iterator, limit))
}

// Take is a convenience method for chaining [it.Take2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Take(limit uint) Iterator2[V, W] {
	return Iterator2[V, W](it.Take2(iterator, limit))
}

// TakeWhile is a convenience method for chaining [it.TakeWhile] on
// [Iterator]s.
func (iterator Iterator[V]) TakeWhile(predicate func(V) bool) Iterator[V] {
	return Iterator[V](it.TakeWhile(iterator, predicate))
}

// TakeWhile is a convenience method for chaining [it.TakeWhile2] on
// [Iterator2]s.
func (iterator Iterator2[V, W]) TakeWhile(predicate func(V, W) bool) Iterator2[V, W] {
	return Iterator2[V, W](it.TakeWhile2(iterator, predicate))
}
