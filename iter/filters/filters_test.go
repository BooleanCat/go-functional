package filters_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/iter/filters"
)

func ExampleIsZero() {
	items := iter.Exclude[int](iter.Lift([]int{1, 2, 3, 0, 4}), filters.IsZero[int])
	fmt.Println(iter.Collect[int](items))
	// Output: [1 2 3 4]
}

func ExampleGreaterThan() {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4, 5, 1}), filters.GreaterThan(2))
	fmt.Println(iter.Collect[int](items))
	// Output: [3 4 5]
}

func ExampleAnd() {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4, 5, 6, 7}), filters.And(
		filters.GreaterThan(2),
		filters.LessThan(7),
	))
	fmt.Println(iter.Collect[int](items))
	// Output: [3 4 5 6]
}

func ExampleLessThan() {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4, 5, 1}), filters.LessThan(2))
	fmt.Println(iter.Collect[int](items))
	// Output: [1 1]
}

func TestIsZero(t *testing.T) {
	items := iter.Exclude[int](iter.Lift([]int{1, 2, 3, 0, 4}), filters.IsZero[int])
	assert.SliceEqual(t, iter.Collect[int](items), []int{1, 2, 3, 4})
}

func TestGreaterThan(t *testing.T) {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4, 5, 1}), filters.GreaterThan(2))
	assert.SliceEqual(t, iter.Collect[int](items), []int{3, 4, 5})
}

func TestLessThan(t *testing.T) {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4, 5, 1}), filters.LessThan(2))
	assert.SliceEqual(t, iter.Collect[int](items), []int{1, 1})
}

func TestAnd(t *testing.T) {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4, 5, 6, 7}), filters.And(
		filters.GreaterThan(2),
		filters.LessThan(7),
	))

	assert.SliceEqual(t, iter.Collect[int](items), []int{3, 4, 5, 6})
}
