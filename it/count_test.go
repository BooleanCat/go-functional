package it_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/it"
)

func ExampleCount() {
	for i := range it.Count[int]() {
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

func TestCountTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull(it.Count[int]())
	stop()
}
