package it_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/v2/it"
)

func ExampleRepeat() {
	for number := range it.Take(it.Repeat(42), 2) {
		fmt.Println(number)
	}

	// Output:
	// 42
	// 42
}

func ExampleRepeat2() {
	for v, w := range it.Take2(it.Repeat2(42, "Life"), 2) {
		fmt.Println(v, w)
	}

	// Output:
	// 42 Life
	// 42 Life
}
