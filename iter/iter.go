//go:build go1.23

package iter

import (
	"iter"

	"github.com/BooleanCat/go-functional/v2/future/slices"
)

type (
	// Iterator is a wrapper around [iter.Seq] that allows for method chaining of
	// most iterators found in this package.
	Iterator[V any] iter.Seq[V]

	// Iterator is a wrapper around [iter.Seq2] that allows for method chaining
	// of most iterators found in this package.
	Iterator2[V, W any] iter.Seq2[V, W]
)

// Collect is a convenience method for chaining [Collect] on [Iterator]s.
func (iterator Iterator[V]) Collect() []V {
	return slices.Collect(iter.Seq[V](iterator))
}

// ForEach consumes an iterator and applies a function to each value yielded.
func ForEach[V any](iter iter.Seq[V], fn func(V)) {
	for item := range iter {
		fn(item)
	}
}

// ForEach is a convenience method for chaining [ForEach] on [Iterator]s.
func (iterator Iterator[V]) ForEach(fn func(V)) {
	ForEach(iter.Seq[V](iterator), fn)
}

// ForEach2 consumes an iterator and applies a function to each pair of values.
func ForEach2[V, W any](iter iter.Seq2[V, W], fn func(V, W)) {
	for v, w := range iter {
		fn(v, w)
	}
}

// ForEach is a convenience method for chaining [ForEach] on [Iterator2]s.
func (iterator Iterator2[V, W]) ForEach(fn func(V, W)) {
	ForEach2(iter.Seq2[V, W](iterator), fn)
}

// Reduce consumes an iterator and applies a function to each value yielded,
// accumulating a single result.
func Reduce[V any, R any](iter iter.Seq[V], fn func(R, V) R, initial R) R {
	result := initial

	for item := range iter {
		result = fn(result, item)
	}

	return result
}

// Reduce2 consumes an iterator and applies a function to each pair of values,
// accumulating a single result.
func Reduce2[V, W any, R any](iter iter.Seq2[V, W], fn func(R, V, W) R, initial R) R {
	result := initial

	for v, w := range iter {
		result = fn(result, v, w)
	}

	return result
}
