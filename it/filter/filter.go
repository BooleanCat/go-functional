// This package contains functions intended for use with [iter.Filter].
package filter

import "cmp"

// IsEven returns true when the provided integer is even.
func IsEven[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64](integer T) bool {
	return integer%2 == 0
}

// IsOdd returns true when the provided integer is odd.
func IsOdd[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64](integer T) bool {
	return integer%2 != 0
}

// IsEqual returns a function that returns true when the provided value is equal
// to some value.
func IsEqual[T comparable](value T) func(T) bool {
	return func(candidate T) bool {
		return candidate == value
	}
}

// NotEqual returns a function that returns true when the provided value is not
// equal to some value.
func NotEqual[T comparable](value T) func(T) bool {
	return func(candidate T) bool {
		return candidate != value
	}
}

// IsZero returns true when the provided value is the zero value for its type.
func IsZero[T comparable](value T) bool {
	var zero T
	return value == zero
}

// GreaterThan returns a function that returns true when the provided value is
// greater than a threshold.
func GreaterThan[T cmp.Ordered](threshold T) func(T) bool {
	return func(value T) bool {
		return value > threshold
	}
}

// LessThan returns a function that returns true when the provided value is less
// than a threshold.
func LessThan[T cmp.Ordered](threshold T) func(T) bool {
	return func(value T) bool {
		return value < threshold
	}
}

// Passthrough returns a function that returns true for any value.
func Passthrough[V any](value V) bool {
	return true
}

// Passthrough2 returns a function that returns true for any pair of values.
func Passthrough2[V any, W any](value1 V, value2 W) bool {
	return true
}

// Not returns a function that inverts the result of the provided function.
func Not[T any](fn func(T) bool) func(T) bool {
	return func(value T) bool {
		return !fn(value)
	}
}

// And returns a function that returns true when all provided functions return
// true.
func And[T any](filters ...func(T) bool) func(T) bool {
	return func(value T) bool {
		for _, filter := range filters {
			if !filter(value) {
				return false
			}
		}

		return true
	}
}

// Or returns a function that returns true when any of the provided functions
// return true.
func Or[T any](filters ...func(T) bool) func(T) bool {
	return func(value T) bool {
		for _, filter := range filters {
			if filter(value) {
				return true
			}
		}

		return false
	}
}
