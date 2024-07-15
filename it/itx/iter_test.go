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
	fmt.Println(maps.Collect(itx.From2(maps.All(map[int]int{1: 2})).Seq()))
	// Output: map[1:2]
}

func ExampleIterator_Collect() {
	fmt.Println(itx.From(slices.Values([]int{1, 2, 3})).Collect())
	// Output: [1 2 3]
}

func ExampleIterator_ForEach() {
	itx.From(slices.Values([]int{1, 2, 3})).ForEach(func(number int) {
		fmt.Println(number)
	})
	// Output:
	// 1
	// 2
	// 3
}

func ExampleIterator2_ForEach() {
	itx.From(slices.Values([]int{1, 2, 3})).Enumerate().ForEach(func(index int, number int) {
		fmt.Println(index, number)
	})
	// Output:
	// 0 1
	// 1 2
	// 2 3
}

func ExampleIterator_Find() {
	fmt.Println(itx.From(slices.Values([]int{1, 2, 3})).Find(func(number int) bool {
		return number == 2
	}))
	// Output: 2 true
}

func ExampleIterator2_Find() {
	fmt.Println(itx.From(slices.Values([]int{1, 2, 3})).Enumerate().Find(func(index int, number int) bool {
		return index == 1
	}))
	// Output: 1 2 true
}
