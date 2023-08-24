package iter

import "github.com/BooleanCat/go-functional/option"

// MapIter iterator, see [Map].
type MapIter[T, U any] struct {
	BaseIter[U]
	iter      Iterator[T]
	fun       func(T) U
	exhausted bool
}

// Map instantiates a [*MapIter] that will apply the provided function to each
// item yielded by the provided [Iterator].
func Map[T, U any](iter Iterator[T], f func(T) U) *MapIter[T, U] {
	iterator := &MapIter[T, U]{iter: iter, fun: f}
	iterator.BaseIter = BaseIter[U]{iterator}
	return iterator
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

// ForEach is a convenience method for [ForEach], providing this iterator as an
// argument.
func (iter *MapIter[T, U]) ForEach(callback func(U)) {
	ForEach[U](iter, callback)
}

// Find is a convenience method for [Find], providing this iterator as an
// argument.
func (iter *MapIter[T, U]) Find(predicate func(U) bool) option.Option[U] {
	return Find[U](iter, predicate)
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
