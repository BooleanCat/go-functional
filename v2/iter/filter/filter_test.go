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
