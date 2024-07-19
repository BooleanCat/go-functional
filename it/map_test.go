package it_test

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/it"
)

func ExampleMap() {
	double := func(n int) int { return n * 2 }

	for number := range it.Map(slices.Values([]int{1, 2, 3}), double) {
		fmt.Println(number)
	}

	// Output:
	// 2
	// 4
	// 6
}

func TestMapEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(it.Map(it.Exhausted[int](), func(int) int { return 0 })))
}

func TestMapYieldFalse(t *testing.T) {
	t.Parallel()

	numbers := it.Map(slices.Values([]int{1, 2, 3, 4, 5}), func(a int) int { return a + 1 })

	values := []int{}
	numbers(func(v int) bool {
		values = append(values, v)
		return false
	})

	assert.SliceEqual(t, []int{2}, values)
}

func ExampleMap2() {
	doubleBoth := func(n, m int) (int, int) {
		return n * 2, m * 2
	}

	pairs := it.Zip(slices.Values([]int{1, 2, 3}), slices.Values([]int{2, 3, 4}))

	for left, right := range it.Map2(pairs, doubleBoth) {
		fmt.Println(left, right)
	}

	// Output:
	// 2 4
	// 4 6
	// 6 8
}

func TestMap2Empty(t *testing.T) {
	t.Parallel()

	doubleBoth := func(n, m int) (int, int) { return n * 2, m * 2 }

	assert.Equal(t, len(maps.Collect(it.Map2(it.Exhausted2[int, int](), doubleBoth))), 0)
}

func TestMap2YieldFalse(t *testing.T) {
	t.Parallel()

	pairs := it.Zip(slices.Values([]int{1, 2, 3}), slices.Values([]int{3, 4, 5}))

	numbers := it.Map2(pairs, func(a, b int) (int, int) {
		return a + 1, b + 2
	})

	numbers(func(v, w int) bool {
		return false
	})
}
