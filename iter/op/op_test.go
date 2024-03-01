package op_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/v2/iter"
	"github.com/BooleanCat/go-functional/v2/iter/op"
)

func ExampleAdd() {
	fmt.Println(iter.Reduce(iter.Lift([]int{1, 2, 3}), op.Add, 0))
	// Output: 6
}

func ExampleAdd_string() {
	fmt.Println(iter.Reduce(iter.Lift([]string{"a", "b", "c"}), op.Add, ""))
	// Output: abc
}
