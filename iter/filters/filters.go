// This package contains functions intended for use with [iter.Filter].
package filters

import (
	"github.com/BooleanCat/go-functional/constraints"
	"github.com/BooleanCat/go-functional/iter"
)

var _ = iter.Filter[struct{}]

// IsZero returns true when the provided value is equal to its zero value.
func IsZero[T comparable](t T) bool {
	var u T
	return t == u
}

// IsEven returns true when the provided integer is even
func IsEven[T constraints.Integer](integer T) bool {
	return integer%2 == 0
}

// IsOdd returns true when the provided integer is even
func IsOdd[T constraints.Integer](integer T) bool {
	return integer%2 != 0
}

// GreaterThan returns a function that returns true when a value is greater
// than a threshold.
func GreaterThan[T constraints.Ordered](t T) func(T) bool {
	return func(s T) bool {
		return s > t
	}
}

// LessThan returns a function that returns true when a value is less than a
// threshold.
func LessThan[T constraints.Ordered](t T) func(T) bool {
	return func(s T) bool {
		return s < t
	}
}

// And returns a function that returns true when all provided functions return
// true.
//
// When no functions are provided the result is always true.
func And[T any](filters ...func(T) bool) func(T) bool {
	return func(t T) bool {
		for _, filter := range filters {
			if !filter(t) {
				return false
			}
		}

		return true
	}
}

// Or returns a function that returns true when one of the provided functions
// returns true.
//
// When no functions are provided the result is always true.
func Or[T any](filters ...func(T) bool) func(T) bool {
	return func(t T) bool {
		if len(filters) == 0 {
			return true
		}

		for _, filter := range filters {
			if filter(t) {
				return true
			}
		}

		return false
	}
}
