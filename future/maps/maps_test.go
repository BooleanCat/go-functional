package maps_test

import (
	"fmt"
	"iter"
	"sort"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/maps"
	"github.com/BooleanCat/go-functional/v2/internal/assert"
)

func ExampleAll() {
	for key, value := range maps.All(map[int]string{1: "one", 2: "two", 3: "three"}) {
		fmt.Println(key, value)
	}
}

type keyValuePair[K comparable, V any] struct {
	key   K
	value V
}

func TestAll(t *testing.T) {
	t.Parallel()

	values := make([]keyValuePair[int, string], 0, 3)

	for key, value := range maps.All(map[int]string{1: "one", 2: "two", 3: "three"}) {
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

func TestAllEmpty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(maps.All(map[int]string{}))), 0)
}

func TestAllTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull2(maps.All(map[int]string{1: "one", 2: "two", 3: "three"}))
	stop()
}

func ExampleCollect() {
	numbers := maps.Collect(maps.All(map[int]string{1: "one", 2: "two"}))

	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	// Output
	// one
	// two
}

func TestCollect(t *testing.T) {
	t.Parallel()

	numbers := maps.Collect(maps.All(map[int]string{1: "one", 2: "two"}))

	assert.Equal(t, numbers[1], "one")
	assert.Equal(t, numbers[2], "two")
}

func TestCollectEmpty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(maps.All(map[int]string{}))), 0)
}
