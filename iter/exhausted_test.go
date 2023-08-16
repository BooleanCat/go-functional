package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
)

func ExampleExhausted() {
	fmt.Println(iter.Exhausted[int]().Next())
	// Output: None
}

func TestExhausted(t *testing.T) {
	assert.True(t, iter.Exhausted[int]().Next().IsNone())
}

func TestExhaustedCollect(t *testing.T) {
	assert.Empty[int](t, iter.Exhausted[int]().Collect())
}
