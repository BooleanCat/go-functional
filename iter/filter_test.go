package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/iter/filters"
	"github.com/BooleanCat/go-functional/option"
)

func ExampleFilter() {
	filtered := iter.Filter[int](iter.Lift([]int{0, 1, 0, 2}), filters.IsZero[int])
	fmt.Println(filtered.Next())
	fmt.Println(filtered.Next())
	fmt.Println(filtered.Next())

	// Output:
	// Some(0)
	// Some(0)
	// None
}

func ExampleFilterMap() {
	selectAndTripleOdds := func(x int) option.Option[int] {
		if x%2 == 0 {
			return option.None[int]()
		}
		return option.Some(x * 3)
	}

	triples := iter.FilterMap[int](
		iter.Take[int](iter.Count(), 6),
		selectAndTripleOdds,
	)

	fmt.Println(iter.Collect(triples))

	// Output: [3 9 15]
}

func ExampleExclude() {
	filtered := iter.Exclude[int](iter.Lift([]int{0, 1, 0, 2}), filters.IsZero[int])
	fmt.Println(filtered.Next())
	fmt.Println(filtered.Next())
	fmt.Println(filtered.Next())

	// Output:
	// Some(1)
	// Some(2)
	// None
}

func TestFilter(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	evens := iter.Filter[int](iter.Count(), isEven)
	assert.Equal(t, evens.Next().Unwrap(), 0)
	assert.Equal(t, evens.Next().Unwrap(), 2)
}

func TestFilterEmpty(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	evens := iter.Filter[int](iter.Exhausted[int](), isEven)
	assert.True(t, evens.Next().IsNone())
}

func TestExclude(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	evens := iter.Exclude[int](iter.Count(), isEven)
	assert.Equal(t, evens.Next().Unwrap(), 1)
	assert.Equal(t, evens.Next().Unwrap(), 3)
}

func TestFilterMap(t *testing.T) {
	selectEvenAndDouble := func(x int) option.Option[int] {
		if x%2 > 0 {
			return option.None[int]()
		}

		return option.Some(x * 2)
	}

	fltMap := iter.FilterMap[int](
		iter.Lift([]int{1, 2, 3, 4, 5, 6}),
		selectEvenAndDouble,
	)
	result := iter.Collect(fltMap)

	assert.SliceEqual(t, result, []int{4, 8, 12})
}

func TestFilterMapEmpty(t *testing.T) {
	selectEvenAndDouble := func(x int) option.Option[int] {
		if x%2 > 0 {
			return option.None[int]()
		}

		return option.Some(x * 2)
	}

	fltMapEmpty := iter.FilterMap[int](
		iter.Exhausted[int](),
		selectEvenAndDouble,
	)

	assert.Empty(t, iter.Collect(fltMapEmpty))
}
