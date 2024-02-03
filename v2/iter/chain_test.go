package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleChain() {
	for value := range iter.Chain(iter.Repeat(42).Take(2), iter.Count().Take(2)) {
		fmt.Println(value)
	}

	// Output:
	// 42
	// 42
	// 0
	// 1
}

func ExampleChain_method() {
	for value := range iter.Repeat(42).Take(2).Chain(iter.Count().Take(2)) {
		fmt.Println(value)
	}

	// Output:
	// 42
	// 42
	// 0
	// 1
}

func TestChain(t *testing.T) {
	numbers := iter.Repeat(42).Take(2).Chain(iter.Count().Take(2)).Collect()
	assert.SliceEqual(t, []int{42, 42, 0, 1}, numbers)
}

func TestChainFirstEmpty(t *testing.T) {
	numbers := iter.Exhausted[int]().Chain(iter.Count().Take(2)).Collect()
	assert.SliceEqual(t, numbers, []int{0, 1})
}

func TestChainNextEmpty(t *testing.T) {
	numbers := iter.Count().Take(2).Chain(iter.Exhausted[int]()).Collect()
	assert.SliceEqual(t, numbers, []int{0, 1})
}

func TestChainNone(t *testing.T) {
	numbers := iter.Chain[int]().Collect()
	assert.Empty[int](t, numbers)
}
