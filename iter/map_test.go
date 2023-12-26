package iter_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/fakes"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/iter/ops"
)

func ExampleMap() {
	double := func(a int) int { return a * 2 }
	items := iter.Map[int](iter.Lift([]int{0, 1, 2, 3}), double).Collect()

	fmt.Println(items)
	// Output: [0 2 4 6]
}

func ExampleTransform() {
	addTwo := func(number int) int { return number + 2 }
	numbers := iter.Transform[int](iter.Count(), addTwo).Take(3).Collect()

	fmt.Println(numbers)
	// Output: [2 3 4]
}

func ExampleMapIter_String() {
	fmt.Println(iter.Map[int, string](iter.Count(), strconv.Itoa))
	// Output: Iterator<Map, type=string>
}

func ExampleTransform_method() {
	addTwo := func(number int) int { return number + 2 }
	numbers := iter.Count().Transform(addTwo).Take(3).Collect()

	fmt.Println(numbers)
	// Output: [2 3 4]
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

func TestTransform(t *testing.T) {
	addTwo := func(number int) int { return number + 2 }
	numbers := iter.Transform[int](iter.Count(), addTwo).Take(3).Collect()

	assert.SliceEqual[int](t, numbers, []int{2, 3, 4})
}

func TestMapIter_String(t *testing.T) {
	assert.Equal(
		t,
		iter.Map[int, string](iter.Count(), strconv.Itoa).String(),
		"Iterator<Map, type=string>",
	)
}
