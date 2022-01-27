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

	iter.limit -= 1

	return iter.iter.Next()
}

var _ Iterator[struct{}] = new(TakeIter[struct{}])
