package iter_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/v2/future/maps"
	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleExhausted() {
	fmt.Println(len(slices.Collect(iter.Exhausted[int]())))
	// Output: 0
}

func ExampleExhausted2() {
	fmt.Println(len(maps.Collect(iter.Exhausted2[int, string]())))
	// Output: 0
}
