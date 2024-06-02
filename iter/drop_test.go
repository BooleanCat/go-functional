package iter_test

import (
	"fmt"
	it "iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleDrop() {
	for value := range iter.Drop(slices.Values([]int{1, 2, 3, 4, 5}), 2) {
		fmt.Println(value)
	}

	// Output:
	// 3
	// 4
	// 5
}

func ExampleDrop_method() {
	for value := range iter.Iterator[int](slices.Values([]int{1, 2, 3, 4, 5})).Drop(2) {
		fmt.Println(value)
	}

	// Output:
	// 3
	// 4
	// 5
}

func TestDropTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull(iter.Drop(slices.Values([]int{1, 2, 3}), 2))
	stop()
}

func TestDropEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(iter.Drop(slices.Values([]int{}), 2)))
}
