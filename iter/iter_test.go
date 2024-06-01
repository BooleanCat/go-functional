package iter_test

import (
	"fmt"
	it "iter"
	"sort"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
	"github.com/BooleanCat/go-functional/v2/iter/op"
)

func ExampleCollect_method() {
	fmt.Println(iter.Iterator[int](slices.Values([]int{1, 2, 3})).Collect())
	// Output: [1 2 3]
}

func ExampleLiftHashMap() {
	for key, value := range iter.LiftHashMap(map[int]string{1: "one", 2: "two", 3: "three"}) {
		fmt.Println(key, value)
	}
}

type keyValuePair[K comparable, V any] struct {
	key   K
	value V
}

func TestLiftHashMap(t *testing.T) {
	t.Parallel()

	values := make([]keyValuePair[int, string], 0, 3)

	for key, value := range iter.LiftHashMap(map[int]string{1: "one", 2: "two", 3: "three"}) {
		values = append(values, keyValuePair[int, string]{key, value})
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i].key < values[j].key
	})

	assert.SliceEqual(t, values, []keyValuePair[int, string]{
		{1, "one"},
		{2, "two"},
		{3, "three"},
	})
}

func TestLiftHashMapEmpty(t *testing.T) {
	t.Parallel()

	for _, _ = range iter.LiftHashMap(map[int]string{}) {
		t.Error("unexpected")
	}
}

func TestLiftHashMapTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull2(it.Seq2[int, string](iter.LiftHashMap(map[int]string{1: "one", 2: "two", 3: "three"})))
	stop()
}

func ExampleForEach() {
	iter.ForEach(iter.Iterator[int](slices.Values([]int{1, 2, 3})), func(number int) {
		fmt.Println(number)
	})
	// Output:
	// 1
	// 2
	// 3
}

func ExampleForEach_method() {
	iter.Iterator[int](slices.Values([]int{1, 2, 3})).ForEach(func(number int) {
		fmt.Println(number)
	})
	// Output:
	// 1
	// 2
	// 3
}

func TestForEachEmpty(t *testing.T) {
	t.Parallel()

	iter.ForEach(iter.Iterator[int](slices.Values([]int{})), func(int) {
		t.Error("unexpected")
	})
}

func ExampleForEach2() {
	iter.ForEach2(iter.Iterator[int](slices.Values([]int{1, 2, 3})).Enumerate(), func(index int, number int) {
		fmt.Println(index, number)
	})
	// Output:
	// 0 1
	// 1 2
	// 2 3
}

func ExampleForEach2_method() {
	iter.Iterator[int](slices.Values([]int{1, 2, 3})).Enumerate().ForEach2(func(index int, number int) {
		fmt.Println(index, number)
	})
	// Output:
	// 0 1
	// 1 2
	// 2 3
}

func TestForEach2Empty(t *testing.T) {
	t.Parallel()

	iter.ForEach2(iter.Iterator[int](slices.Values([]int{})).Enumerate(), func(int, int) {
		t.Error("unexpected")
	})
}

func ExampleReduce() {
	fmt.Println(iter.Reduce(iter.Iterator[int](slices.Values([]int{1, 2, 3})), op.Add, 0))
	// Output: 6
}

func TestReduceEmpty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, iter.Reduce(iter.Iterator[int](slices.Values([]int{})), func(int, int) int { return 0 }, 0), 0)
}

func ExampleReduce2() {
	fmt.Println(iter.Reduce2(iter.Iterator[int](slices.Values([]int{1, 2, 3})).Enumerate(), func(i, a, b int) int {
		return i + 1
	}, 0))

	// Output: 3
}
