# Functional Programming in Go

[![GitHub release (with filter)](https://img.shields.io/github/v/release/BooleanCat/go-functional?sort=semver&logo=Go&color=%23007D9C&include_prereleases)](https://github.com/BooleanCat/go-functional/releases) [![Actions Status](https://github.com/BooleanCat/go-functional/workflows/test/badge.svg)](https://github.com/BooleanCat/go-functional/actions) [![Go Reference](https://pkg.go.dev/badge/github.com/BooleanCat/go-functional/v2.svg)](https://pkg.go.dev/github.com/BooleanCat/go-functional/v2) [![Go Report Card](https://goreportcard.com/badge/github.com/BooleanCat/go-functional/v2)](https://goreportcard.com/report/github.com/BooleanCat/go-functional/v2) [![codecov](https://codecov.io/gh/BooleanCat/go-functional/branch/main/graph/badge.svg?token=N2E43RSR14)](https://codecov.io/gh/BooleanCat/go-functional)

A library of iterators for use with `iter.Seq`. Requires Go 1.23+.

```go
// The first 5 natural numbers
numbers := slices.Collect(
	it.Take(it.NaturalNumbers[int](), 5),
)

// All even numbers
evens := it.Filter(it.NaturalNumbers[int](), filter.IsEven)

// String representations of integers
numbers := it.Map(it.NaturalNumbers[int](), strconv.Itoa)
```

_[Reference documentation](https://pkg.go.dev/github.com/BooleanCat/go-functional/v2)_

## Installation

```terminal
go get github.com/BooleanCat/go-functional/v2@latest
```

## Iterators

This library contains two kinds of iterators in the `it` and `itx` packages. In
most cases you'll find the same iterators in each package, the difference
between them being that the iterators in the `itx` package can be "dot-chained"
(e.g. `iter.Filter(...).Take(3).Collect()`) and those in `it` cannot.

Iterators and functions that work with them come in several varieties and it's
important to be aware of the distinction between them.

- Most iterators are `🔵 finite`, but some are `🔴 infinite` (never terminate)
  and care should be taken when consuming `🔴 infinite` iterators to avoid
  deadlocking.
- Iterators are either `🟣 primary` or `🟡 secondary`. `🟣 primary` iterators
  create new iterators and do not consume other iterators (e.g.
  `it.NaturalNumbers`). `🟡 secondary` iterators consume other iterators (e.g.
  `it.Filter`).
- `🟢 consumer`s are not iterators but either partially or completely consume an
  iterator (e.g. `it.Find`).

Iterators documented below will be tagged with the above information.

### `it.NaturalNumbers` (`🟣 primary`, `🔴 infinite`)

NaturalNumbers yields all non-negative integers in ascending order.

```go
for i := range it.NaturalNumbers[int]() {
	if i >= 3 {
		break
	}

	fmt.Println(i)
}
```

> [!WARNING]
> There is no protection against overflowing whatever integer type is used for
> this iterator.

## Iterator Chaining

The iterators in this package were designed to be used with the native
`iter.Seq` from Go's standard library. In order to facilitate complex
sequences of iterators, the
[`itx`](https://github.com/BooleanCat/go-functional/tree/main/it/itx) package
provides `Iterator` and `Iterator2` as wrappers around `iter.Seq` and
`iter.Seq2` that allow for chaining operations.

Let's take a look at an example:

```go
// The first 10 odd integers
itx.NaturalNumbers[int]().Filter(filter.IsOdd).Take(10).Collect()
```

Most iterators support chaining. A notable exception is `it.Map` which cannot
support chaining due to limitations on Go's type system.
