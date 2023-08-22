package iter

import "github.com/BooleanCat/go-functional/option"

// RepeatIter iterator, see [Repeat].
type RepeatIter[T any] struct {
	item T
}

// Repeat instantiates a [*RepeatIter] always yield the provided item.
//
// This iterator will never be exhausted.
func Repeat[T any](item T) *RepeatIter[T] {
	return &RepeatIter[T]{item}
}

// Next implements the [Iterator] interface.
func (iter *RepeatIter[T]) Next() option.Option[T] {
	return option.Some(iter.item)
}

var _ Iterator[struct{}] = new(RepeatIter[struct{}])

// ForEach is a convenience method for [ForEach], providing this iterator as an
// argument.
func (iter *RepeatIter[T]) ForEach(callback func(T)) {
	ForEach[T](iter, callback)
}

// Find is a convenience method for [Find], providing this iterator as an
// argument.
func (iter *RepeatIter[T]) Find(predicate func(T) bool) option.Option[T] {
	return Find[T](iter, predicate)
}

// Drop is a convenience method for [Drop], providing this iterator as an
// argument.
func (iter *RepeatIter[T]) Drop(n uint) *DropIter[T] {
	return Drop[T](iter, n)
}

// Take is a convenience method for [Take], providing this iterator as an
// argument.
func (iter *RepeatIter[T]) Take(n uint) *TakeIter[T] {
	return Take[T](iter, n)
}
