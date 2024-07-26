package itx_test

import (
	"fmt"
	"maps"
	"slices"
	"strings"

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

func ExampleCollectErr() {
	data := strings.NewReader("one\ntwo\nthree\n")
	lines, err := itx.CollectErr(itx.LinesString(data))
	fmt.Println(err)
	fmt.Println(lines)
	// Output:
	// <nil>
	// [one two three]
}

func ExampleIterator_Len() {
	fmt.Println(itx.FromSlice([]int{1, 2, 3}).Len())
	// Output: 3
}

func ExampleIterator2_Len() {
	fmt.Println(itx.FromSlice([]int{1, 2, 3}).Enumerate().Len())
	// Output: 3
}
