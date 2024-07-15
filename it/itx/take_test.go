package itx_test

import (
	"fmt"
	"maps"
	"slices"

	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleIterator_Take() {
	for number := range itx.From(slices.Values([]int{1, 2, 3, 4, 5})).Take(3) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleIterator2_Take() {
	numbers := maps.Collect(itx.From2(maps.All(map[int]string{1: "one", 2: "two", 3: "three"})).Take(2).Seq())

	fmt.Println(len(numbers))
	// Output: 2
}
