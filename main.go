package main

import (
	"fmt"

	"github.com/BooleanCat/go-functional/reference/fstring"
)

func main() {
	a := []string{"", "hello", "", "", "friend", ""}
	fmt.Println(a)
	exclude := func(value string) bool { return value == "hello" }
	filter := func(value string) bool { return value != "" }
	fmt.Println(fstring.Lift(a).Filter(filter).Exclude(exclude).Collect())
	fmt.Println(fstring.Lift(a).Drop(4).Take(1).Collect())
	op := func(a string) string { return "foo" + a }
	fmt.Println(fstring.Lift(a).Map(op).Collect())

	concat := func(a, b string) string { return a + b }
	fmt.Println(fstring.Lift(a).Fold("", concat))
}
