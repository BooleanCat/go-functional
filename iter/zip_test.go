//go:build go1.23

package iter_test

import (
	"fmt"
	"iter"
	"strconv"
	"sync"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/maps"
	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/internal/assert"
	fn "github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleZip() {
	for left, right := range fn.Zip(slices.Values([]int{1, 2, 3}), slices.Values([]string{"one", "two", "three"})) {
		fmt.Println(left, right)
	}

	// Output:
	// 1 one
	// 2 two
	// 3 three
}

func TestZipEmpty(t *testing.T) {
	t.Parallel()

	assert.Equal(t, len(maps.Collect(fn.Zip(fn.Exhausted[int](), fn.Exhausted[string]()))), 0)
}

func TestZipTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := iter.Pull2(fn.Zip(slices.Values([]int{1, 2}), slices.Values([]string{"one", "two"})))
	stop()
}

func ExampleUnzip() {
	keys, values := fn.Unzip(maps.All(map[int]string{1: "one", 2: "two"}))

	for key := range keys {
		fmt.Println(key)
	}

	for value := range values {
		fmt.Println(value)
	}
}

func ExampleUnzip_method() {
	keys, values := fn.Iterator2[int, string](maps.All(map[int]string{1: "one", 2: "two"})).Unzip()

	for key := range keys {
		fmt.Println(key)
	}

	for value := range values {
		fmt.Println(value)
	}
}

func TestUnzip(t *testing.T) {
	zipped := fn.Zip(slices.Values([]int{1, 2, 3}), slices.Values([]string{"one", "two", "three"}))

	numbers, strings := fn.Unzip(zipped)

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

	zipped := fn.Zip(slices.Values(numbers), slices.Values(strings))

	numbersIter, stringsIter := fn.Unzip(zipped)

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

func TestUnzipTerminateEarly(t *testing.T) {
	t.Parallel()

	zipped := fn.Zip(slices.Values([]int{1, 2}), slices.Values([]string{"one", "two"}))

	numbers, strings := fn.Unzip(zipped)

	_, stop := iter.Pull(numbers)
	stop()

	_, stop = iter.Pull(strings)
	stop()
}

func TestUnzipTerminateLeftEarly(t *testing.T) {
	t.Parallel()

	numbers, strings := fn.Unzip(maps.All(map[int]string{1: "one", 2: "two"}))

	_, stop := iter.Pull(numbers)
	stop()

	assert.EqualElements(t, slices.Collect(strings), []string{"one", "two"})
}

func TestUnzipTerminateRightEarly(t *testing.T) {
	t.Parallel()

	numbers, strings := fn.Unzip(maps.All(map[int]string{1: "one", 2: "two"}))

	_, stop := iter.Pull(strings)
	stop()

	assert.EqualElements(t, slices.Collect(numbers), []int{1, 2})
}

func TestUnzipMethod(t *testing.T) {
	keys, values := fn.Iterator2[int, string](maps.All(map[int]string{1: "one"})).Unzip()

	assert.SliceEqual(t, keys.Collect(), []int{1})
	assert.SliceEqual(t, values.Collect(), []string{"one"})
}
