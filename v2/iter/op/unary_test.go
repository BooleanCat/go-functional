package op_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
	"github.com/BooleanCat/go-functional/v2/iter/op"
)

func ExampleIdentity() {
	for i := range iter.Lift([]int{1, 2, 3}).Transform(op.Identity) {
		fmt.Println(i)
	}

	// Output:
	// 1
	// 2
	// 3
}

func TestIdentity(t *testing.T) {
	numbers := iter.Lift([]int{1, 2, 3}).Transform(op.Identity).Collect()
	assert.SliceEqual(t, []int{1, 2, 3}, numbers)
}
