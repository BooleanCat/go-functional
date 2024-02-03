package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleRepeat() {
	for value := range iter.Repeat(42).Take(3) {
		fmt.Println(value)
	}
	// Output:
	// 42
	// 42
	// 42
}

func TestRepeat(t *testing.T) {
	numbers := iter.Repeat(42).Take(3).Collect()
	assert.SliceEqual(t, []int{42, 42, 42}, numbers)
}
