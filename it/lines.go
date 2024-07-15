package it

import (
	"bufio"
	"io"
)

// Lines yields lines from an io.Reader.
//
// Note: lines longer than 65536 will cauese an error.
func Lines(r io.Reader) func(func([]byte, error) bool) {
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
func LinesString(r io.Reader) func(func(string, error) bool) {
	return Map2(Lines(r), func(b []byte, err error) (string, error) {
		return string(b), err
	})
}
