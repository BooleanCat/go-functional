package iter

import "iter"

// Zip yields pairs of elements from the two provided iterators.
func Zip[V, W any](one Iterator[V], two Iterator[W]) Iterator[Pair[V, W]] {
	return Iterator[Pair[V, W]](iter.Seq[Pair[V, W]](func(yield func(Pair[V, W]) bool) {
		oneNext, oneStop := iter.Pull(iter.Seq[V](one))
		defer oneStop()

		twoNext, twoStop := iter.Pull(iter.Seq[W](two))
		defer twoStop()

		for {
			one, ok := oneNext()
			if !ok {
				return
			}

			two, ok := twoNext()
			if !ok {
				return
			}

			if !yield(Pair[V, W]{one, two}) {
				return
			}
		}
	}))
}
