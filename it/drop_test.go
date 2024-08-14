package it_test

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/BooleanCat/go-functional/v2/it/filter"
)

func ExampleDrop() {
	for value := range it.Drop(slices.Values([]int{1, 2, 3, 4, 5}), 2) {
		fmt.Println(value)
	}

	// Output:
	// 3
	// 4
	// 5
}

func TestDropYieldFalse(t *testing.T) {
	t.Parallel()

	numbers := it.Drop(slices.Values([]int{1, 2, 3, 4, 5}), 2)

	numbers(func(value int) bool {
		return false
	})
}

func TestDropEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(it.Drop(it.Exhausted[int](), 2)))
}

func ExampleDrop2() {
	_, numbers := it.Collect2(it.Drop2(slices.All([]string{"zero", "one", "two"}), 1))

	fmt.Println(numbers)
	// Output: [one two]
}

func TestDrop2(t *testing.T) {
	t.Parallel()

	values := []string{"zero", "one", "two"}

	indices, values := it.Collect2(it.Drop2(slices.All(values), 1))

	assert.SliceEqual(t, indices, []int{1, 2})
	assert.SliceEqual(t, values, []string{"one", "two"})
}

func TestDrop2Empty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(it.Drop2(it.Exhausted2[int, int](), 1))), 0)
}

func TestDrop2Zero(t *testing.T) {
	t.Parallel()

	indices, values := it.Collect2(it.Drop2(slices.All([]string{"zero", "one", "two"}), 0))

	assert.SliceEqual(t, indices, []int{0, 1, 2})
	assert.SliceEqual(t, values, []string{"zero", "one", "two"})
}

func TestDrop2YieldFalse(t *testing.T) {
	t.Parallel()

	numbers := it.Drop2(slices.All([]int{1, 2, 3}), 2)

	numbers(func(v, w int) bool {
		return false
	})
}

func ExampleDropWhile() {
	for value := range it.DropWhile(slices.Values([]int{1, 2, 3, 4, 1}), filter.LessThan(3)) {
		fmt.Println(value)
	}

	// Output:
	// 3
	// 4
	// 1
}

func TestDropWhileYieldFalse(t *testing.T) {
	t.Parallel()

	numbers := it.DropWhile(slices.Values([]int{1, 2, 3, 4, 1}), filter.LessThan(3))

	numbers(func(v int) bool {
		return false
	})
}

func TestDropWhileEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(it.DropWhile(it.Exhausted[int](), filter.Passthrough)))
}

func TestDropWhileNeverDrop(t *testing.T) {
	t.Parallel()

	numbers := slices.Collect(it.DropWhile(slices.Values([]int{1, 2, 3}), func(int) bool { return false }))
	assert.SliceEqual(t, []int{1, 2, 3}, numbers)
}

func TestDropWhileDropAll(t *testing.T) {
	t.Parallel()

	numbers := slices.Collect(it.DropWhile(slices.Values([]int{1, 2, 3}), filter.LessThan(4)))
	assert.Empty[int](t, numbers)
}

func ExampleDropWhile2() {
	_, values := it.Collect2(it.DropWhile2(slices.All([]int{1, 2, 3}), func(int, v int) bool {
		return v < 3
	}))

	fmt.Println(values)
	// Output: [3]
}

func TestDropWhile2YieldFalse(t *testing.T) {
	t.Parallel()

	numbers := it.DropWhile2(slices.All([]int{1, 2, 3}), func(int, v int) bool {
		return v < 3
	})

	numbers(func(int, v int) bool {
		return false
	})
}

func TestDropWhile2Empty(t *testing.T) {
	t.Parallel()

	_, values := it.Collect2(it.DropWhile2(it.Exhausted2[int, int](), filter.Passthrough2))

	assert.Empty[int](t, values)
}

func TestDropWhile2NeverDrop(t *testing.T) {
	t.Parallel()

	_, values := it.Collect2(it.DropWhile2(slices.All([]int{1, 2, 3}), func(int, int) bool { return false }))
	assert.SliceEqual(t, []int{1, 2, 3}, values)
}

func TestDropWhile2DropAll(t *testing.T) {
	t.Parallel()

	_, values := it.Collect2(it.DropWhile2(slices.All([]int{1, 2, 3}), func(int, int) bool { return true }))
	assert.Empty[int](t, values)
}
