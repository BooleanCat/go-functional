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

func TestFilterCollect(t *testing.T) {
	evens := iter.Filter[int](iter.Lift([]int{0, 1, 2, 3}), filters.IsEven[int]).Collect()
	assert.SliceEqual(t, evens, []int{0, 2})
}

func TestFilterForEach(t *testing.T) {
	total := 0

	iter.Filter[int](iter.Lift([]int{1, 2, 3, 4}), filters.IsEven[int]).ForEach(func(number int) {
		total += number
	})

	assert.Equal(t, total, 6)
}

func TestFilterFind(t *testing.T) {
	assert.Equal(t, iter.Filter[int](iter.Count(), filters.IsEven[int]).Find(func(number int) bool {
		return number == 2
	}), option.Some(2))
}

func TestFilterDrop(t *testing.T) {
	zeros := iter.Filter[int](iter.Lift([]int{0, 1, 0, 0}), filters.IsZero[int]).Take(2).Collect()
	assert.SliceEqual(t, zeros, []int{0, 0})
}

func TestFilterTake(t *testing.T) {
	evens := iter.Filter[int](iter.Lift([]int{0, 1, 2, 3}), filters.IsEven[int]).Drop(1).Collect()
	assert.SliceEqual(t, evens, []int{2})
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

func TestExcludeCollect(t *testing.T) {
	odds := iter.Exclude[int](iter.Lift([]int{0, 1, 2, 3}), filters.IsEven[int]).Collect()
	assert.SliceEqual(t, odds, []int{1, 3})
}

func TestExcludeForEach(t *testing.T) {
	total := 0

	iter.Exclude[int](iter.Lift([]int{1, 2, 3, 4}), filters.IsEven[int]).ForEach(func(number int) {
		total += number
	})

	assert.Equal(t, total, 4)
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

func TestFilterMapCollect(t *testing.T) {
	doubles := iter.FilterMap[int](
		iter.Lift([]int{1, 2, 3, 4, 5, 6}),
		selectEvenAndDouble,
	).Collect()

	assert.SliceEqual(t, doubles, []int{4, 8, 12})
}

func TestFilterMapForEach(t *testing.T) {
	total := 0

	iter.FilterMap[int](iter.Lift([]int{1, 2, 3, 4}), selectEvenAndDouble).ForEach(func(number int) {
		total += number
	})

	assert.Equal(t, total, 12)
}

func TestFilterMapDrop(t *testing.T) {
	doubles := iter.FilterMap[int](
		iter.Lift([]int{1, 2, 3, 4, 5, 6}),
		selectEvenAndDouble,
	).Drop(1).Collect()

	assert.SliceEqual(t, doubles, []int{8, 12})
}

func TestFilterMapTake(t *testing.T) {
	doubles := iter.FilterMap[int](
		iter.Lift([]int{1, 2, 3, 4, 5, 6}),
		selectEvenAndDouble,
	).Take(2).Collect()

	assert.SliceEqual(t, doubles, []int{4, 8})
}

func selectEvenAndDouble(x int) option.Option[int] {
	if x%2 > 0 {
		return option.None[int]()
	}

	return option.Some(x * 2)
}
