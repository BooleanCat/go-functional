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
	fmt.Println(items.Collect())
	// Output: [1 2 3 4]
}

func ExampleGreaterThan() {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4, 5, 1}), filters.GreaterThan(2))
	fmt.Println(items.Collect())
	// Output: [3 4 5]
}

func ExampleAnd() {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4, 5, 6, 7}), filters.And(
		filters.GreaterThan(2),
		filters.LessThan(7),
	))
	fmt.Println(items.Collect())
	// Output: [3 4 5 6]
}

func ExampleLessThan() {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4, 5, 1}), filters.LessThan(2))
	fmt.Println(items.Collect())
	// Output: [1 1]
}

func TestIsZero(t *testing.T) {
	items := iter.Exclude[int](iter.Lift([]int{1, 2, 3, 0, 4}), filters.IsZero[int])
	assert.SliceEqual(t, items.Collect(), []int{1, 2, 3, 4})
}

func TestIsEven(t *testing.T) {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4}), filters.IsEven[int])
	assert.SliceEqual(t, items.Collect(), []int{2, 4})
}

func TestIsOdd(t *testing.T) {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4}), filters.IsOdd[int])
	assert.SliceEqual(t, items.Collect(), []int{1, 3})
}

func TestGreaterThan(t *testing.T) {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4, 5, 1}), filters.GreaterThan(2))
	assert.SliceEqual(t, items.Collect(), []int{3, 4, 5})
}

func TestLessThan(t *testing.T) {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4, 5, 1}), filters.LessThan(2))
	assert.SliceEqual(t, items.Collect(), []int{1, 1})
}

func TestAnd(t *testing.T) {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4, 5, 6, 7}), filters.And(
		filters.GreaterThan(2),
		filters.LessThan(7),
	))

	assert.SliceEqual(t, items.Collect(), []int{3, 4, 5, 6})
}

func TestAndEmpty(t *testing.T) {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4, 5, 6, 7}), filters.And[int]())

	assert.SliceEqual(t, items.Collect(), []int{1, 2, 3, 4, 5, 6, 7})
}

func TestOr(t *testing.T) {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4, 5, 6, 7}), filters.Or(
		filters.LessThan(3),
		filters.GreaterThan(6),
	))

	assert.SliceEqual(t, items.Collect(), []int{1, 2, 7})
}

func TestOrEmpty(t *testing.T) {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4, 5, 6, 7}), filters.Or[int]())

	assert.SliceEqual(t, items.Collect(), []int{1, 2, 3, 4, 5, 6, 7})
}
