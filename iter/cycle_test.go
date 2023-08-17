package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/fakes"
	"github.com/BooleanCat/go-functional/iter"
)

func ExampleCycle() {
	numbers := iter.Take[int](iter.Cycle[int](iter.Lift([]int{1, 2})), 5)
	fmt.Println(numbers.Collect())
	// Output: [1 2 1 2 1]
}

func TestCycleIter(t *testing.T) {
	items := iter.Take[int](iter.Cycle[int](iter.Lift([]int{1, 2, 3})), 2).Collect()
	assert.SliceEqual(t, items, []int{1, 2})
}

func TestCycleIterOverflow(t *testing.T) {
	items := iter.Take[int](iter.Cycle[int](iter.Lift([]int{1, 2})), 5).Collect()
	assert.SliceEqual(t, items, []int{1, 2, 1, 2, 1})
}

func TestCycleIterEmpty(t *testing.T) {
	items := iter.Cycle[int](iter.Exhausted[int]())
	assert.True(t, items.Next().IsNone())
}

func TestCycleExhausted(t *testing.T) {
	delegate := new(fakes.Iterator[int])
	ones := iter.Cycle[int](delegate)

	assert.True(t, ones.Next().IsNone())
	assert.True(t, ones.Next().IsNone())
	assert.Equal(t, delegate.NextCallCount(), 1)
}

func TestCycleDrop(t *testing.T) {
	items := iter.Take[int](iter.Cycle[int](iter.Lift([]int{1, 2})).Drop(1), 5).Collect()
	assert.SliceEqual(t, items, []int{2, 1, 2, 1, 2})
}
