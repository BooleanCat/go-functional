package it

import "iter"

// Drop yields all values from a delegate iterator except the first `count`
// values.
func Drop[V any](delegate func(func(V) bool), count int) iter.Seq[V] {
	return func(yield func(V) bool) {
		for value := range delegate {
			if count > 0 {
				count--
				continue
			}

			if !yield(value) {
				return
			}
		}
	}
}

// Drop2 yields all pairs of values from a delegate iterator except the first
// `count` pairs.
func Drop2[V, W any](delegate func(func(V, W) bool), count int) iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {
		for v, w := range delegate {
			if count > 0 {
				count--
				continue
			}

			if !yield(v, w) {
				return
			}
		}
	}
}
