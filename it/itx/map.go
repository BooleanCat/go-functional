package itx

import "github.com/BooleanCat/go-functional/v2/it"

// Transform is a convenience method for chaining [it.Map] on [Iterator]s where
// the provided functions argument type is the same as its return type.
//
// This is a limited version of [it.Map] due to a limitation on Go's type
// system before Go 1.27 whereby new generic type parameters cannot be defined on methods.
//
// Deprecated: This will be removed in later versions and clients should transition to [Iterator.Map].
func (iterator Iterator[V]) Transform(f func(V) V) Iterator[V] {
	return Iterator[V](it.Map(iterator, f))
}

// Transform is a convenience method for chaining [it.Map2] on [Iterator2]s
// where the provided functions argument type is the same as its return type.
//
// This is a limited version of [it.Map2] due to a limitation on Go's type
// system before Go 1.27 whereby new generic type parameters cannot be defined on methods.
//
// Deprecated: This will be removed in later versions and clients should transition to [Iterator2.Map].
func (iterator Iterator2[V, W]) Transform(f func(V, W) (V, W)) Iterator2[V, W] {
	return Iterator2[V, W](it.Map2(iterator, f))
}

// TransformError is a convenience method for chaining [it.MapError] on
// [Iterator]s where the provided functions argument type is the same as its
// return type.
//
// This is a limited version of [it.MapError] due to a limitation on Go's type
// system before Go 1.27 whereby new generic type parameters cannot be defined on methods.
//
// Deprecated: This will be removed in later versions and clients should transition to [Iterator.MapError].
func (iterator Iterator[V]) TransformError(f func(V) (V, error)) Iterator2[V, error] {
	return Iterator2[V, error](it.MapError(iterator, f))
}
