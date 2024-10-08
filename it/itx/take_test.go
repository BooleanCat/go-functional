package itx_test

import (
	"fmt"
	"maps"

	"github.com/BooleanCat/go-functional/v2/it/filter"
	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleIterator_Take() {
	for number := range itx.FromSlice([]int{1, 2, 3, 4, 5}).Take(3) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleIterator2_Take() {
	numbers := maps.Collect(itx.FromMap(map[int]string{1: "one", 2: "two", 3: "three"}).Take(2).Seq())

	fmt.Println(len(numbers))
	// Output: 2
}

func ExampleIterator_TakeWhile() {
	for number := range itx.FromSlice([]int{1, 2, 3, 4, 5}).TakeWhile(filter.LessThan(4)) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleIterator2_TakeWhile() {
	lessThanFour := func(int, v int) bool { return v < 4 }

	_, numbers := itx.FromSlice([]int{1, 2, 3, 4, 5}).Enumerate().TakeWhile(lessThanFour).Collect()
	fmt.Println(numbers)
	// Output: [1 2 3]
}
