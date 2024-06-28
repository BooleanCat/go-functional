//go:build go1.23

package slices

import (
	"iter"
	"slices"
)

// Values returns an iterator over the slice elements, starting with s[0].
func Values[Slice ~[]E, E any](slice Slice) iter.Seq[E] {
	return slices.Values(slice)
}

// Collect collects values from seq into a new slice and returns it.
func Collect[E any](seq iter.Seq[E]) []E {
	return slices.Collect(seq)
}
