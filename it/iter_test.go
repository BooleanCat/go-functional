package it_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/BooleanCat/go-functional/v2/it/op"
)

func ExampleForEach() {
	it.ForEach(slices.Values([]int{1, 2, 3}), func(number int) {
		fmt.Println(number)
	})
	// Output:
	// 1
	// 2
	// 3
}

func TestForEachEmpty(t *testing.T) {
	t.Parallel()

	it.ForEach(slices.Values([]int{}), func(int) {
		t.Error("unexpected")
	})
}

func ExampleForEach2() {
	it.ForEach2(it.Enumerate(slices.Values([]int{1, 2, 3})), func(index int, number int) {
		fmt.Println(index, number)
	})
	// Output:
	// 0 1
	// 1 2
	// 2 3
}

func TestForEach2Empty(t *testing.T) {
	t.Parallel()

	it.ForEach2(it.Enumerate(it.Exhausted[int]()), func(int, int) {
		t.Error("unexpected")
	})
}

func ExampleReduce() {
	fmt.Println(it.Reduce(slices.Values([]int{1, 2, 3}), op.Add, 0))
	// Output: 6
}

func TestReduceEmpty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, it.Reduce(it.Exhausted[int](), func(int, int) int { return 0 }, 0), 0)
}

func ExampleReduce2() {
	fmt.Println(it.Reduce2(it.Enumerate(slices.Values([]int{1, 2, 3})), func(i, a, b int) int {
		return i + 1
	}, 0))

	// Output: 3
}

func ExampleMax() {
	max, ok := it.Max(slices.Values([]int{1, 2, 3}))
	fmt.Println(max, ok)

	// Output: 3 true
}

func TestMaxEmpty(t *testing.T) {
	t.Parallel()

	max, ok := it.Max(it.Exhausted[int]())
	assert.Equal(t, max, 0)
	assert.False(t, ok)
}

func ExampleMin() {
	min, ok := it.Min(slices.Values([]int{4, 2, 1, 3}))
	fmt.Println(min, ok)

	// Output: 1 true
}

func TestMinEmpty(t *testing.T) {
	t.Parallel()

	min, ok := it.Min(it.Exhausted[int]())
	assert.Equal(t, min, 0)
	assert.False(t, ok)
}
