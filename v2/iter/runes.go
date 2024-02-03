package iter

import "iter"

// Runes yields the runes of the provided string or rune slice.
func Runes[V ~string | []rune](runes V) Iterator[rune] {
	return Iterator[rune](iter.Seq[rune](func(yield func(rune) bool) {
		for _, r := range []rune(runes) {
			if !yield(r) {
				return
			}
		}
	}))
}
