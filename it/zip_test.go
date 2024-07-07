package it_test

import (
	"fmt"
	"iter"
	"maps"
	"slices"
	"strconv"
	"sync"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/it"
)

func ExampleZip() {
	for left, right := range it.Zip(slices.Values([]int{1, 2, 3}), slices.Values([]string{"one", "two", "three"})) {
		fmt.Println(left, right)
	}

	// Output:
	// 1 one
	// 2 two
	// 3 three
}

func TestZipEmpty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(it.Zip(it.Exhausted[int](), it.Exhausted[string]()))), 0)
}

func TestZipTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull2(it.Zip(slices.Values([]int{1, 2}), slices.Values([]string{"one", "two"})))
	stop()
}

func ExampleUnzip() {
	keys, values := it.Unzip(maps.All(map[int]string{1: "one", 2: "two"}))

	for key := range keys {
		fmt.Println(key)
	}

	for value := range values {
		fmt.Println(value)
	}
}

func TestUnzip(t *testing.T) {
	zipped := it.Zip(slices.Values([]int{1, 2, 3}), slices.Values([]string{"one", "two", "three"}))

	numbers, strings := it.Unzip(zipped)

	assert.SliceEqual(t, slices.Collect(numbers), []int{1, 2, 3})
	assert.SliceEqual(t, slices.Collect(strings), []string{"one", "two", "three"})
}

func TestUnzipRace(t *testing.T) {
	limit := 100000

	numbers := make([]int, 0, limit)
	strings := make([]string, 0, limit)

	for i := 0; i < limit; i++ {
		numbers = append(numbers, i)
		strings = append(strings, strconv.Itoa(i))
	}

	zipped := it.Zip(slices.Values(numbers), slices.Values(strings))

	numbersIter, stringsIter := it.Unzip(zipped)

	group := sync.WaitGroup{}
	group.Add(2)

	go func() {
		defer group.Done()

		i := 0
		for n := range numbersIter {
			assert.Equal(t, n, i)
			i++
		}
	}()

	go func() {
		defer group.Done()

		i := 0
		for s := range stringsIter {
			assert.Equal(t, s, strconv.Itoa(i))
			i++
		}
	}()

	group.Wait()
}

func TestUnzipYieldFalse(t *testing.T) {
	t.Parallel()

	zipped := it.Zip(slices.Values([]int{1, 2}), slices.Values([]string{"one", "two"}))

	numbers, strings := it.Unzip(zipped)

	numbers(func(int) bool { return false })
	strings(func(string) bool { return false })
}

func TestUnzipTerminateLeftEarly(t *testing.T) {
	t.Parallel()

	numbers, strings := it.Unzip(maps.All(map[int]string{1: "one", 2: "two"}))

	_, stop := iter.Pull(numbers)
	stop()

	assert.EqualElements(t, slices.Collect(strings), []string{"one", "two"})
}

func TestUnzipTerminateRightEarly(t *testing.T) {
	t.Parallel()

	numbers, strings := it.Unzip(maps.All(map[int]string{1: "one", 2: "two"}))

	_, stop := iter.Pull(strings)
	stop()

	assert.EqualElements(t, slices.Collect(numbers), []int{1, 2})
}

func ExampleLeft() {
	for left := range it.Left(maps.All(map[int]string{1: "one"})) {
		fmt.Println(left)
	}

	// Output: 1
}

func ExampleRight() {
	for right := range it.Right(maps.All(map[int]string{1: "one"})) {
		fmt.Println(right)
	}

	// Output: one
}
