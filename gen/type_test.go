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

	Describe("Collapse", func() {
		It("generates", func() {
			Expect(source).To(ContainSubstring(clean(`
				func Collapse(iter Iter) []int {
					return collapse(iter)
				}
			`)))
		})
	})

	Describe("Functor.Collapse", func() {
		It("generates", func() {
			Expect(source).To(ContainSubstring(clean(`
				func (f *Functor) Collapse() []int {
					return collapse(f.iter)
				}
			`)))
		})
	})

	Describe("Fold", func() {
		It("generates", func() {
			Expect(source).To(ContainSubstring(clean(`
				func Fold(iter Iter, initial int, op foldErrFunc) (int, error) {
					result, err := fold(iter, T(initial), op)
					return fromT(result), err
				}
			`)))
		})
	})

	Describe("Functor.Fold", func() {
		It("generates", func() {
			Expect(source).To(ContainSubstring(clean(`
				func (f *Functor) Fold(initial int, op foldErrFunc) (int, error) {
					return Fold(f.iter, initial, op)
				}
			`)))
		})
	})

	Describe("Roll", func() {
		It("generates", func() {
			Expect(source).To(ContainSubstring(clean(`
				func Roll(iter Iter, initial int, op foldFunc) int {
					return fromT(roll(iter, T(initial), op))
				}
			`)))
		})
	})

	Describe("Functor.Roll", func() {
		It("generates", func() {
			Expect(source).To(ContainSubstring(clean(`
				func (f *Functor) Roll(initial int, op foldFunc) int {
					return Roll(f.iter, initial, op)
				}
			`)))
		})
	})

	Describe("Transmute", func() {
		It("generates", func() {
			Expect(source).To(ContainSubstring(clean(`
				func Transmute(v interface{}) int {
					result, ok := v.(int)
					if !ok {
						panic(fmt.Sprintf("could not transmute: %v", v))
					}
					return result
				}
			`)))
		})
	})

	Describe("asMapErrFunc", func() {
		It("generates", func() {
			Expect(source).To(ContainSubstring(clean(`
				func asMapErrFunc(f mapFunc) mapErrFunc {
					return func(v int) (int, error) {
						return f(v), nil
					}
				}
			`)))
		})
	})

	Describe("asFilterErrFunc", func() {
		It("generates", func() {
			Expect(source).To(ContainSubstring(clean(`
				func asFilterErrFunc(f filterFunc) filterErrFunc {
					return func(v int) (bool, error) {
						return f(v), nil
					}
				}
			`)))
		})
	})

	Describe("asFoldErrFunc", func() {
		It("generates", func() {
			Expect(source).To(ContainSubstring(clean(`
				func asFoldErrFunc(f foldFunc) foldErrFunc {
					return func(v int, w int) (int, error) {
						return f(v, w), nil
					}
				}
			`)))
		})
	})
})
