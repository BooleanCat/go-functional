package iter

import (
	"iter"

	"github.com/BooleanCat/go-functional/v2/iter/filter"
)

// Filter yields values from an iterator that satisfy a predicate.
func Filter[V any](delegate iter.Seq[V], predicate func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		delegate, stop := iter.Pull(delegate)
		defer stop()

		for {
			value, ok := delegate()

			if !ok || predicate(value) && !yield(value) {
				return
			}
		}
	}
}

// Exclude yields values from an iterator that do not satisfy a predicate.
func Exclude[V any](delegate iter.Seq[V], predicate func(V) bool) iter.Seq[V] {
	return Filter[V](delegate, filter.Not[V](predicate))
}

// Filter is a convenience method for chaining [Filter] on [Iterator]s.
func (iterator Iterator[V]) Filter(predicate func(V) bool) Iterator[V] {
	return Iterator[V](Filter[V](iter.Seq[V](iterator), predicate))
}

// Exclude is a convenience method for chaining [Exclude] on [Iterator]s.
func (iterator Iterator[V]) Exclude(predicate func(V) bool) Iterator[V] {
	return Iterator[V](Exclude[V](iter.Seq[V](iterator), predicate))
}

// Filter2 yields values from an iterator that satisfy a predicate.
func Filter2[V, W any](delegate iter.Seq2[V, W], predicate func(V, W) bool) iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {
		delegate, stop := iter.Pull2(delegate)
		defer stop()

		for {
			left, right, ok := delegate()

			if !ok || predicate(left, right) && !yield(left, right) {
				return
			}
		}
	}
}

// Exclude2 yields values from an iterator that do not satisfy a predicate.
func Exclude2[V, W any](delegate iter.Seq2[V, W], predicate func(V, W) bool) iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {
		delegate, stop := iter.Pull2(delegate)
		defer stop()

		for {
			left, right, ok := delegate()

			if !ok || !predicate(left, right) && !yield(left, right) {
				return
			}
		}
	}
}

// Filter is a convenience method for chaining [Filter] on [Iterator2]s.
func (iterator Iterator2[V, W]) Filter(predicate func(V, W) bool) Iterator2[V, W] {
	return Iterator2[V, W](Filter2[V, W](iter.Seq2[V, W](iterator), predicate))
}

// Exclude is a convenience method for chaining [Exclude] on [Iterator2]s.
func (iterator Iterator2[V, W]) Exclude(predicate func(V, W) bool) Iterator2[V, W] {
	return Iterator2[V, W](Exclude2[V, W](iter.Seq2[V, W](iterator), predicate))
}
