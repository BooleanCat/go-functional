//go:build go1.23

package iter_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/maps"
	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/internal/assert"
	fn "github.com/BooleanCat/go-functional/v2/iter"
	"github.com/BooleanCat/go-functional/v2/iter/filter"
)

func ExampleFilter() {
	for number := range fn.Filter(slices.Values([]int{1, 2, 3, 4, 5}), filter.IsEven) {
		fmt.Println(number)
	}

	// Output:
	// 2
	// 4
}

func ExampleFilter_method() {
	for number := range fn.Iterator[int](slices.Values([]int{1, 2, 3, 4, 5})).Filter(filter.IsEven) {
		fmt.Println(number)
	}

	// Output:
	// 2
	// 4
}

func TestFilterEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(fn.Filter(fn.Exhausted[int](), filter.IsEven)))
}

func TestFilterTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull(fn.Filter(slices.Values([]int{1, 2, 3}), filter.IsEven))
	stop()
}

func ExampleExclude() {
	for number := range fn.Exclude(slices.Values([]int{1, 2, 3, 4, 5}), filter.IsEven) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 3
	// 5
}

func ExampleExclude_method() {
	for number := range fn.Iterator[int](slices.Values([]int{1, 2, 3, 4, 5})).Exclude(filter.IsEven) {
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

	for key, value := range fn.Filter2(maps.All(numbers), isOne) {
		fmt.Println(key, value)
	}

	// Output: 1 one
}

func ExampleFilter2_method() {
	isOne := func(n int, _ string) bool { return n == 1 }
	numbers := map[int]string{1: "one", 2: "two", 3: "three"}

	for key, value := range fn.Iterator2[int, string](maps.All(numbers)).Filter(isOne) {
		fmt.Println(key, value)
	}

	// Output: 1 one
}

func TestFilter2Empty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(fn.Filter2(fn.Exhausted2[int, int](), filter.Passthrough2))), 0)
}

func TestFilter2TerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull2(fn.Filter2(maps.All(map[int]string{1: "one", 2: "two"}), filter.Passthrough2))
	stop()
}

func ExampleExclude2() {
	isOne := func(n int, _ string) bool { return n == 1 }
	numbers := map[int]string{1: "one", 3: "three"}

	for key, value := range fn.Exclude2(maps.All(numbers), isOne) {
		fmt.Println(key, value)
	}

	// Output: 3 three
}

func ExampleExclude2_method() {
	isOne := func(n int, _ string) bool { return n == 1 }
	numbers := map[int]string{1: "one", 3: "three"}

	for key, value := range fn.Iterator2[int, string](maps.All(numbers)).Exclude(isOne) {
		fmt.Println(key, value)
	}

	// Output: 3 three
}

func TestExclude2Empty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(fn.Exclude2(fn.Exhausted2[int, int](), filter.Passthrough2))), 0)
}

func TestExclude2TerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull2(fn.Exclude2(fn.Exhausted2[int, int](), filter.Passthrough2))
	stop()
}
