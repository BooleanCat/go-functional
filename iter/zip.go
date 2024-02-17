package iter

import (
	"iter"
	"sync"

	"github.com/BooleanCat/go-functional/v2/internal/fifo"
)

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

// Unzip returns two [Iterator]s from a single [Iterator2].
//
// Each returned [Iterator] yields the left and right values from the original
// [Iterator2], respectively.
//
// It is safe to concurrently pull from the returned [Iterator]s.
//
// Both returned [Iterator]s must be stopped, the underlying [Iterator2] is
// stopped when both are stopped. It is safe to stop one of the returned
// [Iterator]s immediately and continue pulling from the other.
func Unzip[V, W any](delegate Iterator2[V, W]) (Iterator[V], Iterator[W]) {
	mutex := sync.Mutex{}

	next, stop := iter.Pull2(iter.Seq2[V, W](delegate))

	queue := fifo.New[V, W]()

	done := sync.WaitGroup{}
	done.Add(2)

	go func() {
		done.Wait()
		stop()
	}()

	return Iterator[V](iter.Seq[V](func(yield func(V) bool) {
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
		})), Iterator[W](iter.Seq[W](func(yield func(W) bool) {
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
		}))
}

// Unzip is a convenience method for chaining [Unzip] on [Iterator2]s.
func (iter Iterator2[V, W]) Unzip() (Iterator[V], Iterator[W]) {
	return Unzip[V, W](iter)
}
