package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/fakes"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/iter/filters"
	"github.com/BooleanCat/go-functional/option"
)

func ExampleZip() {
	evens := iter.Filter[int](iter.Count(), filters.IsEven[int])
	odds := iter.Filter[int](iter.Count(), filters.IsOdd[int])

	fmt.Println(iter.Zip[int, int](evens, odds).Take(3).Collect())
	// Output: [{0 1} {2 3} {4 5}]
}

func TestZip(t *testing.T) {
	evens := iter.Filter[int](iter.Count(), filters.IsEven[int])
	odds := iter.Filter[int](iter.Count(), filters.IsOdd[int])

	zipped := iter.Zip[int, int](evens, odds).Take(3).Collect()
	assert.SliceEqual(t, zipped, []iter.Tuple[int, int]{{0, 1}, {2, 3}, {4, 5}})
}

func TestZipFirstExhausted(t *testing.T) {
	evens := iter.Filter[int](iter.Count(), filters.IsEven[int]).Take(2)
	odds := iter.Filter[int](iter.Count(), filters.IsOdd[int])

	zipped := iter.Zip[int, int](evens, odds).Collect()
	assert.SliceEqual(t, zipped, []iter.Tuple[int, int]{{0, 1}, {2, 3}})
}

func TestZipSecondExhausted(t *testing.T) {
	evens := iter.Filter[int](iter.Count(), filters.IsEven[int])
	odds := iter.Filter[int](iter.Count(), filters.IsOdd[int]).Take(2)

	zipped := iter.Zip[int, int](evens, odds).Collect()
	assert.SliceEqual(t, zipped, []iter.Tuple[int, int]{{0, 1}, {2, 3}})
}

func TestZipFirstExhaustedDelegate(t *testing.T) {
	delegate := new(fakes.Iterator[int])
	odds := iter.Filter[int](iter.Count(), filters.IsOdd[int])

	zipped := iter.Zip[int, int](delegate, odds)
	assert.True(t, zipped.Next().IsNone())
	assert.True(t, zipped.Next().IsNone())
	assert.Equal(t, delegate.NextCallCount(), 1)
}

func TestZipSecondExhaustedDelegate(t *testing.T) {
	delegate := new(fakes.Iterator[int])
	odds := iter.Filter[int](iter.Count(), filters.IsOdd[int])

	zipped := iter.Zip[int, int](odds, delegate)
	assert.True(t, zipped.Next().IsNone())
	assert.True(t, zipped.Next().IsNone())
	assert.Equal(t, delegate.NextCallCount(), 1)
}

func TestZipCollect(t *testing.T) {
	evens := iter.Filter[int](iter.Count(), filters.IsEven[int])
	odds := iter.Filter[int](iter.Count(), filters.IsOdd[int]).Take(2)

	items := iter.Zip[int, int](evens, odds).Collect()
	assert.SliceEqual(t, items, []iter.Tuple[int, int]{{0, 1}, {2, 3}})
}

func TestZipForEach(t *testing.T) {
	evens := iter.Filter[int](iter.Count(), filters.IsEven[int])
	odds := iter.Filter[int](iter.Count(), filters.IsOdd[int]).Take(2)

	totalOdd := 0

	iter.Zip[int, int](evens, odds).ForEach(func(pairs iter.Tuple[int, int]) {
		totalOdd += pairs.Two
	})

	assert.Equal(t, totalOdd, 4)
}

func TestZipFind(t *testing.T) {
	evens := iter.Filter[int](iter.Count(), filters.IsEven[int])
	odds := iter.Filter[int](iter.Count(), filters.IsOdd[int]).Take(2)

	item := iter.Zip[int, int](evens, odds).Find(func(pairs iter.Tuple[int, int]) bool {
		return pairs.One == 2
	})

	assert.Equal(t, item, option.Some(iter.Tuple[int, int]{2, 3}))
}

func TestZipDrop(t *testing.T) {
	evens := iter.Filter[int](iter.Count(), filters.IsEven[int])
	odds := iter.Filter[int](iter.Count(), filters.IsOdd[int]).Take(2)

	items := iter.Zip[int, int](evens, odds).Drop(1).Collect()
	assert.SliceEqual(t, items, []iter.Tuple[int, int]{{2, 3}})
}

func TestZipTake(t *testing.T) {
	evens := iter.Filter[int](iter.Count(), filters.IsEven[int])
	odds := iter.Filter[int](iter.Count(), filters.IsOdd[int])

	items := iter.Zip[int, int](evens, odds).Take(1).Collect()
	assert.SliceEqual(t, items, []iter.Tuple[int, int]{{0, 1}})
}
