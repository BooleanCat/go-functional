package iter_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/iter/ops"
	"github.com/BooleanCat/go-functional/option"
)

func ExampleCollect() {
	fmt.Println(iter.Collect[int](iter.Count().Take(3)))
	// Output: [0 1 2]
}

func ExampleCollect_method() {
	fmt.Println(iter.Count().Take(3).Collect())
	// Output: [0 1 2]
}

func ExampleFold() {
	sum := iter.Fold[int](iter.Count().Take(4), 0, ops.Add[int])

	fmt.Println(sum)
	// Output: 6
}

func ExampleToChannel() {
	for number := range iter.ToChannel[int](iter.Lift([]int{1, 2, 3})) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleFind() {
	values := iter.Lift([]string{"foo", "bar", "baz"})
	bar := iter.Find[string](values, func(v string) bool { return v == "bar" })

	fmt.Println(bar)
	// Output: Some(bar)
}

func TestCollect(t *testing.T) {
	items := iter.Collect[int](iter.Count().Take(5))
	assert.SliceEqual(t, items, []int{0, 1, 2, 3, 4})
}

func TestCollectEmpty(t *testing.T) {
	items := iter.Collect[int](iter.Exhausted[int]())
	assert.Empty[int](t, items)
}

func TestFold(t *testing.T) {
	add := func(a, b int) int { return a + b }
	sum := iter.Fold[int](iter.Count().Take(11), 0, add)
	assert.Equal(t, sum, 55)

	concat := func(path string, part int) string {
		return path + strconv.Itoa(part) + "/"
	}
	result := iter.Fold[int](iter.Count().Take(3), "/", concat)
	assert.Equal(t, result, "/0/1/2/")
}

func TestToChannel(t *testing.T) {
	expected := 0
	for number := range iter.ToChannel[int](iter.Lift([]int{1, 2, 3, 4})) {
		expected += 1
		assert.Equal(t, number, expected)
	}
}

func TestToChannelEmpty(t *testing.T) {
	for range iter.ToChannel[int](iter.Exhausted[int]()) {
		t.Fail()
	}
}

func TestForEach(t *testing.T) {
	words := iter.Lift([]string{"foo", "bar", "baz"})
	sum := ""
	iter.ForEach[string](words, func(word string) {
		sum += word
	})
	assert.Equal(t, "foobarbaz", sum)
}

func TestForEachEmpty(t *testing.T) {
	words := iter.Lift([]string{})
	sum := ""
	iter.ForEach[string](words, func(word string) {
		sum += word
	})

	assert.Empty[string](t, sum)
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
