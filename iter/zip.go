package iter

import "github.com/BooleanCat/go-functional/option"

// Tuples are pairs of values.
type Tuple[T, U any] struct {
	One T
	Two U
}

// ZipIter iterator, see [Zip].
type ZipIter[T, U any] struct {
	iter1     Iterator[T]
	iter2     Iterator[U]
	exhausted bool
}

// Zip instantiates a [*ZipIter] yielding a [Tuple] containing the result of a
// call to each provided [Iterator]'s Next. This iterator is exhausted when one
// of the provided iterators is exhausted.
func Zip[T, U any](iter1 Iterator[T], iter2 Iterator[U]) *ZipIter[T, U] {
	return &ZipIter[T, U]{iter1, iter2, false}
}

// Next implements the [Iterator] interface.
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

// Collect is a convenience method for [Collect], providing this iterator as
// an argument.
func (iter *ZipIter[T, U]) Collect() []Tuple[T, U] {
	return Collect[Tuple[T, U]](iter)
}

// ForEach is a convenience method for [ForEach], providing this iterator as an
// argument.
func (iter *ZipIter[T, U]) ForEach(callback func(Tuple[T, U])) {
	ForEach[Tuple[T, U]](iter, callback)
}

// Find is a convenience method for [Find], providing this iterator as an
// argument.
func (iter *ZipIter[T, U]) Find(predicate func(Tuple[T, U]) bool) option.Option[Tuple[T, U]] {
	return Find[Tuple[T, U]](iter, predicate)
}

// Drop is a convenience method for [Drop], providing this iterator as an
// argument.
func (iter *ZipIter[T, U]) Drop(n uint) *DropIter[Tuple[T, U]] {
	return Drop[Tuple[T, U]](iter, n)
}

// Take is a convenience method for [Take], providing this iterator as an
// argument.
func (iter *ZipIter[T, U]) Take(n uint) *TakeIter[Tuple[T, U]] {
	return Take[Tuple[T, U]](iter, n)
}
