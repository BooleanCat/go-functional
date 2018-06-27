package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BooleanCat/go-functional/gen"
	"github.com/dave/jennifer/jen"
)

func main() {
	args, err := parseArgs()
	if err != nil {
		os.Exit(1)
	}

	err = os.Mkdir("f"+args.Positional.TypeName, 0755)
	if err != nil {
		os.Exit(1)
	}

	files := map[string]func(string) *jen.File{
		"drop.go":    gen.Drop,
		"exclude.go": gen.Exclude,
		"filter.go":  gen.Filter,
		"functor.go": gen.Functor,
		"iter.go":    gen.Iter,
		"map.go":     gen.Map,
		"option.go":  gen.Option,
		"take.go":    gen.Take,
	}

	for f, generate := range files {
		path := filepath.Join("f"+args.Positional.TypeName, f)
		content := fmt.Sprintf("%#v", generate(args.Positional.TypeName))
		err = ioutil.WriteFile(path, []byte(content), 0755)
		if err != nil {
			os.Exit(1)
		}
	}
}
