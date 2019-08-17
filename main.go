package main

import (
	"os"
	"strings"

	"github.com/BooleanCat/go-functional/gen"
)

func main() {
	args, err := parseArgs()
	if isErrHelp(err) {
		os.Exit(0)
	}
	exitOn(err)

	p := packageName(args.Positional.TypeName)

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

func packageName(typeName string) string {
	if typeName == "interface{}" {
		return "finterface"
	}

	if strings.HasPrefix(typeName, "*") {
		return "fp" + typeName[1:len(typeName)]
	}

	return "f" + typeName
}
