package itx_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleIterator_Take() {
	for number := range itx.SlicesValues([]int{1, 2, 3, 4, 5}).Take(3) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleIterator2_Take() {
	numbers := itx.MapsCollect(itx.MapsAll(map[int]string{1: "one", 2: "two", 3: "three"}).Take(2))

	fmt.Println(len(numbers))
	// Output: 2
}
