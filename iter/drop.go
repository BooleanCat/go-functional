package iter

import (
	"fmt"
	"reflect"

	"github.com/BooleanCat/go-functional/option"
)

// DropIter iterator, see [Drop].
type DropIter[T any] struct {
	BaseIter[T]
	iter      Iterator[T]
	count     uint
	dropped   bool
	exhausted bool
}

// Drop instantiates a [*DropIter] that will skip the number of items of its
// wrapped iterator by the provided count.
func Drop[T any](iter Iterator[T], count uint) *DropIter[T] {
	iterator := &DropIter[T]{iter: iter, count: count}
	iterator.BaseIter = BaseIter[T]{iterator}
	return iterator
}

// Next implements the [Iterator] interface.
func (iter *DropIter[T]) Next() option.Option[T] {
	if iter.exhausted {
		return option.None[T]()
	}

	if !iter.dropped {
		iter.dropped = true

		for i := uint(0); i < iter.count; i++ {
			if iter.delegateNext().IsNone() {
				return option.None[T]()
			}
		}
	}

	return iter.delegateNext()
}

func (iter *DropIter[T]) delegateNext() option.Option[T] {
	next := iter.iter.Next()
	if next.IsNone() {
		iter.exhausted = true
	}

	return next
}

// String implements the [fmt.Stringer] interface
func (iter DropIter[T]) String() string {
	var instanceOfT T
	return fmt.Sprintf("Iterator<Drop, type=%s>", reflect.TypeOf(instanceOfT))
}

var _ Iterator[struct{}] = new(DropIter[struct{}])
