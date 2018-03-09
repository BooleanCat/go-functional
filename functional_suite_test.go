package functional_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoFunctional(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "go-functional suite")
}

func doubleInt(i int) int {
	return i * 2
}

func squareInt(i int) int {
	return i * i
}

func isEvenInt(i int) bool {
	return i%2 == 0
}

func sumInt(a, b int) int {
	return a + b
}

func isNotEmptyString(s string) bool {
	return s != ""
}

func concatString(s, t string) string {
	return s + t
}
