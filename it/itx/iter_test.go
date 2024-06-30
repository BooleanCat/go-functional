package itx_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleIterator_Collect() {
	fmt.Println(itx.SlicesValues([]int{1, 2, 3}).Collect())
	// Output: [1 2 3]
}

func ExampleIterator_ForEach() {
	itx.SlicesValues([]int{1, 2, 3}).ForEach(func(number int) {
		fmt.Println(number)
	})
	// Output:
	// 1
	// 2
	// 3
}

func ExampleIterator2_ForEach() {
	itx.SlicesValues([]int{1, 2, 3}).Enumerate().ForEach(func(index int, number int) {
		fmt.Println(index, number)
	})
	// Output:
	// 0 1
	// 1 2
	// 2 3
}
