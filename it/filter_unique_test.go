package it_test

import (
	"fmt"
	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/it"
	"slices"
	"testing"
)

func ExampleFilterUnique() {
	for number := range it.FilterUnique(slices.Values([]int{1, 2, 2, 3, 3, 3, 4, 5})) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func TestFilterUniqueEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(it.Chain[int]()))
}

func TestFilterUniqueYieldFalse(t *testing.T) {
	t.Parallel()

	iterator := it.FilterUnique(slices.Values([]int{100, 200, 300}))

	var value int
	iterator(func(v int) bool {
		value = v
		return false
	})
	assert.Equal(t, 100, value)
}

func TestFilterUniqueWithNoDuplicates(t *testing.T) {
	t.Parallel()

	numbers := slices.Collect(it.FilterUnique(slices.Values([]int{1, 2, 3})))
	assert.SliceEqual(t, []int{1, 2, 3}, numbers)
}

func TestFilterUniqueWithDuplicates(t *testing.T) {
	t.Parallel()

	strings := slices.Collect(it.FilterUnique(slices.Values([]string{"hello", "world", "hello", "world", "hello"})))
	assert.SliceEqual(t, []string{"hello", "world"}, strings)
}
