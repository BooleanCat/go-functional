package it_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/v2/it"
)

func ExampleNaturalNumbers() {
	for i := range it.NaturalNumbers[int]() {
		if i >= 3 {
			break
		}

		fmt.Println(i)
	}

	// Output:
	// 0
	// 1
	// 2
}

func ExampleIntegers() {
	for i := range it.Integers[uint](0, 5, 2) {
		fmt.Println(i)
	}

	// Output:
	// 0
	// 2
	// 4
}

func TestIntegersYieldFalse(t *testing.T) {
	t.Parallel()

	numbers := it.Integers[uint](0, 3, 1)

	numbers(func(value uint) bool {
		return false
	})
}
