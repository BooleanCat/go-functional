package iter

import (
	"fmt"
	"reflect"

	"github.com/BooleanCat/go-functional/option"
)

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

// String implements the [fmt.Stringer] interface
func (iter ChannelIter[T]) String() string {
	var instanceOfT T
	return fmt.Sprintf("Iterator<Channel, type=%s>", reflect.TypeOf(instanceOfT))
}

var (
	_ fmt.Stringer       = new(ChannelIter[struct{}])
	_ Iterator[struct{}] = new(ChannelIter[struct{}])
)
