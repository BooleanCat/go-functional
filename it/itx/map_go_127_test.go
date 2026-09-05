//go:build go1.27

package itx_test

import (
	"fmt"
	"maps"

	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleIterator_Map() {
	fmt.Println(itx.FromSlice([]string{"", "a", "aa"}).Map(func(v string) int {
		return len(v)
	}).Collect())
	// Output: [0 1 2]
}

func ExampleIterator2_Map() {
	count := func(a, b string) (string, int) {
		return a, len(b)
	}

	fmt.Println(maps.Collect(itx.FromMap(map[string]string{"k1": "a", "k2": "aa"}).Map(count).Seq()))
	// Output: map[k1:1 k2:2]
}

func ExampleIterator_MapError() {
	fmt.Println(it.TryCollect(itx.FromSlice([]string{"", "a", "aa"}).MapError(func(v string) (int, error) {
		return len(v), nil
	})))
	// Output: [0 1 2] <nil>
}
