package slices_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/internal/assert"
)

func ExampleValues() {
	for number := range slices.Values([]int{1, 2, 3}) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func TestValuesEmpty(t *testing.T) {
	t.Parallel()

	for number := range slices.Values([]int{}) {
		t.Error("unexpected", number)
	}
}

func TestValuesTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull(slices.Values([]int{1, 2, 3}))
	stop()
}

func ExampleCollect() {
	fmt.Println(slices.Collect(slices.Values([]int{1, 2, 3})))
	// Output: [1 2 3]
}

func TestCollectEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(slices.Values([]int{})))
}
