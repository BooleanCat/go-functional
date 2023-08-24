package iter

import "github.com/BooleanCat/go-functional/option"

// EnumerateIter iterator, see [Enumerate].
type EnumerateIter[T any] struct {
	BaseIter[T]
	counter   uint
	exhausted bool
}

// Drop instantiates an [*EnumerateIter] that yield Pairs of the index of
// iteration and values for a given iterator.
func Enumerate[T any](iterator Iterator[T]) *EnumerateIter[T] {
	return &EnumerateIter[T]{BaseIter: BaseIter[T]{iterator}}
}

// Next implements the [Iterator] interface.
func (iter *EnumerateIter[T]) Next() option.Option[Pair[uint, T]] {
	if iter.exhausted {
		return option.None[Pair[uint, T]]()
	}

	value, ok := iter.BaseIter.Next().Value()
	if !ok {
		iter.exhausted = true
		return option.None[Pair[uint, T]]()
	}

	next := Pair[uint, T]{iter.counter, value}
	iter.counter++

	return option.Some(next)
}

var _ Iterator[Pair[uint, struct{}]] = new(EnumerateIter[struct{}])
