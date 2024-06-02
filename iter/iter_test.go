package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
	"github.com/BooleanCat/go-functional/v2/iter/op"
)

func ExampleCollect_method() {
	fmt.Println(iter.Iterator[int](slices.Values([]int{1, 2, 3})).Collect())
	// Output: [1 2 3]
}

func ExampleForEach() {
	iter.ForEach(slices.Values([]int{1, 2, 3}), func(number int) {
		fmt.Println(number)
	})
	// Output:
	// 1
	// 2
	// 3
}

func ExampleForEach_method() {
	iter.Iterator[int](slices.Values([]int{1, 2, 3})).ForEach(func(number int) {
		fmt.Println(number)
	})
	// Output:
	// 1
	// 2
	// 3
}

func TestForEachEmpty(t *testing.T) {
	t.Parallel()

	iter.ForEach(slices.Values([]int{}), func(int) {
		t.Error("unexpected")
	})
}

func ExampleForEach2() {
	iter.ForEach2(iter.Enumerate[int](slices.Values([]int{1, 2, 3})), func(index int, number int) {
		fmt.Println(index, number)
	})
	// Output:
	// 0 1
	// 1 2
	// 2 3
}

func ExampleForEach2_method() {
	iter.Iterator[int](slices.Values([]int{1, 2, 3})).Enumerate().ForEach(func(index int, number int) {
		fmt.Println(index, number)
	})
	// Output:
	// 0 1
	// 1 2
	// 2 3
}

func TestForEach2Empty(t *testing.T) {
	t.Parallel()

	iter.ForEach2(iter.Enumerate[int](slices.Values([]int{})), func(int, int) {
		t.Error("unexpected")
	})
}

func ExampleReduce() {
	fmt.Println(iter.Reduce(slices.Values([]int{1, 2, 3}), op.Add, 0))
	// Output: 6
}

func TestReduceEmpty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, iter.Reduce(slices.Values([]int{}), func(int, int) int { return 0 }, 0), 0)
}

func ExampleReduce2() {
	fmt.Println(iter.Reduce2(iter.Enumerate[int](slices.Values([]int{1, 2, 3})), func(i, a, b int) int {
		return i + 1
	}, 0))

	// Output: 3
}
