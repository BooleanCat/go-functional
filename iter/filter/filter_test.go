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

func ExampleIsEqual() {
	fmt.Println(iter.Lift([]int{1, 2, 3, 4}).Filter(filter.IsEqual(2)).Collect())
	// Output: [2]
}

func ExampleIsZero() {
	fmt.Println(iter.Lift([]int{0, 1, 2, 3}).Filter(filter.IsZero).Collect())
	// Output: [0]
}

func ExampleGreaterThan() {
	fmt.Println(iter.Lift([]int{1, 2, 3, 4}).Filter(filter.GreaterThan(2)).Collect())
	// Output: [3 4]
}

func ExampleLessThan() {
	fmt.Println(iter.Lift([]int{1, 2, 3, 4}).Filter(filter.LessThan(3)).Collect())
	// Output: [1 2]
}

func ExampleNot() {
	fmt.Println(iter.Lift([]int{1, 2, 3, 4}).Filter(filter.Not[int](filter.IsEven)).Collect())
	// Output: [1 3]
}

func ExampleAnd() {
	fmt.Println(iter.Lift([]int{1, 2, 3, 4}).Filter(filter.And(filter.IsOdd, filter.GreaterThan(2))).Collect())
	// Output: [3]
}

func ExampleOr() {
	fmt.Println(iter.Lift([]int{1, 2, 3, 4}).Filter(filter.Or(filter.IsEven, filter.LessThan(3))).Collect())
	// Output: [1 2 4]
}
