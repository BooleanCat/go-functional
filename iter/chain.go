package iter

import "github.com/BooleanCat/go-functional/option"

// ChainIter implements `Chain`. See `Chain`'s documentation.
type ChainIter[T any] struct {
	iterators     []Iterator[T]
	iteratorIndex int
}

// Chain instantiates a `ChainIter` that will yield all items in the provided
// iterators to exhaustion, from left to right.
func Chain[T any](iterators ...Iterator[T]) *ChainIter[T] {
	return &ChainIter[T]{iterators, 0}
}

// Next implements the Iterator interface for `Chain`.
func (iter *ChainIter[T]) Next() option.Option[T] {
	for {
		if iter.iteratorIndex == len(iter.iterators) {
			return option.None[T]()
		}

		if value, ok := iter.iterators[iter.iteratorIndex].Next().Value(); ok {
			return option.Some(value)
		}

		iter.iteratorIndex++
	}

}

var _ Iterator[struct{}] = new(ChainIter[struct{}])

// Collect is an alternative way of invoking Collect(iter)
func (iter *ChainIter[T]) Collect() []T {
	return Collect[T](iter)
}

// Drop is an alternative way of invoking Drop(iter)
func (iter *ChainIter[T]) Drop(n uint) *DropIter[T] {
	return Drop[T](iter, n)
}
