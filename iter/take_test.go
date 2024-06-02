package iter_test

import (
	"fmt"
	it "iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleTake() {
	for number := range iter.Take(slices.Values([]int{1, 2, 3, 4, 5}), 3) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleTake_method() {
	for number := range iter.Iterator[int](slices.Values([]int{1, 2, 3, 4, 5})).Take(3) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func TestTakeTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull(iter.Take(slices.Values([]int{1, 2, 3}), 2))
	stop()
}

func TestTakeZero(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(iter.Take(slices.Values([]int{1, 2, 3}), 0)))
}

func TestTakeEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(iter.Take(slices.Values([]int{}), 2)))
}
