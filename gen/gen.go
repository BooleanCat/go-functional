package gen

import (
	"io/ioutil"
	"path/filepath"
)

func Generate(typeName, path string) error {
	if err := generateSourceFiles(typeName); err != nil {
		return err
	}

	destination := filepath.Join("f"+typeName, "type.go")
	content := []byte(TypeFileContent(typeName).GoString())
	return writeFile(destination, content)
}

func generateSourceFiles(typeName string) error {
	sourceFiles, err := NewSourceFiles("f" + typeName)
	if err != nil {
		return err
	}

	for _, f := range templateFiles() {
		content, err := sourceFiles.Generate(f)
		if err != nil {
			return err
		}

		destination := filepath.Join("f"+typeName, f)
		if err = writeFile(destination, content); err != nil {
			return err
		}
	}

	return nil
}

func templateFiles() []string {
	return []string{
		"drop.go",
		"exclude.go",
		"filter.go",
		"functor.go",
		"iter.go",
		"map.go",
		"option.go",
		"take.go",
		"chain.go",
	}
}

func writeFile(destination string, content []byte) error {
	return ioutil.WriteFile(destination, content, 0755)
}
