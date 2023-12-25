package iter

import (
	"fmt"

	"github.com/BooleanCat/go-functional/option"
)

// CountIter iterator, see [Count].
type CountIter struct {
	BaseIter[int]
	index int
}

// Count instantiates a [*CountIter] that will iterate over 0 and the
// natural numbers. Count is functionally "unlimited" although it does not
// protect against the integer limit.
func Count() *CountIter {
	iter := new(CountIter)
	iter.BaseIter = BaseIter[int]{iter}
	return iter
}

// Next implements the [Iterator] interface.
func (c *CountIter) Next() option.Option[int] {
	c.index++
	return option.Some(c.index - 1)
}

// String implements the [fmt.Stringer] interface
func (c *CountIter) String() string {
	return "Iterator<Count>"
}

var _ fmt.Stringer = new(CountIter)
var _ Iterator[int] = new(CountIter)
