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

func ExampleZip() {
	for left, right := range iter.Zip(iter.Lift([]int{1, 2, 3}), iter.Lift([]string{"one", "two", "three"})) {
		fmt.Println(left, right)
	}

	// Output:
	// 1 one
	// 2 two
	// 3 three
}

func TestZipEmpty(t *testing.T) {
	t.Parallel()

	for _, _ = range iter.Zip(iter.Lift([]int{}), iter.Lift([]string{})) {
		t.Error("unexpected")
	}
}

func TestZipTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull2(it.Seq2[int, string](iter.Zip(iter.Lift([]int{1, 2}), iter.Lift([]string{"one", "two"}))))
	stop()
}
