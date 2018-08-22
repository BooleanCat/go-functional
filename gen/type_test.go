package gen_test

import (
	"github.com/BooleanCat/go-functional/gen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type.go", func() {
	var (
		source   string
		typeName string
	)

	BeforeEach(func() {
		typeName = "int"
	})

	JustBeforeEach(func() {
		source = normaliseSource(gen.NewTypeFileGen(typeName).File().GoString())
	})

	Describe("type declarations", func() {
		It("generates", func() {
			Expect(source).To(ContainSubstring(clean(`
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
			`)))
		})
	})

	Describe("fromT", func() {
		It("generates", func() {
			Expect(source).To(ContainSubstring(clean(`
				func fromT(t T) int {
					return int(t)
				}
			`)))
		})

		When("provided a pointer", func() {
			BeforeEach(func() {
				typeName = "*int"
			})

			It("generates", func() {
				Expect(source).To(ContainSubstring(clean(`
					func fromT(t T) *int {
						return t
					}
				`)))
			})
		})
	})

	Describe("Collect", func() {
		It("generates", func() {
			Expect(source).To(ContainSubstring(clean(`
				func Collect(iter Iter) ([]int, error) {
					return collect(iter)
				}
			`)))
		})
	})

	Describe("Functor.Collect", func() {
		It("generates", func() {
			Expect(source).To(ContainSubstring(clean(`
				func (f *Functor) Collect() ([]int, error) {
					return collect(f.iter)
				}
			`)))
		})
	})
})
