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

	numbers(func(value int) bool {
		return false
	})
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
		a int
		b string
	)
	numbers(func(key int, value string) bool {
		a, b = key, value
		return false
	})
	assert.Equal(t, a, 1)
	assert.Equal(t, b, "one")
}
