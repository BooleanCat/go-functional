package iter_test

import (
	"fmt"
	"iter"
	"testing"

	fn "github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleCount() {
	for i := range fn.Count[int]() {
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

	_, stop := iter.Pull(fn.Count[int]())
	stop()
}
