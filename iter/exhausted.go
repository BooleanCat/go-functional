package iter

import "github.com/BooleanCat/go-functional/option"

// ExhaustedIter iterator, see [Exhausted].
type ExhaustedIter[T any] struct {
	BaseIter[T]
}

// Exhausted instantiates an [*ExhaustedIter] that will immediately be
// exhausted (Next will always return a None variant).
func Exhausted[T any]() *ExhaustedIter[T] {
	iter := new(ExhaustedIter[T])
	iter.BaseIter = BaseIter[T]{iter}
	return iter
}

// Next implements the [Iterator] interface.
func (iter *ExhaustedIter[T]) Next() option.Option[T] {
	return option.None[T]()
}

var _ Iterator[struct{}] = new(ExhaustedIter[struct{}])
