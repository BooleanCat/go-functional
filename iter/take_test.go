package iter_test

import (
	"fmt"
	it "iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleTake() {
	for number := range iter.Take(iter.Lift([]int{1, 2, 3, 4, 5}), 3) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleTake_method() {
	for number := range iter.Lift([]int{1, 2, 3, 4, 5}).Take(3) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func TestTakeTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull(it.Seq[int](iter.Take(iter.Lift([]int{1, 2, 3}), 2)))
	stop()
}

func TestTakeZero(t *testing.T) {
	t.Parallel()

	for _ = range iter.Take(iter.Lift([]int{1, 2, 3}), 0) {
		t.Error("unexpected")
	}
}
