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
func Unzip[V, W any](delegate Iterator2[V, W]) (Iterator[V], Iterator[W]) {
	mutex := sync.Mutex{}

	next, stop := iter.Pull2(iter.Seq2[V, W](delegate))

	queue := fifo.New[V, W]()

	return Iterator[V](iter.Seq[V](func(yield func(V) bool) {
			for {
				mutex.Lock()
				left, ok := queue.DequeueLeft()
				if !ok {
					v, w, ok := next()

					if !ok {
						stop()
						mutex.Unlock()
						return
					}

					queue.Enqueue(v, w)
					mutex.Unlock()
					continue
				}

				mutex.Unlock()

				if !yield(left) {
					mutex.Lock()
					stop()
					mutex.Unlock()
					return
				}
			}
		})), Iterator[W](iter.Seq[W](func(yield func(W) bool) {
			for {
				mutex.Lock()
				right, ok := queue.DequeueRight()
				if !ok {
					v, w, ok := next()

					if !ok {
						stop()
						mutex.Unlock()
						return
					}

					queue.Enqueue(v, w)
					mutex.Unlock()
					continue
				}

				mutex.Unlock()

				if !yield(right) {
					mutex.Lock()
					stop()
					mutex.Unlock()
					return
				}
			}
		}))
}

// Unzip is a convenience method for chaining [Unzip] on [Iterator2]s.
func (iter Iterator2[V, W]) Unzip() (Iterator[V], Iterator[W]) {
	return Unzip[V, W](iter)
}
