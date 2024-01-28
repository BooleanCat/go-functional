package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/fakes"
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

func ExampleFilter_method() {
	filtered := iter.Lift([]int{0, 1, 0, 2}).Filter(filters.IsZero[int])
	fmt.Println(filtered.Next())
	fmt.Println(filtered.Next())
	fmt.Println(filtered.Next())

	// Output:
	// Some(0)
	// Some(0)
	// None
}

func ExampleFilterIter_String() {
	fmt.Println(iter.Filter[int](iter.Exhausted[int](), filters.IsZero[int]))
	// Output: Iterator<Filter, type=int>
}

func ExampleFilterMap() {
	selectAndTripleOdds := func(x int) option.Option[int] {
		if x%2 == 0 {
			return option.None[int]()
		}
		return option.Some(x * 3)
	}

	triples := iter.FilterMap[int](
		iter.Count().Take(6),
		selectAndTripleOdds,
	)

	fmt.Println(triples.Collect())
	// Output: [3 9 15]
}

func ExampleFilterMap_method() {
	selectAndTripleOdds := func(x int) option.Option[int] {
		if x%2 == 0 {
			return option.None[int]()
		}
		return option.Some(x * 3)
	}

	triples := iter.Count().Take(6).FilterMap(selectAndTripleOdds)

	fmt.Println(triples.Collect())
	// Output: [3 9 15]
}

func ExampleFilterMapIter_String() {
	fmt.Println(iter.FilterMap[int](iter.Exhausted[int](), func(_ int) option.Option[int] {
		return option.Some(42)
	}))
	// Output: Iterator<FilterMap, type=int>
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

func ExampleExclude_method() {
	filtered := iter.Lift([]int{0, 1, 0, 2}).Exclude(filters.IsZero[int])
	fmt.Println(filtered.Next())
	fmt.Println(filtered.Next())
	fmt.Println(filtered.Next())

	// Output:
	// Some(1)
	// Some(2)
	// None
}

func TestFilter(t *testing.T) {
	evens := iter.Filter[int](iter.Count(), filters.IsEven[int])
	assert.Equal(t, evens.Next().Unwrap(), 0)
	assert.Equal(t, evens.Next().Unwrap(), 2)
}

func TestFilterEmpty(t *testing.T) {
	evens := iter.Filter[int](iter.Exhausted[int](), filters.IsEven[int])
	assert.True(t, evens.Next().IsNone())
}

func TestFilterExhausted(t *testing.T) {
	delegate := new(fakes.Iterator[int])
	ones := iter.Filter[int](delegate, func(_ int) bool { return true })

	assert.True(t, ones.Next().IsNone())
	assert.True(t, ones.Next().IsNone())
	assert.Equal(t, delegate.NextCallCount(), 1)
}

func TestFilter_String(t *testing.T) {
	filter := iter.Filter[int](iter.Count(), filters.IsEven[int])
	assert.Equal(t, fmt.Sprintf("%s", filter), "Iterator<Filter, type=int>")  //nolint:gosimple
	assert.Equal(t, fmt.Sprintf("%s", *filter), "Iterator<Filter, type=int>") //nolint:gosimple
}

func TestExclude(t *testing.T) {
	evens := iter.Exclude[int](iter.Count(), filters.IsEven[int])
	assert.Equal(t, evens.Next().Unwrap(), 1)
	assert.Equal(t, evens.Next().Unwrap(), 3)
}

func TestExcludeExhausted(t *testing.T) {
	delegate := new(fakes.Iterator[int])
	ones := iter.Exclude[int](delegate, func(_ int) bool { return false })

	assert.True(t, ones.Next().IsNone())
	assert.True(t, ones.Next().IsNone())
	assert.Equal(t, delegate.NextCallCount(), 1)
}

func TestExcludeFind(t *testing.T) {
	assert.Equal(t, iter.Exclude[int](iter.Count(), filters.IsEven[int]).Find(func(number int) bool {
		return number == 3
	}), option.Some(3))
}

func TestExcludeDrop(t *testing.T) {
	odds := iter.Exclude[int](iter.Lift([]int{0, 1, 2, 3}), filters.IsEven[int]).Drop(1).Collect()
	assert.SliceEqual(t, odds, []int{3})
}

func TestExcludeTake(t *testing.T) {
	evens := iter.Exclude[int](iter.Lift([]int{0, 1, 2, 3}), filters.IsEven[int]).Take(2).Collect()
	assert.SliceEqual(t, evens, []int{1, 3})
}

func TestFilterMap(t *testing.T) {
	numbers := iter.FilterMap[int](
		iter.Lift([]int{1, 2, 3, 4, 5, 6}),
		selectEvenAndDouble,
	)

	assert.SliceEqual(t, numbers.Collect(), []int{4, 8, 12})
}

func TestFilterMapEmpty(t *testing.T) {
	numbers := iter.FilterMap[int](
		iter.Exhausted[int](),
		selectEvenAndDouble,
	)

	assert.Empty[int](t, numbers.Collect())
}

func TestFilterMapExhausted(t *testing.T) {
	delegate := new(fakes.Iterator[int])
	ones := iter.FilterMap[int](delegate, func(_ int) option.Option[int] { return option.Some(1) })

	assert.True(t, ones.Next().IsNone())
	assert.True(t, ones.Next().IsNone())
	assert.Equal(t, delegate.NextCallCount(), 1)
}

func TestFilterMap_String(t *testing.T) {
	filterMap := iter.FilterMap[int](
		iter.Lift([]int{1, 2, 3, 4, 5, 6}),
		selectEvenAndDouble,
	)

	assert.Equal(t, fmt.Sprintf("%s", filterMap), "Iterator<FilterMap, type=int>")  //nolint:gosimple
	assert.Equal(t, fmt.Sprintf("%s", *filterMap), "Iterator<FilterMap, type=int>") //nolint:gosimple
}

func selectEvenAndDouble(x int) option.Option[int] {
	if x%2 > 0 {
		return option.None[int]()
	}

	return option.Some(x * 2)
}
