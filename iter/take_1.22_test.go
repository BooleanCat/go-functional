//go:build go1.22 && goexperiment.rangefunc

package iter_test

import (
	"fmt"
	"iter"
	sl "slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/maps"
	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/internal/assert"
	fn "github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleTake() {
	for number := range fn.Take(slices.Values([]int{1, 2, 3, 4, 5}), 3) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleTake_method() {
	for number := range fn.Iterator[int](slices.Values([]int{1, 2, 3, 4, 5})).Take(3) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func TestTakeTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull(fn.Take(slices.Values([]int{1, 2, 3}), 2))
	stop()
}

func TestTakeZero(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(fn.Take(slices.Values([]int{1, 2, 3}), 0)))
}

func TestTakeEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(fn.Take(fn.Exhausted[int](), 2)))
}

func ExampleTake2() {
	numbers := maps.Collect(fn.Take2(maps.All(map[int]string{1: "one", 2: "two", 3: "three"}), 2))

	fmt.Println(len(numbers))
	// Output: 2
}

func ExampleTake2_method() {
	numbers := maps.Collect(iter.Seq2[int, string](fn.Iterator2[int, string](maps.All(map[int]string{1: "one", 2: "two", 3: "three"})).Take(2)))

	fmt.Println(len(numbers))
	// Output: 2
}

func TestTake2(t *testing.T) {
	t.Parallel()

	keys := []int{1, 2, 3}
	values := []string{"one", "two", "three"}

	numbers := maps.Collect(fn.Take2(fn.Zip(slices.Values(keys), slices.Values(values)), 2))

	assert.Equal(t, len(numbers), 2)

	for key := range numbers {
		assert.True(t, sl.Contains(keys, key))
	}
}

func TestTake2Zero(t *testing.T) {
	t.Parallel()

	numbers := maps.Collect(fn.Take2(maps.All(map[int]string{1: "one", 2: "two", 3: "three"}), 0))
	assert.Equal(t, len(numbers), 0)
}

func TestTake2Empty(t *testing.T) {
	t.Parallel()

	numbers := maps.Collect(fn.Take2(fn.Exhausted2[int, int](), 2))
	assert.Equal(t, len(numbers), 0)
}

func TestTake2TerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull2(fn.Take2(maps.All(map[int]string{1: "one", 2: "two"}), 1))
	stop()
}
