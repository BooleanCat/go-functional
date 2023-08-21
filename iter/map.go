package iter

import "github.com/BooleanCat/go-functional/option"

// MapIter iterator, see [Map].
type MapIter[T, U any] struct {
	iter      Iterator[T]
	fun       func(T) U
	exhausted bool
}

// Map instantiates a [*MapIter] that will apply the provided function to each
// item yielded by the provided [Iterator].
func Map[T, U any](iter Iterator[T], f func(T) U) *MapIter[T, U] {
	return &MapIter[T, U]{iter, f, false}
}

// Next implements the [Iterator] interface.
func (iter *MapIter[T, U]) Next() option.Option[U] {
	if iter.exhausted {
		return option.None[U]()
	}

	value, ok := iter.iter.Next().Value()
	if !ok {
		iter.exhausted = true
		return option.None[U]()
	}

	return option.Some(iter.fun(value))
}

var _ Iterator[struct{}] = new(MapIter[struct{}, struct{}])

// Collect is a convenience method for [Collect], providing this iterator as
// an argument.
func (iter *MapIter[T, U]) Collect() []U {
	return Collect[U](iter)
}

// Drop is a convenience method for [Drop], providing this iterator as an
// argument.
func (iter *MapIter[T, U]) Drop(n uint) *DropIter[U] {
	return Drop[U](iter, n)
}

// Take is a convenience method for [Take], providing this iterator as an
// argument.
func (iter *MapIter[T, U]) Take(n uint) *TakeIter[U] {
	return Take[U](iter, n)
}
