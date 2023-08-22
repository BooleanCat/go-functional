package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/fakes"
	"github.com/BooleanCat/go-functional/iter"
)

func ExampleCycle() {
	numbers := iter.Cycle[int](iter.Lift([]int{1, 2})).Take(5)
	fmt.Println(numbers.Collect())
	// Output: [1 2 1 2 1]
}

func TestCycleIter(t *testing.T) {
	items := iter.Cycle[int](iter.Lift([]int{1, 2, 3})).Take(2).Collect()
	assert.SliceEqual(t, items, []int{1, 2})
}

func TestCycleIterOverflow(t *testing.T) {
	items := iter.Cycle[int](iter.Lift([]int{1, 2})).Take(5).Collect()
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

func TestCycleForEach(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "oops")
	}()

	iter.Cycle[int](iter.Lift([]int{1, 2})).ForEach(func(_ int) {
		panic("oops")
	})

	t.Error("did not panic")
}

func TestCycleDrop(t *testing.T) {
	items := iter.Cycle[int](iter.Lift([]int{1, 2})).Drop(1).Take(5).Collect()
	assert.SliceEqual(t, items, []int{2, 1, 2, 1, 2})
}

func TestCycleTake(t *testing.T) {
	items := iter.Cycle[int](iter.Lift([]int{1, 2})).Take(5).Collect()
	assert.SliceEqual(t, items, []int{1, 2, 1, 2, 1})
}
