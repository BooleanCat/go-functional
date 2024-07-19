package itx_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/v2/it/filter"
	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleIterator_Filter() {
	for number := range itx.FromSlice([]int{1, 2, 3, 4, 5}).Filter(filter.IsEven) {
		fmt.Println(number)
	}

	// Output:
	// 2
	// 4
}

func ExampleIterator_Exclude() {
	for number := range itx.FromSlice([]int{1, 2, 3, 4, 5}).Exclude(filter.IsEven) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 3
	// 5
}

func ExampleIterator2_Filter() {
	isOne := func(n int, _ string) bool { return n == 1 }
	numbers := map[int]string{1: "one", 2: "two", 3: "three"}

	for key, value := range itx.FromMap(numbers).Filter(isOne) {
		fmt.Println(key, value)
	}

	// Output: 1 one
}

func ExampleIterator2_Exclude() {
	isOne := func(n int, _ string) bool { return n == 1 }
	numbers := map[int]string{1: "one", 3: "three"}

	for key, value := range itx.FromMap(numbers).Exclude(isOne) {
		fmt.Println(key, value)
	}

	// Output: 3 three
}
