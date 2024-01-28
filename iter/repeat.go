package iter

import (
	"fmt"
	"reflect"

	"github.com/BooleanCat/go-functional/option"
)

// RepeatIter iterator, see [Repeat].
type RepeatIter[T any] struct {
	BaseIter[T]
	item T
}

// Repeat instantiates a [*RepeatIter] always yield the provided item.
//
// This iterator will never be exhausted.
func Repeat[T any](item T) *RepeatIter[T] {
	iter := &RepeatIter[T]{item: item}
	iter.BaseIter = BaseIter[T]{iter}
	return iter
}

// Next implements the [Iterator] interface.
func (iter *RepeatIter[T]) Next() option.Option[T] {
	return option.Some(iter.item)
}

// String implements the [fmt.Stringer] interface
func (iter RepeatIter[T]) String() string {
	var instanceOfT T
	return fmt.Sprintf("Iterator<Map, type=%s>", reflect.TypeOf(instanceOfT))
}

var (
	_ fmt.Stringer       = new(RepeatIter[struct{}])
	_ Iterator[struct{}] = new(RepeatIter[struct{}])
)
