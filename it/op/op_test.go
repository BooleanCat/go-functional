package op_test

import (
	"fmt"
	"slices"

	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/BooleanCat/go-functional/v2/it/op"
)

func ExampleAdd() {
	fmt.Println(it.Fold(slices.Values([]int{1, 2, 3}), op.Add, 0))
	// Output: 6
}

func ExampleAdd_string() {
	fmt.Println(it.Fold(slices.Values([]string{"a", "b", "c"}), op.Add, ""))
	// Output: abc
}

func ExampleRef() {
	refs := slices.Collect(it.Map(slices.Values([]int{5, 6, 7}), op.Ref))
	fmt.Println(*refs[0], *refs[1], *refs[2])
	// Output: 5 6 7
}

func ExampleDeref() {
	intRef := func(a int) *int {
		return &a
	}

	values := slices.Values([]*int{intRef(4), intRef(5), intRef(6)})

	fmt.Println(slices.Collect(it.Map(values, op.Deref)))
	// Output: [4 5 6]
}
