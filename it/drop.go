package it

import "iter"

// Drop yields all values from a delegate iterator except the first `count`
// values.
func Drop[V any](delegate func(func(V) bool), count uint) iter.Seq[V] {
	return func(yield func(V) bool) {
		for value := range delegate {
			if count == 0 {
				if !yield(value) {
					return
				}
			} else {
				count--
			}
		}
	}
}

// Drop2 yields all pairs of values from a delegate iterator except the first
// `count` pairs.
func Drop2[V, W any](delegate func(func(V, W) bool), count uint) iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {
		for v, w := range delegate {
			if count == 0 {
				if !yield(v, w) {
					return
				}
			} else {
				count--
			}
		}
	}
}
