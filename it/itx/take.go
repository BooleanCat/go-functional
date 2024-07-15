package itx

import "github.com/BooleanCat/go-functional/v2/it"

// Take is a convenience method for chaining [it.Take] on [Iterator]s.
func (iterator Iterator[V]) Take(limit int) Iterator[V] {
	return Iterator[V](it.Take(iterator, limit))
}

// Take is a convenience method for chaining [it.Take2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Take(limit int) Iterator2[V, W] {
	return Iterator2[V, W](it.Take2(iterator, limit))
}
