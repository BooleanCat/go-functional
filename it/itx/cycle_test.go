package itx_test

import (
	"fmt"
	"maps"

	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleIterator_Cycle() {
	numbers := itx.FromSlice([]int{1, 2}).Cycle().Take(5).Collect()

	fmt.Println(numbers)
	// Output: [1 2 1 2 1]
}

func ExampleIterator2_Cycle() {
	numbers := itx.FromMap(map[int]string{1: "one"}).Cycle().Take(5)

	fmt.Println(maps.Collect(numbers.Seq()))
	// Output: map[1:one]
}
