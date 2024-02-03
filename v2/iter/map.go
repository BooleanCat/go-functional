package iter

import "iter"

// Map applies a transformation to each value from delegate.
//
// Unlike other iterators, Map cannot be chained as a method on iterators
// because of a limitation of Go's type system; new type parameters cannot be
// defined on methods. A limited version of this method is available as the
// [Transform] method on iterators where the argument and returned value for
// the operation are of the same type.
func Map[V, W any](delegate Iterator[V], transform func(V) W) Iterator[W] {
	return Iterator[W](iter.Seq[W](func(yield func(W) bool) {
		next, stop := iter.Pull(iter.Seq[V](delegate))
		defer stop()

		for {
			v, ok := next()
			if !ok {
				return
			}

			if !yield(transform(v)) {
				return
			}
		}
	}))
}

// Transform is a convenience method for chaining [Map] after an iterator.
//
// See [Map] for a method chaining caveat.
func (iter Iterator[V]) Transform(transform func(V) V) Iterator[V] {
	return Map(iter, transform)
}
