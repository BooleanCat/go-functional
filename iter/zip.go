package iter

import "iter"

// Zip yields pairs of values from two iterators.
func Zip[V, W any](left Iterator[V], right Iterator[W]) Iterator2[V, W] {
	return Iterator2[V, W](iter.Seq2[V, W](func(yield func(V, W) bool) {
		left, stop := iter.Pull(iter.Seq[V](left))
		defer stop()

		right, stop := iter.Pull(iter.Seq[W](right))
		defer stop()

		for {
			leftValue, leftOk := left()
			rightValue, rightOk := right()

			if !leftOk || !rightOk {
				return
			}

			if !yield(leftValue, rightValue) {
				return
			}
		}
	}))
}
