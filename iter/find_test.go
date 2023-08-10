package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/option"
)

func ExampleFind() {
	values := iter.Lift([]string{"foo", "bar", "baz"})
	bar := iter.Find[string](values, func(v string) bool { return v == "bar" })

	fmt.Println(bar)
	// Output: Some(bar)
}

func TestFind(t *testing.T) {
	values := iter.Lift([]string{"foo", "bar", "baz"})
	bar := iter.Find[string](values, func(v string) bool { return v == "bar" })

	assert.Equal(t, bar, option.Some("bar"))
	assert.Equal(t, values.Next().Unwrap(), "baz")
}

func TestFindEmpty(t *testing.T) {
	values := iter.Exhausted[int]()
	found := iter.Find[int](values, func(v int) bool { return v == 0 })

	assert.True(t, found.IsNone())
}
