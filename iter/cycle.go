package iter

import (
	"fmt"
	"reflect"

	"github.com/BooleanCat/go-functional/option"
)

// CycleIter iterator, see [Cycle].
type CycleIter[T any] struct {
	BaseIter[T]
	iter  Iterator[T]
	items []T
	index int
}

// Cycle instantiates a [*CycleIter] yielding all items from the provided
// iterator until exhaustion, and then yields them all over again on repeat.
//
// Note that [CycleIter] stores the members from the underlying iterator so
// will grow in size as items are yielded until the underlying iterator is
// exhausted.
//
// In most cases this iterator is infinite, except when the underlying iterator
// is exhausted before the first call to Next() in which case this iterator
// will always yield None.
func Cycle[T any](iter Iterator[T]) *CycleIter[T] {
	iterator := &CycleIter[T]{iter: iter, items: make([]T, 0)}
	iterator.BaseIter = BaseIter[T]{iterator}
	return iterator
}

// Next implements the [Iterator] interface.
func (iter *CycleIter[T]) Next() option.Option[T] {
	if iter.iter != nil {
		if value, ok := iter.iter.Next().Value(); ok {
			iter.items = append(iter.items, value)
			return option.Some(value)
		} else {
			iter.iter = nil
		}
	}

	if len(iter.items) == 0 {
		return option.None[T]()
	}

	if iter.index == len(iter.items) {
		iter.index = 0
	}

	next := iter.items[iter.index]
	iter.index++
	return option.Some(next)
}

// String implements the [fmt.Stringer] interface
func (iter CycleIter[T]) String() string {
	var instanceOfT T
	return fmt.Sprintf("Iterator<Cycle, type=%s>", reflect.TypeOf(instanceOfT))
}

var (
	_ fmt.Stringer       = new(CycleIter[struct{}])
	_ Iterator[struct{}] = new(CycleIter[struct{}])
)
