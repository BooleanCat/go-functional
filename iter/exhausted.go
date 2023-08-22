package iter

import "github.com/BooleanCat/go-functional/option"

// ExhaustedIter iterator, see [Exhausted].
type ExhaustedIter[T any] struct{}

// Exhausted instantiates an [*ExhaustedIter] that will immediately be
// exhausted (Next will always return a None variant).
func Exhausted[T any]() *ExhaustedIter[T] {
	return new(ExhaustedIter[T])
}

// Next implements the [Iterator] interface.
func (iter *ExhaustedIter[T]) Next() option.Option[T] {
	return option.None[T]()
}

var _ Iterator[struct{}] = new(ExhaustedIter[struct{}])

// Collect is a convenience method for [Collect], providing this iterator as
// an argument.
func (iter *ExhaustedIter[T]) Collect() []T {
	return Collect[T](iter)
}

// ForEach is a convenience method for [ForEach], providing this iterator as an
// argument.
func (iter *ExhaustedIter[T]) ForEach(callback func(T)) {
	ForEach[T](iter, callback)
}

// Drop is a convenience method for [Drop], providing this iterator as an
// argument.
func (iter *ExhaustedIter[T]) Drop(n uint) *DropIter[T] {
	return Drop[T](iter, n)
}

// Take is a convenience method for [Take], providing this iterator as an
// argument.
func (iter *ExhaustedIter[T]) Take(n uint) *TakeIter[T] {
	return Take[T](iter, n)
}
