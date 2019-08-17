package gen_test

import (
	"bytes"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"
	"testing"

	"github.com/lithammer/dedent"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGen(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gen Suite")
}

func clean(text string) string {
	return strings.TrimSpace(dedent.Dedent(text))
}

func normaliseSource(source string) string {
	fset := token.NewFileSet()
	config := printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 4}
	buffer := new(bytes.Buffer)

	file, err := parser.ParseFile(fset, "donotcare.go", source, 0)
	Expect(err).NotTo(HaveOccurred())
	Expect(config.Fprint(buffer, fset, file)).To(Succeed())
	return buffer.String()
}
