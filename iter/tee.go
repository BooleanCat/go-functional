package iter

import (
	"sync"

	"github.com/BooleanCat/go-functional/internal/utils.go"
	"github.com/BooleanCat/go-functional/option"
)

type teeIters[T any] struct {
	iter      Iterator[T]
	buffer    []T
	lock      sync.Mutex
	oneIndex  int
	twoIndex  int
	exhausted bool
}

// Tee instantiates a [*TeeIters] that provides two iterators that yield all
// items from the provided iterator. The two iterators are independent, so
// consuming one will not affect the other.
//
// It is safe to consume the two iterators in parallel.
func Tee[T any](iter Iterator[T]) (*TeeIterOutput[T], *TeeIterOutput[T]) {
	teeIters := &teeIters[T]{iter: iter}

	one := &TeeIterOutput[T]{tee: teeIters, id: 1}
	one.BaseIter = BaseIter[T]{one}

	two := &TeeIterOutput[T]{tee: teeIters, id: 2}
	two.BaseIter = BaseIter[T]{two}

	return one, two
}

func (iter *teeIters[T]) take(id int) option.Option[T] {
	iter.lock.Lock()
	defer iter.lock.Unlock()

	if iter.oneIndex > 0 && iter.twoIndex > 0 {
		smaller := utils.Min(iter.oneIndex, iter.twoIndex)
		iter.oneIndex -= smaller
		iter.twoIndex -= smaller
		iter.buffer = iter.buffer[smaller:]
	}

	if id == 1 {
		if iter.oneIndex < len(iter.buffer) {
			next := iter.buffer[iter.oneIndex]
			iter.oneIndex++
			return option.Some(next)
		}

		if iter.exhausted && len(iter.buffer) == 0 {
			return option.None[T]()
		}

		if next := iter.iter.Next(); next.IsSome() {
			iter.buffer = append(iter.buffer, next.Unwrap())
			iter.oneIndex++
			return next
		}

		iter.exhausted = true

		return option.None[T]()
	}

	if iter.twoIndex < len(iter.buffer) {
		next := iter.buffer[iter.twoIndex]
		iter.twoIndex++
		return option.Some(next)
	}

	if iter.exhausted && len(iter.buffer) == 0 {
		return option.None[T]()
	}

	if next := iter.iter.Next(); next.IsSome() {
		iter.buffer = append(iter.buffer, next.Unwrap())
		iter.twoIndex++
		return next
	}

	iter.exhausted = true

	return option.None[T]()
}

// TeeIterOutput iterator, see [Tee].
type TeeIterOutput[T any] struct {
	BaseIter[T]
	tee *teeIters[T]
	id  int
}

// Next implements the [Iterator] interface.
func (iter *TeeIterOutput[T]) Next() option.Option[T] {
	return iter.tee.take(iter.id)
}
