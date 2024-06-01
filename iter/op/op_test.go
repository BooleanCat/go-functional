package op_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/iter"
	"github.com/BooleanCat/go-functional/v2/iter/op"
)

func ExampleAdd() {
	fmt.Println(iter.Reduce(iter.Iterator[int](slices.Values([]int{1, 2, 3})), op.Add, 0))
	// Output: 6
}

func ExampleAdd_string() {
	fmt.Println(iter.Reduce(iter.Iterator[string](slices.Values([]string{"a", "b", "c"})), op.Add, ""))
	// Output: abc
}
