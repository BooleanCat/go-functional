package itx_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleIterator2_Unzip() {
	keys, values := itx.MapsAll(map[int]string{1: "one", 2: "two"}).Unzip()

	for key := range keys {
		fmt.Println(key)
	}

	for value := range values {
		fmt.Println(value)
	}
}

func TestUnzipMethod(t *testing.T) {
	keys, values := itx.MapsAll(map[int]string{1: "one"}).Unzip()

	assert.SliceEqual(t, keys.Collect(), []int{1})
	assert.SliceEqual(t, values.Collect(), []string{"one"})
}

func ExampleIterator2_Left() {
	for key := range itx.MapsAll(map[int]string{1: "one"}).Left() {
		fmt.Println(key)
	}
	// Output: 1
}

func ExampleIterator2_Right() {
	for value := range itx.MapsAll(map[int]string{1: "one"}).Right() {
		fmt.Println(value)
	}
	// Output: one
}
