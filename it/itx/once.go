package itx

import "github.com/BooleanCat/go-functional/v2/it"

// Once yields the provided value once.
func Once[V any](value V) Iterator[V] {
	return Iterator[V](it.Once(value))
}

// Once2 yields the provided value pair once.
func Once2[V, W any](v V, w W) Iterator2[V, W] {
	return Iterator2[V, W](it.Once2(v, w))
}
