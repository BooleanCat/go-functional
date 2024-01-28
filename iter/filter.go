package iter

import (
	"fmt"
	"reflect"

	"github.com/BooleanCat/go-functional/option"
)

// FilterIter iterator, see [Filter].
type FilterIter[T any] struct {
	BaseIter[T]
	iter      Iterator[T]
	fun       func(T) bool
	exhausted bool
}

// Filter instantiates a [*FilterIter] that selectively yields only results
// that cause the provided function to return `true`.
func Filter[T any](iter Iterator[T], fun func(T) bool) *FilterIter[T] {
	iterator := &FilterIter[T]{iter: iter, fun: fun}
	iterator.BaseIter = BaseIter[T]{iterator}
	return iterator
}

// Next implements the [Iterator] interface.
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

// String implements the [fmt.Stringer] interface
func (iter FilterIter[T]) String() string {
	var instanceOfT T
	return fmt.Sprintf("Iterator<Filter, type=%s>", reflect.TypeOf(instanceOfT))
}

var (
	_ fmt.Stringer       = new(FilterIter[struct{}])
	_ Iterator[struct{}] = new(FilterIter[struct{}])
)

// Exclude instantiates a [*FilterIter] that selectively yields only results
// that cause the provided function to return `false`.
func Exclude[T any](iter Iterator[T], fun func(T) bool) *FilterIter[T] {
	inverse := func(t T) bool { return !fun(t) }
	iterator := &FilterIter[T]{iter: iter, fun: inverse}
	iterator.BaseIter = BaseIter[T]{iterator}
	return iterator
}

// FilterMapIter iterator, see [FilterMap].
type FilterMapIter[T any, U any] struct {
	BaseIter[U]
	iter      Iterator[T]
	fn        func(T) option.Option[U]
	exhausted bool
}

// Next implements the [Iterator] interface.
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

// String implements the [fmt.Stringer] interface
func (iter FilterMapIter[T, U]) String() string {
	var instanceOfU U
	return fmt.Sprintf("Iterator<FilterMap, type=%s>", reflect.TypeOf(instanceOfU))
}

var _ Iterator[struct{}] = new(FilterMapIter[struct{}, struct{}])

// FilterMap instantiates a [*FilterMapIter] that selectively yields only
// results that cause the provided function to return `true` with a map
// operation performed on them.
//
// Unlike other iterators (e.g. [Filter]), it is not possible to call FilterMap
// as a method on iterators defined in this package. This is due to a
// limitation of Go's type system; new type parameters cannot be defined on
// methods.
func FilterMap[T any, U any](iter Iterator[T], fun func(T) option.Option[U]) *FilterMapIter[T, U] {
	iterator := &FilterMapIter[T, U]{iter: iter, fn: fun}
	iterator.BaseIter = BaseIter[U]{iterator}
	return iterator
}
