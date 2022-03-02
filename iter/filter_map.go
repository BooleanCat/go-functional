package iter

import (
	"github.com/BooleanCat/go-functional/option"
)

// FilterMapIter implements `FilterMap`. See `FilterMap`'s documentation.
type FilterMapIter[T any, U any] struct {
	iter Iterator[T]
	fun  func(T) option.Option[U]
}

// FilterMap instantiates a `FilterMapIter` that selectively yields the inner value of `Some` variants
// from the provided function. See `Option`s documentation..
func FilterMap[T any, U any](iter Iterator[T], fun func(T) option.Option[U]) *FilterMapIter[T, U] {
	return &FilterMapIter[T, U]{iter, fun}
}

// Next implements the Iterator interface for `FilterMap`.
func (iter *FilterMapIter[T, U]) Next() option.Option[U] {
	for {
		value, ok := iter.iter.Next().Value()
		if !ok {
			return option.None[U]()
		}

		opt := iter.fun(value)
		if opt.IsSome() {
			return opt
		}
	}
}

var _ Iterator[struct{}] = new(FilterMapIter[struct{}, struct{}])
