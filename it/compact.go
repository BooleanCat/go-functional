package it

import "iter"

// Compact yields all values from a delegate iterator that are not zero values.
func Compact[V comparable](delegate func(func(V) bool)) iter.Seq[V] {
	return func(yield func(V) bool) {
		for value := range delegate {
			var zero V
			if zero == value {
				continue
			}

			if !yield(value) {
				return
			}
		}
	}
}
