package it_test

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/it"
)

func ExampleTake() {
	for number := range it.Take(slices.Values([]int{1, 2, 3, 4, 5}), 3) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func TestTakeZero(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(it.Take(slices.Values([]int{1, 2, 3}), 0)))
}

func TestTakeMoreThanAvailable(t *testing.T) {
	t.Parallel()

	numbers := slices.Collect(it.Take(slices.Values([]int{1, 2, 3}), 5))
	assert.SliceEqual(t, []int{1, 2, 3}, numbers)
}

func TestTakeYieldFalse(t *testing.T) {
	t.Parallel()

	seq := it.Take(slices.Values([]int{1, 2, 3, 4, 5}), 3)

	values := []int{}
	seq(func(v int) bool {
		values = append(values, v)
		return false
	})

	expected := []int{1}
	assert.SliceEqual(t, expected, values)
}

func TestTakeEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(it.Take(it.Exhausted[int](), 2)))
}

func ExampleTake2() {
	numbers := maps.Collect(it.Take2(maps.All(map[int]string{1: "one", 2: "two", 3: "three"}), 2))

	fmt.Println(len(numbers))
	// Output: 2
}

func TestTake2(t *testing.T) {
	t.Parallel()

	keys := []int{1, 2, 3}
	values := []string{"one", "two", "three"}

	numbers := maps.Collect(it.Take2(it.Zip(slices.Values(keys), slices.Values(values)), 2))

	assert.Equal(t, len(numbers), 2)

	for key := range numbers {
		assert.True(t, slices.Contains(keys, key))
	}
}

func TestTake2Zero(t *testing.T) {
	t.Parallel()

	numbers := maps.Collect(it.Take2(maps.All(map[int]string{1: "one", 2: "two", 3: "three"}), 0))
	assert.Equal(t, len(numbers), 0)
}

func TestTake2Empty(t *testing.T) {
	t.Parallel()

	numbers := maps.Collect(it.Take2(it.Exhausted2[int, int](), 2))
	assert.Equal(t, len(numbers), 0)
}

func TestTake2MoreThanAvailable(t *testing.T) {
	t.Parallel()

	numbers := maps.Collect(it.Take2(maps.All(map[int]string{1: "one", 2: "two"}), 3))
	assert.Equal(t, len(numbers), 2)
}

func TestTake2YieldFalse(t *testing.T) {
	t.Parallel()

	seq := it.Take2(maps.All(map[int]string{1: "one", 2: "two", 3: "three"}), 2)

	values := make(map[int]string)
	seq(func(k int, v string) bool {
		values[k] = v
		return false
	})

	assert.Equal(t, len(values), 1)
}
