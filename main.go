package main

import (
	"os"
	"path/filepath"

	"github.com/BooleanCat/go-functional/gen"
	"github.com/dave/jennifer/jen"
)

type generator struct {
	name     string
	generate func(string) *jen.File
}

func (g generator) do(typeName string) error {
	src, err := os.Create(filepath.Join("f"+typeName, g.name))
	if err != nil {
		return err
	}
	defer src.Close()

	return g.generate(typeName).Render(src)
}

func main() {
	args, err := parseArgs()
	if isErrHelp(err) {
		os.Exit(0)
	}
	exitOn(err)

	err = os.Mkdir("f"+args.Positional.TypeName, 0755)
	exitOn(err)

	files := []generator{
		{"drop.go", gen.Drop},
		{"exclude.go", gen.Exclude},
		{"filter.go", gen.Filter},
		{"functor.go", gen.Functor},
		{"iter.go", gen.Iter},
		{"map.go", gen.Map},
		{"option.go", gen.Option},
		{"take.go", gen.Take},
	}

	for _, g := range files {
		err = g.do(args.Positional.TypeName)
		exitOn(err)
	}
}

func exitOn(err error) {
	if err == nil {
		return
	}

	os.Stderr.WriteString(err.Error())
	os.Exit(1)
}
