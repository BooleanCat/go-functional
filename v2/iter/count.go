package iter

import "iter"

// Count yields an infinite sequence of integers, starting from 0.
func Count() Iterator[int] {
	return Iterator[int](iter.Seq[int](func(yield func(int) bool) {
		for i := 0; ; i++ {
			if !yield(i) {
				return
			}
		}
	}))
}
