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

func TestExhaustedForEach(t *testing.T) {
	total := 0

	iter.Exhausted[int]().ForEach(func(number int) {
		total += number
	})

	assert.Equal(t, total, 0)
}

func TestExhaustedFind(t *testing.T) {
	assert.True(t, iter.Exhausted[int]().Find(func(number int) bool {
		return number == 0
	}).IsNone())
}

func TestExhaustedDrop(t *testing.T) {
	assert.True(t, iter.Exhausted[int]().Drop(1).Next().IsNone())
}

func TestExhaustedTake(t *testing.T) {
	assert.Empty[int](t, iter.Exhausted[int]().Take(1).Collect())
}
