package iter

import "iter"

// Count yields all non-negative integers in ascending order.
func Count() Iterator[int] {
	return Iterator[int](iter.Seq[int](func(yield func(int) bool) {
		for i := 0; ; i++ {
			if !yield(i) {
				return
			}
		}
	}))
}
