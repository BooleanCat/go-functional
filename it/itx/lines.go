package itx

import (
	"io"

	"github.com/BooleanCat/go-functional/v2/it"
)

// Lines yields lines from an io.Reader.
//
// Note: lines longer than 65536 will cauese an error.
func Lines(r io.Reader) Iterator2[[]byte, error] {
	return Iterator2[[]byte, error](it.Lines(r))
}

// LinesString yields lines from an io.Reader as strings.
//
// Note: lines longer than 65536 will cauese an error.
func LinesString(r io.Reader) Iterator2[string, error] {
	return Iterator2[string, error](it.LinesString(r))
}
