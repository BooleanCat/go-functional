package iter

import "github.com/BooleanCat/go-functional/option"

// DropIter implements `Drop`. See `Drop`'s documentation.
type DropIter[T any] struct {
	iter      Iterator[T]
	count     uint
	dropped   bool
	exhausted bool
}

// Drop instantiates a `DropIter` that will skip the number of items of its
// wrapped iterator by the provided count.
func Drop[T any](iter Iterator[T], count uint) *DropIter[T] {
	return &DropIter[T]{iter, count, false, false}
}

// Next implements the Iterator interface for `DropIter`.
func (iter *DropIter[T]) Next() option.Option[T] {
	if iter.exhausted {
		return option.None[T]()
	}

	if !iter.dropped {
		iter.dropped = true

		for i := uint(0); i < iter.count; i++ {
			if iter.delegateNext().IsNone() {
				return option.None[T]()
			}
		}
	}

	return iter.delegateNext()
}

func (iter *DropIter[T]) delegateNext() option.Option[T] {
	next := iter.iter.Next()
	if next.IsNone() {
		iter.exhausted = true
	}

	return next
}

var _ Iterator[struct{}] = new(DropIter[struct{}])

// Collect is an alternative way of invoking Collect(iter)
func (iter *DropIter[T]) Collect() []T {
	return Collect[T](iter)
}
