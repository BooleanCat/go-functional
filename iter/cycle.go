package iter

import "github.com/BooleanCat/go-functional/option"

// CycleIter iterator, see [Cycle].
type CycleIter[T any] struct {
	iter  Iterator[T]
	items []T
	index int
}

// Cycle instantiates a [*CycleIter] yielding all items from the provided
// iterator until exhaustion, and then yields them all over again on repeat.
//
// Note that [CycleIter] stores the members from the underlying iterator so
// will grow in size as items are yielded until the underlying iterator is
// exhausted.
//
// In most cases this iterator is infinite, except when the underlying iterator
// is exhausted before the first call to Next() in which case this iterator
// will always yield None.
func Cycle[T any](iter Iterator[T]) *CycleIter[T] {
	return &CycleIter[T]{iter, make([]T, 0), 0}
}

// Filter istantiates a [*FilterIter] for filtering by a chosen function.
func (iter *CycleIter[T]) Filter(fun func(T) bool) *FilterIter[T] {
	return &FilterIter[T]{iter, fun, false}
}

// Next implements the [Iterator] interface.
func (iter *CycleIter[T]) Next() option.Option[T] {
	if iter.iter != nil {
		if value, ok := iter.iter.Next().Value(); ok {
			iter.items = append(iter.items, value)
			return option.Some(value)
		} else {
			iter.iter = nil
		}
	}

	if len(iter.items) == 0 {
		return option.None[T]()
	}

	if iter.index == len(iter.items) {
		iter.index = 0
	}

	next := iter.items[iter.index]
	iter.index++
	return option.Some(next)
}

var _ Iterator[struct{}] = new(CycleIter[struct{}])

// ForEach is a convenience method for [ForEach], providing this iterator as an
// argument.
func (iter *CycleIter[T]) ForEach(callback func(T)) {
	ForEach[T](iter, callback)
}

// Find is a convenience method for [Find], providing this iterator as an
// argument.
func (iter *CycleIter[T]) Find(predicate func(T) bool) option.Option[T] {
	return Find[T](iter, predicate)
}

// Drop is a convenience method for [Drop], providing this iterator as an
// argument.
func (iter *CycleIter[T]) Drop(n uint) *DropIter[T] {
	return Drop[T](iter, n)
}

// Take is a convenience method for [Take], providing this iterator as an
// argument.
func (iter *CycleIter[T]) Take(n uint) *TakeIter[T] {
	return Take[T](iter, n)
}
