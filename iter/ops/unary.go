package ops

import (
	"github.com/BooleanCat/go-functional/option"
	"github.com/BooleanCat/go-functional/result"
)

// UnwrapOption calls Unwrap on an [option.Option].
func UnwrapOption[T any](o option.Option[T]) T {
	return o.Unwrap()
}

// UnwrapResult calls Unwrap on a [result.Result].
func UnwrapResult[T any](r result.Result[T]) T {
	return r.Unwrap()
}

// Passthrough returns the provided value.
func Passthrough[T any](t T) T {
	return t
}
