package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleTake() {
	for i := range iter.Take(iter.Lift([]int{1, 2, 3}), 2) {
		fmt.Println(i)
	}

	// Output:
	// 1
	// 2
}

func ExampleTake_method() {
	for i := range iter.Lift([]int{1, 2, 3}).Take(2) {
		fmt.Println(i)
	}

	// Output:
	// 1
	// 2
}

func TestTake(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3}).Take(2).Collect()
	assert.SliceEqual(t, []int{1, 2}, numbers)
}

func TestTakeEmpty(t *testing.T) {
	numbers := iter.Lift([]int{}).Take(2).Collect()
	assert.Empty[int](t, numbers)
}

func TestTakeZero(t *testing.T) {
	numbers := iter.Lift([]int{1, 2}).Take(0).Collect()
	assert.Empty[int](t, numbers)
}

func TestTakeNegative(t *testing.T) {
	numbers := iter.Lift([]int{1, 2}).Take(-2).Collect()
	assert.Empty[int](t, numbers)
}
