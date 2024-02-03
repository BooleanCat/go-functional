package iter_test

import (
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleExhausted() {
	for _ = range iter.Exhausted[int]() {
	}
	// Output:
}

func TestExhausted(t *testing.T) {
	numbers := iter.Exhausted[int]().Collect()
	assert.Empty[int](t, numbers)
}
