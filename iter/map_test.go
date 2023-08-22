package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/fakes"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/iter/ops"
	"github.com/BooleanCat/go-functional/option"
)

func ExampleMap() {
	double := func(a int) int { return a * 2 }
	items := iter.Map[int](iter.Lift([]int{0, 1, 2, 3}), double).Collect()

	fmt.Println(items)
	// Output: [0 2 4 6]
}

func TestMap(t *testing.T) {
	double := func(a int) int { return a * 2 }
	items := iter.Map[int](iter.Count(), double).Take(4).Collect()
	assert.SliceEqual(t, items, []int{0, 2, 4, 6})
}

func TestMapEmpty(t *testing.T) {
	double := func(a int) int { return a * 2 }
	items := iter.Map[int](iter.Exhausted[int](), double).Collect()
	assert.Empty[int](t, items)
}

func TestMapExhausted(t *testing.T) {
	delegate := new(fakes.Iterator[int])
	iter := iter.Map[int](delegate, ops.Passthrough[int])

	assert.True(t, iter.Next().IsNone())
	assert.True(t, iter.Next().IsNone())
	assert.Equal(t, delegate.NextCallCount(), 1)
}

func TestMapCollect(t *testing.T) {
	double := func(a int) int { return a * 2 }
	items := iter.Map[int, int](iter.Lift([]int{0, 1, 2, 3}), double).Collect()
	assert.SliceEqual(t, items, []int{0, 2, 4, 6})
}

func TestMapForEach(t *testing.T) {
	total := 0

	double := func(a int) int { return a * 2 }

	iter.Map[int, int](iter.Lift([]int{0, 1, 2, 3}), double).ForEach(func(number int) {
		total += number
	})

	assert.Equal(t, total, 12)
}

func TestMapFind(t *testing.T) {
	double := func(a int) int { return a * 2 }
	assert.Equal(t, iter.Map[int, int](iter.Lift([]int{0, 1, 2, 3}), double).Find(func(number int) bool {
		return number == 4
	}), option.Some(4))
}

func TestMapDrop(t *testing.T) {
	double := func(a int) int { return a * 2 }
	items := iter.Map[int, int](iter.Lift([]int{0, 1, 2, 3}), double).Drop(2).Collect()
	assert.SliceEqual(t, items, []int{4, 6})
}

func TestMapTake(t *testing.T) {
	double := func(a int) int { return a * 2 }
	items := iter.Map[int, int](iter.Lift([]int{0, 1, 2, 3}), double).Take(2).Collect()
	assert.SliceEqual(t, items, []int{0, 2})
}
