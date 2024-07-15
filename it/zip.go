package it

import (
	"iter"
	"sync"

	"github.com/BooleanCat/go-functional/v2/internal/fifo"
)

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

// Unzip returns two iterators yielding single values from an iterator yielding
// pairs of values.
//
// Each returned iterator yields the left and right values from the original
// iterator, respectively.
//
// It is safe to concurrently pull from the returned iterators.
//
// Both returned iterators must be stopped, the underlying iterators is stopped
// when both are stopped. It is safe to stop one of the returned iterators
// immediately and continue pulling from the other.
func Unzip[V, W any](delegate func(func(V, W) bool)) (iter.Seq[V], iter.Seq[W]) {
	mutex := sync.Mutex{}

	next, stop := iter.Pull2(delegate)

	queue := fifo.New[V, W]()

	done := sync.WaitGroup{}
	done.Add(2)

	go func() {
		done.Wait()
		stop()
	}()

	return func(yield func(V) bool) {
			defer done.Done()

			for {
				mutex.Lock()
				left, ok := queue.DequeueLeft()
				if !ok {
					v, w, ok := next()

					if !ok {
						mutex.Unlock()
						return
					}

					queue.Enqueue(v, w)
					mutex.Unlock()
					continue
				}

				mutex.Unlock()

				if !yield(left) {
					return
				}
			}
		}, func(yield func(W) bool) {
			defer done.Done()

			for {
				mutex.Lock()
				right, ok := queue.DequeueRight()
				if !ok {
					v, w, ok := next()

					if !ok {
						mutex.Unlock()
						return
					}

					queue.Enqueue(v, w)
					mutex.Unlock()
					continue
				}

				mutex.Unlock()

				if !yield(right) {
					return
				}
			}
		}
}

// Left is a convenience function that unzips an iterator and returns the left
// iterator, closing the right iterator.
func Left[V, W any](delegate func(func(V, W) bool)) iter.Seq[V] {
	left, right := Unzip(delegate)

	_, stop := iter.Pull(right)
	stop()

	return left
}

// Right is a convenience function that unzips an iterator and returns the
// right iterator, closing the left iterator.
func Right[V, W any](delegate func(func(V, W) bool)) iter.Seq[W] {
	left, right := Unzip(delegate)

	_, stop := iter.Pull(left)
	stop()

	return right
}
