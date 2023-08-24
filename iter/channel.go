package iter

import "github.com/BooleanCat/go-functional/option"

// ChannelIter iterator, see [FromChannel].
type ChannelIter[T any] struct {
	BaseIter[T]
	item chan T
}

// FromChannel instantiates a [*ChannelIter] that will yield each value from
// the provided channel.
func FromChannel[T any](ch chan T) *ChannelIter[T] {
	iter := &ChannelIter[T]{item: ch}
	iter.BaseIter = BaseIter[T]{iter}
	return iter
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
