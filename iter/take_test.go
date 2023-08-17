package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/fakes"
	"github.com/BooleanCat/go-functional/iter"
)

func ExampleTake() {
	iter := iter.Take[int](iter.Count(), 2)
	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	// Output:
	// Some(0)
	// Some(1)
	// None
}

func TestTakeIter(t *testing.T) {
	iter := iter.Take[int](iter.Count(), 2)
	assert.Equal(t, iter.Next().Unwrap(), 0)
	assert.Equal(t, iter.Next().Unwrap(), 1)
	assert.True(t, iter.Next().IsNone())
}

func TestTakeIterEmpty(t *testing.T) {
	iter := iter.Take[int](iter.Count(), 0)
	assert.True(t, iter.Next().IsNone())
}

func TestTakeExhausted(t *testing.T) {
	delegate := new(fakes.Iterator[int])
	iter := iter.Take[int](delegate, 10)

	assert.True(t, iter.Next().IsNone())
	assert.True(t, iter.Next().IsNone())
	assert.Equal(t, delegate.NextCallCount(), 1)
}

func TestTakeCollect(t *testing.T) {
	items := iter.Take[int](iter.Count(), 3).Collect()
	assert.SliceEqual(t, items, []int{0, 1, 2})
}

func TestTakeDrop(t *testing.T) {
	items := iter.Take[int](iter.Count(), 3).Drop(1).Collect()
	assert.SliceEqual(t, items, []int{1, 2})
}
