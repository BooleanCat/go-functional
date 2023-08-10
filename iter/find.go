package iter

import "github.com/BooleanCat/go-functional/option"

// Find searches for the first occurance of a value that satisfies the
// predicate and returns that value. If no value satisfies the predicate, it
// returns `None`.
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
