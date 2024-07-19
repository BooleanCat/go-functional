package it_test

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/it"
)

func ExampleZip() {
	numbers := slices.Values([]int{1, 2, 3})
	strings := slices.Values([]string{"one", "two", "three"})

	for left, right := range it.Zip(numbers, strings) {
		fmt.Println(left, right)
	}

	// Output:
	// 1 one
	// 2 two
	// 3 three
}

func TestZipEmpty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(it.Zip(it.Exhausted[int](), it.Exhausted[string]()))), 0)
}

func ExampleLeft() {
	for left := range it.Left(maps.All(map[int]string{1: "one"})) {
		fmt.Println(left)
	}

	// Output: 1
}

func ExampleRight() {
	for right := range it.Right(maps.All(map[int]string{1: "one"})) {
		fmt.Println(right)
	}

	// Output: one
}

func TestLeftYieldFalse(t *testing.T) {
	t.Parallel()

	numbers := it.Left(maps.All(map[int]string{1: "one"}))

	numbers(func(value int) bool {
		return false
	})
}

func TestRightYieldFalse(t *testing.T) {
	t.Parallel()

	strings := it.Right(maps.All(map[int]string{1: "one"}))

	strings(func(value string) bool {
		return false
	})
}
