package iter_test

import (
	"fmt"
	it "iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleDrop() {
	for value := range iter.Drop(iter.Lift([]int{1, 2, 3, 4, 5}), 2) {
		fmt.Println(value)
	}

	// Output:
	// 3
	// 4
	// 5
}

func ExampleDrop_method() {
	for value := range iter.Lift([]int{1, 2, 3, 4, 5}).Drop(2) {
		fmt.Println(value)
	}

	// Output:
	// 3
	// 4
	// 5
}

func TestDropTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull(it.Seq[int](iter.Drop(iter.Lift([]int{1, 2, 3}), 2)))
	stop()
}

func TestDropEmpty(t *testing.T) {
	t.Parallel()

	for _ = range iter.Drop(iter.Lift([]int{}), 2) {
		t.Error("unexpected")
	}
}
