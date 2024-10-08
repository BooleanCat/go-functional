package itx_test

import (
	"fmt"
	"maps"
	"slices"

	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleIterator_Seq() {
	fmt.Println(slices.Collect(itx.NaturalNumbers[int]().Take(3).Seq()))
	// Output: [0 1 2]
}

func ExampleIterator2_Seq() {
	fmt.Println(maps.Collect(itx.FromMap(map[int]int{1: 2}).Seq()))
	// Output: map[1:2]
}

func ExampleIterator_Collect() {
	fmt.Println(itx.FromSlice([]int{1, 2, 3}).Collect())
	// Output: [1 2 3]
}

func ExampleIterator_ForEach() {
	itx.FromSlice([]int{1, 2, 3}).ForEach(func(number int) {
		fmt.Println(number)
	})
	// Output:
	// 1
	// 2
	// 3
}

func ExampleIterator2_ForEach() {
	itx.FromSlice([]int{1, 2, 3}).Enumerate().ForEach(func(index int, number int) {
		fmt.Println(index, number)
	})
	// Output:
	// 0 1
	// 1 2
	// 2 3
}

func ExampleIterator_Find() {
	fmt.Println(itx.FromSlice([]int{1, 2, 3}).Find(func(number int) bool {
		return number == 2
	}))
	// Output: 2 true
}

func ExampleIterator2_Find() {
	fmt.Println(itx.FromSlice([]int{1, 2, 3}).Enumerate().Find(func(index int, number int) bool {
		return index == 1
	}))
	// Output: 1 2 true
}

func ExampleFrom() {
	fmt.Println(itx.From(slices.Values([]int{1, 2, 3})).Collect())
	// Output: [1 2 3]
}

func ExampleFrom2() {
	numbers := maps.All(map[int]string{1: "one"})
	fmt.Println(maps.Collect(itx.From2(numbers).Seq()))
	// Output: map[1:one]
}

func ExampleFromSlice() {
	fmt.Println(itx.FromSlice([]int{1, 2, 3}).Collect())
	// Output: [1 2 3]
}

func ExampleFromMap() {
	fmt.Println(itx.FromMap(map[int]int{1: 2}).Right().Collect())
	// Output: [2]
}

func ExampleIterator2_Collect() {
	indicies, values := itx.FromSlice([]int{1, 2, 3}).Enumerate().Collect()
	fmt.Println(values)
	fmt.Println(indicies)

	// Output:
	// [1 2 3]
	// [0 1 2]
}

func ExampleIterator_Len() {
	fmt.Println(itx.FromSlice([]int{1, 2, 3}).Len())
	// Output: 3
}

func ExampleIterator2_Len() {
	fmt.Println(itx.FromSlice([]int{1, 2, 3}).Enumerate().Len())
	// Output: 3
}

func ExampleIterator_Drain() {
	itx.From(it.Map(slices.Values([]int{1, 2, 3}), func(n int) int {
		fmt.Println(n)
		return n
	})).Drain()

	// Output:
	// 1
	// 2
	// 3
}

func ExampleIterator2_Drain() {
	itx.From2(it.Map2(slices.All([]int{1, 2, 3}), func(i, n int) (int, int) {
		fmt.Println(n)
		return i, n
	})).Drain()

	// Output:
	// 1
	// 2
	// 3
}
