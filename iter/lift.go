package iter

import (
	"sync"

	"github.com/BooleanCat/go-functional/option"
)

// LiftIter implements `Lift`. See `Lift`'s documentation.
type LiftIter[T any] struct {
	items []T
	index int
}

// Lift instantiates a `LiftIter` that will yield all items in the provided
// slice.
func Lift[T any](items []T) *LiftIter[T] {
	return &LiftIter[T]{items, 0}
}

// Next implements the Iterator interface for `Lift`.
func (iter *LiftIter[T]) Next() option.Option[T] {
	if iter.index >= len(iter.items) {
		return option.None[T]()
	}

	iter.index++

	return option.Some(iter.items[iter.index-1])
}

var _ Iterator[struct{}] = new(LiftIter[struct{}])

// LiftHashMapIter implements `LiftHashMap`. See `LiftHashMap`'s documentation.
type LiftHashMapIter[T comparable, U any] struct {
	hashmap  map[T]U
	items    chan Tuple[T, U]
	stopOnce sync.Once
	stop     chan struct{}
}

// LiftHashMap instantiates a `LiftHashMapIter` that will yield all items in
// the provided map in the form iter.Tuple[key, value].
//
// Unlike most iterators, `LiftHashMap` should be closed after usage (because
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
	iter := &LiftHashMapIter[T, U]{hashmap, make(chan Tuple[T, U]), sync.Once{}, make(chan struct{})}

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

// Close the iterator. See `LiftHashMap` documentation for details.
func (iter *LiftHashMapIter[T, U]) Close() {
	iter.stopOnce.Do(func() {
		iter.stop <- struct{}{}
		close(iter.stop)
	})
}

// Next implements the Iterator interface for `LiftHashMap`.
func (iter *LiftHashMapIter[T, U]) Next() option.Option[Tuple[T, U]] {
	pair, ok := <-iter.items
	if !ok {
		return option.None[Tuple[T, U]]()
	}

	return option.Some(pair)
}

var _ Iterator[Tuple[struct{}, struct{}]] = new(LiftHashMapIter[struct{}, struct{}])

// LiftHashMapKeysIter implements `LiftHashMapKeys`. See `LiftHashMapKeys`'
// documentation.
type LiftHashMapKeysIter[T comparable, U any] struct {
	delegate    *LiftHashMapIter[T, U]
	delegateMap *MapIter[Tuple[T, U], T]
	exhausted   bool
}

// LiftHashMapKeys instantiates a `LiftHashMapKeysIter` that will yield all
// keys in the provided map.
//
// See documentation on `LiftHashMap` for information on closing this iterator.
func LiftHashMapKeys[T comparable, U any](hashmap map[T]U) *LiftHashMapKeysIter[T, U] {
	delegate := LiftHashMap(hashmap)

	return &LiftHashMapKeysIter[T, U]{delegate, Map[Tuple[T, U]](delegate, func(pair Tuple[T, U]) T { return pair.One }), false}
}

// Close the iterator. See `LiftHashMapKeys` documentation for details.
func (iter *LiftHashMapKeysIter[T, U]) Close() {
	iter.delegate.Close()
}

// Next implements the Iterator interface for `LiftHashMapKeys`.
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

var _ Iterator[struct{}] = new(LiftHashMapKeysIter[struct{}, struct{}])

// LiftHashMapValuesIter implements `LiftHashMapValues`. See
// `LiftHashMapValues`' documentation.
type LiftHashMapValuesIter[T comparable, U any] struct {
	delegate    *LiftHashMapIter[T, U]
	delegateMap *MapIter[Tuple[T, U], U]
	exhausted   bool
}

// LiftHashMapValues instantiates a `LiftHashMapValuesIter` that will yield all
// values in the provided map.
//
// See documentation on `LiftHashMap` for information on closing this iterator.
func LiftHashMapValues[T comparable, U any](hashmap map[T]U) *LiftHashMapValuesIter[T, U] {
	delegate := LiftHashMap(hashmap)

	return &LiftHashMapValuesIter[T, U]{delegate, Map[Tuple[T, U]](delegate, func(pair Tuple[T, U]) U { return pair.Two }), false}
}

// Close the iterator. See `LiftHashMapKeys` documentation for details.
func (iter *LiftHashMapValuesIter[T, U]) Close() {
	iter.delegate.Close()
}

// Next implements the Iterator interface for `LiftHashMapValuesIter`.
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

var _ Iterator[struct{}] = new(LiftHashMapValuesIter[struct{}, struct{}])
