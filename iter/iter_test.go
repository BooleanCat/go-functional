package iter_test

import (
	"fmt"
	it "iter"
	"sort"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleCollect() {
	fmt.Println(iter.Collect(iter.Lift([]int{1, 2, 3})))
	// Output: [1 2 3]
}

func ExampleCollect_method() {
	fmt.Println(iter.Lift([]int{1, 2, 3}).Collect())
	// Output: [1 2 3]
}

func TestCollectEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, iter.Collect(iter.Lift([]int{})))
}

func ExampleLift() {
	for number := range iter.Lift([]int{1, 2, 3}) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func TestLiftEmpty(t *testing.T) {
	t.Parallel()

	for number := range iter.Lift([]int{}) {
		t.Error("unexpected", number)
	}
}

func TestLiftTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull(it.Seq[int](iter.Lift([]int{1, 2, 3})))
	stop()
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

	assert.SliceEqual[keyValuePair[int, string]](t, values, []keyValuePair[int, string]{
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
	iter.ForEach(iter.Lift([]int{1, 2, 3}), func(number int) {
		fmt.Println(number)
	})
	// Output:
	// 1
	// 2
	// 3
}

func ExampleForEach_method() {
	iter.Lift([]int{1, 2, 3}).ForEach(func(number int) {
		fmt.Println(number)
	})
	// Output:
	// 1
	// 2
	// 3
}

func TestForEachEmpty(t *testing.T) {
	t.Parallel()

	iter.ForEach(iter.Lift([]int{}), func(int) {
		t.Error("unexpected")
	})
}

func ExampleForEach2() {
	iter.ForEach2(iter.Lift([]int{1, 2, 3}).Enumerate(), func(index int, number int) {
		fmt.Println(index, number)
	})
	// Output:
	// 0 1
	// 1 2
	// 2 3
}

func ExampleForEach2_method() {
	iter.Lift([]int{1, 2, 3}).Enumerate().ForEach2(func(index int, number int) {
		fmt.Println(index, number)
	})
	// Output:
	// 0 1
	// 1 2
	// 2 3
}

func TestForEach2Empty(t *testing.T) {
	t.Parallel()

	iter.ForEach2(iter.Lift([]int{}).Enumerate(), func(int, int) {
		t.Error("unexpected")
	})
}

func ExampleReduce() {
	fmt.Println(iter.Reduce(iter.Lift([]int{1, 2, 3}), func(a, b int) int {
		return a + b
	}, 0))

	// Output: 6
}

func TestReduceEmpty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, iter.Reduce(iter.Lift([]int{}), func(int, int) int { return 0 }, 0), 0)
}

func ExampleReduce2() {
	fmt.Println(iter.Reduce2(iter.Lift([]int{1, 2, 3}).Enumerate(), func(i, a, b int) int {
		return i + 1
	}, 0))

	// Output: 3
}
