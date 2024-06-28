//go:build go1.22 && goexperiment.rangefunc

package iter_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/maps"
	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/internal/assert"
	fn "github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleMap() {
	double := func(n int) int { return n * 2 }

	for number := range fn.Map(slices.Values([]int{1, 2, 3}), double) {
		fmt.Println(number)
	}

	// Output:
	// 2
	// 4
	// 6
}

func TestMapEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(fn.Map(fn.Exhausted[int](), func(int) int { return 0 })))
}

func TestMapTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull(fn.Map(slices.Values([]int{1, 2, 3}), func(int) int { return 0 }))
	stop()
}

func ExampleMap2() {
	doubleBoth := func(n, m int) (int, int) { return n * 2, m * 2 }

	for left, right := range fn.Map2(fn.Zip(slices.Values([]int{1, 2, 3}), slices.Values([]int{2, 3, 4})), doubleBoth) {
		fmt.Println(left, right)
	}

	// Output:
	// 2 4
	// 4 6
	// 6 8
}

func TestMap2Empty(t *testing.T) {
	t.Parallel()

	doubleBoth := func(n, m int) (int, int) { return n * 2, m * 2 }

	assert.Equal(t, len(maps.Collect(fn.Map2(fn.Exhausted2[int, int](), doubleBoth))), 0)
}

func TestMap2TerminateEarly(t *testing.T) {
	t.Parallel()

	doubleBoth := func(n, m int) (int, int) { return n * 2, m * 2 }

	_, stop := iter.Pull2(fn.Map2(maps.All(map[int]int{1: 2}), doubleBoth))
	stop()
}
