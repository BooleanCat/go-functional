package itx_test

import (
	"fmt"
	"maps"
	"slices"

	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleIterator_Seq() {
	fmt.Println(slices.Collect(itx.Count[int]().Take(3).Seq()))
	// Output: [0 1 2]
}

func ExampleIterator2_Seq() {
	fmt.Println(maps.Collect(itx.MapsAll(map[int]int{1: 2}).Seq()))
	// Output: map[1:2]
}

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
