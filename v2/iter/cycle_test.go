package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleCycle() {
	for i := range iter.Cycle(iter.Lift([]int{1, 2})).Take(5) {
		fmt.Println(i)
	}

	// Output:
	// 1
	// 2
	// 1
	// 2
	// 1
}

func ExampleCycle_method() {
	for i := range iter.Lift([]int{1, 2}).Cycle().Take(5) {
		fmt.Println(i)
	}

	// Output:
	// 1
	// 2
	// 1
	// 2
	// 1
}

func TestCycle(t *testing.T) {
	numbers := iter.Count().Take(2).Cycle().Take(5).Collect()
	assert.SliceEqual(t, []int{0, 1, 0, 1, 0}, numbers)
}

func TestCycleEmpty(t *testing.T) {
	numbers := iter.Exhausted[int]().Cycle().Collect()
	assert.Empty[int](t, numbers)
}
