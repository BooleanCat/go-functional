package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
)

func ExampleRepeat() {
	iter := iter.Repeat[int](42)
	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	// Output:
	// Some(42)
	// Some(42)
}

func TestRepeat(t *testing.T) {
	numbers := iter.Repeat[int](42)
	assert.Equal(t, numbers.Next().Unwrap(), 42)
	assert.Equal(t, numbers.Next().Unwrap(), 42)
	assert.Equal(t, numbers.Next().Unwrap(), 42)
}

func TestRepeatDrop(t *testing.T) {
	numbers := iter.Repeat[int](42).Drop(1)
	assert.Equal(t, numbers.Next().Unwrap(), 42)
}
