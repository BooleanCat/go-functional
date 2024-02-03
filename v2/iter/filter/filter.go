// This package contains functions intended for use with [iter.Filter].
package filter

import (
	"cmp"

	"github.com/BooleanCat/go-functional/v2/iter"
)

var _ = iter.Filter[struct{}]

// IsZero returns true when the provided value is equal to its zero value.
func IsZero[V comparable](value V) bool {
	var zero V
	return value == zero
}

// IsEven returns true when the provided integer is even.
func IsEven[V ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr](integer V) bool {
	return integer%2 == 0
}

// IsOdd returns true when the provided integer is odd.
func IsOdd[V ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr](integer V) bool {
	return integer%2 != 0
}

// GreaterThan returns a function that returns true when a value is greater
// than a threshold.
func GreaterThan[V cmp.Ordered](value V) func(V) bool {
	return func(s V) bool {
		return s > value
	}
}

// GreaterThanEqual returns a function that returns true when a value is
// greater than or equal to a threshold.
func GreaterThanEqual[T cmp.Ordered](t T) func(T) bool {
	return func(s T) bool {
		return s >= t
	}
}

// LessThan returns a function that returns true when a value is less than a
// threshold.
func LessThan[T cmp.Ordered](t T) func(T) bool {
	return func(s T) bool {
		return s < t
	}
}

// LessThanEqual returns a function that returns true when a value is less than
// or equal to a threshold.
func LessThanEqual[T cmp.Ordered](t T) func(T) bool {
	return func(s T) bool {
		return s <= t
	}
}
