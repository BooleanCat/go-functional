//go:build go1.22 && goexperiment.rangefunc

package iter

import "iter"

// Count yields all non-negative integers in ascending order.
func Count[V ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64]() iter.Seq[V] {
	return func(yield func(V) bool) {
		var i V

		for {
			if !yield(i) {
				return
			}
			i++
		}
	}
}
