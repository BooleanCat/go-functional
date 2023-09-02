package iter

import (
	"fmt"

	"github.com/BooleanCat/go-functional/option"
)

// Pairs of values.
type Pair[T, U any] struct {
	One T
	Two U
}

func (p Pair[T, U]) String() string {
	one := fmt.Sprintf("%+v", p.One)
	two := fmt.Sprintf("%+v", p.Two)

	if val, ok := interface{}(p.One).(fmt.Stringer); ok {
		one = val.String()
	}

	if val, ok := interface{}(p.Two).(fmt.Stringer); ok {
		two = val.String()
	}

	return fmt.Sprintf("(%s, %s)", one, two)
}

// ZipIter iterator, see [Zip].
type ZipIter[T, U any] struct {
	BaseIter[Pair[T, U]]
	iter1     Iterator[T]
	iter2     Iterator[U]
	exhausted bool
}

// Zip instantiates a [*ZipIter] yielding a [Pair] containing the result of a
// call to each provided [Iterator]'s Next. This iterator is exhausted when one
// of the provided iterators is exhausted.
func Zip[T, U any](iter1 Iterator[T], iter2 Iterator[U]) *ZipIter[T, U] {
	iter := &ZipIter[T, U]{iter1: iter1, iter2: iter2}
	iter.BaseIter = BaseIter[Pair[T, U]]{iter}
	return iter
}

// Next implements the [Iterator] interface.
func (iter *ZipIter[T, U]) Next() option.Option[Pair[T, U]] {
	if iter.exhausted {
		return option.Option[Pair[T, U]]{}
	}

	v1, ok1 := iter.iter1.Next().Value()
	v2, ok2 := iter.iter2.Next().Value()

	if !ok1 || !ok2 {
		iter.exhausted = true
		return option.None[Pair[T, U]]()
	}

	return option.Some(Pair[T, U]{v1, v2})
}

var _ Iterator[Pair[struct{}, struct{}]] = new(ZipIter[struct{}, struct{}])
