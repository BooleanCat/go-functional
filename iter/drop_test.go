package iter_test

import (
	"fmt"
	it "iter"
	sl "slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/maps"
	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleDrop() {
	for value := range iter.Drop(slices.Values([]int{1, 2, 3, 4, 5}), 2) {
		fmt.Println(value)
	}

	// Output:
	// 3
	// 4
	// 5
}

func ExampleDrop_method() {
	for value := range iter.Iterator[int](slices.Values([]int{1, 2, 3, 4, 5})).Drop(2) {
		fmt.Println(value)
	}

	// Output:
	// 3
	// 4
	// 5
}

func TestDropTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull(iter.Drop(slices.Values([]int{1, 2, 3}), 2))
	stop()
}

func TestDropEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(iter.Drop(slices.Values([]int{}), 2)))
}

func ExampleDrop2() {
	numbers := maps.Collect(iter.Drop2(maps.All(map[int]string{1: "one", 2: "two", 3: "three"}), 1))

	fmt.Println(len(numbers))
	// Output: 2
}

func ExampleDrop2_method() {
	numbers := iter.Iterator2[int, string](maps.All(map[int]string{1: "one", 2: "two", 3: "three"})).Drop(1)

	fmt.Println(len(maps.Collect(it.Seq2[int, string](numbers))))
	// Output: 2
}

func TestDrop2(t *testing.T) {
	t.Parallel()

	keys := []int{1, 2, 3}
	values := []string{"one", "two", "three"}

	numbers := maps.Collect(iter.Drop2(iter.Zip(slices.Values(keys), slices.Values(values)), 1))

	assert.Equal(t, len(numbers), 2)

	for key := range numbers {
		assert.True(t, sl.Contains(keys, key))
	}
}

func TestDrop2Empty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(iter.Drop2(maps.All(map[int]string{}), 1))), 0)
}

func TestDrop2Zero(t *testing.T) {
	t.Parallel()

	numbers := maps.Collect(iter.Drop2(maps.All(map[int]string{1: "one", 2: "two", 3: "three"}), 0))

	assert.Equal(t, len(numbers), 3)
}

func TestDrop2TerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull2(iter.Drop2(maps.All(map[int]string{1: "one", 2: "two", 3: "three"}), 1))
	stop()
}
