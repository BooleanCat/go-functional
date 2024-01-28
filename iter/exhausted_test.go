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

func ExampleExhaustedIter_String() {
	fmt.Println(iter.Exhausted[int]())
	// Output: Iterator<Exhausted, type=int>
}

func TestExhausted(t *testing.T) {
	assert.True(t, iter.Exhausted[int]().Next().IsNone())
}

func TestExhaustedIter_String(t *testing.T) {
	exhausted := iter.Exhausted[int]()
	assert.Equal(t, fmt.Sprintf("%s", exhausted), "Iterator<Exhausted, type=int>") //nolint:gosimple
	assert.Equal(t, fmt.Sprintf("%s", exhausted), "Iterator<Exhausted, type=int>") //nolint:gosimple
}
