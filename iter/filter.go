package iter

import "github.com/BooleanCat/go-functional/option"

// FilterIter implements `Filter`. See `Filter`'s documentation.
type FilterIter[T any] struct {
	iter      Iterator[T]
	fun       func(T) bool
	exhausted bool
}

// Filter instantiates a `FilterIter` that selectively yields only results that
// cause the provided function to return `true`.
func Filter[T any](iter Iterator[T], fun func(T) bool) *FilterIter[T] {
	return &FilterIter[T]{iter, fun, false}
}

// Next implements the Iterator interface for `Filter`.
func (iter *FilterIter[T]) Next() option.Option[T] {
	if iter.exhausted {
		return option.None[T]()
	}

	for {
		value, ok := iter.iter.Next().Value()
		if !ok {
			iter.exhausted = true
			return option.None[T]()
		}

		if iter.fun(value) {
			return option.Some(value)
		}
	}
}

var _ Iterator[struct{}] = new(FilterIter[struct{}])

// Exclude instantiates a `FilterIter` that selectively yields only results that
// cause the provided function to return `false`.
func Exclude[T any](iter Iterator[T], fun func(T) bool) *FilterIter[T] {
	inverse := func(t T) bool { return !fun(t) }
	return &FilterIter[T]{iter, inverse, false}
}

type FilterMapIter[T any, U any] struct {
	iter      Iterator[T]
	fn        func(T) option.Option[U]
	exhausted bool
}

func (iter *FilterMapIter[T, U]) Next() option.Option[U] {
	if iter.exhausted {
		return option.None[U]()
	}

	for {
		val, ok := iter.iter.Next().Value()
		if !ok {
			iter.exhausted = true
			return option.None[U]()
		}

		result := iter.fn(val)
		if result.IsSome() {
			return result
		}
	}
}

var _ Iterator[struct{}] = new(FilterMapIter[struct{}, struct{}])

// Accepts an underlying iterator as data source and a filtering + mapping function
// it allows the user to filter elements by returning a None variant and to transform
// elements by returning a Some variant.
func FilterMap[T any, U any](itr Iterator[T], fun func(T) option.Option[U]) Iterator[U] {
	return &FilterMapIter[T, U]{itr, fun, false}
}
