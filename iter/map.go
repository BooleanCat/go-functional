package iter

import "github.com/BooleanCat/go-functional/option"

// MapIter implements `Map`. See `Map`'s documentation.
type MapIter[T, U any] struct {
	iter Iterator[T]
	fun  func(T) U
}

// Take instantiates a `MapIter` that will apply the provided function to each
// item yielded by the provided Iterator.
func Map[T, U any](iter Iterator[T], f func(T) U) *MapIter[T, U] {
	return &MapIter[T, U]{iter, f}
}

// Next implements the Iterator interface for `MapIter`.
func (iter *MapIter[T, U]) Next() option.Option[U] {
	value, ok := iter.iter.Next().Value()
	if !ok {
		return option.None[U]()
	}

	return option.Some(iter.fun(value))
}

var _ Iterator[struct{}] = new(MapIter[struct{}, struct{}])
