package filter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
	"github.com/BooleanCat/go-functional/v2/iter/filter"
)

func ExampleIsZero() {
	numbers := iter.Exclude[int](iter.Lift([]int{1, 2, 3, 0, 4}), filter.IsZero)
	fmt.Println(numbers.Collect())
	// Output: [1 2 3 4]
}

func TestIsZero(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3, 0, 4}).Exclude(filter.IsZero).Collect()
	assert.SliceEqual(t, []int{1, 2, 3, 4}, numbers)
}

func ExampleIsEven() {
	numbers := iter.Exclude[int](iter.Lift([]int{1, 2, 3, 4, 5}), filter.IsEven)
	fmt.Println(numbers.Collect())
	// Output: [1 3 5]
}

func TestIsEven(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5}).Exclude(filter.IsEven).Collect()
	assert.SliceEqual(t, []int{1, 3, 5}, numbers)
}

func ExampleIsOdd() {
	numbers := iter.Exclude[int](iter.Lift([]int{1, 2, 3, 4, 5}), filter.IsOdd)
	fmt.Println(numbers.Collect())
	// Output: [2 4]
}

func TestIsOdd(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5}).Exclude(filter.IsOdd).Collect()
	assert.SliceEqual(t, []int{2, 4}, numbers)
}

func ExampleGreaterThan() {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5}).Filter(filter.GreaterThan(2))
	fmt.Println(numbers.Collect())
	// Output: [3 4 5]
}

func TestGreaterThan(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5}).Filter(filter.GreaterThan(2)).Collect()
	assert.SliceEqual(t, []int{3, 4, 5}, numbers)
}

func ExampleLessThan() {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5}).Filter(filter.LessThan(3))
	fmt.Println(numbers.Collect())
	// Output: [1 2]
}

func TestLessThan(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5}).Filter(filter.LessThan(3)).Collect()
	assert.SliceEqual(t, []int{1, 2}, numbers)
}

func ExampleGreaterThanEqual() {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5}).Filter(filter.GreaterThanEqual(3))
	fmt.Println(numbers.Collect())
	// Output: [3 4 5]
}

func TestGreaterThanEqual(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5}).Filter(filter.GreaterThanEqual(3)).Collect()
	assert.SliceEqual(t, []int{3, 4, 5}, numbers)
}

func ExampleLessThanEqual() {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5}).Filter(filter.LessThanEqual(3))
	fmt.Println(numbers.Collect())
	// Output: [1 2 3]
}

func TestLessThanEqual(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3, 4, 5}).Filter(filter.LessThanEqual(3)).Collect()
	assert.SliceEqual(t, []int{1, 2, 3}, numbers)
}
