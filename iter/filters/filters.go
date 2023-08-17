package filters

import "github.com/BooleanCat/go-functional/constraints"

// IsZero is a filter intended for use with iter.Filter that returns true when
// the provided value is equal to its zero value.
func IsZero[T comparable](t T) bool {
	var u T
	return t == u
}

// GreaterThan creates a filter for use with iter.Filter that returns true when
// a value is greater than the provided threshold.
func GreaterThan[T constraints.Ordered](t T) func(T) bool {
	return func(s T) bool {
		return s > t
	}
}

// LessThan creates a filter for use with iter.Filter that returns true when a
// value is less than the provided threshold.
func LessThan[T constraints.Ordered](t T) func(T) bool {
	return func(s T) bool {
		return s < t
	}
}

// And aggregates multiple filters for use with iter.Filter (and iter.Exclude)
// to create a filter that returns true when all provided filters return true.
//
// If no filters are provided the result is always true.
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

// And aggregates multiple filters for use with iter.Filter (and iter.Exclude)
// to create a filter that returns true when one of the provided filters return
// true.
//
// If no filters are provided the result is always true.
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
