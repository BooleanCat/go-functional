package iter

import "github.com/BooleanCat/go-functional/option"

// Iterator declares that each Iterator must implement a Next method.
// Successive calls to the next method shall return the next item in the
// Iterator, wrapped in an `Option.Some` variant.
//
// Exhausted Iterators shall return a `None` variant of `Option` on every
// subsequent call.
type Iterator[T any] interface {
	Next() option.Option[T]
}

// Collect consumes an Iterator and returns all remaining items within a slice.
// It does not protect against infinite Iterators.
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

// Fold consumes an Iterator and returns the final result of applying the accumulator function to each element.
// The accumulator function accepts two arguments - an accumulator and an element and returns a new  accumulator.
// The initial value is the accumulator for the first call.
// Fold does not protect against infinite Iterators.
func Fold[T any, U any](iter Iterator[T], initial U, biop func(U, T) U) U {
	for {
		if value, ok := iter.Next().Value(); ok {
			initial = biop(initial, value)
		} else {
			return initial
		}
	}
}

// ToChannel consumes an iterator and returns a channel that will receive all
// values from the provided iterator. The channel is closed once the iterator
// is exhausted.
func ToChannel[T any](iter Iterator[T]) chan T {
	ch := make(chan T)

	go func() {
		for {
			next := iter.Next()
			if next.IsNone() {
				close(ch)
				return
			}

			ch <- next.Unwrap()
		}
	}()

	return ch
}
