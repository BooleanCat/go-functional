package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleLift() {
	for i := range iter.Lift([]int{1, 2}) {
		fmt.Println(i)
	}

	// Output:
	// 1
	// 2
}

func TestLift(t *testing.T) {
	number := 0

	for i := range iter.Lift([]int{1, 2}) {
		assert.Equal(t, number+1, i)
		number++
	}
}

func TestLiftEmpty(t *testing.T) {
	for _ = range iter.Lift([]int{}) {
		t.Error("expected no iteration")
	}
}

func ExampleCollect() {
	fmt.Println(iter.Collect(iter.Lift([]int{1, 2})))
	// Output: [1 2]
}

func ExampleCollect_method() {
	fmt.Println(iter.Lift([]int{1, 2}).Collect())
	// Output: [1 2]
}

func TestCollect(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3}).Collect()
	assert.SliceEqual(t, []int{1, 2, 3}, numbers)
}

func TestCollectEmpty(t *testing.T) {
	numbers := iter.Lift([]int{}).Collect()
	assert.Empty[int](t, numbers)
}
