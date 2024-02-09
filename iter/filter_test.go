package iter_test

import (
	"fmt"
	it "iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/iter"
	"github.com/BooleanCat/go-functional/v2/iter/filter"
)

func ExampleFilter() {
	isEven := func(n int) bool { return n%2 == 0 }

	for number := range iter.Filter(iter.Lift([]int{1, 2, 3, 4, 5}), isEven) {
		fmt.Println(number)
	}

	// Output:
	// 2
	// 4
}

func ExampleFilter_method() {
	for number := range iter.Lift([]int{1, 2, 3, 4, 5}).Filter(filter.IsEven) {
		fmt.Println(number)
	}

	// Output:
	// 2
	// 4
}

func TestFilterEmpty(t *testing.T) {
	t.Parallel()

	for _ = range iter.Filter(iter.Lift([]int{}), filter.IsEven) {
		t.Error("unexpected")
	}
}

func TestFilterTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull(it.Seq[int](iter.Filter(iter.Lift([]int{1, 2, 3}), filter.IsEven)))
	stop()
}

func ExampleFilter2() {
	isOne := func(n int, _ string) bool { return n == 1 }
	numbers := map[int]string{1: "one", 2: "two", 3: "three"}

	for key, value := range iter.Filter2(iter.LiftHashMap(numbers), isOne) {
		fmt.Println(key, value)
	}

	// Output:
	// 1 one
}

func ExampleFilter2_method() {
	isOne := func(n int, _ string) bool { return n == 1 }
	numbers := map[int]string{1: "one", 2: "two", 3: "three"}

	for key, value := range iter.LiftHashMap(numbers).Filter2(isOne) {
		fmt.Println(key, value)
	}

	// Output:
	// 1 one
}

func TestFilter2Empty(t *testing.T) {
	t.Parallel()

	for _, _ = range iter.Filter2(iter.LiftHashMap(map[int]string{}), func(int, string) bool { return true }) {
		t.Error("unexpected")
	}
}

func TestFilter2TerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull2(it.Seq2[int, string](iter.Filter2(iter.LiftHashMap(map[int]string{
		1: "one",
		2: "two",
	}), func(int, string) bool { return true })))
	stop()
}
