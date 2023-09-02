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
	// Output:[(0, 1) (2, 3) (4, 5)]
}

func TestPairStringer(t *testing.T) {
	foo := map[string]interface{}{
		"text": "Random Text",
	}
	pair1 := iter.Pair[string, interface{}]{One: "1", Two: foo}
	pair2 := iter.Pair[int, interface{}]{One: 2, Two: pair1}

	assert.Equal(t, pair1.String(), "(1, map[text:Random Text])")
	assert.Equal(t, pair2.String(), "(2, (1, map[text:Random Text]))")
}

func TestZip(t *testing.T) {
	evens := iter.Filter[int](iter.Count(), filters.IsEven[int])
	odds := iter.Filter[int](iter.Count(), filters.IsOdd[int])

	zipped := iter.Zip[int, int](evens, odds).Take(3).Collect()
	assert.SliceEqual(t, zipped, []iter.Pair[int, int]{{0, 1}, {2, 3}, {4, 5}})
}

func TestZipFirstExhausted(t *testing.T) {
	evens := iter.Filter[int](iter.Count(), filters.IsEven[int]).Take(2)
	odds := iter.Filter[int](iter.Count(), filters.IsOdd[int])

	zipped := iter.Zip[int, int](evens, odds).Collect()
	assert.SliceEqual(t, zipped, []iter.Pair[int, int]{{0, 1}, {2, 3}})
}

func TestZipSecondExhausted(t *testing.T) {
	evens := iter.Filter[int](iter.Count(), filters.IsEven[int])
	odds := iter.Filter[int](iter.Count(), filters.IsOdd[int]).Take(2)

	zipped := iter.Zip[int, int](evens, odds).Collect()
	assert.SliceEqual(t, zipped, []iter.Pair[int, int]{{0, 1}, {2, 3}})
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
