package itx

import (
	"iter"
	"slices"

	"github.com/BooleanCat/go-functional/v2/it"
)

type (
	// Iterator is a wrapper around [iter.Seq] that allows for method chaining of
	// most iterators found in the `it` package.
	Iterator[V any] iter.Seq[V]

	// Iterator2 is a wrapper around [iter.Seq2] that allows for method chaining
	// of most iterators found in the `it` package.
	Iterator2[V, W any] iter.Seq2[V, W]
)

// Collect is a convenience method for chaining [slices.Collect] on
// [Iterator]s.
func (iterator Iterator[V]) Collect() []V {
	return slices.Collect(iter.Seq[V](iterator))
}

// ForEach is a convenience method for chaining [it.ForEach] on [Iterator]s.
func (iterator Iterator[V]) ForEach(fn func(V)) {
	it.ForEach(iter.Seq[V](iterator), fn)
}

// ForEach is a convenience method for chaining [it.ForEach2] on [Iterator2]s.
func (iterator Iterator2[V, W]) ForEach(fn func(V, W)) {
	it.ForEach2(iter.Seq2[V, W](iterator), fn)
}
