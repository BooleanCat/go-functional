package iter

import "github.com/BooleanCat/go-functional/option"

// TakeIter implements `Take`. See `Take`'s documentation.
type TakeIter[T any] struct {
	iter  Iterator[T]
	limit int
}

// Take instantiates a `TakeIter` that will limit the number of items of its
// wrapped iterator to a maximum limit.
func Take[T any](iter Iterator[T], limit int) *TakeIter[T] {
	return &TakeIter[T]{iter, limit}
}

// Next implements the Iterator interface for `TakeIter`.
func (iter *TakeIter[T]) Next() option.Option[T] {
	if iter.limit == 0 {
		return option.None[T]()
	}

	next := iter.iter.Next()
	if next.IsNone() {
		iter.limit = 0
	} else {
		iter.limit -= 1
	}

	return next
}

var _ Iterator[struct{}] = new(TakeIter[struct{}])

// Collect is an alternative way of invoking Collect(iter)
func (iter *TakeIter[T]) Collect() []T {
	return Collect[T](iter)
}

// Drop is an alternative way of invoking Drop(iter)
func (iter *TakeIter[T]) Drop(n uint) *DropIter[T] {
	return Drop[T](iter, n)
}
