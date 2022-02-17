package ops

import (
	"github.com/BooleanCat/go-functional/constraints"
	"github.com/BooleanCat/go-functional/option"
	"github.com/BooleanCat/go-functional/result"
)

// UnwrapOption may be used as an operation for iter.Map in order to unwrap
// all options in an iterator.
func UnwrapOption[T any](o option.Option[T]) T {
	return o.Unwrap()
}

// UnwrapResult may be used as an operation for iter.Map in order to unwrap
// all results in an iterator.
func UnwrapResult[T any](r result.Result[T]) T {
	return r.Unwrap()
}

// Add performs the `+` operation for the two inputs, returning the result.
func Add[T constraints.Integer | constraints.Float | ~string](a, b T) T {
	return a + b
}
