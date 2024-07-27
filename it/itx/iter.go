package itx

import (
	"iter"
	"maps"
	"slices"

	"github.com/BooleanCat/go-functional/v2/it"
)

type (
	// Iterator is a wrapper around [iter.Seq] that allows for method chaining of
	// most iterators found in the `it` package.
	Iterator[V any] func(func(V) bool)

	// Iterator2 is a wrapper around [iter.Seq2] that allows for method chaining
	// of most iterators found in the `it` package.
	Iterator2[V, W any] func(func(V, W) bool)
)

// From converts an iterator in an [Iterator] to support method chaining.
func From[V any](iterator func(func(V) bool)) Iterator[V] {
	return Iterator[V](iterator)
}

// From2 converts an iterator in an [Iterator2] to support method chaining.
func From2[V, W any](iterator func(func(V, W) bool)) Iterator2[V, W] {
	return Iterator2[V, W](iterator)
}

// FromSlice converts a slice to an [Iterator].
func FromSlice[V any](slice []V) Iterator[V] {
	return Iterator[V](slices.Values(slice))
}

// FromMap converts a map to an [Iterator2].
func FromMap[V comparable, W any](m map[V]W) Iterator2[V, W] {
	return Iterator2[V, W](maps.All(m))
}

// Seq converts an [Iterator] to an [iter.Seq].
func (iterator Iterator[V]) Seq() iter.Seq[V] {
	return iter.Seq[V](iterator)
}

// Seq converts an [Iterator2] to an [iter.Seq2].
func (iterator Iterator2[V, W]) Seq() iter.Seq2[V, W] {
	return iter.Seq2[V, W](iterator)
}

// Collect is a convenience method for chaining [slices.Collect] on
// [Iterator]s.
func (iterator Iterator[V]) Collect() []V {
	return slices.Collect(iter.Seq[V](iterator))
}

// ForEach is a convenience method for chaining [it.ForEach] on [Iterator]s.
func (iterator Iterator[V]) ForEach(fn func(V)) {
	it.ForEach(iterator, fn)
}

// ForEach is a convenience method for chaining [it.ForEach2] on [Iterator2]s.
func (iterator Iterator2[V, W]) ForEach(fn func(V, W)) {
	it.ForEach2(iterator, fn)
}

// Find is a convenience method for chaining [it.Find] on [Iterator]s.
func (iterator Iterator[V]) Find(predicate func(V) bool) (V, bool) {
	return it.Find(iterator, predicate)
}

// Find is a convenience method for chaining [it.Find2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Find(predicate func(V, W) bool) (V, W, bool) {
	return it.Find2(iterator, predicate)
}

// TryCollect consumes an [iter.Seq2] where the right side yields errors and
// returns a slice of values and the first error encountered. Iteration stops
// at the first error.
func TryCollect[V any](iterator func(func(V, error) bool)) ([]V, error) {
	return it.TryCollect(iterator)
}

// Collect2 consumes an [iter.Seq2] and returns two slices of values.
func (iterator Iterator2[V, W]) Collect() ([]V, []W) {
	return it.Collect2(iterator)
}

// Len is a convenience method for chaining [it.Len] on [Iterator]s.
func (iterator Iterator[V]) Len() int {
	return it.Len(iterator)
}

// Len is a convenience method for chaining [it.Len2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Len() int {
	return it.Len2(iterator)
}
