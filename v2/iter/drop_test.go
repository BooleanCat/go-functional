package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleDrop() {
	for i := range iter.Drop(iter.Lift([]int{1, 2, 3}), 1) {
		fmt.Println(i)
	}

	// Output:
	// 2
	// 3
}

func ExampleDrop_method() {
	for i := range iter.Lift([]int{1, 2, 3}).Drop(1) {
		fmt.Println(i)
	}

	// Output:
	// 2
	// 3
}

func TestDrop(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3}).Drop(1).Collect()
	assert.SliceEqual(t, []int{2, 3}, numbers)
}

func TestDropEmpty(t *testing.T) {
	numbers := iter.Exhausted[int]().Drop(1).Collect()
	assert.Empty[int](t, numbers)
}

func TestDropZero(t *testing.T) {
	numbers := iter.Lift([]int{1, 2}).Drop(0).Collect()
	assert.SliceEqual(t, []int{1, 2}, numbers)
}

func TestDropNegative(t *testing.T) {
	numbers := iter.Lift([]int{1, 2}).Drop(-2).Collect()
	assert.SliceEqual(t, []int{1, 2}, numbers)
}
