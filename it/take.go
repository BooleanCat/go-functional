package it

import "iter"

// Take yields the first `limit` values from a delegate iterator.
func Take[V any](delegate func(func(V) bool), limit uint) iter.Seq[V] {
	return func(yield func(V) bool) {
		for value := range delegate {
			if limit > 0 {
				limit--
				if !yield(value) {
					return
				}
			} else {
				return
			}
		}
	}
}

// Take2 yields the first `limit` pairs of values from a delegate iterator.
func Take2[V, W any](delegate func(func(V, W) bool), limit uint) iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {
		for v, w := range delegate {
			if limit > 0 {
				limit--
				if !yield(v, w) {
					return
				}
			} else {
				return
			}
		}
	}
}

// TakeWhile yields values from a delegate iterator until the predicate returns
// false.
func TakeWhile[V any](delegate func(func(V) bool), predicate func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for value := range delegate {
			if predicate(value) {
				if !yield(value) {
					return
				}
			} else {
				return
			}
		}
	}
}

// TakeWhile2 yields pairs of values from a delegate iterator until the
// predicate returns false.
func TakeWhile2[V, W any](delegate func(func(V, W) bool), predicate func(V, W) bool) iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {
		for v, w := range delegate {
			if predicate(v, w) {
				if !yield(v, w) {
					return
				}
			} else {
				return
			}
		}
	}
}
