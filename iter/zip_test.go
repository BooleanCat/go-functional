package iter_test

import (
	"fmt"
	it "iter"
	"strconv"
	"sync"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleZip() {
	for left, right := range iter.Zip(iter.Lift([]int{1, 2, 3}), iter.Lift([]string{"one", "two", "three"})) {
		fmt.Println(left, right)
	}

	// Output:
	// 1 one
	// 2 two
	// 3 three
}

func TestZipEmpty(t *testing.T) {
	t.Parallel()

	for _, _ = range iter.Zip(iter.Lift([]int{}), iter.Lift([]string{})) {
		t.Error("unexpected")
	}
}

func TestZipTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull2(it.Seq2[int, string](iter.Zip(iter.Lift([]int{1, 2}), iter.Lift([]string{"one", "two"}))))
	stop()
}

func ExampleUnzip() {
	keys, values := iter.Unzip(iter.LiftHashMap(map[int]string{1: "one", 2: "two"}))

	for key := range keys {
		fmt.Println(key)
	}

	for value := range values {
		fmt.Println(value)
	}
}

func ExampleUnzip_method() {
	keys, values := iter.LiftHashMap(map[int]string{1: "one", 2: "two"}).Unzip()

	for key := range keys {
		fmt.Println(key)
	}

	for value := range values {
		fmt.Println(value)
	}
}

func TestUnzip(t *testing.T) {
	zipped := iter.Zip(iter.Lift([]int{1, 2, 3}), iter.Lift([]string{"one", "two", "three"}))

	numbers, strings := iter.Unzip(zipped)

	assert.SliceEqual[int](t, numbers.Collect(), []int{1, 2, 3})
	assert.SliceEqual[string](t, strings.Collect(), []string{"one", "two", "three"})
}

func TestUnzipRace(t *testing.T) {
	limit := 100000

	numbers := make([]int, 0, limit)
	strings := make([]string, 0, limit)

	for i := 0; i < limit; i++ {
		numbers = append(numbers, i)
		strings = append(strings, strconv.Itoa(i))
	}

	zipped := iter.Zip(iter.Lift(numbers), iter.Lift(strings))

	numbersIter, stringsIter := iter.Unzip(zipped)

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

	zipped := iter.Zip(iter.Lift([]int{1, 2}), iter.Lift([]string{"one", "two"}))

	numbers, strings := iter.Unzip(zipped)

	_, stop := it.Pull(it.Seq[int](numbers))
	stop()

	_, stop = it.Pull(it.Seq[string](strings))
	stop()
}
