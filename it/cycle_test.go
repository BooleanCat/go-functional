package it_test

import (
	"fmt"
	"iter"
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

func ExampleCycle_method() {
	numbers := it.Iterator[int](slices.Values([]int{1, 2})).Cycle().Take(5).Collect()

	fmt.Println(numbers)
	// Output: [1 2 1 2 1]
}

func TestCycleYieldFalse(t *testing.T) {
	t.Parallel()

	numbers := it.Cycle(slices.Values([]int{1, 2}))
	var a int
	numbers(func(value int) bool {
		a = value
		return false
	})
	assert.Equal(t, a, 1)
}

func ExampleCycle2() {
	numbers := maps.Collect(it.Take2(it.Cycle2(maps.All(map[int]string{1: "one"})), 5))

	fmt.Println(numbers)
	// Output: map[1:one]
}

func ExampleCycle2_method() {
	numbers := it.Iterator2[int, string](maps.All(map[int]string{1: "one"})).Cycle().Take(5)

	fmt.Println(maps.Collect(iter.Seq2[int, string](numbers)))
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
