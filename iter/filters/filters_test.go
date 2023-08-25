package filters_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/iter/filters"
)

func ExampleIsZero() {
	numbers := iter.Exclude[int](iter.Lift([]int{1, 2, 3, 0, 4}), filters.IsZero[int])
	fmt.Println(numbers.Collect())
	// Output: [1 2 3 4]
}

func ExampleGreaterThan() {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5, 1}).Filter(filters.GreaterThan(2))
	fmt.Println(numbers.Collect())
	// Output: [3 4 5]
}

func ExampleGreaterThanEqual() {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5}).Filter(filters.GreaterThanEqual(3))
	fmt.Println(numbers.Collect())
	// Output: [3 4 5]
}

func ExampleAnd() {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5, 6, 7}).Filter(filters.And(
		filters.GreaterThan(2),
		filters.LessThan(7),
	))
	fmt.Println(numbers.Collect())
	// Output: [3 4 5 6]
}

func ExampleOr() {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5, 6, 7}).Filter(filters.Or(
		filters.LessThan(3),
		filters.GreaterThan(6),
	))

	fmt.Println(numbers.Collect())
	// Output: [1 2 7]
}

func ExampleLessThan() {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5, 1}).Filter(filters.LessThan(2))
	fmt.Println(numbers.Collect())
	// Output: [1 1]
}

func ExampleLessThanEqual() {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5}).Filter(filters.LessThanEqual(3))
	fmt.Println(numbers.Collect())
	// Output: [1 2 3]
}

func ExampleIsEven() {
	numbers := iter.Lift([]int{1, 2, 3, 4}).Filter(filters.IsEven[int])
	fmt.Println(numbers.Collect())
	// Output: [2 4]
}

func ExampleIsOdd() {
	numbers := iter.Lift([]int{1, 2, 3, 4}).Filter(filters.IsOdd[int])
	fmt.Println(numbers.Collect())
	// Output: [1 3]
}

func TestIsZero(t *testing.T) {
	numbers := iter.Exclude[int](iter.Lift([]int{1, 2, 3, 0, 4}), filters.IsZero[int])
	assert.SliceEqual(t, numbers.Collect(), []int{1, 2, 3, 4})
}

func TestIsEven(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3, 4}).Filter(filters.IsEven[int])
	assert.SliceEqual(t, numbers.Collect(), []int{2, 4})
}

func TestIsOdd(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3, 4}).Filter(filters.IsOdd[int])
	assert.SliceEqual(t, numbers.Collect(), []int{1, 3})
}

func TestGreaterThan(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5, 1}).Filter(filters.GreaterThan(2))
	assert.SliceEqual(t, numbers.Collect(), []int{3, 4, 5})
}

func TestGreaterThanEqual(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5}).Filter(filters.GreaterThanEqual(3))
	assert.SliceEqual(t, numbers.Collect(), []int{3, 4, 5})
}

func TestLessThan(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5, 1}).Filter(filters.LessThan(2))
	assert.SliceEqual(t, numbers.Collect(), []int{1, 1})
}

func TestLessThanEqual(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5}).Filter(filters.LessThanEqual(3))
	assert.SliceEqual(t, numbers.Collect(), []int{1, 2, 3})
}

func TestAnd(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5, 6, 7}).Filter(filters.And(
		filters.GreaterThan(2),
		filters.LessThan(7),
	))

	assert.SliceEqual(t, numbers.Collect(), []int{3, 4, 5, 6})
}

func TestAndEmpty(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5, 6, 7}).Filter(filters.And[int]())
	assert.SliceEqual(t, numbers.Collect(), []int{1, 2, 3, 4, 5, 6, 7})
}

func TestOr(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5, 6, 7}).Filter(filters.Or(
		filters.LessThan(3),
		filters.GreaterThan(6),
	))

	assert.SliceEqual(t, numbers.Collect(), []int{1, 2, 7})
}

func TestOrEmpty(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5, 6, 7}).Filter(filters.Or[int]())
	assert.SliceEqual(t, numbers.Collect(), []int{1, 2, 3, 4, 5, 6, 7})
}
