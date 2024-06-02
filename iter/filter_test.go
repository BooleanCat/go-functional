package iter_test

import (
	"fmt"
	it "iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/maps"
	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
	"github.com/BooleanCat/go-functional/v2/iter/filter"
)

func ExampleFilter() {
	for number := range iter.Filter(slices.Values([]int{1, 2, 3, 4, 5}), filter.IsEven) {
		fmt.Println(number)
	}

	// Output:
	// 2
	// 4
}

func ExampleFilter_method() {
	for number := range iter.Iterator[int](slices.Values([]int{1, 2, 3, 4, 5})).Filter(filter.IsEven) {
		fmt.Println(number)
	}

	// Output:
	// 2
	// 4
}

func TestFilterEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(iter.Filter(slices.Values([]int{}), filter.IsEven)))
}

func TestFilterTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull(iter.Filter(slices.Values([]int{1, 2, 3}), filter.IsEven))
	stop()
}

func ExampleExclude() {
	for number := range iter.Exclude(slices.Values([]int{1, 2, 3, 4, 5}), filter.IsEven) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 3
	// 5
}

func ExampleExclude_method() {
	for number := range iter.Iterator[int](slices.Values([]int{1, 2, 3, 4, 5})).Exclude(filter.IsEven) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 3
	// 5
}

func ExampleFilter2() {
	isOne := func(n int, _ string) bool { return n == 1 }
	numbers := map[int]string{1: "one", 2: "two", 3: "three"}

	for key, value := range iter.Filter2(maps.All(numbers), isOne) {
		fmt.Println(key, value)
	}

	// Output: 1 one
}

func ExampleFilter2_method() {
	isOne := func(n int, _ string) bool { return n == 1 }
	numbers := map[int]string{1: "one", 2: "two", 3: "three"}

	for key, value := range iter.Iterator2[int, string](maps.All(numbers)).Filter2(isOne) {
		fmt.Println(key, value)
	}

	// Output: 1 one
}

func TestFilter2Empty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(iter.Filter2(maps.All(map[int]string{}), func(int, string) bool {
		return true
	}))), 0)
}

func TestFilter2TerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull2(iter.Filter2(maps.All(map[int]string{
		1: "one",
		2: "two",
	}), func(int, string) bool { return true }))
	stop()
}

func ExampleExclude2() {
	isOne := func(n int, _ string) bool { return n == 1 }
	numbers := map[int]string{1: "one", 3: "three"}

	for key, value := range iter.Exclude2(maps.All(numbers), isOne) {
		fmt.Println(key, value)
	}

	// Output: 3 three
}

func ExampleExclude2_method() {
	isOne := func(n int, _ string) bool { return n == 1 }
	numbers := map[int]string{1: "one", 3: "three"}

	for key, value := range iter.Iterator2[int, string](maps.All(numbers)).Exclude2(isOne) {
		fmt.Println(key, value)
	}

	// Output: 3 three
}

func TestExclude2Empty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(iter.Exclude2(maps.All(map[int]string{}), func(int, string) bool {
		return true
	}))), 0)
}

func TestExclude2TerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull2(iter.Exclude2(maps.All(map[int]string{}), func(int, string) bool { return true }))
	stop()
}
