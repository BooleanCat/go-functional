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
func And[T constraints.Ordered](filters ...func(T) bool) func(T) bool {
	return func(t T) bool {
		for _, filter := range filters {
			if !filter(t) {
				return false
			}
		}

		return true
	}
}
