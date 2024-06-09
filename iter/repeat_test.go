package iter_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/slices"
	fn "github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleRepeat() {
	numbers := slices.Collect(fn.Take(fn.Repeat(42), 2))

	fmt.Println(numbers)
	// Output: [42 42]
}

func TestRepeatTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull(fn.Repeat(42))
	stop()
}
