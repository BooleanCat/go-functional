package ops

import (
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

// Passthrough may be used an an operation for iter.Map. It returns the
// provided value without modification.
func Passthrough[T any](t T) T {
	return t
}
