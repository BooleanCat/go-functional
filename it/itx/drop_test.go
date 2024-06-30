package itx_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleIterator_Drop() {
	for value := range itx.SlicesValues([]int{1, 2, 3, 4, 5}).Drop(2) {
		fmt.Println(value)
	}

	// Output:
	// 3
	// 4
	// 5
}

func ExampleIterator2_Drop() {
	numbers := itx.MapsAll(map[int]string{1: "one", 2: "two", 3: "three"}).Drop(1)

	fmt.Println(len(itx.MapsCollect(numbers)))
	// Output: 2
}
