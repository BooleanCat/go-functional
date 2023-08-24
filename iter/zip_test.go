package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/fakes"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/iter/filters"
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
