package itx_test

import (
	"fmt"
	"iter"
	"maps"

	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleIterator_Chain() {
	numbers := itx.SlicesValues([]int{1, 2}).Chain(itx.SlicesValues([]int{3, 4})).Collect()

	fmt.Println(numbers)
	// Output: [1 2 3 4]
}

func ExampleIterator2_Chain() {
	pairs := itx.MapsAll(map[string]int{"a": 1}).Chain(itx.MapsAll(map[string]int{"b": 2}))

	fmt.Println(len(maps.Collect(iter.Seq2[string, int](pairs))))
	// Output: 2
}
