package iter_test

import (
	"fmt"
	it "iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleEnumerate() {
	for index, value := range iter.Enumerate(iter.Iterator[int](slices.Values([]int{1, 2, 3}))) {
		fmt.Println(index, value)
	}

	// Output:
	// 0 1
	// 1 2
	// 2 3
}

func ExampleEnumerate_method() {
	for index, value := range iter.Iterator[int](slices.Values([]int{1, 2, 3})).Enumerate() {
		fmt.Println(index, value)
	}

	// Output:
	// 0 1
	// 1 2
	// 2 3
}

func TestEnumerateTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull2(it.Seq2[int, int](iter.Enumerate(iter.Iterator[int](slices.Values([]int{1, 2})))))
	stop()
}
