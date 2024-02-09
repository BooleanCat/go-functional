package iter_test

import (
	"fmt"
	it "iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleCollect() {
	fmt.Println(iter.Collect(iter.Lift([]int{1, 2, 3})))
	// Output: [1 2 3]
}

func ExampleCollect_method() {
	fmt.Println(iter.Lift([]int{1, 2, 3}).Collect())
	// Output: [1 2 3]
}

func TestCollectEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, iter.Collect(iter.Lift([]int{})))
}

func ExampleLift() {
	for number := range iter.Lift([]int{1, 2, 3}) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func TestLiftEmpty(t *testing.T) {
	t.Parallel()

	for number := range iter.Lift([]int{}) {
		t.Error("unexpected", number)
	}
}

func TestLiftTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull(it.Seq[int](iter.Lift([]int{1, 2, 3})))
	stop()
}
