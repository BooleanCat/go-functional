package it_test

import (
	"fmt"
	"iter"
	"maps"
	"slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/it"
)

func ExampleChain() {
	numbers := slices.Collect(it.Chain(slices.Values([]int{1, 2}), slices.Values([]int{3, 4})))

	fmt.Println(numbers)
	// Output: [1 2 3 4]
}

func TestChainEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(it.Chain[int]()))
}

func TestChainMany(t *testing.T) {
	t.Parallel()

	numbers := slices.Collect(it.Chain(
		slices.Values([]int{1, 2}),
		it.Take(it.Drop(it.Count[int](), 3), 2),
		slices.Values([]int{5, 6}),
	))

	assert.SliceEqual(t, []int{1, 2, 3, 4, 5, 6}, numbers)
}

func TestChainTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull(it.Chain(slices.Values([]int{1, 2}), slices.Values([]int{3, 4})))
	stop()
}

func ExampleChain2() {
	pairs := maps.Collect(it.Chain2(maps.All(map[string]int{"a": 1}), maps.All(map[string]int{"b": 2})))

	fmt.Println(len(pairs))
	// Output: 2
}

func TestChain2(t *testing.T) {
	t.Parallel()

	pairs := maps.Collect(it.Chain2(maps.All(map[string]int{"a": 1}), maps.All(map[string]int{"b": 2})))

	assert.Equal(t, 2, len(pairs))
	assert.Equal(t, pairs["a"], 1)
	assert.Equal(t, pairs["b"], 2)
}

func TestChain2Empty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(it.Chain2[int, int]())), 0)
}

func TestChain2Many(t *testing.T) {
	t.Parallel()

	pairs := maps.Collect(it.Chain2(
		maps.All(map[string]int{"a": 1}),
		maps.All(map[string]int{"b": 2}),
		maps.All(map[string]int{"c": 3}),
	))

	assert.Equal(t, 3, len(pairs))
	assert.Equal(t, pairs["a"], 1)
	assert.Equal(t, pairs["b"], 2)
	assert.Equal(t, pairs["c"], 3)
}

func TestChain2TerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull2(it.Chain2(maps.All(map[string]int{"a": 1})))
	stop()
}
