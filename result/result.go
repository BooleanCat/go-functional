package result

import "fmt"

// Result represents failure or success. The [Ok] variant represents success
// and contains a value. The [Err] variant represent a failure and contains an
// error.
type Result[T any] struct {
	value T
	err   error
}

// Ok instantiates a [Result] with a value.
func Ok[T any](value T) Result[T] {
	return Result[T]{value, nil}
}

// Err instantiates a [Result] with an error.
func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

// String implements the [fmt.Stringer] interface.
func (r Result[T]) String() string {
	if r.err == nil {
		return fmt.Sprintf("Ok(%v)", r.value)
	}

	return fmt.Sprintf("Err(%s)", r.err.Error())
}

var _ fmt.Stringer = Result[struct{}]{}

// Unwrap returns the underlying value of an [Ok] variant, or panics if called
// on an [Err] variant.
func (r Result[T]) Unwrap() T {
	if r.err == nil {
		return r.value
	}

	panic("called `Result.Unwrap()` on an `Err` value")
}

// UnwrapOr returns the underlying value of an [Ok] variant, or the provided
// value on an [Err] variant.
func (r Result[T]) UnwrapOr(value T) T {
	if r.err == nil {
		return r.value
	}

	return value
}

// UnwrapOrElse returns the underlying value of an [Ok] variant, or the result
// of calling the provided function on an [Err] variant.
func (r Result[T]) UnwrapOrElse(f func() T) T {
	if r.err == nil {
		return r.value
	}

	return f()
}

// UnwrapOrZero returns the underlying value of an [Ok] variant, or the zero
// value of an [Err] variant.
func (r Result[T]) UnwrapOrZero() T {
	if r.err == nil {
		return r.value
	}

	var value T
	return value
}

// IsOk returns true if the [Result] is an [Ok] variant.
func (r Result[T]) IsOk() bool {
	return r.err == nil
}

// IsErr returns true if the [Result] is an [Err] variant.
func (r Result[T]) IsErr() bool {
	return r.err != nil
}

// Value returns the underlying value and nil for an [Ok] variant, or the zero
// value and an error for an [Err] variant.
func (r Result[T]) Value() (T, error) {
	return r.value, r.err
}

// UnwrapErr returns the underlying error of an [Err] variant, or panics if called
// on an [Ok] variant.
func (r Result[T]) UnwrapErr() error {
	if r.IsOk() {
		panic("called `Result.UnwrapErr()` on an `Ok` value")
	}

	return r.err
}

// Expect returns the underlying value of an [Ok] variant, or panics with the
// provided message if called on an [Err] variant.
func (r Result[T]) Expect(message string) T {
	if r.err == nil {
		return r.value
	}

	panic(message)
}
