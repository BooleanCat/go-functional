package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleCount() {
	for i := range iter.Count().Take(3) {
		fmt.Println(i)
	}

	// Output:
	// 0
	// 1
	// 2
}

func TestCount(t *testing.T) {
	numbers := iter.Count().Take(3).Collect()
	assert.SliceEqual(t, []int{0, 1, 2}, numbers)
}
