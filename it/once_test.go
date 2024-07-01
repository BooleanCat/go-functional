package it_test

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/it"
)

func ExampleOnce() {
	fmt.Println(slices.Collect(it.Once(42)))
	// Output: [42]
}

func TestOnceYieldsFalse(t *testing.T) {
	t.Parallel()

	once := it.Once(42)

	once(func(int) bool {
		return false
	})
}

func ExampleOnce2() {
	fmt.Println(maps.Collect(it.Once2(1, 2)))
	// Output: map[1:2]
}

func TestOnce2YieldsFalse(t *testing.T) {
	t.Parallel()

	once := it.Once2(1, 2)

	once(func(int, int) bool {
		return false
	})
}
