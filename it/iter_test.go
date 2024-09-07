package it_test

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/BooleanCat/go-functional/v2/it/op"
)

func ExampleForEach() {
	it.ForEach(slices.Values([]int{1, 2, 3}), func(number int) {
		fmt.Println(number)
	})
	// Output:
	// 1
	// 2
	// 3
}

func TestForEachEmpty(t *testing.T) {
	t.Parallel()

	it.ForEach(slices.Values([]int{}), func(int) {
		t.Error("unexpected")
	})
}

func ExampleForEach2() {
	it.ForEach2(slices.All([]int{1, 2, 3}), func(index int, number int) {
		fmt.Println(index, number)
	})
	// Output:
	// 0 1
	// 1 2
	// 2 3
}

func TestForEach2Empty(t *testing.T) {
	t.Parallel()

	it.ForEach2(it.Enumerate(it.Exhausted[int]()), func(int, int) {
		t.Error("unexpected")
	})
}

func ExampleFold() {
	fmt.Println(it.Fold(slices.Values([]int{1, 2, 3}), op.Add, 0))
	// Output: 6
}

func TestFoldEmpty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, it.Fold(it.Exhausted[int](), func(int, int) int { return 0 }, 0), 0)
}

func ExampleFold2() {
	fmt.Println(it.Fold2(slices.All([]int{1, 2, 3}), func(i, a, b int) int {
		return i + 1
	}, 0))

	// Output: 3
}

func ExampleMax() {
	max, ok := it.Max(slices.Values([]int{1, 2, 3}))
	fmt.Println(max, ok)

	// Output: 3 true
}

func TestMaxEmpty(t *testing.T) {
	t.Parallel()

	max, ok := it.Max(it.Exhausted[int]())
	assert.Equal(t, max, 0)
	assert.False(t, ok)
}

func ExampleMin() {
	min, ok := it.Min(slices.Values([]int{4, 2, 1, 3}))
	fmt.Println(min, ok)

	// Output: 1 true
}

func TestMinEmpty(t *testing.T) {
	t.Parallel()

	min, ok := it.Min(it.Exhausted[int]())
	assert.Equal(t, min, 0)
	assert.False(t, ok)
}

func ExampleFind() {
	found, ok := it.Find(slices.Values([]int{1, 2, 3}), func(i int) bool {
		return i == 2
	})
	fmt.Println(found, ok)

	// Output: 2 true
}

func ExampleFind_notFound() {
	found, ok := it.Find(slices.Values([]int{1, 2, 3}), func(i int) bool {
		return i == 4
	})
	fmt.Println(found, ok)

	// Output: 0 false
}

func ExampleFind2() {
	index, value, ok := it.Find2(slices.All([]int{1, 2, 3}), func(i, v int) bool {
		return i == 2
	})
	fmt.Println(index, value, ok)

	// Output: 2 3 true
}

func ExampleFind2_notFound() {
	index, value, ok := it.Find2(slices.All([]int{1, 2, 3}), func(i, v int) bool {
		return i == 4
	})

	fmt.Println(index, value, ok)
	// Output: 0 0 false
}

func ExampleCollect2() {
	indicies, values := it.Collect2(slices.All([]int{1, 2, 3}))
	fmt.Println(values)
	fmt.Println(indicies)

	// Output:
	// [1 2 3]
	// [0 1 2]
}

func ExampleTryCollect() {
	text := strings.NewReader("one\ntwo\nthree\n")

	lines, err := it.TryCollect(it.LinesString(text))
	fmt.Println(err)
	fmt.Println(lines)

	// Output:
	// <nil>
	// [one two three]
}

func TestTryCollectError(t *testing.T) {
	t.Parallel()

	text := new(failSecondTime)
	lines, err := it.TryCollect(it.LinesString(text))

	assert.Equal(t, err.Error(), "read error")
	assert.SliceEqual(t, lines, []string{"o"})
}

func ExampleLen() {
	fmt.Println(it.Len(slices.Values([]int{1, 2, 3})))

	// Output: 3
}

func TestLenEmpty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, it.Len(it.Exhausted[int]()), 0)
}

func ExampleLen2() {
	fmt.Println(it.Len2(slices.All([]int{1, 2, 3})))

	// Output: 3
}

func TestLen2Empty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, it.Len2(it.Enumerate(it.Exhausted[int]())), 0)
}

func ExampleContains() {
	numbers := slices.Values([]int{1, 2, 3})
	fmt.Println(it.Contains(numbers, 2))
	// Output: true
}

func TestContainsFalse(t *testing.T) {
	assert.False(t, it.Contains(slices.Values([]int{1, 2, 3}), 4))
}
