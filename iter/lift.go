package iter

import (
	"context"
	"fmt"
	"io"
	"reflect"

	"github.com/BooleanCat/go-functional/option"
)

// LiftIter iterator, see [Lift].
type LiftIter[T any] struct {
	BaseIter[T]
	items []T
	index int
}

// Lift instantiates a [*LiftIter] that will yield all items in the provided
// slice.
func Lift[T any](items []T) *LiftIter[T] {
	iterator := &LiftIter[T]{items: items}
	iterator.BaseIter = BaseIter[T]{iterator}
	return iterator
}

// Next implements the [Iterator] interface.
func (iter *LiftIter[T]) Next() option.Option[T] {
	if iter.index >= len(iter.items) {
		return option.None[T]()
	}

	iter.index++

	return option.Some(iter.items[iter.index-1])
}

// String implements the [fmt.Stringer] interface
func (iter LiftIter[T]) String() string {
	var instanceOfT T
	return fmt.Sprintf("Iterator<Lift, type=%s>", reflect.TypeOf(instanceOfT))
}

var (
	_ fmt.Stringer       = new(LiftIter[struct{}])
	_ Iterator[struct{}] = new(LiftIter[struct{}])
)

// LiftHashMapIter iterator, see [LiftHashMap].
type LiftHashMapIter[T comparable, U any] struct {
	BaseIter[Pair[T, U]]
	hashmap map[T]U
	items   chan Pair[T, U]
	cancel  func()
}

// LiftHashMap instantiates a [*LiftHashMapIter] that will yield all items in
// the provided map as a [Pair].
//
// Unlike most iterators, [LiftHashMap] should be closed after usage (because
// range order is non-deterministic and the iterator needs to preserve its
// progress). This restriction may be removed if/when Go has a "yield" keyword.
//
// The iterator is closed when any of the two conditions are met.
//
// 1. The caller explicitly invokes the `Close` method.
// 2. The iterator is exhausted.
//
// It is safe to call Close multiple times or after exhaustion. It is not
// necessary to call Close if exhaustion is guaranteed, but may be wise to
// redundantly call Close if you're unsure.
func LiftHashMap[T comparable, U any](hashmap map[T]U) *LiftHashMapIter[T, U] {
	ctx, cancel := context.WithCancel(context.Background())

	iter := &LiftHashMapIter[T, U]{
		hashmap: hashmap,
		items:   make(chan Pair[T, U]),
		cancel:  cancel,
	}

	iter.BaseIter = BaseIter[Pair[T, U]]{iter}

	go func() {
		defer close(iter.items)

		for k, v := range hashmap {
			select {
			case <-ctx.Done():
				return
			case iter.items <- Pair[T, U]{k, v}:
			}
		}
	}()

	return iter
}

// Close the iterator. See [LiftHashMap]'s documentation for details.
//
// This function implements the [io.Closer] interface.
//
// This function can never fail and the error can be ignored.
func (iter *LiftHashMapIter[T, U]) Close() error {
	iter.cancel()
	<-iter.items

	return nil
}

// Next implements the [Iterator] interface.
func (iter *LiftHashMapIter[T, U]) Next() option.Option[Pair[T, U]] {
	pair, ok := <-iter.items
	if !ok {
		return option.None[Pair[T, U]]()
	}

	return option.Some(pair)
}

// String implements the [fmt.Stringer] interface
func (iter LiftHashMapIter[T, U]) String() string {
	var (
		instanceOfT T
		instanceOfU U
	)

	return fmt.Sprintf("Iterator<LiftHashMap, type=Pair<%s, %s>>", reflect.TypeOf(instanceOfT), reflect.TypeOf(instanceOfU))
}

var (
	_ fmt.Stringer                       = new(LiftHashMapIter[struct{}, struct{}])
	_ Iterator[Pair[struct{}, struct{}]] = new(LiftHashMapIter[struct{}, struct{}])
	_ io.Closer                          = new(LiftHashMapIter[struct{}, struct{}])
)

// LiftHashMapKeysIter iterator, see [LiftHashMapKeys].
type LiftHashMapKeysIter[T comparable, U any] struct {
	BaseIter[T]
	delegate    *LiftHashMapIter[T, U]
	delegateMap *MapIter[Pair[T, U], T]
	exhausted   bool
}

// LiftHashMapKeys instantiates a [*LiftHashMapKeysIter] that will yield all
// keys in the provided map.
//
// See [LiftHashMap] for information on closing this iterator.
func LiftHashMapKeys[T comparable, U any](hashmap map[T]U) *LiftHashMapKeysIter[T, U] {
	delegate := LiftHashMap(hashmap)

	iter := &LiftHashMapKeysIter[T, U]{
		delegate:    delegate,
		delegateMap: Map[Pair[T, U]](delegate, func(pair Pair[T, U]) T { return pair.One }),
	}

	iter.BaseIter = BaseIter[T]{iter}

	return iter
}

// Close the iterator. See [LiftHashMap]'s documentation for details.
//
// This function implements the [io.Closer] interface.
//
// This function can never fail and the error can be ignored.
func (iter *LiftHashMapKeysIter[T, U]) Close() error {
	return iter.delegate.Close()
}

// Next implements the [Iterator] interface.
func (iter *LiftHashMapKeysIter[T, U]) Next() option.Option[T] {
	if iter.exhausted {
		return option.None[T]()
	}

	next := iter.delegateMap.Next()
	if next.IsNone() {
		iter.exhausted = true
	}

	return next
}

// String implements the [fmt.Stringer] interface
func (iter LiftHashMapKeysIter[T, U]) String() string {
	var instanceOfT T
	return fmt.Sprintf("Iterator<LiftHashMapKeys, type=%s>", reflect.TypeOf(instanceOfT))
}

var (
	_ fmt.Stringer       = new(LiftHashMapKeysIter[struct{}, struct{}])
	_ Iterator[struct{}] = new(LiftHashMapKeysIter[struct{}, struct{}])
	_ io.Closer          = new(LiftHashMapKeysIter[struct{}, struct{}])
)

// LiftHashMapValuesIter iterator, see [LiftHashMapValues].
type LiftHashMapValuesIter[T comparable, U any] struct {
	BaseIter[U]
	delegate    *LiftHashMapIter[T, U]
	delegateMap *MapIter[Pair[T, U], U]
	exhausted   bool
}

// LiftHashMapValues instantiates a [*LiftHashMapValuesIter] that will yield
// all keys in the provided map.
//
// See [LiftHashMap] for information on closing this iterator.
func LiftHashMapValues[T comparable, U any](hashmap map[T]U) *LiftHashMapValuesIter[T, U] {
	delegate := LiftHashMap(hashmap)

	iter := &LiftHashMapValuesIter[T, U]{
		delegate:    delegate,
		delegateMap: Map[Pair[T, U]](delegate, func(pair Pair[T, U]) U { return pair.Two }),
	}

	iter.BaseIter = BaseIter[U]{iter}

	return iter
}

// Close the iterator. See [LiftHashMap]'s documentation for details.
//
// This function implements the [io.Closer] interface.
//
// This function can never fail and the error can be ignored.
func (iter *LiftHashMapValuesIter[T, U]) Close() error {
	return iter.delegate.Close()
}

// Next implements the [Iterator] interface.
func (iter *LiftHashMapValuesIter[T, U]) Next() option.Option[U] {
	if iter.exhausted {
		return option.None[U]()
	}

	next := iter.delegateMap.Next()
	if next.IsNone() {
		iter.exhausted = true
	}

	return next
}

// String implements the [fmt.Stringer] interface
func (iter LiftHashMapValuesIter[T, U]) String() string {
	var instanceOfU U
	return fmt.Sprintf("Iterator<LiftHashMapValues, type=%s>", reflect.TypeOf(instanceOfU))
}

var (
	_ fmt.Stringer       = new(LiftHashMapValuesIter[struct{}, struct{}])
	_ Iterator[struct{}] = new(LiftHashMapValuesIter[struct{}, struct{}])
	_ io.Closer          = new(LiftHashMapValuesIter[struct{}, struct{}])
)
