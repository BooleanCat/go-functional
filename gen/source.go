package gen

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

type SourceFiles struct {
	packageName string
	templateDir string
}

func NewSourceFiles(packageName string) (SourceFiles, error) {
	dir, err := templateDir()
	if err != nil {
		return SourceFiles{}, err
	}

	return SourceFiles{packageName, dir}, nil
}

func (s SourceFiles) Generate(name string) ([]byte, error) {
	path := filepath.Join(s.templateDir, name)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read template file: %v", err)
	}

	return s.replacePackageName(content), nil
}

func (s SourceFiles) replacePackageName(content []byte) []byte {
	original := []byte("package template")
	replacement := []byte("package " + s.packageName)
	return bytes.Replace(content, original, replacement, 1)
}

func templateDir() (string, error) {
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("could not get template dir")
	}

	rootDir := filepath.Dir(filepath.Dir(path))
	return filepath.Join(rootDir, "template"), nil
}
