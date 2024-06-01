package iter_test

import (
	"fmt"
	it "iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/slices"
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

	for _ = range iter.Filter(slices.Values([]int{}), filter.IsEven) {
		t.Error("unexpected")
	}
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

	for key, value := range iter.Filter2(it.Seq2[int, string](iter.LiftHashMap(numbers)), isOne) {
		fmt.Println(key, value)
	}

	// Output: 1 one
}

func ExampleFilter2_method() {
	isOne := func(n int, _ string) bool { return n == 1 }
	numbers := map[int]string{1: "one", 2: "two", 3: "three"}

	for key, value := range iter.LiftHashMap(numbers).Filter2(isOne) {
		fmt.Println(key, value)
	}

	// Output: 1 one
}

func TestFilter2Empty(t *testing.T) {
	t.Parallel()

	for _, _ = range iter.Filter2(it.Seq2[int, string](iter.LiftHashMap(map[int]string{})), func(int, string) bool { return true }) {
		t.Error("unexpected")
	}
}

func TestFilter2TerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull2(iter.Filter2(it.Seq2[int, string](iter.LiftHashMap(map[int]string{
		1: "one",
		2: "two",
	})), func(int, string) bool { return true }))
	stop()
}

func ExampleExclude2() {
	isOne := func(n int, _ string) bool { return n == 1 }
	numbers := map[int]string{1: "one", 3: "three"}

	for key, value := range iter.Exclude2(it.Seq2[int, string](iter.LiftHashMap(numbers)), isOne) {
		fmt.Println(key, value)
	}

	// Output: 3 three
}

func ExampleExclude2_method() {
	isOne := func(n int, _ string) bool { return n == 1 }
	numbers := map[int]string{1: "one", 3: "three"}

	for key, value := range iter.LiftHashMap(numbers).Exclude2(isOne) {
		fmt.Println(key, value)
	}

	// Output: 3 three
}

func TestExclude2Empty(t *testing.T) {
	t.Parallel()

	for _, _ = range iter.Exclude2(it.Seq2[int, string](iter.LiftHashMap(map[int]string{})), func(int, string) bool { return true }) {
		t.Error("unexpected")
	}
}

func TestExclude2TerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull2(iter.Exclude2(it.Seq2[int, string](iter.LiftHashMap(map[int]string{})), func(int, string) bool { return true }))
	stop()
}
