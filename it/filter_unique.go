package it

import "iter"

// FilterUnique yields all the unique values from an iterator.
//
// Note: All unique values seen from an iterator are stored in memory.
func FilterUnique[V comparable](iterator func(func(V) bool)) iter.Seq[V] {
	return func(yield func(V) bool) {
		seen := make(map[V]struct{})

		for value := range iterator {
			if _, ok := seen[value]; ok {
				continue
			}

			seen[value] = struct{}{}
			if !yield(value) {
				return
			}
		}
	}
}
