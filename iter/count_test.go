package iter_test

import (
	"fmt"
	it "iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleCount() {
	for i := range iter.Count[int]() {
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

	_, stop := it.Pull(iter.Count[int]())
	stop()
}
