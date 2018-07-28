package integration_test

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go-functional", func() {
	var (
		workDir     string
		someBinPath string
	)

	BeforeEach(func() {
		workDir = tempDir()
		mkdirAt(workDir, "src", "somebin")
		someBinPath = filepath.Join(workDir, "src", "somebin")
	})

	AfterEach(func() {
		Expect(os.RemoveAll(workDir)).To(Succeed())
	})

	It("succeeds", func() {
		cmd := goFunctionalCommand(someBinPath, "string")
		Expect(cmd.Run()).To(Succeed())
	})

	When("the type name is omitted", func() {
		It("fails", func() {
			cmd := goFunctionalCommand(someBinPath)
			Expect(cmd.Run()).NotTo(Succeed())
		})
	})

	It("creates a new package in the working directory", func() {
		cmd := goFunctionalCommand(someBinPath, "string")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import "somebin/fstring"

			func main() {
				_ = fstring.Some("foo")
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	When("the -h flag is provided", func() {
		It("succeeds", func() {
			Expect(goFunctionalCommand(workDir, "-h").Run()).To(Succeed())
		})
	})
})
