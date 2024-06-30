package itx

import (
	"iter"

	"github.com/BooleanCat/go-functional/v2/it"
)

// Chain is a convenience method for chaining [it.Chain] on [Iterator]s.
func (iterator Iterator[V]) Chain(iterators ...Iterator[V]) Iterator[V] {
	seqs := make([]iter.Seq[V], 0, len(iterators))

	for _, iterator := range iterators {
		seqs = append(seqs, iter.Seq[V](iterator))
	}

	return Iterator[V](it.Chain[V](append([]iter.Seq[V]{iter.Seq[V](iterator)}, seqs...)...))
}

// Chain is a convenience method for chaining [it.Chain2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Chain(iterators ...Iterator2[V, W]) Iterator2[V, W] {
	seqs := make([]iter.Seq2[V, W], 0, len(iterators))

	for _, iterator := range iterators {
		seqs = append(seqs, iter.Seq2[V, W](iterator))
	}

	return Iterator2[V, W](it.Chain2(append([]iter.Seq2[V, W]{iter.Seq2[V, W](iterator)}, seqs...)...))
}
