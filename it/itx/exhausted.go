package itx

import "github.com/BooleanCat/go-functional/v2/it"

// Exhausted is an iterator that yields no values.
func Exhausted[V any]() Iterator[V] {
	return Iterator[V](it.Exhausted[V]())
}

// Exhausted2 is an iterator that yields no values.
func Exhausted2[V, W any]() Iterator2[V, W] {
	return Iterator2[V, W](it.Exhausted2[V, W]())
}
