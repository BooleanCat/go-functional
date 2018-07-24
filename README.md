# go-functional

go-functional is a code generation tool that outputs functional helpers for
golang types.

## Show me

Let's imagine you want to find the first 100 prime numbers. You could generate
functional helpers for the `int` type like so:

```
$ go-functional int
```

You could then use the generated `fint` package to find those primes:

```go
type Counter struct {
  i int
}

func (c *Counter) Next() fint.Option {
  option := fint.Some(c.i)
  c.i++
  return option
}

func getPrimes() []int {
  return fint.New(new(Counter)).Filter(isPrime).Take(100).Collect()
}
```

## Tell me more

The basic building blocks of go-functional are `Iterator`s and `Option`s. An
`Option` is simply a struct that *may* hold some value, or it *may* hold no
value. For example, for functional helpers generated for the `string` type:

```go
a := fstring.Some("foo")  // holds the value `foo`
b := fstring.None()  // holds no value
a.Present()  # true
b.Present()  # false
a.Value  // The underlying value of the option.
```

Options should always be checked for the presence of a value before using
`Option.Value`.

`Iterators`s are defined as types that yield Options of a particular type by
implementing:

```go
type Iterator interface {
  Next() Option
}
```

An example could be an iterator that yields each letter of the alphabet:

```go
type Alphabet struct {
  letter int
}

func (a *Alphabet) Next() fstring.Option {
  if a.letter > 0x7A {
    return fstring.None()
  }

  next := a.letter + 0x61
  a.letter++
  return fstring.Some(string(next))
}
```

An iterator is single use. Iterators *may* be exhausted after some yields. In
the example above, calling `Next()` will return each letter in turn, and finally
will only ever return `fstring.None()`. Iterators may be infinite, an example
would be an infinite counter:

```go
type Counter struct {
  i int
}

func (c *Counter) Next() fint.Option {
  next := c.i
  c.i++
  return fint.Some(next)
}
```

Note that attempting to consume all values in an infinite iterater is undefined,
it is up to the programmer to ensure they don't do that.

## The Functor

The `Functor` type is a way to chain operations over an iterator easily, and
lazily. Functors are initialised from iterators, and may be collected into
slices, an example is the prime number finder above. Slices may be "lifted" to
`Functors`; let's choose only even numbers from a slice of integers:

```go
isEven := func(value int) bool { return value % 2 == 0 }
a := []int{1, 2, 3, 4, 5, 6, 7}
a = fint.Lift(a).Filter(isEven).Collect()
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

a := fint.New(new(Counter)).Map(expensiveOp).Take(2).Collect()
```

## Functor operations

### Drop

Drop the first `n` values from the Functor.

```go
a := []int{1, 2, 3, 4, 5}
a = fint.Lift(a).Drop(2).Collect()
a == []int{3, 4, 5}
```

### Take

Take the first `n` values from the Functor:

```go
a := []int{1, 2, 3, 4, 5}
a = fint.Lift(a).Take(3).Collect()
a == []int{1, 2, 3}
```

### Filter

Keep only values that satisfying `filter(value) == true`:

```go
a := []int{1, 2, 3, 4, 5}
a = fint.Lift(a).Filter(func(value) bool { return value < 3 }).Collect()
a == []int{1, 2}
```

### Exclude

Drop all values that satisfy `exclude(value) == true`:

```go
a := []int{1, 2, 3, 4, 5}
a = fint.Lift(a).Exclude(func(value) bool { return value < 3 }).Collect()
a == []int{3, 4, 5}
```

### Map

Apply an operation to each value yielded by the Functor's iterator:

```go
double := func(value int) int { return value * 2 }
a := []int{1, 2, 3, 4, 5}
a = fint.Lift(a).Map(double).Collect()
a == []int{2, 4, 6, 8, 10}
```

### Fold

Apply an operation to each value yielded by Functor's iterator and the previous
Fold result. For example, adding the first 100 integers:

```go
sum := func(a, b int) int { return a + b }
a := fint.New(new(Counter)).Take(100).Fold(0, sum)
```

The first argument to Fold is the value to use as the first "previous fold
result".

## Upcoming features

### Custom code generation dir

Currenty, `go-functional int` will create a package `fint` in the current
directory. I want to allow `go-functional -o /some/other/path int` creating
`/some/other/path/fint`.

### Support for non-builtin types

This should be allowed: `go-functional os.File`.

### Support for pointers to types

This should also be allowed: `go-functional *os.File`.

### Implement help

`go-functional [-h|--help]` should print usage. As should incorrect invocations.

### Write contributing guide

Show how to install test dependencies, run tests and PR.

### Write talk about this tool

Instructional, show real life examples. Best practises (write those too!).

### Error Functors

Looking at the usage of map, it's entirely feasible and in fact likely that
instead of an operation like `func(int) int`, you'll have
`func(int) (int, error)`. I want to allow for another kind of Functor that
handles allows for error flavours of functions. The `Collect` or `Fold`
invocation will then propogate the first error encountered. Should any calls
to operations return an error, all future operations on that functor are no ops.
For example (naming and usage subject to change):

```go
failingDouble := func(value int) (int, error) { return nil, errors.New("Oops.") }
a := fint.Lift([]int{1, 2, 3, 4, 5}).WithErrors()
_, err := a.Map(failingDouble).Collect()
err != nil
```

### Type Transformation

Let's say we wanted the first 3 integers as strings, we could re-use the
counter to define this with Type Transformation

```go
transformer := func(value int) interface{} { return interface(strconv.Atoi(value)) }
a := fint.New(new(Counter)).Take(3).Transform()
b := fstring.From(a).Collect()
b == []string{"1", "2", "3"}
```

My current thinking is that it would be up to the programmer to ensure type
safety, and `From`'s `Next` would look something like:

```go
func (f Foo) Next() Option {
  a := f.Next()
  if !a.Present() {
    return None()
  }

  if a.Value == nil {
    panic("attempt to yield nil transformatation")
  }

  v, ok := a.Value.(string)
  if !ok {
    panic("attempt to type assert non-string to string")
  }

  return Some(v)
}
```

### Consider offering pre-generated builtin helpers

Maybe I could pre-generate helpers for int, string, bool and float? Maybe attach
a "standard library" of iterators for those types? Things like fint.Counter(),
fstring.Alphabet()...
