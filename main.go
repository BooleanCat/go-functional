package main

import (
	"os"

	"github.com/BooleanCat/go-functional/gen"
	"github.com/BooleanCat/go-functional/pkgname"
)

func main() {
	args, err := parseArgs()
	if isErrHelp(err) {
		os.Exit(0)
	}
	exitOn(err)

	p := pkgname.Name(args.Positional.TypeName)

	err = os.Mkdir(p, 0755)
	exitOn(err)

	err = gen.Generate(args.Positional.TypeName, p)
	exitOn(err)
}

func exitOn(err error) {
	if err == nil {
		return
	}

	os.Stderr.WriteString(err.Error())
	os.Exit(1)
}
