package iter_test

import (
	"fmt"
	"iter"
	"maps"
	"slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	fn "github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleDrop() {
	for value := range fn.Drop(slices.Values([]int{1, 2, 3, 4, 5}), 2) {
		fmt.Println(value)
	}

	// Output:
	// 3
	// 4
	// 5
}

func ExampleDrop_method() {
	for value := range fn.Iterator[int](slices.Values([]int{1, 2, 3, 4, 5})).Drop(2) {
		fmt.Println(value)
	}

	// Output:
	// 3
	// 4
	// 5
}

func TestDropYieldFalse(t *testing.T) {
	t.Parallel()

	numbers := fn.Drop(slices.Values([]int{1, 2, 3, 4, 5}), 2)
	var a int
	numbers(func(value int) bool {
		a = value
		return false
	})
	assert.Equal(t, a, 3)
}

func TestDropEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(fn.Drop(fn.Exhausted[int](), 2)))
}

func ExampleDrop2() {
	numbers := maps.Collect(fn.Drop2(maps.All(map[int]string{1: "one", 2: "two", 3: "three"}), 1))

	fmt.Println(len(numbers))
	// Output: 2
}

func ExampleDrop2_method() {
	numbers := fn.Iterator2[int, string](maps.All(map[int]string{1: "one", 2: "two", 3: "three"})).Drop(1)

	fmt.Println(len(maps.Collect(iter.Seq2[int, string](numbers))))
	// Output: 2
}

func TestDrop2(t *testing.T) {
	t.Parallel()

	keys := []int{1, 2, 3}
	values := []string{"one", "two", "three"}

	numbers := maps.Collect(fn.Drop2(fn.Zip(slices.Values(keys), slices.Values(values)), 1))

	assert.Equal(t, len(numbers), 2)

	for key := range numbers {
		assert.True(t, slices.Contains(keys, key))
	}
}

func TestDrop2Empty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(fn.Drop2(fn.Exhausted2[int, int](), 1))), 0)
}

func TestDrop2Zero(t *testing.T) {
	t.Parallel()

	numbers := maps.Collect(fn.Drop2(maps.All(map[int]string{1: "one", 2: "two", 3: "three"}), 0))

	assert.Equal(t, len(numbers), 3)
}

func TestDrop2YieldFalse(t *testing.T) {
	t.Parallel()

	numbersZipped := fn.Zip(slices.Values([]int{1, 2, 3}), slices.Values([]int{3, 4, 5}))
	numbers := fn.Drop2(numbersZipped, 2)
	var a, b int
	numbers(func(v, w int) bool {
		a, b = v, w
		return false
	})
	assert.Equal(t, a, 3)
	assert.Equal(t, b, 5)
}
