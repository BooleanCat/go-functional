package iter

import "github.com/BooleanCat/go-functional/result"

// CollectResults consumes an [Iterator] with a yield type of [result.Result].
// It returns all items as a slice. It does not protect against infinite
// iterators.
//
// When a [result.Err] is encountered, collection stops and the [result.Err] is
// returned immediately.
func CollectResults[T any](iter Iterator[result.Result[T]]) result.Result[[]T] {
	values := make([]T, 0)

	for value, ok := iter.Next().Value(); ok; value, ok = iter.Next().Value() {
		if value.IsErr() {
			return result.Err[[]T](value.UnwrapErr())
		}

		values = append(values, value.Unwrap())
	}

	return result.Ok(values)
}
