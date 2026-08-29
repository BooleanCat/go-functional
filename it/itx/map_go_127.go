//go:build go1.27

package itx

import "github.com/BooleanCat/go-functional/v2/it"

// Map is a convenience method for chaining [it.Map] on [Iterator]s.
func (iterator Iterator[V]) Map[W any](f func(V) W) Iterator[W] {
	return Iterator[W](it.Map(iterator, f))
}

// Map is a convenience method for chaining [it.Map2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Map[X, Y any](f func(V, W) (X, Y)) Iterator2[X, Y] {
	return Iterator2[X, Y](it.Map2(iterator, f))
}

// MapError is a convenience method for chaining [it.MapError] on [Iterator]s.
func (iterator Iterator[V]) MapError[W any](f func(V) (W, error)) Iterator2[W, error] {
	return Iterator2[W, error](it.MapError(iterator, f))
}
