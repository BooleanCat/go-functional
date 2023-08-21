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

// Collect consumes an [Iterator] and returns all remaining items within a
// slice. It does not protect against infinite Iterators.
func Collect[T any](iter Iterator[T]) []T {
	items := make([]T, 0)

	for {
		if value, ok := iter.Next().Value(); ok {
			items = append(items, value)
		} else {
			return items
		}
	}
}

// Fold consumes an [Iterator] and returns the final result of applying the
// accumulator function to each element. The accumulator function accepts two
// arguments - an accumulator and an initial value and returns a new value for
// the next accumulation. Fold does not protect against infinite Iterators.
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
