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
	var a int
	numbers(func(value int) bool {
		a = value
		return false
	})
	assert.Equal(t, a, 3)
}

func TestDropEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(it.Drop(it.Exhausted[int](), 2)))
}

func ExampleDrop2() {
	numbers := maps.Collect(it.Drop2(maps.All(map[int]string{1: "one", 2: "two", 3: "three"}), 1))

	fmt.Println(len(numbers))
	// Output: 2
}

func TestDrop2(t *testing.T) {
	t.Parallel()

	keys := []int{1, 2, 3}
	values := []string{"one", "two", "three"}

	numbers := maps.Collect(it.Drop2(it.Zip(slices.Values(keys), slices.Values(values)), 1))

	assert.Equal(t, len(numbers), 2)

	for key := range numbers {
		assert.True(t, slices.Contains(keys, key))
	}
}

func TestDrop2Empty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(it.Drop2(it.Exhausted2[int, int](), 1))), 0)
}

func TestDrop2Zero(t *testing.T) {
	t.Parallel()

	numbers := maps.Collect(it.Drop2(maps.All(map[int]string{1: "one", 2: "two", 3: "three"}), 0))

	assert.Equal(t, len(numbers), 3)
}

func TestDrop2YieldFalse(t *testing.T) {
	t.Parallel()

	numbersZipped := it.Zip(slices.Values([]int{1, 2, 3}), slices.Values([]int{3, 4, 5}))
	numbers := it.Drop2(numbersZipped, 2)
	var a, b int
	numbers(func(v, w int) bool {
		a, b = v, w
		return false
	})
	assert.Equal(t, a, 3)
	assert.Equal(t, b, 5)
}
