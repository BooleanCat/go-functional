package gen

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

func Generate(typeName, path string) error {
	files := []string{
		"drop.go",
		"exclude.go",
		"filter.go",
		"functor.go",
		"iter.go",
		"map.go",
		"option.go",
		"take.go",
	}

	templateDir, err := templateDir()
	if err != nil {
		return err
	}

	for _, f := range files {
		if err = createFile(filepath.Join(templateDir, f), filepath.Join("f"+typeName, f), "f"+typeName); err != nil {
			return err
		}
	}

	return createTypeFile(filepath.Join("f"+typeName, "type.go"), typeName)
}

func createFile(source, destination, packageName string) error {
	content, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	content = bytes.Replace(content, []byte("package template"), []byte("package "+packageName), 1)

	return ioutil.WriteFile(destination, content, 0755)
}

func createTypeFile(destination, typeName string) error {
	f := Type(typeName)
	typeFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer typeFile.Close()

	return f.Render(typeFile)
}

func templateDir() (string, error) {
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("no caller info")
	}

	return filepath.Join(filepath.Dir(filepath.Dir(path)), "template"), nil
}
