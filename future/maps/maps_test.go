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

	for _, _ = range maps.All(map[int]string{}) {
		t.Error("unexpected")
	}
}

func TestAllTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull2(maps.All(map[int]string{1: "one", 2: "two", 3: "three"}))
	stop()
}
