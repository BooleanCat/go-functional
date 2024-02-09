package filter_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/v2/iter"
	"github.com/BooleanCat/go-functional/v2/iter/filter"
)

func ExampleIsEven() {
	fmt.Println(iter.Lift([]int{1, 2, 3, 4}).Filter(filter.IsEven).Collect())
	// Output: [2 4]
}

func ExampleIsOdd() {
	fmt.Println(iter.Lift([]int{1, 2, 3, 4}).Filter(filter.IsOdd).Collect())
	// Output: [1 3]
}
