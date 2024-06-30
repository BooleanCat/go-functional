package op

// Add returns the sum of `a` and `b`.
func Add[V ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~string | ~float32 | ~float64](a, b V) V {
	return a + b
}
