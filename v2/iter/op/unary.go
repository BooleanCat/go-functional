package op

// Identity returns the provided value.
func Identity[V any](value V) V {
	return value
}
