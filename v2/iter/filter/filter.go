// This package contains functions intended for use with [iter.Filter].
package filter

import (
	"github.com/BooleanCat/go-functional/constraints"
	"github.com/BooleanCat/go-functional/v2/iter"
)

var _ = iter.Filter[struct{}]

// IsZero returns true when the provided value is equal to its zero value.
func IsZero[V comparable](value V) bool {
	var zero V
	return value == zero
}

// IsEven returns true when the provided integer is even.
func IsEven[T constraints.Integer](integer T) bool {
	return integer%2 == 0
}

// IsOdd returns true when the provided integer is odd.
func IsOdd[T constraints.Integer](integer T) bool {
	return integer%2 != 0
}
