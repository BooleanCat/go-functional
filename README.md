# go-functional

![](https://github.com/BooleanCat/go-functional/workflows/test/badge.svg)

go-functional is a code generation tool that outputs functional helpers for
golang types.

## Show me

Let's imagine you want to find the first 100 prime numbers. Functional helpers
for the `int` type can be generated like so:

```
$ go-functional int
```

The generated `fint` package can be leveraged to find primes:

```go
type Counter struct {
  i fint.T
}

func (c *Counter) Next() fint.OptionalResult {
  next := fint.Some(c.i)
  c.i++
  return fint.Success(next)
}

func getPrimes() []int {
  return fint.New(new(Counter)).Filter(isPrime).Take(100).Collapse()
}
```

## Tell me more

The basic building blocks of go-functional are iterators. Iterators are defined
as such:

```go
type Iter interface {
  Next() OptionalResult
}
```

That is, they are implementations that define a method `Next` that yields
a value, wrapped in an `OptionalResult`, each time it is invoked.

An `OptionalResult` is a type that may hold an `Option` or an error. The
presence of an error can be checked with `result.Error()`. If `result.Error()`
is nil, then `result.Value()` holds a meaningful optional value.

An `Option` is a type that *may* hold some value. The presence of a value can
be checked with `option.Present()`. If `option.Present()` is true, then
`option.Value()` holds a meaningful value.

Here is an example of an iterator that yields each letter of the alphabet:

```go
type Alphabet struct {
  letter int
}

func (a *Alphabet) Next() fstring.OptionalResult {
  if a.letter > 25 {
    return fstring.Success(fstring.None())
  }

  next := fstring.Some(fstring.T(a.letter + 0x61))
  a.letter++
  return fstring.Success(next)
}
```

Iterators *may* be exhausted after some yields. In the example above, calling
`Next()` will return each letter in turn, and finally will only ever return
`fstring.Success(fstring.None())` (a successful empty value). An iterator is
considered exhausted when `Next()` yields an empty value.

Iterators may be infinite, an example would be an infinite counter:

```go
type Counter struct {
  i fint.T
}

func (c *Counter) Next() fint.OptionalResult {
  next := fint.Some(c.i)
  c.i++
  return fint.Success(next)
}
```

## The Functor

The `Functor` type is a way to chain operations over an iterator more
conveniently. Functors are initialised from iterators, and may be collapsed into
slices, an example is the prime number finder above. Slices may be "lifted" to
`Functors`; let's choose only even numbers from a slice of integers:

```go
isEven := func(value int) bool { return value%2 == 0 }
numbers := []int{1, 2, 3, 4, 5, 6, 7}
numbers = fint.Lift(numbers).Filter(isEven).Collapse()
```

## Laziness?

Yep, functor operations are lazy. In the example below, `expensiveOp` is only
computed twice:

```go
func expensiveOp(value int) int {
  time.Sleep(time.Second)
  return value * 2
}

...

_ = fint.New(new(Counter)).Map(expensiveOp).Take(2).Collapse()
```

## Functor operations

### Drop

Drop the first `n` values from the Functor.

```go
numbers := []int{1, 2, 3, 4, 5}
numbers = fint.Lift(numbers).Drop(2).Collapse()
numbers == []int{3, 4, 5}
```

### Take

Take the first `n` values from the Functor:

```go
numbers := []int{1, 2, 3, 4, 5}
numbers = fint.Lift(numbers).Take(3).Collapse()
numbers == []int{1, 2, 3}
```

### Filter

Keep only values that satisfying `filter(value) == true`:

```go
numbers := []int{1, 2, 3, 4, 5}
numbers = fint.Lift(numbers).Filter(func(value int) bool { return value < 3 }).Collapse()
numbers == []int{1, 2}
```

### Exclude

Drop all values that satisfy `exclude(value) == true`:

```go
numbers := []int{1, 2, 3, 4, 5}
numbers = fint.Lift(numbers).Exclude(func(value int) bool { return value < 3 }).Collapse()
numbers == []int{3, 4, 5}
```

### Map

Apply an operation to each value yielded by the Functor's iterator:

```go
double := func(value int) int { return value * 2 }
numbers := []int{1, 2, 3, 4, 5}
numbers = fint.Lift(numbers).Map(double).Collapse()
numbers == []int{2, 4, 6, 8, 10}
```

### Roll

Apply an operation to each value yielded by Functor's iterator and the previous
Roll result. For example, adding the first 100 integers:

```go
sum := func(a, b int) int { return a + b }
numbers := fint.New(new(Counter)).Take(100).Roll(0, sum)
```

The first argument to Fold is the value to use as the first "previous fold
result".

## Error handling

It is common in Go for functions to return errors. Let's image we have a Map
operation who's signature is `func(i int) (int, error)`, Functors provide a way
to handle such cases:

```go
numbers, err := fint.New(iter).MapErr(op).Collect()
```

Notice that we used `Collect` rather than `Collapse` to turn the functor into
a slice - collecting an iterator has the possibility of returning an error, and
we the programmers chose to collect rather than collapse because we admit the
possibility of errors occurring. Collapsing a functor which encounters an error
during evaluation of its members will result in a runtime panic. It is up to
the programmer to choose the appropriate method.

Other familiar functor operations provide error flavours.

It is also possible to Roll an iterator into a result with error admission, you
would instead invoke:

```go
number, err := fint.New(iter).MapErr(op).Fold(0, sum)
```

## Type Transformation

Let's imagine you have a counter iterator, yielding integers incrementally. Now
you want them as strings rather than integers, is it possible to to do so using
a family of functions: Transform, Transmute and Blur.

We would first "blur" our int iterator in to a "GenericIter":

```go
iter := fint.New(NewCounter()).Blur()
```

Then, we transform the iterator into a string iterator, providing a function to
instruct transform how that should happen:

```go
toString := func(v interface{}) (string, error) {
  i := fint.Transmute(v)  // a helper for v.(int) that panics if the type assertion fails
  return strconv.Itoa(i), nil
}

numbers := fstring.New(fstring.Transform(iter, toString)).Take(5).Collapse()
```

## Upcoming features

### Custom code generation dir

Currenty, `go-functional int` will create a package `fint` in the current
directory. I want to allow `go-functional -o /some/other/path int` creating
`/some/other/path/fint`.

### Support for non-builtin types

This should be allowed: `go-functional os.File`.

### Implement help

`go-functional [-h|--help]` should print usage. As should incorrect invocations.

### Write contributing guide

Show how to install test dependencies, run tests and PR.

### Write talk about this tool

Instructional, show real life examples. Best practises (write those too!).

### Consider offering pre-generated builtin helpers

Maybe I could pre-generate helpers for int, string, bool and float? Maybe attach
a "standard library" of iterators for those types? Things like fint.Counter(),
fstring.Alphabet()...
