package integration_test

import (
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go-functional", func() {
	var (
		cmd     *exec.Cmd
		workDir string
	)

	BeforeEach(func() {
		workDir = tempDir()
		cmd = exec.Command("go-functional", "int")
		cmd.Stdout = GinkgoWriter
		cmd.Stderr = GinkgoWriter
		cmd.Dir = workDir
	})

	AfterEach(func() {
		Expect(os.RemoveAll(workDir)).To(Succeed())
	})

	It("succeeds", func() {
		Expect(cmd.Run()).To(Succeed())
	})

	When("the type name is omitted", func() {
		BeforeEach(func() {
			cmd.Args = cmd.Args[:1]
		})

		It("fails", func() {
			Expect(cmd.Run()).NotTo(Succeed())
		})
	})

	When("the -h flag is provided", func() {
		BeforeEach(func() {
			cmd.Args = append(cmd.Args[:1], "-h")
		})

		It("succeeds", func() {
			Expect(cmd.Run()).To(Succeed())
		})
	})
})
