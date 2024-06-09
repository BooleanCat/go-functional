package filter_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/v2/future/maps"
	"github.com/BooleanCat/go-functional/v2/future/slices"
	fn "github.com/BooleanCat/go-functional/v2/iter"
	"github.com/BooleanCat/go-functional/v2/iter/filter"
)

func ExampleIsEven() {
	fmt.Println(fn.Iterator[int](slices.Values([]int{1, 2, 3, 4})).Filter(filter.IsEven).Collect())
	// Output: [2 4]
}

func ExampleIsOdd() {
	fmt.Println(fn.Iterator[int](slices.Values([]int{1, 2, 3, 4})).Filter(filter.IsOdd).Collect())
	// Output: [1 3]
}

func ExampleIsEqual() {
	fmt.Println(fn.Iterator[int](slices.Values([]int{1, 2, 3, 4})).Filter(filter.IsEqual(2)).Collect())
	// Output: [2]
}

func ExampleNotEqual() {
	fmt.Println(fn.Iterator[int](slices.Values([]int{1, 2, 3, 4})).Filter(filter.NotEqual(2)).Collect())
	// Output: [1 3 4]
}

func ExampleIsZero() {
	fmt.Println(fn.Iterator[int](slices.Values([]int{0, 1, 2, 3})).Filter(filter.IsZero).Collect())
	// Output: [0]
}

func ExampleGreaterThan() {
	fmt.Println(fn.Iterator[int](slices.Values([]int{1, 2, 3, 4})).Filter(filter.GreaterThan(2)).Collect())
	// Output: [3 4]
}

func ExampleLessThan() {
	fmt.Println(fn.Iterator[int](slices.Values([]int{1, 2, 3, 4})).Filter(filter.LessThan(3)).Collect())
	// Output: [1 2]
}

func ExamplePassthrough() {
	fmt.Println(fn.Iterator[int](slices.Values([]int{1, 2, 3, 4})).Filter(filter.Passthrough).Collect())
	// Output: [1 2 3 4]

}

func ExamplePassthrough2() {
	numbers := maps.Collect(fn.Filter2(maps.All(map[int]string{1: "two"}), filter.Passthrough2))

	fmt.Println(numbers[1])
	// Output: two

}

func ExampleNot() {
	fmt.Println(fn.Iterator[int](slices.Values([]int{1, 2, 3, 4})).Filter(filter.Not[int](filter.IsEven)).Collect())
	// Output: [1 3]
}

func ExampleAnd() {
	fmt.Println(fn.Iterator[int](slices.Values([]int{1, 2, 3, 4})).Filter(filter.And(filter.IsOdd, filter.GreaterThan(2))).Collect())
	// Output: [3]
}

func ExampleOr() {
	fmt.Println(fn.Iterator[int](slices.Values([]int{1, 2, 3, 4})).Filter(filter.Or(filter.IsEven, filter.LessThan(3))).Collect())
	// Output: [1 2 4]
}
