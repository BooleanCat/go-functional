package iter

import (
	"io"
	"sync"

	"github.com/BooleanCat/go-functional/option"
)

// LiftIter iterator, see [Lift].
type LiftIter[T any] struct {
	items []T
	index int
}

// Filter istantiates a [*FilterIter] for filtering by a chosen function.
func (iter *LiftIter[T]) Filter(fun func(T) bool) *FilterIter[T] {
	return &FilterIter[T]{iter, fun, false}
}

// Lift instantiates a [*LiftIter] that will yield all items in the provided
// slice.
func Lift[T any](items []T) *LiftIter[T] {
	return &LiftIter[T]{items, 0}
}

// Next implements the [Iterator] interface.
func (iter *LiftIter[T]) Next() option.Option[T] {
	if iter.index >= len(iter.items) {
		return option.None[T]()
	}

	iter.index++

	return option.Some(iter.items[iter.index-1])
}

var _ Iterator[struct{}] = new(LiftIter[struct{}])

// Collect is a convenience method for [Collect], providing this iterator as
// an argument.
func (iter *LiftIter[T]) Collect() []T {
	return Collect[T](iter)
}

// ForEach is a convenience method for [ForEach], providing this iterator as an
// argument.
func (iter *LiftIter[T]) ForEach(callback func(T)) {
	ForEach[T](iter, callback)
}

// Find is a convenience method for [Find], providing this iterator as an
// argument.
func (iter *LiftIter[T]) Find(predicate func(T) bool) option.Option[T] {
	return Find[T](iter, predicate)
}

// Drop is a convenience method for [Drop], providing this iterator as an
// argument.
func (iter *LiftIter[T]) Drop(n uint) *DropIter[T] {
	return Drop[T](iter, n)
}

// Take is a convenience method for [Take], providing this iterator as an
// argument.
func (iter *LiftIter[T]) Take(n uint) *TakeIter[T] {
	return Take[T](iter, n)
}

// LiftHashMapIter iterator, see [LiftHashMap].
type LiftHashMapIter[T comparable, U any] struct {
	hashmap  map[T]U
	items    chan Tuple[T, U]
	stopOnce sync.Once
	stop     chan struct{}
}

// LiftHashMap instantiates a [*LiftHashMapIter] that will yield all items in
// the provided map as a [Tuple].
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
	iter := &LiftHashMapIter[T, U]{hashmap, make(chan Tuple[T, U]), sync.Once{}, make(chan struct{}, 1)}

	go func() {
		defer close(iter.items)
		defer iter.stopOnce.Do(func() { close(iter.stop) })
	outer:
		for k, v := range hashmap {
			select {
			case iter.items <- Tuple[T, U]{k, v}:
				continue
			case <-iter.stop:
				break outer
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
	iter.stopOnce.Do(func() {
		iter.stop <- struct{}{}
		close(iter.stop)
	})

	return nil
}

// Next implements the [Iterator] interface.
func (iter *LiftHashMapIter[T, U]) Next() option.Option[Tuple[T, U]] {
	pair, ok := <-iter.items
	if !ok {
		return option.None[Tuple[T, U]]()
	}

	return option.Some(pair)
}

var (
	_ Iterator[Tuple[struct{}, struct{}]] = new(LiftHashMapIter[struct{}, struct{}])
	_ io.Closer                           = new(LiftHashMapIter[struct{}, struct{}])
)

// Collect is a convenience method for [Collect], providing this iterator as
// an argument.
func (iter *LiftHashMapIter[T, U]) Collect() []Tuple[T, U] {
	return Collect[Tuple[T, U]](iter)
}

// ForEach is a convenience method for [ForEach], providing this iterator as an
// argument.
func (iter *LiftHashMapIter[T, U]) ForEach(callback func(Tuple[T, U])) {
	ForEach[Tuple[T, U]](iter, callback)
}

// Find is a convenience method for [Find], providing this iterator as an
// argument.
func (iter *LiftHashMapIter[T, U]) Find(predicate func(Tuple[T, U]) bool) option.Option[Tuple[T, U]] {
	return Find[Tuple[T, U]](iter, predicate)
}

// Drop is a convenience method for [Drop], providing this iterator as an
// argument.
func (iter *LiftHashMapIter[T, U]) Drop(n uint) *DropIter[Tuple[T, U]] {
	return Drop[Tuple[T, U]](iter, n)
}

// Take is a convenience method for [Take], providing this iterator as an
// argument.
func (iter *LiftHashMapIter[T, U]) Take(n uint) *TakeIter[Tuple[T, U]] {
	return Take[Tuple[T, U]](iter, n)
}

// LiftHashMapKeysIter iterator, see [LiftHashMapKeys].
type LiftHashMapKeysIter[T comparable, U any] struct {
	delegate    *LiftHashMapIter[T, U]
	delegateMap *MapIter[Tuple[T, U], T]
	exhausted   bool
}

// LiftHashMapKeys instantiates a [*LiftHashMapKeysIter] that will yield all
// keys in the provided map.
//
// See [LiftHashMap] for information on closing this iterator.
func LiftHashMapKeys[T comparable, U any](hashmap map[T]U) *LiftHashMapKeysIter[T, U] {
	delegate := LiftHashMap(hashmap)

	return &LiftHashMapKeysIter[T, U]{delegate, Map[Tuple[T, U]](delegate, func(pair Tuple[T, U]) T { return pair.One }), false}
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

var (
	_ Iterator[struct{}] = new(LiftHashMapKeysIter[struct{}, struct{}])
	_ io.Closer          = new(LiftHashMapKeysIter[struct{}, struct{}])
)

// Collect is a convenience method for [Collect], providing this iterator as
// an argument.
func (iter *LiftHashMapKeysIter[T, U]) Collect() []T {
	return Collect[T](iter)
}

// ForEach is a convenience method for [ForEach], providing this iterator as an
// argument.
func (iter *LiftHashMapKeysIter[T, U]) ForEach(callback func(T)) {
	ForEach[T](iter, callback)
}

// Find is a convenience method for [Find], providing this iterator as an
// argument.
func (iter *LiftHashMapKeysIter[T, U]) Find(predicate func(T) bool) option.Option[T] {
	return Find[T](iter, predicate)
}

// Drop is a convenience method for [Drop], providing this iterator as an
// argument.
func (iter *LiftHashMapKeysIter[T, U]) Drop(n uint) *DropIter[T] {
	return Drop[T](iter, n)
}

// Take is a convenience method for [Take], providing this iterator as an
// argument.
func (iter *LiftHashMapKeysIter[T, U]) Take(n uint) *TakeIter[T] {
	return Take[T](iter, n)
}

// LiftHashMapValuesIter iterator, see [LiftHashMapValues].
type LiftHashMapValuesIter[T comparable, U any] struct {
	delegate    *LiftHashMapIter[T, U]
	delegateMap *MapIter[Tuple[T, U], U]
	exhausted   bool
}

// LiftHashMapValues instantiates a [*LiftHashMapValuesIter] that will yield
// all keys in the provided map.
//
// See [LiftHashMap] for information on closing this iterator.
func LiftHashMapValues[T comparable, U any](hashmap map[T]U) *LiftHashMapValuesIter[T, U] {
	delegate := LiftHashMap(hashmap)

	return &LiftHashMapValuesIter[T, U]{delegate, Map[Tuple[T, U]](delegate, func(pair Tuple[T, U]) U { return pair.Two }), false}
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

var (
	_ Iterator[struct{}] = new(LiftHashMapValuesIter[struct{}, struct{}])
	_ io.Closer          = new(LiftHashMapValuesIter[struct{}, struct{}])
)

// Collect is a convenience method for [Collect], providing this iterator as
// an argument.
func (iter *LiftHashMapValuesIter[T, U]) Collect() []U {
	return Collect[U](iter)
}

// ForEach is a convenience method for [ForEach], providing this iterator as an
// argument.
func (iter *LiftHashMapValuesIter[T, U]) ForEach(callback func(U)) {
	ForEach[U](iter, callback)
}

// Find is a convenience method for [Find], providing this iterator as an
// argument.
func (iter *LiftHashMapValuesIter[T, U]) Find(predicate func(U) bool) option.Option[U] {
	return Find[U](iter, predicate)
}

// Drop is a convenience method for [Drop], providing this iterator as an
// argument.
func (iter *LiftHashMapValuesIter[T, U]) Drop(n uint) *DropIter[U] {
	return Drop[U](iter, n)
}

// Take is a convenience method for [Take], providing this iterator as an
// argument.
func (iter *LiftHashMapValuesIter[T, U]) Take(n uint) *TakeIter[U] {
	return Take[U](iter, n)
}
