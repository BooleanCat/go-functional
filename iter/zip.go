package iter

import "github.com/BooleanCat/go-functional/option"

type Tuple[T, U any] struct {
	One T
	Two U
}

type ZipIter[T, U any] struct {
	iter1 Iterator[T]
	iter2 Iterator[U]
}

func Zip[T, U any](iter1 Iterator[T], iter2 Iterator[U]) *ZipIter[T, U] {
	return &ZipIter[T, U]{iter1, iter2}
}

func (iter *ZipIter[T, U]) Next() option.Option[Tuple[T, U]] {
	v1, ok := iter.iter1.Next().Value()
	if !ok {
		return option.None[Tuple[T, U]]()
	}

	v2, ok := iter.iter2.Next().Value()
	if !ok {
		return option.None[Tuple[T, U]]()
	}

	return option.Some(Tuple[T, U]{v1, v2})
}

var _ Iterator[Tuple[struct{}, struct{}]] = new(ZipIter[struct{}, struct{}])
