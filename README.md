# Functional Programming in Go

[![GitHub release (with filter)](https://img.shields.io/github/v/release/BooleanCat/go-functional?sort=semver&logo=Go&color=%23007D9C&include_prereleases)](https://github.com/BooleanCat/go-functional/releases)
[![Actions Status](https://github.com/BooleanCat/go-functional/workflows/test/badge.svg)](https://github.com/BooleanCat/go-functional/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/BooleanCat/go-functional/v2.svg)](https://pkg.go.dev/github.com/BooleanCat/go-functional/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/BooleanCat/go-functional/v2)](https://goreportcard.com/report/github.com/BooleanCat/go-functional/v2)
[![codecov](https://codecov.io/gh/BooleanCat/go-functional/branch/main/graph/badge.svg?token=N2E43RSR14)](https://codecov.io/gh/BooleanCat/go-functional)

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

[Consumers](#consumers) will iterate over an iterator and completely or partially drain them of
values and (in most cases) collect the values into a data type.

[Iterators](#iterators) are functions that yield new values and can be ranged over. See Go's
documentation for iterators for more details.

<h2 id="consumers">Consumers</h2>

The standard libary provides functions to collect iterators in the `slices` and `maps` packages that
should satisfy most cases where collection is needed.

This package provides additional collection methods and makes existing consumers from the standard
library chainable.

<!-- prettier-ignore -->
> [!WARNING]
> Attempting to collect infinite iterators will cause an infinite loop and likely deadlock. Consider
> bounding infinite iterators before collect (for example using `Take`).

### Collect

In most cases `slices.Collect` from the standard library may be used to collect items from an
iterator into a slice. There are several other variants of collect available for use for different
use cases.

```go
// Chainable
numbers := itx.NaturalNumbers[int]().Take(5).Collect()

// Collect an `iter.Seq2[V, W] into two slices
keys, values := it.Collect2(maps.All(map[string]int{"one": 1, "two": 2}))

// As above, but chainable
keys, values := itx.FromMap(map[string]int{"one": 1, "two": 2}).Collect()
```

### TryCollect

Dealing with iterators that return `T, error` can involve the boilerplate of checking that the
returned slice of errors only contains `nil`. `TryCollect` solves this by collecting all values into
a slice and returning a single error: the first one encountered.

```go
text := strings.NewReader("one\ntwo\nthree\n")

if lines, err := it.TryCollect(it.LinesString(text)); err != nil {
	fmt.Println(lines)
}
```

<!-- prettier-ignore -->
> [!NOTE]
> If an error is encountered, collection stops. This means the iterator being collected may not be
> fully drained.

<!-- prettier-ignore -->
> [!NOTE]
> The `itx` package does not contain `TryCollect` due to limitations with Go's type system.

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

### Fold

Fold every element into an accumulator by applying a function and passing an initial value.

```go
it.Fold(slices.Values([]int{1, 2, 3}), op.Add, 0)

// Fold an iter.Seq2
it.Fold2(slices.All([]int{1, 2, 3}), func(i, a, b int) int {
	return i + 1
}, 0)
```

<!-- prettier-ignore -->
> [!TIP]
> The [op package](it/op/op.go) contains some simple, pre-defined operation functions.

<!-- prettier-ignore -->
> [!NOTE]
> The `itx` package does not contain `Fold` due to limitations with Go's type system.

### Max & Min

Max and Min consume an iterator and return the maximum or minimum value yielded and true if the
iterator contained at least one value, or the zero value and false if the iterator was empty.

The type of the value yielded by the iterator must be `comparable`.

```go
max, ok := it.Max(slices.Values([]int{1, 2, 3}))
min, ok := it.Min(slices.Values([]int{1, 2, 3}))
```

<!-- prettier-ignore -->
> [!NOTE]
> The `itx` package does not contain `Fold` due to limitations with Go's type system.

### Len

Len consumes an iterator and returns the number of values yielded.

```go
it.Len(slices.Values([]int{1, 2, 3}))

// Chainable
itx.FromSlice([]int{1, 2, 3}).Len()

// Len of an iter.Seq2
it.Len2(slices.All([]int{1, 2, 3}))

// As above, but chainable
itx.FromSlice([]int{1, 2, 3}).Enumerate().Len()
```

### Find

Find consumes an iterator until a value is found that satisfies a predicate. It returns the value
and true if one was found, or the zero value and false if the iterator was exhausted before a value
was found.

```go
found, ok := it.Find(slices.Values([]int{1, 2, 3}), func(i int) bool {
	return i == 2
})

// Chainable
value, ok := itx.FromSlice([]int{1, 2, 3}).Find(func(number int) bool {
	return number == 2
})

// Finding within an iter.Seq2
index, value, ok := it.Find2(slices.All([]int{1, 2, 3}), func(i, v int) bool {
	return i == 2
})

// As above, but chainable
index, value, ok := itx.FromSlice([]int{1, 2, 3}).Enumerate().Find(func(index int, number int) bool {
	return index == 1
})
```

<!-- prettier-ignore -->
> [!TIP]
> The [filter package](it/filter/filter.go) contains some simple, pre-defined predicate functions.

### From, FromSlice, FromMap & Seq

The itx package contains some helper functions to convert iterators, slices or maps directly into
chainable iterators to avoid some boilerplate.

```go
itx.From(slices.Values([]int{1, 2, 3})).Collect()

itx.From2(maps.All(map[int]string{1: "one"}))

itx.FromSlice([]int{1, 2, 3}).Collect()

itx.FromMap(map[int]int{1: 2}).Collect()
```

The `itx` package also contains a helper function Seq that will convert a chainable iterator into an
`iter.Seq` so that it can be used in functions that accept that type (such as the standard library).

The standard library functions that work with iterators (such as `slices.Collect`) accept the
`iter.Seq` family of types. This precludes those functions from accepting types with another alias
(such as `itx.Iterator`) with the same type definition. This means it is necessary to "covert" an
`itx.Iterator` into an `iter.Seq` before passing the iterator into those functions.

```go
slices.Collect(itx.NaturalNumbers[int]().Take(3).Seq())
```

<!-- prettier-ignore -->
> [!TIP]
> go-functional's functions that accept iterators always accept `func(func(V) bool)` or
> `func(func(V, W) bool)` rather than any specific type alias so that they can accept any type alias
> with the definitions `iter.Seq`, `iter.Seq2`, `itx.Iterator`, `itx.Iterator2` or any other
> third-party types aliased to the same type.

### ToChannel

ToChannel sends yielded values to a channel.

The channel is closed when the iterator is exhausted. Beware of leaked go routines when using this
function with an infinite iterator.

```go
channel := it.ToChannel(slices.Values([]int{1, 2, 3}))

for number := range channel {
	fmt.Println(number)
}

// Chainable
channel := itx.FromSlice([]int{1, 2, 3}).ToChannel()

for number := range channel {
	fmt.Println(number)
}
```

<!-- prettier-ignore -->
> [!NOTE]
> Unlike most consumers, the iterator is not immediately consumed by ToChannel. Instead is it
> consumed as values are pulled from the channel.

<h2 id="iterators">Iterators</h2>

This library contains two kinds of iterators in the `it` and `itx` packages. In most cases you'll
find the same iterators in each package, the difference between them being that the iterators in the
`itx` package can be "dot-chained" (e.g. `iter.Filter(...).Take(3).Collect()`) and those in `it`
cannot.

Iterators within the `it` package are of the type `iter.Seq[V]` or `iter.Seq2[V, W]` (from the
standard library). Iterators within the `itx` package are of the type `itx.Iterator[V]` or
`itx.Iterator2[V, W]`.

There are two important factors to consider when using iterators:

1. Some iterators yield an infinite number of values and care should be taken to avoid consuming
   (using functions such as `slices.Collect`) otherwise it's likely to cause an infinite while loop.
2. Many iterators take another iterator as an argument (such as Filter or Map). Avoid using an
   iterator after it has been passed to another iterator otherwise you'll risk multiple functions
   consuming a single (likely not thread-safe) iterator and causing confusing and difficult to debug
   behaviour.

### Chain

Chain yields values from multiple iterators in the sequence they are provided in. Think of it as
glueing many iterators together.

When provided zero iterators it will behave like `it.Exhausted`.

```go
numbers := it.Chain(slices.Values([]int{1, 2}), slices.Values([]int{3, 4}))

pairs := itx.FromSlice([]int{1, 2}).Chain(slices.Values([]int{3, 4}))

pairs := it.Chain2(maps.All(map[string]int{"a": 1}), maps.All(map[string]int{"b": 2}))

pairs := itx.FromMap(map[string]int{"a": 1}).Chain(maps.All(map[string]int{"b": 2}))
```

### FromChannel

FromChannel pulls values from a channel and yields them via an iterator. The usual concerns around
channel deadlocks apply here.

The iterator is exhausted when the channel is closed and it is the responsibility of the caller to
close the channel.

```go
items := make(chan int)

go func() {
	defer close(items)
	items <- 1
	items <- 2
}()

for number := range it.FromChannel(items) {
	fmt.Println(number)
}

// Chainable
items := make(chan int)

go func() {
	defer close(items)
	items <- 1
	items <- 0
}()

for number := range itx.FromChannel(items).Exclude(filter.IsZero) {
	fmt.Println(number)
}
```

<!-- prettier-ignore -->
> [!WARNING]
> In order to prevent a deadlock, the channel must be closed before attemping to stop the iterator
> when it's used in a pull style. See `iter.Pull`.

### Cycle

Cycle yields all values from an iterator before returning to the beginning and yielding all values
again (indefinitely).

```go
numbers := it.Take(it.Cycle(slices.Values([]int{1, 2})), 5)

// Chainable
numbers := itx.FromSlice([]int{1, 2}).Cycle().Take(5)

// Cycling an iter.Seq2
numbers := it.Take2(it.Cycle2(maps.All(map[int]string{1: "one"})), 5)

// As above, but chainable
numbers := itx.FromMap(map[int]string{1: "one"}).Cycle().Take(5)
```

<!-- prettier-ignore -->
> [!NOTE]
> Since cycle needs to store all values yielded in memory, its memory usage will grow as the first
> cycle is consumed and remain at a constant size on subsequent cycles.

<!-- prettier-ignore -->
> [!WARNING]
> This iterator yields an infinite number of values and care should be taken when consuming it
> otherwise it's likely to result in an infinite while loop. Consider bounding the size of the
> iterator before consuming (e.g. using Take).

### Drop

Drop yields all values from a delegate iterator after dropping a number of values from the
beginning. Values are not dropped immediately, but when consumption begins.

When dropping a number of values larger than the length of the iterator, it behaves like
`it.Exhausted`.

```go
numbers := it.Drop(slices.Values([]int{1, 2, 3, 4, 5}), 2)

// Chainable
numbers := itx.FromSlice([]int{1, 2, 3, 4, 5}).Drop(2)

// Dropping on iter.Seq2
numbers := it.Drop2(maps.All(map[int]string{1: "one", 2: "two", 3: "three"}), 1)

// As above, but chainable
numbers := itx.FromMap(map[int]string{1: "one", 2: "two", 3: "three"}).Drop(1)
```

### Enumerate

Enumerating an `iter.Seq` like iterator returns an `iter.Seq2` like iterator yielding the index of
each value and the value.

```go
indexedValues := it.Enumerate(slices.Values([]int{1, 2, 3}))

// Chainable
indexedValues := itx.FromSlice([]int{1, 2, 3}).Enumerate()
```

<!-- prettier-ignore -->
> [!TIP]
> When iterating over a slice and immediately enumerating, consider instead using the standard
> library's `slices.All` function rather than this.

### Integers

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

### NaturalNumbers

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

<!-- prettier-ignore -->
> [!WARNING]
> This iterator yields an infinite number of values and care should be taken when consuming it
> otherwise it's likely to result in an infinite while loop. Consider bounding the size of the
> iterator before consuming (e.g. using Take).

<!-- prettier-ignore -->
> [!WARNING]
> There is no protection against overflowing whatever integer type is used for this iterator.
