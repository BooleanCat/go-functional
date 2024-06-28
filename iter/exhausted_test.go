package iter_test

import (
	"fmt"
	"maps"
	"slices"

	fn "github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleExhausted() {
	fmt.Println(len(slices.Collect(fn.Exhausted[int]())))
	// Output: 0
}

func ExampleExhausted2() {
	fmt.Println(len(maps.Collect(fn.Exhausted2[int, string]())))
	// Output: 0
}
