//go:build go1.23

package iter_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/slices"
	fn "github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleEnumerate() {
	for index, value := range fn.Enumerate(slices.Values([]int{1, 2, 3})) {
		fmt.Println(index, value)
	}

	// Output:
	// 0 1
	// 1 2
	// 2 3
}

func ExampleEnumerate_method() {
	for index, value := range fn.Iterator[int](slices.Values([]int{1, 2, 3})).Enumerate() {
		fmt.Println(index, value)
	}

	// Output:
	// 0 1
	// 1 2
	// 2 3
}

func TestEnumerateTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull2(fn.Enumerate(slices.Values([]int{1, 2})))
	stop()
}
