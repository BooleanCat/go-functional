package iter

import "github.com/BooleanCat/go-functional/option"

// TakeIter iterator, see [Take].
type TakeIter[T any] struct {
	BaseIter[T]
	iter  Iterator[T]
	limit uint
}

// Take instantiates a [*TakeIter] that will limit the number of items of its
// wrapped iterator to a maximum limit.
func Take[T any](iter Iterator[T], limit uint) *TakeIter[T] {
	iterator := &TakeIter[T]{iter: iter, limit: limit}
	iterator.BaseIter = BaseIter[T]{iterator}
	return iterator
}

// Next implements the [Iterator] interface.
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

// ForEach is a convenience method for [ForEach], providing this iterator as an
// argument.
func (iter *TakeIter[T]) ForEach(callback func(T)) {
	ForEach[T](iter, callback)
}

// Find is a convenience method for [Find], providing this iterator as an
// argument.
func (iter *TakeIter[T]) Find(predicate func(T) bool) option.Option[T] {
	return Find[T](iter, predicate)
}

// Drop is a convenience method for [Drop], providing this iterator as an
// argument.
func (iter *TakeIter[T]) Drop(n uint) *DropIter[T] {
	return Drop[T](iter, n)
}

// Take is a convenience method for [Take], providing this iterator as an
// argument.
func (iter *TakeIter[T]) Take(n uint) *TakeIter[T] {
	return Take[T](iter, n)
}
