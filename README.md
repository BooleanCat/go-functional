# Functional Programming in Go [![Actions Status](https://github.com/BooleanCat/go-functional/workflows/test/badge.svg)](https://github.com/BooleanCat/go-functional/actions) [![Go Reference](https://pkg.go.dev/badge/github.com/BooleanCat/go-functional.svg)](https://pkg.go.dev/github.com/BooleanCat/go-functional)

A general purpose library offering functional tooling for Golang.

## What do you mean?

```go
isEven := func(a int) bool { return a%2 == 0 }
evens := iter.Take[int[(iter.Filter[int](iter.Count(), isEven), 3)
assert.SliceEqual(t, iter.Collect[int](evens), []int{0, 2, 4})
```

## Not intended for production use (yet?)

This library was written to demonstrate what is possible using generics in the
realm of functional programming in Go.

I've yet to formalise many things about the rules and types offered by this
library.

## The Iterator

The core idea of this library is the Iterator. It's an interface that describes
type associated with a single receiver: `Next() option.Option[T]`. Simply put,
calling `Next()` will yield the next value from an iterator. This library
implements a fair number of common Iterators such as `Map`, `Filter`, `Count`
and more described in the documentation below.

To use an example to demonstrate how Iterators can be powerfully combined,
here's how you'd use a Filter and an infinite Counter to generate every
possible prime number:

```go
primes := iter.Filter[int](iter.Count(), isPrime)
```

## Options and Results

Two core types were implemented for this package: `Option[T]` and `Result[T]`.

Options represent the presence or absence of a value and come in two flavours:
`Some(value)` and `None`. Iterators yield `Some` variants until they are
exhuasted and then they yield `None` variants. For example, an iterator
yielding the first three naturals numbers would yield `Some(1)`, `Some(2)`,
`Some(3)` and `None` on subsequent calls to `Next()`.

Results represent success or failure and are likewise represented by two
variants: `Ok(value)` and `Err(error)`.

## To Do

- Write a better README
- Learn how to use go doc and generate docs
- Write some Example specs for auto-docs
- Formalise iterator rules for things like repeated calls and chain-call rules