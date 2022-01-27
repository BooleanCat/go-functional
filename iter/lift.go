package iter

import "github.com/BooleanCat/go-functional/option"

// LiftIter implements `Lift`. See `Lift`'s documentation.
type LiftIter[T any] struct {
	items []T
	index int
}

// Lift instantiates a `LiftIter` that will yield all items in the provided
// slice.
func Lift[T any](items []T) *LiftIter[T] {
	return &LiftIter[T]{items, 0}
}

// Next implements the Iterator interface for `Lift`.
func (iter *LiftIter[T]) Next() option.Option[T] {
	if iter.index >= len(iter.items) {
		return option.None[T]()
	}

	iter.index++

	return option.Some(iter.items[iter.index-1])
}

var _ Iterator[struct{}] = new(LiftIter[struct{}])
