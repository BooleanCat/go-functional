package iter

import "github.com/BooleanCat/go-functional/option"

// FilterIter implements `Filter`. See `Filter`'s documentation.
type FilterIter[T any] struct {
	iter Iterator[T]
	fun  func(T) bool
}

// Filter instantiates a `FilterIter` that selectively yields only results that
// cause the provided function to return `true`.
func Filter[T any](iter Iterator[T], fun func(T) bool) *FilterIter[T] {
	return &FilterIter[T]{iter, fun}
}

// Next implements the Iterator interface for `Filter`.
func (iter *FilterIter[T]) Next() option.Option[T] {
	for {
		value, ok := iter.iter.Next().Value()
		if !ok {
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
	return &FilterIter[T]{iter, inverse}
}
