package it

import "iter"

// Filter yields values from an iterator that satisfy a predicate.
func Filter[V any](delegate func(func(V) bool), predicate func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for value := range delegate {
			if predicate(value) {
				if !yield(value) {
					return
				}
			}
		}
	}
}

// Exclude yields values from an iterator that do not satisfy a predicate.
func Exclude[V any](delegate func(func(V) bool), predicate func(V) bool) iter.Seq[V] {
	return Filter(delegate, not(predicate))
}

// Filter2 yields values from an iterator that satisfy a predicate.
func Filter2[V, W any](delegate func(func(V, W) bool), predicate func(V, W) bool) iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {
		for k, v := range delegate {
			if predicate(k, v) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// Exclude2 yields values from an iterator that do not satisfy a predicate.
func Exclude2[V, W any](delegate func(func(V, W) bool), predicate func(V, W) bool) iter.Seq2[V, W] {
	return Filter2(delegate, not2(predicate))
}

// FilterError yields values from an iterator that satisfy a predicate where
// the predicate can return an error.
func FilterError[V any](delegate func(func(V) bool), predicate func(V) (bool, error)) iter.Seq2[V, error] {
	return func(yield func(V, error) bool) {
		for value := range delegate {
			if ok, err := predicate(value); err != nil {
				var zero V
				if !yield(zero, err) {
					return
				}
			} else if ok {
				if !yield(value, nil) {
					return
				}
			}
		}
	}
}

// ExcludeError yields values from an iterator that do not satisfy a predicate
// where the predicate can return an error.
func ExcludeError[V any](delegate func(func(V) bool), predicate func(V) (bool, error)) iter.Seq2[V, error] {
	return FilterError(delegate, notError(predicate))
}

func not[V any](predicate func(V) bool) func(V) bool {
	return func(value V) bool {
		return !predicate(value)
	}
}

func not2[V, W any](predicate func(V, W) bool) func(V, W) bool {
	return func(k V, v W) bool {
		return !predicate(k, v)
	}
}

func notError[V any](predicate func(V) (bool, error)) func(V) (bool, error) {
	return func(value V) (bool, error) {
		ok, err := predicate(value)
		return !ok, err
	}
}
