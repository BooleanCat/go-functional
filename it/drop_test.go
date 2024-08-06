package it_test

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/it"
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
