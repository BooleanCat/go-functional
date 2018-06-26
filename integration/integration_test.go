package integration_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go-functional", func() {
	It("succeeds", func() {
		cmd := exec.Command(goFunctionalBin, "string")
		Expect(cmd.Run()).To(Succeed())
	})

	When("the type name is omitted", func() {
		It("fails", func() {
			cmd := exec.Command(goFunctionalBin)
			Expect(cmd.Run()).NotTo(Succeed())
		})
	})
})
