package itx

import "github.com/BooleanCat/go-functional/v2/it"

// Repeat yields the same value indefinitely.
func Repeat[V any](value V) Iterator[V] {
	return Iterator[V](it.Repeat[V](value))
}

// Repeat2 yields the same two values indefinitely.
func Repeat2[V, W any](value1 V, value2 W) Iterator2[V, W] {
	return Iterator2[V, W](it.Repeat2(value1, value2))
}
