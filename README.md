# Functional Programming in Go

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

## To Do

- Write a better README
- Learn how to use go doc and generate docs
- Write some Example specs for auto-docs
- Formalise iterator rules for things like repeated calls and chain-call rules