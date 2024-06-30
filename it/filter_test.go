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

func ExampleFilter() {
	for number := range it.Filter(slices.Values([]int{1, 2, 3, 4, 5}), filter.IsEven) {
		fmt.Println(number)
	}

	// Output:
	// 2
	// 4
}

func TestFilterEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(it.Filter(it.Exhausted[int](), filter.IsEven)))
}

func TestFilterYieldsFalse(t *testing.T) {
	t.Parallel()

	seq := it.Filter(slices.Values([]int{1, 2, 3, 4, 5}), filter.IsEven)

	var value int

	seq(func(v int) bool {
		value = v
		return false
	})

	assert.Equal(t, value, 2)
}

func ExampleExclude() {
	for number := range it.Exclude(slices.Values([]int{1, 2, 3, 4, 5}), filter.IsEven) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 3
	// 5
}

func ExampleFilter2() {
	isOne := func(n int, _ string) bool { return n == 1 }
	numbers := map[int]string{1: "one", 2: "two", 3: "three"}

	for key, value := range it.Filter2(maps.All(numbers), isOne) {
		fmt.Println(key, value)
	}

	// Output: 1 one
}

func TestFilter2Empty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(it.Filter2(it.Exhausted2[int, int](), filter.Passthrough2))), 0)
}

func TestFilter2YieldsFalse(t *testing.T) {
	t.Parallel()

	seq := it.Filter2(maps.All(map[int]string{1: "one", 2: "two"}), filter.Passthrough2)

	var (
		key   int
		value string
	)

	seq(func(k int, v string) bool {
		key = k
		value = v
		return false
	})

	assert.Equal(t, key, 1)
	assert.Equal(t, value, "one")

}

func ExampleExclude2() {
	isOne := func(n int, _ string) bool { return n == 1 }
	numbers := map[int]string{1: "one", 3: "three"}

	for key, value := range it.Exclude2(maps.All(numbers), isOne) {
		fmt.Println(key, value)
	}

	// Output: 3 three
}

func TestExclude2Empty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(it.Exclude2(it.Exhausted2[int, int](), filter.Passthrough2))), 0)
}
