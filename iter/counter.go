package iter

import "github.com/BooleanCat/go-functional/option"

// CountIter iterator, see [Count].
type CountIter struct {
	index int
}

// Count instantiates a [*CountIter] that will iterate over 0 and the
// natural numbers. Count is functionally "unlimited" although it does not
// protect against the integer limit.
func Count() *CountIter {
	return new(CountIter)
}

// Next implements the [Iterator] interface.
func (c *CountIter) Next() option.Option[int] {
	c.index++
	return option.Some(c.index - 1)
}

var _ Iterator[int] = new(CountIter)

// Drop is a convenience method for [Drop], providing this iterator as an
// argument.
func (c *CountIter) Drop(n uint) *DropIter[int] {
	return Drop[int](c, n)
}

// Take is a convenience method for [Take], providing this iterator as an
// argument.
func (iter *CountIter) Take(n uint) *TakeIter[int] {
	return Take[int](iter, n)
}
