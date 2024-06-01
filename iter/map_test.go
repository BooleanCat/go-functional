package iter_test

import (
	"fmt"
	it "iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleMap() {
	double := func(n int) int { return n * 2 }

	for number := range iter.Map(iter.Iterator[int](slices.Values([]int{1, 2, 3})), double) {
		fmt.Println(number)
	}

	// Output:
	// 2
	// 4
	// 6
}

func TestMapEmpty(t *testing.T) {
	t.Parallel()

	for _ = range iter.Map(iter.Iterator[int](slices.Values([]int{})), func(int) int { return 0 }) {
		t.Error("unexpected")
	}
}

func TestMapTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull(it.Seq[int](iter.Map(iter.Iterator[int](slices.Values([]int{1, 2, 3})), func(int) int { return 0 })))
	stop()
}
