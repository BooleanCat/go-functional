package integration_test

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("usage", func() {
	var (
		workDir string
		libPath string
	)

	BeforeEach(func() {
		workDir = tempDir()
		mkdirAt(workDir, "src", "lib")
		libPath = filepath.Join(workDir, "src", "lib")
	})

	AfterEach(func() {
		Expect(os.RemoveAll(workDir)).To(Succeed())
	})

	Describe("help", func() {
		It("succeeds", func() {
			Expect(goFunctionalCommand(libPath, "-h").Run()).To(Succeed())
		})
	})
})
