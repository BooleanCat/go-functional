package itx

import (
	"iter"

	"github.com/BooleanCat/go-functional/v2/it"
)

// Filter is a convenience method for chaining [it.Filter] on [Iterator]s.
func (iterator Iterator[V]) Filter(predicate func(V) bool) Iterator[V] {
	return Iterator[V](it.Filter(iter.Seq[V](iterator), predicate))
}

// Exclude is a convenience method for chaining [it.Exclude] on [Iterator]s.
func (iterator Iterator[V]) Exclude(predicate func(V) bool) Iterator[V] {
	return Iterator[V](it.Exclude(iter.Seq[V](iterator), predicate))
}

// Filter is a convenience method for chaining [it.Filter2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Filter(predicate func(V, W) bool) Iterator2[V, W] {
	return Iterator2[V, W](it.Filter2(iter.Seq2[V, W](iterator), predicate))
}

// Exclude is a convenience method for chaining [it.Exclude2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Exclude(predicate func(V, W) bool) Iterator2[V, W] {
	return Iterator2[V, W](it.Exclude2[V, W](iter.Seq2[V, W](iterator), predicate))
}
