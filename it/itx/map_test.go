package itx_test

import (
	"fmt"
	"maps"

	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleIterator_Transform() {
	fmt.Println(itx.FromSlice([]int{0, 1, 2}).Transform(func(v int) int {
		return v + 1
	}).Collect())
	// Output: [1 2 3]
}

func ExampleIterator2_Transform() {
	addOne := func(a, b int) (int, int) {
		return a + 1, b + 1
	}

	fmt.Println(maps.Collect(itx.FromMap(map[int]int{1: 2}).Transform(addOne).Seq()))
	// Output: map[2:3]
}

func ExampleIterator_TransformError() {
	fmt.Println(it.TryCollect(itx.FromSlice([]int{0, 1, 2}).TransformError(func(v int) (int, error) {
		return v + 1, nil
	})))
	// Output: [1 2 3] <nil>
}
