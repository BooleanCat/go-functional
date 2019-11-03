package integration_test

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

func tempDir() string {
	name, err := ioutil.TempDir("", "")
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return name
}

func writeFileString(filename, data string, perm os.FileMode) {
	ExpectWithOffset(1, ioutil.WriteFile(filename, []byte(data), perm)).To(Succeed())
}

func mkdir(filepath string, perm os.FileMode) {
	ExpectWithOffset(1, os.Mkdir(filepath, perm)).To(Succeed())
}

func open(filepath string) *os.File {
	file, err := os.Open(filepath)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return file
}

func newInt(i int) *int {
	return &i
}

func newString(s string) *string {
	return &s
}
