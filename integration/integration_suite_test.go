package integration_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/renstrom/dedent"
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

func goFunctionalCommand(workDir string, args ...string) *exec.Cmd {
	cmd := exec.Command(goFunctionalBin, args...)
	cmd.Dir = workDir
	cmd.Stdout = GinkgoWriter
	cmd.Stderr = GinkgoWriter
	return cmd
}

func tempDir() string {
	name, err := ioutil.TempDir("", "")
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return name
}

func clean(text string) string {
	return strings.TrimSpace(dedent.Dedent(text))
}

func mkdirAt(dir ...string) {
	ExpectWithOffset(1, os.MkdirAll(filepath.Join(dir...), 0755)).To(Succeed())
}

func writeFile(path, content string) {
	ExpectWithOffset(1, ioutil.WriteFile(path, []byte(content), 0755)).To(Succeed())
}

func makeFunctionalSample(workDir, name, src string) *exec.Cmd {
	mainGo := filepath.Join(workDir, "src", name, "main.go")
	Expect(ioutil.WriteFile(mainGo, []byte(src), 0755)).To(Succeed())
	cmd := exec.Command("go", "run", mainGo)
	cmd.Env = []string{"GOPATH=" + workDir}
	cmd.Stdout = GinkgoWriter
	cmd.Stderr = GinkgoWriter
	return cmd
}
