package it_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/it"
)

func ExampleEnumerate() {
	for index, value := range it.Enumerate(slices.Values([]int{1, 2, 3})) {
		fmt.Println(index, value)
	}

	// Output:
	// 0 1
	// 1 2
	// 2 3
}

func TestEnumerateYieldFalse(t *testing.T) {
	t.Parallel()

	iterator := it.Enumerate(slices.Values([]int{1, 2, 3, 4, 5}))

	iterator(func(i int, n int) bool {
		return false
	})
}
