package integration_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var goFunctionalBin string

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var _ = BeforeSuite(func() {
	goFunctionalBin = gexecBuild("github.com/BooleanCat/go-functional")
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func gexecBuild(packagePath string) string {
	path, err := gexec.Build(packagePath)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return path
}
