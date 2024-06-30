package itx

import (
	"github.com/BooleanCat/go-functional/v2/it"
)

// Count yields all non-negative integers in ascending order.
func Count[V ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64]() Iterator[V] {
	return Iterator[V](it.Count[V]())
}
