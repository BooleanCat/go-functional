package iter

import "github.com/BooleanCat/go-functional/option"

// ChannelIter iterator, see [FromChannel].
type ChannelIter[T any] struct {
	item chan T
}

// FromChannel instantiates a [*ChannelIter] that will yield each value from
// the provided channel.
func FromChannel[T any](ch chan T) *ChannelIter[T] {
	return &ChannelIter[T]{ch}
}

// Filter istantiates a [*FilterIter] for filtering by a chosen function.
func (iter *ChannelIter[T]) Filter(fun func(T) bool) *FilterIter[T] {
	return &FilterIter[T]{iter, fun, false}
}

// Next implements the [Iterator] interface.
func (iter *ChannelIter[T]) Next() option.Option[T] {

	value, ok := <-iter.item
	if !ok {
		return option.None[T]()
	}

	return option.Some(value)
}

var _ Iterator[struct{}] = new(ChannelIter[struct{}])

// Collect is a convenience method for [Collect], providing this iterator as an
// argument.
func (iter *ChannelIter[T]) Collect() []T {
	return Collect[T](iter)
}

// ForEach is a convenience method for [ForEach], providing this iterator as an
// argument.
func (iter *ChannelIter[T]) ForEach(callback func(T)) {
	ForEach[T](iter, callback)
}

// Find is a convenience method for [Find], providing this iterator as an
// argument.
func (iter *ChannelIter[T]) Find(predicate func(T) bool) option.Option[T] {
	return Find[T](iter, predicate)
}

// Drop is a convenience method for [Drop], providing this iterator as an
// argument.
func (iter *ChannelIter[T]) Drop(n uint) *DropIter[T] {
	return Drop[T](iter, n)
}

// Take is a convenience method for [Take], providing this iterator as an
// argument.
func (iter *ChannelIter[T]) Take(n uint) *TakeIter[T] {
	return Take[T](iter, n)
}
