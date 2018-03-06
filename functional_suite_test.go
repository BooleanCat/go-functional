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
