package iter

import "github.com/BooleanCat/go-functional/option"

// CycleIter implements `Cycle`. See `Cycle`'s documentation.
type CycleIter[T any] struct {
	iter  Iterator[T]
	items []T
	index int
}

// Cycle instantiates a `CycleIter` yielding all items from the provided
// iterator until exhaustion, and then yields them all over again on repeat.
//
// Note that CycleIter stores the members from the underlying iterator so will
// grow in size as items are yielded until the underlying iterator is
// exhausted.
//
// In most cases this iterator is infinite, except when the underlying iterator
// is exhausted before the first call to Next() in which case this iterator
// will always yield None.
func Cycle[T any](iter Iterator[T]) *CycleIter[T] {
	return &CycleIter[T]{iter, make([]T, 0), 0}
}

// Next implements the Iterator interface for `CycleIter`.
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

// Drop is an alternative way of invoking Drop(iter)
func (iter *CycleIter[T]) Drop(n uint) *DropIter[T] {
	return Drop[T](iter, n)
}
