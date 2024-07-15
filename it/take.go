package it

import "iter"

// Take yields the first `limit` values from a delegate iterator.
func Take[V any](delegate func(func(V) bool), limit int) iter.Seq[V] {
	return func(yield func(V) bool) {
		if limit <= 0 {
			return
		}

		for value := range delegate {
			if !yield(value) {
				return
			}

			limit--
			if limit <= 0 {
				return
			}
		}
	}
}

// Take2 yields the first `limit` pairs of values from a delegate iterator.
func Take2[V, W any](delegate func(func(V, W) bool), limit int) iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {
		if limit <= 0 {
			return
		}

		for v, w := range delegate {
			if !yield(v, w) {
				return
			}

			limit--
			if limit <= 0 {
				return
			}
		}
	}
}
