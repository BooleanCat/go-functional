// Package slices provides early implementations of slice-related functions
// available in Go 1.23+.
package slices

import "iter"

// Values returns an iterator over the slice elements, starting with s[0].
func Values[Slice ~[]E, E any](slice Slice) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, value := range slice {
			if !yield(value) {
				return
			}
		}
	}
}

// Collect collects values from seq into a new slice and returns it.
func Collect[E any](seq iter.Seq[E]) []E {
	slice := make([]E, 0)

	for item := range seq {
		slice = append(slice, item)
	}

	return slice
}
