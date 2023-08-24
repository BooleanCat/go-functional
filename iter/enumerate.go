package iter

import "github.com/BooleanCat/go-functional/option"

// EnumerateIter iterator, see [Enumerate].
type EnumerateIter[T any] struct {
	BaseIter[T]
	counter   uint
	exhausted bool
}

// Drop instantiates an [*EnumerateIter] that yield Tuples of the index of
// iteration and values for a given iterator.
func Enumerate[T any](iterator Iterator[T]) *EnumerateIter[T] {
	return &EnumerateIter[T]{BaseIter: BaseIter[T]{iterator}}
}

// Next implements the [Iterator] interface.
func (iter *EnumerateIter[T]) Next() option.Option[Tuple[uint, T]] {
	if iter.exhausted {
		return option.None[Tuple[uint, T]]()
	}

	value, ok := iter.BaseIter.Next().Value()
	if !ok {
		iter.exhausted = true
		return option.None[Tuple[uint, T]]()
	}

	next := Tuple[uint, T]{iter.counter, value}
	iter.counter++

	return option.Some(next)
}

var _ Iterator[Tuple[uint, struct{}]] = new(EnumerateIter[struct{}])
