package itx

import (
	"iter"
	"maps"
)

// MapsAll is a wrapper around maps.All returning an [Iterator2] rather than an
// [iter.Seq2].
func MapsAll[Map ~map[K]V, K comparable, V any](m Map) Iterator2[K, V] {
	return Iterator2[K, V](maps.All(m))
}

// MapsCollect is a wrapper around maps.Collect accepting an [Iterator2] rather
// than an [iter.Seq2].
//
// Note: there is no convenience method for chaining Collect on an
// [it.Iterator2] because Go's type system will not allow for specifying a
// comparable constraint on a type parameter in the method definition.
func MapsCollect[K comparable, V any](iterator Iterator2[K, V]) map[K]V {
	return maps.Collect(iter.Seq2[K, V](iterator))
}
