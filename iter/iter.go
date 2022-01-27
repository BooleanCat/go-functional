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
