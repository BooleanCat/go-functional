package main

import (
	"os"

	"github.com/BooleanCat/go-functional/gen"
)

func main() {
	args, err := parseArgs()
	if isErrHelp(err) {
		os.Exit(0)
	}
	exitOn(err)

	err = os.Mkdir(packageName(args.Positional.TypeName), 0755)
	exitOn(err)

	err = gen.Generate(args.Positional.TypeName, "f"+args.Positional.TypeName)
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
	return "f" + typeName
}
