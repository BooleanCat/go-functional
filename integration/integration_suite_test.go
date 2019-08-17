package integration_test

import (
	"io/ioutil"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var goFunctionalBin string

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

func tempDir() string {
	name, err := ioutil.TempDir("", "")
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return name
}

func newInt(i int) *int {
	return &i
}

func newString(s string) *string {
	return &s
}
