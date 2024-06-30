package itx

import (
	"iter"

	"github.com/BooleanCat/go-functional/v2/it"
)

// Unzip is a convenience method for chaining [it.Unzip] on [Iterator2]s.
func (iterator Iterator2[V, W]) Unzip() (Iterator[V], Iterator[W]) {
	left, right := it.Unzip[V, W](iter.Seq2[V, W](iterator))
	return Iterator[V](left), Iterator[W](right)
}
