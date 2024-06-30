package it_test

import (
	"fmt"
	"maps"
	"slices"

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
