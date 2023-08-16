package iter

import "github.com/BooleanCat/go-functional/option"

// ChannelIter implements `FromChannel`. See `FromChannel`'s documentation.
type ChannelIter[T any] struct {
	item chan T
}

// FromChannel instantiates a `ChannelIter` that will yield each value from the provided channel.
func FromChannel[T any](ch chan T) *ChannelIter[T] {
	return &ChannelIter[T]{ch}
}

// Next implements the Iterator interface for `ChannelIter`.
func (iter *ChannelIter[T]) Next() option.Option[T] {

	value, ok := <-iter.item
	if !ok {
		return option.None[T]()
	}

	return option.Some(value)
}

var _ Iterator[struct{}] = new(ChannelIter[struct{}])

// Collect is an alternative way of invoking Collect(iter)
func (iter *ChannelIter[T]) Collect() []T {
	return Collect[T](iter)
}
