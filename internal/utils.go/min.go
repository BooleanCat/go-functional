package utils

import "github.com/BooleanCat/go-functional/constraints"

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}

	return b
}
