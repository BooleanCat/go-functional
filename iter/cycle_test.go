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

func ExampleCycle() {
	numbers := slices.Collect(fn.Take(fn.Cycle(slices.Values([]int{1, 2})), 5))

	fmt.Println(numbers)
	// Output: [1 2 1 2 1]
}

func ExampleCycle_method() {
	numbers := fn.Iterator[int](slices.Values([]int{1, 2})).Cycle().Take(5).Collect()

	fmt.Println(numbers)
	// Output: [1 2 1 2 1]
}

func TestCycleYieldFalse(t *testing.T) {
	t.Parallel()

	numbers := fn.Cycle(slices.Values([]int{1, 2}))
	var a int
	numbers(func(value int) bool {
		a = value
		return false
	})
	assert.Equal(t, a, 1)
}

func ExampleCycle2() {
	numbers := maps.Collect(fn.Take2(fn.Cycle2(maps.All(map[int]string{1: "one"})), 5))

	fmt.Println(numbers)
	// Output: map[1:one]
}

func ExampleCycle2_method() {
	numbers := fn.Iterator2[int, string](maps.All(map[int]string{1: "one"})).Cycle().Take(5)

	fmt.Println(maps.Collect(iter.Seq2[int, string](numbers)))
	// Output: map[1:one]
}

func TestCycle2YieldFalse(t *testing.T) {
	t.Parallel()

	numbers := fn.Cycle2(maps.All(map[int]string{1: "one"}))
	var (
		a int
		b string
	)
	numbers(func(key int, value string) bool {
		a, b = key, value
		return false
	})
	assert.Equal(t, a, 1)
	assert.Equal(t, b, "one")
}
