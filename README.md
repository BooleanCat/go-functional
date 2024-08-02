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

## Overview

Most functions offered by this package are either consumers or iterators.

[Consumers](#consumers) will iterate over an iterator and completely or partially drain them
of values and (in most cases) collect the values into a data type.

[Iterators](#iterators) are functions that yield new values and can be ranged over. See Go's
documentation for iterators for more details.

<h2 id="consumers">Consumers</h2>

The standard libary provides functions to collect iterators in the `slices` and
`maps` packages that should satisfy most cases where collection is needed.

This package provides additional collection methods and makes existing
consumers from the standard library chainable.

> [!WARNING]
> Attempting to collect infinite iterators will cause an infinite loop and
> likely deadlock. Consider bounding infinite iterators before collect (for
> example using `Take`).

### Collect

In most cases `slices.Collect` from the standard library may be used to collect
items from an iterator into a slice. There are several other variants of
collect available for use for different use cases.

```go
// Chainable
numbers := itx.NaturalNumbers[int]().Take(5).Collect()

// Collect an `iter.Seq2[V, W] into two slices
keys, values := it.Collect2(maps.All(map[string]int{"one": 1, "two": 2}))

// As above, but chainable
keys, values := itx.FromMap(map[string]int{"one": 1, "two": 2}).Collect()
```

### TryCollect

Dealing with iterators that return `T, error` can involve the boilerplate of
checking that the returned slice of errors only contains `nil`. `TryCollect`
solves this by collecting all values into a slice and returning a single error:
the first one encountered.

```go
text := strings.NewReader("one\ntwo\nthree\n")

if lines, err := it.TryCollect(it.LinesString(text)); err != nil {
	fmt.Println(lines)
}
```

> [!NOTE]
> If an error is encountered, collection stops. This means the iterator being
> collected may not be fully drained.

> [!NOTE]
> Although the `itx` package also contains `TryCollect`, it is not chainable
> due to limitations with Go's type system.

### ForEach

ForEach consumes an iterator and applies a function to each value yielded.

```go
it.ForEach(slices.Values([]int{1, 2, 3}), func(number int) {
	fmt.Println(number)
})

// Chainable
itx.FromSlice([]int{1, 2, 3}).ForEach(func(number int) {
	fmt.Println(number)
})

// For each member of an iter.Seq2
it.ForEach2(slices.All([]int{1, 2, 3}), func(index int, number int) {
	fmt.Println(index, number)
})

// As above, but chainable
itx.FromSlice([]int{1, 2, 3}).Enumerate().ForEach(func(index int, number int) {
	fmt.Println(index, number)
})
```

<h2 id="iterators">Iterators</h2>

This library contains two kinds of iterators in the `it` and `itx` packages. In
most cases you'll find the same iterators in each package, the difference
between them being that the iterators in the `itx` package can be "dot-chained"
(e.g. `iter.Filter(...).Take(3).Collect()`) and those in `it` cannot.

Iterators within the `it` package are of the type `iter.Seq[V]` or
`iter.Seq2[V, W]` (from the standard library). Iterators within the `itx`
package are of the type `itx.Iterator[V]` or `itx.Iterator2[V, W]`.

Iterators come in several varieties and it's important to be aware of the
distinction between them.

- Most iterators are `🔵 finite`, but some are `🔴 infinite` (never terminate)
  and care should be taken when consuming `🔴 infinite` iterators to avoid
  deadlocking.
- Iterators are either `🟣 primary` or `🟡 secondary`. `🟣 primary` iterators
  create new iterators and do not consume other iterators (e.g.
  `it.NaturalNumbers`). `🟡 secondary` iterators consume other iterators (e.g.
  `it.Filter`).

Iterators documented below will be tagged with the above information.

### NaturalNumbers <sup>`🟣 primary` `🔴 infinite`</sup>

NaturalNumbers yields all non-negative integers in ascending order.

```go
for i := range it.Take(it.NaturalNumbers[int](), 3) {
	fmt.Println(i)
}

// Chainable
for i := range itx.NaturalNumbers[int]().Take(3) {
	fmt.Println(i)
}
```

> [!WARNING]
> There is no protection against overflowing whatever integer type is used for
> this iterator.

### Integers <sup>`🟣 primary` `🔵 finite`</sup>

Integers yields all integers in the range [start, stop) with the given step.

```go
for i := range it.Integers[uint](0, 3, 1) {
	fmt.Println(i)
}

// Chainable
for i := range itx.Integers[uint](0, 3, 1).Drop(1) {
	fmt.Println(i)
}
```
