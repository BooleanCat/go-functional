package itx

import "github.com/BooleanCat/go-functional/v2/it"

// Enumerate is a convenience method for chaining [it.Enumerate] on
// [Iterator]s.
func (iterator Iterator[V]) Enumerate() Iterator2[int, V] {
	return Iterator2[int, V](it.Enumerate(iterator))
}
