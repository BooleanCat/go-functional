package iter

import "github.com/BooleanCat/go-functional/option"

// ChainIter iterator, see [Chain].
type ChainIter[T any] struct {
	iterators     []Iterator[T]
	iteratorIndex int
}

// Chain instantiates a [*ChainIter] that will yield all items in the provided
// iterators to exhaustion first to last.
func Chain[T any](iterators ...Iterator[T]) *ChainIter[T] {
	return &ChainIter[T]{iterators, 0}
}

// Filter istantiates a [*FilterIter] for filtering by a chosen function.
func (iter *ChainIter[T]) Filter(fun func(T) bool) *FilterIter[T] {
	return &FilterIter[T]{iter, fun, false}
}

// Next implements the [Iterator] interface.
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

// Collect is a convenience method for [Collect], providing this iterator as an
// argument.
func (iter *ChainIter[T]) Collect() []T {
	return Collect[T](iter)
}

// ForEach is a convenience method for [ForEach], providing this iterator as an
// argument.
func (iter *ChainIter[T]) ForEach(callback func(T)) {
	ForEach[T](iter, callback)
}

// Find is a convenience method for [Find], providing this iterator as an
// argument.
func (iter *ChainIter[T]) Find(predicate func(T) bool) option.Option[T] {
	return Find[T](iter, predicate)
}

// Drop is a convenience method for [Drop], providing this iterator as an
// argument.
func (iter *ChainIter[T]) Drop(n uint) *DropIter[T] {
	return Drop[T](iter, n)
}

// Take is a convenience method for [Take], providing this iterator as an
// argument.
func (iter *ChainIter[T]) Take(n uint) *TakeIter[T] {
	return Take[T](iter, n)
}
