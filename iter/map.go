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
//
// Unlike other iterators (e.g. [Filter]), it is not possible to call Map as a
// method on iterators defined in this package. This is due to a limitation of
// Go's type system; new type parameters cannot be defined on methods.
func Map[T, U any](iter Iterator[T], f func(T) U) *MapIter[T, U] {
	iterator := &MapIter[T, U]{iter: iter, fun: f}
	iterator.BaseIter = BaseIter[U]{iterator}
	return iterator
}

// Transform instantiates a [*MapIter] that will apply the provided function to
// each item yielded by the provided [Iterator].
//
// Transform is a special case of [Map] where the argument and returned value
// for the operation are of the same type. This makes it possible to call
// Transform as a method on other iterators.
func Transform[T any](iter Iterator[T], op func(T) T) *MapIter[T, T] {
	return Map[T, T](iter, op)
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
