package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type FunctorTemplate struct {
	Type        string
	TypeCapital string
	ShortVar    string
}

func NewFunctorTemplate(t string) FunctorTemplate {
	return FunctorTemplate{
		Type:        strings.ToLower(t),
		TypeCapital: strings.ToUpper(t[0:1]) + strings.ToLower(t[1:]),
		ShortVar:    strings.ToLower(t[0:1]),
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
