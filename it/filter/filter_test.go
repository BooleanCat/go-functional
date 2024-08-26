package filter_test

import (
	"fmt"
	"maps"
	"regexp"
	"slices"

	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/BooleanCat/go-functional/v2/it/filter"
)

func ExampleIsEven() {
	for number := range it.Filter(slices.Values([]int{1, 2, 3, 4}), filter.IsEven) {
		fmt.Println(number)
	}

	// Output:
	// 2
	// 4
}

func ExampleIsOdd() {
	for number := range it.Filter(slices.Values([]int{1, 2, 3, 4}), filter.IsOdd) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 3
}

func ExampleIsEqual() {
	for number := range it.Filter(slices.Values([]int{1, 2, 3, 4}), filter.IsEqual(2)) {
		fmt.Println(number)
	}

	// Output: 2
}

func ExampleNotEqual() {
	for number := range it.Filter(slices.Values([]int{1, 2, 3, 4}), filter.NotEqual(2)) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 3
	// 4
}

func ExampleIsZero() {
	for number := range it.Filter(slices.Values([]int{0, 1, 2, 3}), filter.IsZero) {
		fmt.Println(number)
	}

	// Output: 0
}

func ExampleGreaterThan() {
	for number := range it.Filter(slices.Values([]int{1, 2, 3, 4}), filter.GreaterThan(2)) {
		fmt.Println(number)
	}

	// Output:
	// 3
	// 4
}

func ExampleLessThan() {
	for number := range it.Filter(slices.Values([]int{1, 2, 3, 4}), filter.LessThan(3)) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
}

func ExamplePassthrough() {
	for number := range it.Filter(slices.Values([]int{1, 2, 3}), filter.Passthrough) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExamplePassthrough2() {
	for key, value := range it.Filter2(maps.All(map[int]string{1: "two"}), filter.Passthrough2) {
		fmt.Println(key, value)
	}

	// Output: 1 two
}

func ExampleNot() {
	numbers := slices.Values([]int{1, 2, 3, 4})

	for number := range it.Filter(numbers, filter.Not[int](filter.IsEven)) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 3
}

func ExampleAnd() {
	numbers := slices.Values([]int{1, 2, 3, 4})

	for number := range it.Filter(numbers, filter.And(filter.IsOdd, filter.GreaterThan(2))) {
		fmt.Println(number)
	}

	// Output: 3
}

func ExampleOr() {
	numbers := slices.Values([]int{1, 2, 3, 4})

	for number := range it.Filter(numbers, filter.Or(filter.IsEven, filter.LessThan(3))) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 4
}

func ExampleMatch_string() {
	pattern := regexp.MustCompile(`^foo`)

	strings := slices.Values([]string{"foobar", "barfoo"})

	for match := range it.Filter(strings, filter.Match[string](pattern)) {
		fmt.Println(match)
	}

	// Output: foobar
}

func ExampleMatch_bytes() {
	pattern := regexp.MustCompile(`^foo`)

	strings := slices.Values([][]byte{[]byte("foobar"), []byte("barfoo")})

	for match := range it.Filter(strings, filter.Match[[]byte](pattern)) {
		fmt.Println(string(match))
	}

	// Output: foobar
}

func ExampleContains_string() {
	strings := slices.Values([]string{"foobar", "barfoo"})

	for element := range it.Filter(strings, filter.Contains("rfoo")) {
		fmt.Println(element)
	}

	// Output: barfoo
}

func ExampleContains_bytes() {
	strings := slices.Values([][]byte{[]byte("foobar"), []byte("barfoo")})

	for element := range it.Filter(strings, filter.Contains([]byte("rfoo"))) {
		fmt.Println(string(element))
	}

	// Output: barfoo
}
