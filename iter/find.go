package iter

import "github.com/BooleanCat/go-functional/option"

// Find the first occurrence of a value that satisfies the predicate and return
// that value. If no value satisfies the predicate, return `None`.
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
