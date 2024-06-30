package it_test

import (
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/it"
)

func ExampleRepeat() {
	numbers := slices.Collect(it.Take(it.Repeat(42), 2))

	fmt.Println(numbers)
	// Output: [42 42]
}

func TestRepeatTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull(it.Repeat(42))
	stop()
}

func ExampleRepeat2() {
	numbers := it.Take2(it.Repeat2(42, "Life"), 2)

	for v, w := range numbers {
		fmt.Println(v, w)
	}

	// Output:
	// 42 Life
	// 42 Life
}

func TestRepeat2TerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull2(it.Repeat2(42, "Life"))
	stop()
}
