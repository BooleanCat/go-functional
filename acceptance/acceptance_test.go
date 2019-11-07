package acceptance_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go-functional", func() {
	var cmd *exec.Cmd

	BeforeEach(func() {
		cmd = exec.Command("go", "run", "github.com/BooleanCat/go-functional", "int")
		cmd.Stdout = GinkgoWriter
		cmd.Stderr = GinkgoWriter
	})

	When("the type name is omitted", func() {
		BeforeEach(func() {
			cmd.Args = cmd.Args[:len(cmd.Args)-1]
		})

		It("fails", func() {
			Expect(cmd.Run()).NotTo(Succeed())
		})
	})

	When("the -h flag is provided", func() {
		BeforeEach(func() {
			cmd.Args = append(cmd.Args[:len(cmd.Args)-1], "-h")
		})

		It("succeeds", func() {
			Expect(cmd.Run()).To(Succeed())
		})
	})
})
