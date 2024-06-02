package iter_test

import (
	"fmt"
	it "iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/maps"
	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleChain() {
	numbers := slices.Collect(iter.Chain(slices.Values([]int{1, 2}), slices.Values([]int{3, 4})))

	fmt.Println(numbers)
	// Output: [1 2 3 4]
}

func ExampleChain_method() {
	numbers := iter.Iterator[int](slices.Values([]int{1, 2})).Chain(slices.Values([]int{3, 4})).Collect()

	fmt.Println(numbers)
	// Output: [1 2 3 4]
}

func TestChainEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(iter.Chain[int]()))
}

func TestChainMany(t *testing.T) {
	t.Parallel()

	numbers := slices.Collect(iter.Chain(
		slices.Values([]int{1, 2}),
		iter.Take(iter.Drop(iter.Count[int](), 3), 2),
		slices.Values([]int{5, 6}),
	))

	assert.SliceEqual(t, []int{1, 2, 3, 4, 5, 6}, numbers)
}

func TestChainTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull(iter.Chain(slices.Values([]int{1, 2}), slices.Values([]int{3, 4})))
	stop()
}

func ExampleChain2() {
	pairs := maps.Collect(iter.Chain2(maps.All(map[string]int{"a": 1}), maps.All(map[string]int{"b": 2})))

	fmt.Println(len(pairs))
	// Output: 2
}

func ExampleChain2_method() {
	pairs := iter.Iterator2[string, int](maps.All(map[string]int{"a": 1})).Chain(maps.All(map[string]int{"b": 2}))

	fmt.Println(len(maps.Collect(it.Seq2[string, int](pairs))))
	// Output: 2
}

func TestChain2(t *testing.T) {
	t.Parallel()

	pairs := maps.Collect(iter.Chain2(maps.All(map[string]int{"a": 1}), maps.All(map[string]int{"b": 2})))

	assert.Equal(t, 2, len(pairs))
	assert.Equal(t, pairs["a"], 1)
	assert.Equal(t, pairs["b"], 2)
}

func TestChain2Empty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(iter.Chain2[int, int]())), 0)
}

func TestChain2Many(t *testing.T) {
	t.Parallel()

	pairs := maps.Collect(iter.Chain2(
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

	_, stop := it.Pull2(iter.Chain2(maps.All(map[string]int{"a": 1})))
	stop()
}
