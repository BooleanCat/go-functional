package it_test

import (
	"fmt"
	"slices"

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
