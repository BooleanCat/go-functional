package it_test

import (
	"errors"
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/BooleanCat/go-functional/v2/it/filter"
)

func ExampleFilter() {
	for number := range it.Filter(slices.Values([]int{1, 2, 3, 4, 5}), filter.IsEven) {
		fmt.Println(number)
	}

	// Output:
	// 2
	// 4
}

func TestFilterEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty[int](t, slices.Collect(it.Filter(it.Exhausted[int](), filter.IsEven)))
}

func TestFilterYieldsFalse(t *testing.T) {
	t.Parallel()

	seq := it.Filter(slices.Values([]int{1, 2, 3, 4, 5}), filter.IsEven)

	var value int

	seq(func(v int) bool {
		value = v
		return false
	})

	assert.Equal(t, value, 2)
}

func ExampleExclude() {
	for number := range it.Exclude(slices.Values([]int{1, 2, 3, 4, 5}), filter.IsEven) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 3
	// 5
}

func ExampleFilter2() {
	isOne := func(n int, _ string) bool { return n == 1 }
	numbers := map[int]string{1: "one", 2: "two", 3: "three"}

	for key, value := range it.Filter2(maps.All(numbers), isOne) {
		fmt.Println(key, value)
	}

	// Output: 1 one
}

func TestFilter2Empty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(it.Filter2(it.Exhausted2[int, int](), filter.Passthrough2))), 0)
}

func TestFilter2YieldsFalse(t *testing.T) {
	t.Parallel()

	seq := it.Filter2(maps.All(map[int]string{1: "one", 2: "two"}), filter.Passthrough2)

	seq(func(k int, v string) bool {
		return false
	})
}

func ExampleExclude2() {
	isOne := func(n int, _ string) bool { return n == 1 }
	numbers := map[int]string{1: "one", 3: "three"}

	for key, value := range it.Exclude2(maps.All(numbers), isOne) {
		fmt.Println(key, value)
	}

	// Output: 3 three
}

func TestExclude2Empty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(it.Exclude2(it.Exhausted2[int, int](), filter.Passthrough2))), 0)
}

func ExampleFilterError() {
	isFoo := func(s string) (bool, error) { return s == "foo", nil }

	values := slices.Values([]string{"foo", "bar", "foo"})

	foos, err := it.TryCollect(it.FilterError(values, isFoo))
	fmt.Println(foos, err)

	// Output: [foo foo] <nil>
}

func TestFilterErrorYieldsFalse(t *testing.T) {
	t.Parallel()

	passthrough := func(s string) (bool, error) { return true, nil }

	seq := it.FilterError(slices.Values([]string{"foo", "bar", "foo"}), passthrough)

	seq(func(v string, _ error) bool {
		return false
	})
}

func TestFilterErrorError(t *testing.T) {
	t.Parallel()

	alwaysError := func(string) (bool, error) { return false, errors.New("oops") }

	_, err := it.TryCollect(it.FilterError(slices.Values([]string{"foo"}), alwaysError))
	assert.Equal(t, err.Error(), "oops")
}

func TestFilterErrorErrorSometimes(t *testing.T) {
	t.Parallel()

	sometimesError := func(s string) (bool, error) {
		if s == "foo" {
			return true, errors.New("oops")
		}

		return true, nil
	}

	values, errs := it.Collect2(it.FilterError(slices.Values([]string{"foo", "bar"}), sometimesError))
	assert.SliceEqual(t, values, []string{"", "bar"})

	assert.Equal(t, len(errs), 2)
	assert.Equal(t, errs[0].Error(), "oops")
	assert.Equal(t, errs[1], nil)
}

func ExampleExcludeError() {
	isFoo := func(s string) (bool, error) { return s == "foo", nil }

	values := slices.Values([]string{"foo", "bar", "foo"})

	bars, err := it.TryCollect(it.ExcludeError(values, isFoo))
	fmt.Println(bars, err)

	// Output: [bar] <nil>
}
