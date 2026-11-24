package it_test

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/it"
)

func ExampleCycle() {
	numbers := slices.Collect(it.Take(it.Cycle(slices.Values([]int{1, 2})), 5))

	fmt.Println(numbers)
	// Output: [1 2 1 2 1]
}

func TestCycleYieldFalse(t *testing.T) {
	t.Parallel()

	numbers := it.Cycle(slices.Values([]int{1, 2}))
	calls := 0

	numbers(func(value int) bool {
		calls++
		return false
	})
	assert.Equal(t, calls, 1)
}

func TestCycleBreak(t *testing.T) {
	t.Parallel()

	for _, limit := range []int{1, 3, 4} {
		t.Run(fmt.Sprintf("%d_values", limit), func(t *testing.T) {
			t.Parallel()

			var values []int
			for value := range it.Cycle(slices.Values([]int{1, 2, 3})) {
				values = append(values, value)
				if len(values) == limit {
					break
				}
			}

			assert.SliceEqual(t, values, []int{1, 2, 3, 1}[:limit])
		})
	}
}

func ExampleCycle2() {
	numbers := maps.Collect(it.Take2(it.Cycle2(maps.All(map[int]string{1: "one"})), 5))

	fmt.Println(numbers)
	// Output: map[1:one]
}

func TestCycle2YieldFalse(t *testing.T) {
	t.Parallel()

	numbers := it.Cycle2(maps.All(map[int]string{1: "one"}))
	var (
		a     int
		b     string
		calls int
	)
	numbers(func(key int, value string) bool {
		a, b = key, value
		calls++
		return false
	})
	assert.Equal(t, a, 1)
	assert.Equal(t, b, "one")
	assert.Equal(t, calls, 1)
}

func TestCycle2Break(t *testing.T) {
	t.Parallel()

	for _, limit := range []int{1, 3, 4} {
		t.Run(fmt.Sprintf("%d_values", limit), func(t *testing.T) {
			t.Parallel()

			var keys []int
			var values []string
			for key, value := range it.Cycle2(slices.All([]string{"one", "two", "three"})) {
				keys = append(keys, key)
				values = append(values, value)
				if len(values) == limit {
					break
				}
			}

			assert.SliceEqual(t, keys, []int{0, 1, 2, 0}[:limit])
			assert.SliceEqual(t, values, []string{"one", "two", "three", "one"}[:limit])
		})
	}
}
