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

	seq(func(v int) bool {
		return false
	})
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

	seq(func(k int, v string) bool {
		return false
	})
}

func ExampleTakeWhile() {
	for number := range it.TakeWhile(slices.Values([]int{1, 2, 3, 4, 5}), filter.LessThan(4)) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func TestTakeWhileYieldsFalse(t *testing.T) {
	t.Parallel()

	seq := it.TakeWhile(slices.Values([]int{1, 2, 3, 4, 5}), filter.LessThan(4))

	seq(func(n int) bool {
		return false
	})
}

func TestTakeWhileEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(it.TakeWhile(it.Exhausted[int](), filter.Passthrough)))
}

func TestTakeWhileNeverTake(t *testing.T) {
	t.Parallel()

	numbers := slices.Collect(it.TakeWhile(slices.Values([]int{1, 2, 3}), func(int) bool { return false }))
	assert.Empty[int](t, numbers)
}

func TestTakeWhileTakeAll(t *testing.T) {
	t.Parallel()

	numbers := slices.Collect(it.TakeWhile(slices.Values([]int{1, 2, 3}), filter.Passthrough))
	assert.SliceEqual(t, []int{1, 2, 3}, numbers)
}

func ExampleTakeWhile2() {
	_, values := it.Collect2(it.TakeWhile2(slices.All([]int{1, 2, 3}), func(i int, v int) bool {
		return v < 3
	}))

	fmt.Println(values)
	// Output: [1 2]
}

func TestTakeWhile2YieldsFalse(t *testing.T) {
	t.Parallel()

	seq := it.TakeWhile2(slices.All([]int{1, 2, 3}), func(i int, v int) bool {
		return v < 3
	})

	seq(func(i int, v int) bool {
		return false
	})
}

func TestTakeWhile2Empty(t *testing.T) {
	t.Parallel()

	_, values := it.Collect2(it.TakeWhile2(it.Exhausted2[int, int](), filter.Passthrough2))
	assert.Empty[int](t, values)
}

func TestTakeWhile2NeverTake(t *testing.T) {
	t.Parallel()

	_, values := it.Collect2(it.TakeWhile2(slices.All([]int{1, 2, 3}), func(int, int) bool { return false }))
	assert.Empty[int](t, values)
}

func TestTakeWhile2TakeAll(t *testing.T) {
	t.Parallel()

	_, values := it.Collect2(it.TakeWhile2(slices.All([]int{1, 2, 3}), filter.Passthrough2))
	assert.SliceEqual(t, []int{1, 2, 3}, values)
}
