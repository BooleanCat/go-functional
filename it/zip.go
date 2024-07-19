package it

import "iter"

// Zip yields pairs of values from two iterators.
func Zip[V, W any](left func(func(V) bool), right func(func(W) bool)) iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {
		left, stop := iter.Pull(left)
		defer stop()

		right, stop := iter.Pull(right)
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
	}
}

// Left is a convenience function that unzips an iterator and returns the left
// iterator, closing the right iterator.
func Left[V, W any](delegate func(func(V, W) bool)) iter.Seq[V] {
	return func(yield func(V) bool) {
		for left := range delegate {
			if !yield(left) {
				return
			}
		}
	}
}

// Right is a convenience function that unzips an iterator and returns the
// right iterator, closing the left iterator.
func Right[V, W any](delegate func(func(V, W) bool)) iter.Seq[W] {
	return func(yield func(W) bool) {
		for _, right := range delegate {
			if !yield(right) {
				return
			}
		}
	}
}
