package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/option"
)

func ExampleFilterMap() {
	lengthIfGTFive := func(s string) option.Option[int] {
		if len(s) > 5 {
			return option.Some(len(s))
		} else {
			return option.None[int]()
		}
	}
	filtered := iter.FilterMap[string](iter.Lift([]string{"foo", "axolotl", "bort", "cowabunga"}), lengthIfGTFive)
	fmt.Println(filtered.Next())
	fmt.Println(filtered.Next())
	fmt.Println(filtered.Next())

	// Output:
	// Some(7)
	// Some(9)
	// None
}

func TestFilterMap(t *testing.T) {
	lookupAge := func(name string) option.Option[int] {
		if name == "william" {
			return option.Some(32)
		} else if name == "methuselah" {
			return option.Some(969)
		} else {
			return option.None[int]()
		}
	}
	ages := iter.FilterMap[string](iter.Lift([]string{"william", "bort", "methuselah"}), lookupAge)
	assert.Equal(t, ages.Next().Unwrap(), 32)
	assert.Equal(t, ages.Next().Unwrap(), 969)
}

func TestFilterMapEmpty(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	evens := iter.FilterMap[int](iter.Exhausted[int](), option.FromPredicate(isEven))
	assert.True(t, evens.Next().IsNone())
}
