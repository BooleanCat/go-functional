package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type FunctorTemplate struct {
	Type                  string
	FunctorName           string
	FunctorErrName        string
	LiftFuncName          string
	NegateFuncName        string
	NegateWithErrFuncName string
	TypeVar               string
}

func NewFunctorTemplate(t string) FunctorTemplate {
	title := strings.Title(t)
	return FunctorTemplate{
		Type:                  strings.ToLower(t),
		FunctorName:           title + "SliceFunctor",
		FunctorErrName:        title + "SliceErrFunctor",
		LiftFuncName:          "Lift" + title + "Slice",
		NegateFuncName:        "negate" + title + "Op",
		NegateWithErrFuncName: "negate" + title + "OpWithErr",
		TypeVar:               strings.ToLower(t[0:1]),
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "go-functional <type-name>")
	}

	src, err := ioutil.ReadFile("/Users/tom/go/src/github.com/BooleanCat/go-functional/generate/template.txt")
	if err != nil {
		panic(err)
	}

	templateVars := NewFunctorTemplate(os.Args[1])

	template, err := template.New("int").Parse(string(src))
	if err != nil {
		panic(err)
	}
	err = template.Execute(os.Stdout, templateVars)
	if err != nil {
		panic(err)
	}
}
