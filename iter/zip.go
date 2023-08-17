package iter

import "github.com/BooleanCat/go-functional/option"

type Tuple[T, U any] struct {
	One T
	Two U
}

// ZipIter implements `Zip`. See `Zip`'s documentation.
type ZipIter[T, U any] struct {
	iter1     Iterator[T]
	iter2     Iterator[U]
	exhausted bool
}

// Zip instantiates a `Zip` yield `Tuples` containing the result of a call to
// each provided Iterator's `Next`. The Iterator is exhausted when one of the
// provided Iterators is exhausted.
func Zip[T, U any](iter1 Iterator[T], iter2 Iterator[U]) *ZipIter[T, U] {
	return &ZipIter[T, U]{iter1, iter2, false}
}

// Next implements the Iterator interface for `ZipIter`.
func (iter *ZipIter[T, U]) Next() option.Option[Tuple[T, U]] {
	if iter.exhausted {
		return option.Option[Tuple[T, U]]{}
	}

	v1, ok1 := iter.iter1.Next().Value()
	v2, ok2 := iter.iter2.Next().Value()

	if !ok1 || !ok2 {
		iter.exhausted = true
		return option.None[Tuple[T, U]]()
	}

	return option.Some(Tuple[T, U]{v1, v2})
}

var _ Iterator[Tuple[struct{}, struct{}]] = new(ZipIter[struct{}, struct{}])

// Collect is an alternative way of invoking Collect(iter)
func (iter *ZipIter[T, U]) Collect() []Tuple[T, U] {
	return Collect[Tuple[T, U]](iter)
}

// Drop is an alternative way of invoking Drop(iter)
func (iter *ZipIter[T, U]) Drop(n uint) *DropIter[Tuple[T, U]] {
	return Drop[Tuple[T, U]](iter, n)
}
