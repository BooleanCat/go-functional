package gen_test

import (
	"bytes"
	"go/parser"
	"go/printer"
	"go/token"

	"github.com/BooleanCat/go-functional/gen"
	"github.com/dave/jennifer/jen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type.go", func() {
	It("generates Defs", func() {
		defs := generate(gen.Defs("int"))
		expected := normaliseSource(`
			package foo

			type (
				T             int
				tSlice        []int
				mapFunc       func(int) int
				mapErrFunc    func(int) (int, error)
				foldFunc      func(int, int) int
				foldErrFunc   func(int, int) (int, error)
				filterFunc    func(int) bool
				filterErrFunc func(int) (bool, error)
				transformFunc func(interface{}) (int, error)
			)
		`)

		Expect(defs).To(Equal(expected))
	})

	It("generates fromT", func() {
		defs := generate(gen.FromT("int"))
		expected := normaliseSource(`
			package foo

			func fromT(t T) int {
				return int(t)
			}
		`)

		Expect(defs).To(Equal(expected))
	})
})

func generate(code jen.Code) string {
	file := jen.NewFile("foo")
	file.Add(code)
	return normaliseSource(file.GoString())
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
