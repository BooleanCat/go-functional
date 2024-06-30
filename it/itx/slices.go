package itx

import "slices"

// SlicesValues is a wrapper around slices.Values returning an [Iterator]
// rather than an [iter.Seq].
func SlicesValues[Slice ~[]E, E any](s Slice) Iterator[E] {
	return Iterator[E](slices.Values(s))
}
