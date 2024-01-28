package iter

import "github.com/BooleanCat/go-functional/option"

// Iterator declares that each Iterator must implement a Next method.
// Successive calls to the next method shall return the next item in the
// Iterator, wrapped in an [option.Some] variant.
//
// Exhausted Iterators shall return a [option.None] variant on every subsequent
// call.
type Iterator[T any] interface {
	Next() option.Option[T]
}

// Collect consumes an [Iterator] and returns all items within a slice. It does
// not protect against infinite Iterators.
func Collect[T any](iter Iterator[T]) []T {
	values := make([]T, 0)

	for value, ok := iter.Next().Value(); ok; value, ok = iter.Next().Value() {
		values = append(values, value)
	}

	return values
}

// Fold consumes an [Iterator] and returns the final result of applying the
// accumulator function to each element. The accumulator function accepts two
// arguments - an accumulator and an initial value and returns a new value for
// the next accumulation. Fold does not protect against infinite Iterators.
//
// Unlike other [Iterator] consumers (e.g. [Collect]), it is not possible to
// call Fold as a method on iterators defined in this package. This is due to a
// limitation of Go's type system; new type parameters cannot be defined on
// methods.
func Fold[T any, U any](iter Iterator[T], initial U, biop func(U, T) U) U {
	for {
		if value, ok := iter.Next().Value(); ok {
			initial = biop(initial, value)
		} else {
			return initial
		}
	}
}

// ToChannel consumes an [Iterator] and returns a channel that will receive all
// values from the provided [Iterator]. The channel is closed once the
// [Iterator] is exhausted.
func ToChannel[T any](iter Iterator[T]) chan T {
	ch := make(chan T)

	go func() {
		for {
			value, ok := iter.Next().Value()
			if !ok {
				close(ch)
				break
			}

			ch <- value
		}
	}()

	return ch
}

// ForEach consumes an [Iterator] and executes callback function on each item.
func ForEach[T any](iter Iterator[T], callback func(T)) {
	for {
		if value, ok := iter.Next().Value(); ok {
			callback(value)
		} else {
			break
		}
	}
}

// Find the first occurrence of a value that satisfies the predicate and return
// that value. If no value satisfies the predicate, return [option.None].
func Find[T any](iter Iterator[T], predicate func(v T) bool) option.Option[T] {
	for {
		if value, ok := iter.Next().Value(); ok {
			if predicate(value) {
				return option.Some(value)
			}
		} else {
			return option.None[T]()
		}
	}
}

// BaseIter is intended to be embedded in other iterators to expose method
// chaining.
type BaseIter[T any] struct {
	Iterator[T]
}

// Collect is a convenience method for [Collect], providing this iterator as
// an argument.
func (iter *BaseIter[T]) Collect() []T {
	return Collect[T](iter.Iterator)
}

// ForEach is a convenience method for [ForEach], providing this iterator as an
// argument.
func (iter *BaseIter[T]) ForEach(callback func(T)) {
	ForEach[T](iter, callback)
}

// Find is a convenience method for [Find], providing this iterator as an
// argument.
func (iter *BaseIter[T]) Find(predicate func(T) bool) option.Option[T] {
	return Find[T](iter, predicate)
}

// Drop is a convenience method for [Drop], providing this iterator as an
// argument.
func (iter *BaseIter[T]) Drop(n uint) *DropIter[T] {
	return Drop[T](iter, n)
}

// Take is a convenience method for [Take], providing this iterator as an
// argument.
func (iter *BaseIter[T]) Take(n uint) *TakeIter[T] {
	return Take[T](iter, n)
}

// Filter is a convenience method for [Filter], providing this iterator
// as an argument.
func (iter *BaseIter[T]) Filter(fun func(T) bool) *FilterIter[T] {
	return Filter[T](iter, fun)
}

// Chain is a convenience method for [Chain], providing this iterator as an
// argument.
func (iter *BaseIter[T]) Chain(iterators ...Iterator[T]) *ChainIter[T] {
	return Chain[T](append([]Iterator[T]{iter}, iterators...)...)
}

// ToChannel is a convenience method for [ToChannel], providing this iterator
// as an argument.
func (iter *BaseIter[T]) ToChannel() chan T {
	return ToChannel[T](iter)
}

// Enumerate is a convenience method for [Enumerate], providing this iterator
// as an argument.
func (iter *BaseIter[T]) Enumerate() *EnumerateIter[T] {
	return Enumerate[T](iter)
}

// Transform is a convenience method for [Transform], providing this iterator
// as an argument.
func (iter *BaseIter[T]) Transform(op func(T) T) *MapIter[T, T] {
	return Transform[T](iter, op)
}

// Exclude is a convenience method for [Exclude], providing this iterator
// as an argument.
func (iter *BaseIter[T]) Exclude(fun func(T) bool) *FilterIter[T] {
	return Exclude[T](iter, fun)
}

// FilterMap is a convenience method for [FilterMap], providing this iterator
// as an argument.
func (iter *BaseIter[T]) FilterMap(fun func(T) option.Option[T]) *FilterMapIter[T, T] {
	return FilterMap[T, T](iter, fun)
}
