package gen

import (
	"io/ioutil"
	"path/filepath"

	"github.com/BooleanCat/go-functional/pkgname"
)

func Generate(typeName, path string) error {
	if err := generateSourceFiles(typeName); err != nil {
		return err
	}

	destination := filepath.Join(pkgname.Name(typeName), "type.go")
	content := []byte(NewTypeFileGen(typeName).File().GoString())
	return writeFile(destination, content)
}

func generateSourceFiles(typeName string) error {
	sourceFiles, err := NewSourceFiles(pkgname.Name(typeName))
	if err != nil {
		return err
	}

	for _, f := range templateFiles() {
		content, err := sourceFiles.Generate(f)
		if err != nil {
			return err
		}

		destination := filepath.Join(pkgname.Name(typeName), f)
		if err = writeFile(destination, content); err != nil {
			return err
		}
	}

	return nil
}

func templateFiles() []string {
	return []string{
		"chain.go",
		"drop.go",
		"exclude.go",
		"filter.go",
		"functor.go",
		"iter.go",
		"map.go",
		"option.go",
		"repeat.go",
		"result.go",
		"take.go",
		"transform.go",
	}
}

func writeFile(destination string, content []byte) error {
	return ioutil.WriteFile(destination, content, 0666)
}
