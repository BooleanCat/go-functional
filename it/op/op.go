package op

// Add returns the sum of `a` and `b`.
func Add[V ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~string | ~float32 | ~float64](a, b V) V {
	return a + b
}

// Ref returns a reference to a copy of the provided value.
//
// This may be useful when interacting with packages that use pointers as
// proxies for optional values.
func Ref[V any](v V) *V {
	return &v
}

// Deref returns the value pointed to by the provided pointer.
func Deref[V any](v *V) V {
	return *v
}

func Must[V any](v V, err error) V {
	if err != nil {
		panic(err)
	}

	return v
}
