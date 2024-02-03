package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleMap() {
	for i := range iter.Map(iter.Lift([]int{1, 2, 3}), func(i int) int { return i * 2 }) {
		fmt.Println(i)
	}

	// Output:
	// 2
	// 4
	// 6
}

func ExampleMap_method() {
	for i := range iter.Lift([]int{1, 2, 3}).Transform(func(i int) int { return i * 2 }) {
		fmt.Println(i)
	}

	// Output:
	// 2
	// 4
	// 6
}

func TestMap(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3}).Transform(func(i int) int { return i * 2 }).Collect()
	assert.SliceEqual(t, []int{2, 4, 6}, numbers)
}

func TestMapEmpty(t *testing.T) {
	numbers := iter.Exhausted[int]().Transform(func(i int) int { return i * 2 }).Collect()
	assert.Empty[int](t, numbers)
}
