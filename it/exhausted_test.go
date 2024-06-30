package it_test

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/it"
)

func ExampleExhausted() {
	fmt.Println(len(slices.Collect(it.Exhausted[int]())))
	// Output: 0
}

func ExampleExhausted2() {
	fmt.Println(len(maps.Collect(it.Exhausted2[int, string]())))
	// Output: 0
}

func TestEnumerateYieldFalse(t *testing.T) {
	t.Parallel()

	iterator := it.Enumerate(slices.Values([]int{1, 2, 3, 4, 5}))

	var (
		index  int
		number int
	)

	iterator(func(i int, n int) bool {
		index = i
		number = n
		return false
	})

	assert.Equal(t, index, 0)
	assert.Equal(t, number, 1)
}
