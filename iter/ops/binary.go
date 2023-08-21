// This package contains functions intended for use with [iter.Map].
package ops

import (
	"github.com/BooleanCat/go-functional/constraints"
	"github.com/BooleanCat/go-functional/iter"
)

var _ = iter.Map[struct{}, struct{}]

// Add performs the `+` operation for the two inputs, returning the result.
func Add[T constraints.Integer | constraints.Float | ~string](a, b T) T {
	return a + b
}

// Multiply performs the `*` operation for the two inputs, returning the result.
func Multiply[T constraints.Integer | constraints.Float](a, b T) T {
	return a * b
}

// BitwiseAnd performs the `&` operation for the two inputs, returning the result.
func BitwiseAnd[T constraints.Integer](a, b T) T {
	return a & b
}

// BitwiseOr performs the `|` operation for the two inputs, returning the result.
func BitwiseOr[T constraints.Integer](a, b T) T {
	return a | b
}

// BitwiseXor performs the `^` operation for the two inputs, returning the result.
func BitwiseXor[T constraints.Integer](a, b T) T {
	return a ^ b
}
