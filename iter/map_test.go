package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/fakes"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/iter/ops"
)

func ExampleMap() {
	double := func(a int) int { return a * 2 }
	items := iter.Collect[int](iter.Map[int](iter.Lift([]int{0, 1, 2, 3}), double))

	fmt.Println(items)
	// Output: [0 2 4 6]
}

func TestMap(t *testing.T) {
	double := func(a int) int { return a * 2 }
	items := iter.Collect[int](iter.Take[int](
		iter.Map[int](iter.Count(), double),
		4,
	))
	assert.SliceEqual(t, items, []int{0, 2, 4, 6})
}

func TestMapEmpty(t *testing.T) {
	double := func(a int) int { return a * 2 }
	items := iter.Collect[int](iter.Map[int](iter.Exhausted[int](), double))
	assert.Empty[int](t, items)
}

func TestMapExhausted(t *testing.T) {
	delegate := new(fakes.Iterator[int])
	iter := iter.Map[int](delegate, ops.Passthrough[int])

	assert.True(t, iter.Next().IsNone())
	assert.True(t, iter.Next().IsNone())
	assert.Equal(t, delegate.NextCallCount(), 1)
}
