package it

import "iter"

// Repeat yields the same value indefinitely.
func Repeat[V any](value V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for {
			if !yield(value) {
				return
			}
		}
	}
}

// Repeat2 yields the same two values indefinitely.
func Repeat2[V, W any](value1 V, value2 W) iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {
		for {
			if !yield(value1, value2) {
				return
			}
		}
	}
}
