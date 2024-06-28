//go:build go1.23

package iter_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/maps"
	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/internal/assert"
	fn "github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleChain() {
	numbers := slices.Collect(fn.Chain(slices.Values([]int{1, 2}), slices.Values([]int{3, 4})))

	fmt.Println(numbers)
	// Output: [1 2 3 4]
}

func ExampleChain_method() {
	numbers := fn.Iterator[int](slices.Values([]int{1, 2})).Chain(slices.Values([]int{3, 4})).Collect()

	fmt.Println(numbers)
	// Output: [1 2 3 4]
}

func TestChainEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(fn.Chain[int]()))
}

func TestChainMany(t *testing.T) {
	t.Parallel()

	numbers := slices.Collect(fn.Chain(
		slices.Values([]int{1, 2}),
		fn.Take(fn.Drop(fn.Count[int](), 3), 2),
		slices.Values([]int{5, 6}),
	))

	assert.SliceEqual(t, []int{1, 2, 3, 4, 5, 6}, numbers)
}

func TestChainTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull(fn.Chain(slices.Values([]int{1, 2}), slices.Values([]int{3, 4})))
	stop()
}

func ExampleChain2() {
	pairs := maps.Collect(fn.Chain2(maps.All(map[string]int{"a": 1}), maps.All(map[string]int{"b": 2})))

	fmt.Println(len(pairs))
	// Output: 2
}

func ExampleChain2_method() {
	pairs := fn.Iterator2[string, int](maps.All(map[string]int{"a": 1})).Chain(maps.All(map[string]int{"b": 2}))

	fmt.Println(len(maps.Collect(iter.Seq2[string, int](pairs))))
	// Output: 2
}

func TestChain2(t *testing.T) {
	t.Parallel()

	pairs := maps.Collect(fn.Chain2(maps.All(map[string]int{"a": 1}), maps.All(map[string]int{"b": 2})))

	assert.Equal(t, 2, len(pairs))
	assert.Equal(t, pairs["a"], 1)
	assert.Equal(t, pairs["b"], 2)
}

func TestChain2Empty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(fn.Chain2[int, int]())), 0)
}

func TestChain2Many(t *testing.T) {
	t.Parallel()

	pairs := maps.Collect(fn.Chain2(
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

	_, stop := iter.Pull2(fn.Chain2(maps.All(map[string]int{"a": 1})))
	stop()
}
