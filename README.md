# Functional Programming in Go [![Actions Status](https://github.com/BooleanCat/go-functional/workflows/test/badge.svg)](https://github.com/BooleanCat/go-functional/actions) [![Go Reference](https://pkg.go.dev/badge/github.com/BooleanCat/go-functional.svg)](https://pkg.go.dev/github.com/BooleanCat/go-functional)

A general purpose library offering functional helpers for Golang.

_*Note: this library required Go 1.18, which is currently in beta.*_

```go
// Find the first 5 prime numbers
primes := iter.Take[int[(iter.Filter[int](iter.Count(), isPrime), 5)
assert.SliceEqual(t, iter.Collect[int](primes), []int{2, 3, 5, 7, 11})
```

_[Read the docs.](https://pkg.go.dev/github.com/BooleanCat/go-functional)_

## Core concepts

This library introduces two core concepts, the `Iterator` and the `Option`.
Using these two concepts this library derives many pre-defined iterators for
use.

### The `Option`

The `Option` is simply a type that represents the presence or absence of a
value. Options behave somewhat like enumerations with two variants:
`Some(value)` and `None`.

### The `Iterator`

```go
type Iterator[T any] interface {
	Next() option.Option[T]
}
```

The `Iterator` is an interface that defines the behaviour of some type that
"yields" values. Each call to `Next()` on an `Iterator` will yield another
value as defined by that specific `Iterator`. For example, the iterator
`iter.Count()` yields 0, 1, 2, 3, etc. and onwards to infinity (or the integer
limit!).

Iterators will yield `Some(value)` for as long as they have values to yield.
Iterators that have exhausted all their values will then always yield `None`.

This library defines many iterators and a few are demonstrated below. For the
entire set simply visit the
[package documentation](https://pkg.go.dev/github.com/BooleanCat/go-functional/iter).

### Iterator showcase

Here are a few trivial example of what's possible using the iterators in this
library.

```go
// All even natural numbers (2, 4, 6, 8...)
isEven := func(n int) bool { return n%2 == 0 }
evens := iter.Filter[int](iter.Drop[int](iter.Count(), 1), isEven)
```

```go
// All non-empty lines from a file
lines := iter.Exclude[string](iter.LinesString(file), filters.IsZero[string])
```

```go
// String representations of numbers
numbers := iter.Map[int, string](iter.Count(), strconv.Itoa)
```
