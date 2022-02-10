package filters

import "constraints"

func IsZero[T comparable](t T) bool {
	var u T
	return t == u
}

func GreaterThan[T constraints.Ordered](t T) func(T) bool {
	return func(s T) bool {
		return s > t
	}
}

func LessThan[T constraints.Ordered](t T) func(T) bool {
	return func(s T) bool {
		return s < t
	}
}

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
