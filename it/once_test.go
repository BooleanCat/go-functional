package it_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/v2/it"
)

func ExampleOnce() {
	for number := range it.Once(42) {
		fmt.Println(number)
	}

	// Output: 42
}

func TestOnceYieldsFalse(t *testing.T) {
	t.Parallel()

	once := it.Once(42)

	once(func(int) bool {
		return false
	})
}

func ExampleOnce2() {
	for key, value := range it.Once2(1, 2) {
		fmt.Println(key, value)
	}

	// Output: 1 2
}

func TestOnce2YieldsFalse(t *testing.T) {
	t.Parallel()

	once := it.Once2(1, 2)

	once(func(int, int) bool {
		return false
	})
}
