//go:build go1.22 && goexperiment.rangefunc

package iter

import "iter"

// Chain yields values from multiple iterators in sequence.
func Chain[V any](iterators ...iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, iterator := range iterators {
			iterator(yield)
		}
	}
}

// Chain is a convenience method for chaining [Chain] on [Iterator]s.
func (iterator Iterator[V]) Chain(iterators ...iter.Seq[V]) Iterator[V] {
	return Iterator[V](Chain[V](append([]iter.Seq[V]{iter.Seq[V](iterator)}, iterators...)...))
}

// Chain2 yields values from multiple iterators in sequence.
func Chain2[V, W any](iterators ...iter.Seq2[V, W]) iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {
		for _, iterator := range iterators {
			iterator(yield)
		}
	}
}

// Chain is a convenience method for chaining [Chain2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Chain(iterators ...iter.Seq2[V, W]) Iterator2[V, W] {
	return Iterator2[V, W](Chain2(append([]iter.Seq2[V, W]{iter.Seq2[V, W](iterator)}, iterators...)...))
}
