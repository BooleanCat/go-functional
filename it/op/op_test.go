package op_test

import (
	"fmt"
	"slices"

	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/BooleanCat/go-functional/v2/it/op"
)

func ExampleAdd() {
	fmt.Println(it.Reduce(slices.Values([]int{1, 2, 3}), op.Add, 0))
	// Output: 6
}

func ExampleAdd_string() {
	fmt.Println(it.Reduce(slices.Values([]string{"a", "b", "c"}), op.Add, ""))
	// Output: abc
}
