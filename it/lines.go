package it

import (
	"bufio"
	"io"
	"iter"
)

// Lines yields lines from an io.Reader.
//
// Note: lines longer than 65536 will cauese an error.
func Lines(r io.Reader) iter.Seq2[[]byte, error] {
	return func(yield func([]byte, error) bool) {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			if !yield(scanner.Bytes(), nil) {
				return
			}
		}

		if err := scanner.Err(); err != nil {
			if !yield(nil, err) {
				return
			}
		}
	}
}

// LinesString yields lines from an io.Reader as strings.
//
// Note: lines longer than 65536 will cauese an error.
func LinesString(r io.Reader) iter.Seq2[string, error] {
	return Map2(Lines(r), func(b []byte, err error) (string, error) {
		return string(b), err
	})
}
