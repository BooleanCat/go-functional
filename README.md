# Functional Programming in Go

[![GitHub release (with filter)](https://img.shields.io/github/v/release/BooleanCat/go-functional?sort=semver&logo=Go&color=%23007D9C)](https://github.com/BooleanCat/go-functional/releases) [![Actions Status](https://github.com/BooleanCat/go-functional/workflows/test/badge.svg)](https://github.com/BooleanCat/go-functional/actions) [![Go Reference](https://pkg.go.dev/badge/github.com/BooleanCat/go-functional.svg)](https://pkg.go.dev/github.com/BooleanCat/go-functional/v2) [![Go Report Card](https://goreportcard.com/badge/github.com/BooleanCat/go-functional)](https://goreportcard.com/report/github.com/BooleanCat/go-functional) [![codecov](https://codecov.io/gh/BooleanCat/go-functional/branch/main/graph/badge.svg?token=N2E43RSR14)](https://codecov.io/gh/BooleanCat/go-functional)

A library of iterators for use with `iter.Seq`.

```go
// The first 5 natural numbers
numbers := slices.Collect(
	it.Take(it.Count(), 5),
)

// All even numbers
evens := it.Filter(it.Count(), filter.IsEven)

// String representations of integers
numbers := it.Map(it.Count(), strconv.Itoa)
```

_[Read the docs](https://pkg.go.dev/github.com/BooleanCat/go-functional/v2)_ to see the full iterator library.
