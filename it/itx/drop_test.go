package itx_test

import (
	"fmt"
	"maps"

	"github.com/BooleanCat/go-functional/v2/it/filter"
	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleIterator_Drop() {
	for value := range itx.FromSlice([]int{1, 2, 3, 4, 5}).Drop(2) {
		fmt.Println(value)
	}

	// Output:
	// 3
	// 4
	// 5
}

func ExampleIterator2_Drop() {
	numbers := itx.FromMap(map[int]string{1: "one", 2: "two", 3: "three"}).Drop(1)

	fmt.Println(len(maps.Collect(numbers.Seq())))
	// Output: 2
}

func ExampleIterator_DropWhile() {
	for value := range itx.FromSlice([]int{1, 2, 3, 4, 5}).DropWhile(filter.LessThan(3)) {
		fmt.Println(value)
	}

	// Output:
	// 3
	// 4
	// 5
}

func ExampleIterator2_DropWhile() {
	lessThanThree := func(int, v int) bool { return v < 3 }

	_, numbers := itx.FromSlice([]int{1, 2, 3, 4, 5}).Enumerate().DropWhile(lessThanThree).Collect()
	fmt.Println(numbers)
	// Output: [3 4 5]
}
