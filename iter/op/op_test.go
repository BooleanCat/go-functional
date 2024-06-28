package op_test

import (
	"fmt"
	"slices"

	fn "github.com/BooleanCat/go-functional/v2/iter"
	"github.com/BooleanCat/go-functional/v2/iter/op"
)

func ExampleAdd() {
	fmt.Println(fn.Reduce(slices.Values([]int{1, 2, 3}), op.Add, 0))
	// Output: 6
}

func ExampleAdd_string() {
	fmt.Println(fn.Reduce(slices.Values([]string{"a", "b", "c"}), op.Add, ""))
	// Output: abc
}
