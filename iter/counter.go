package iter

import "github.com/BooleanCat/go-functional/option"

// CountIter implements `Count`. See `Count`'s documentation.
type CountIter struct {
	index int
}

// Count instantiates a `CountIter` that will iterate over 0 and the
// natural numbers. Count is functionally "unlimited" although it does not
// protect against the integer limit.
func Count() *CountIter {
	return new(CountIter)
}

// Next implements the Iterator interface for `CountIter`.
func (c *CountIter) Next() option.Option[int] {
	c.index += 1
	return option.Some(c.index - 1)
}

var _ Iterator[int] = new(CountIter)
