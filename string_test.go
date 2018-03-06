package functional_test

import (
	functional "BooleanCat/go-functional"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoFunctional", func() {
	Describe("StringSliceFunctor", func() {
		It("can be initialised", func() {
			slice := []string{"foo", "bar"}
			functional.LiftStringSlice(slice)
		})

		Describe("#Collect", func() {
			It("returns the string slice", func() {
				slice := []string{"foo", "bar"}
				functor := functional.LiftStringSlice(slice)
				Expect(functor.Collect()).To(Equal(slice))
			})
		})

		Describe("#Map", func() {
			var (
				slice   []string
				functor functional.StringSliceFunctor
			)

			BeforeEach(func() {
				slice = []string{"foo", "bar"}
			})

			JustBeforeEach(func() {
				functor = functional.LiftStringSlice(slice).Map(strings.ToUpper)
			})

			It("applies an operation to all members of a slice", func() {
				Expect(functor.Collect()).To(Equal([]string{"FOO", "BAR"}))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []string{}
				})

				It("collects to an empty slice", func() {
					Expect(functor.Collect()).To(BeEmpty())
				})
			})

			Context("when the input slice is arbitrarily large", func() {
				BeforeEach(func() {
					slice = make([]string, 50000)
					for i := range slice {
						slice[i] = "a"
					}
				})

				It("applies an operation to all members of a slice", func() {
					expected := make([]string, 50000)
					for i := range expected {
						expected[i] = "A"
					}
					Expect(functor.Collect()).To(Equal(expected))
				})
			})
		})

		Describe("#Filter", func() {
			var (
				slice   []string
				functor functional.StringSliceFunctor
			)

			BeforeEach(func() {
				slice = []string{"foo", "", "bar"}
			})

			JustBeforeEach(func() {
				functor = functional.LiftStringSlice(slice).Filter(isNotEmpty)
			})

			It("applies a filter to all members of a slice", func() {
				Expect(functor.Collect()).To(Equal([]string{"foo", "bar"}))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []string{}
				})

				It("collects to an empty slice", func() {
					Expect(functor.Collect()).To(BeEmpty())
				})
			})

			Context("when the input slice is arbitrarily large", func() {
				BeforeEach(func() {
					slice = make([]string, 50000)
					for i := range slice {
						if i%2 == 0 {
							slice[i] = ""
						} else {
							slice[i] = "foo"
						}
					}
				})

				It("applies a filter to all members of a slice", func() {
					expected := make([]string, 25000)
					for i := range expected {
						expected[i] = "foo"
					}
					Expect(functor.Collect()).To(Equal(expected))
				})
			})
		})

		Describe("#Fold", func() {
			var (
				slice  []string
				folded string
			)

			BeforeEach(func() {
				slice = []string{"join", "these", "words"}
			})

			JustBeforeEach(func() {
				folded = functional.LiftStringSlice(slice).Fold("please", concat)
			})

			It("applies a fold over all members of a slice", func() {
				Expect(folded).To(Equal("pleasejointhesewords"))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []string{}
				})

				It("returns the initial value", func() {
					Expect(folded).To(Equal("please"))
				})
			})

			Context("when the input slice is arbitrarily large", func() {
				BeforeEach(func() {
					slice = make([]string, 10000)
					for i := range slice {
						slice[i] = "foo"
					}
				})

				It("applies a fold over all members of a slice", func() {
					expected := "please"
					for i := 0; i < 10000; i++ {
						expected += "foo"
					}
					Expect(folded).To(Equal(expected))
				})
			})
		})

		Describe("a complicated chain of operations", func() {
			It("can filter out empty strings, convert the remaining to lower case and concatenate them", func() {
				slice := []string{"", "a", "FEW", "strings", "", "", "to", "CoNsIdEr"}
				result := functional.LiftStringSlice(slice).Filter(isNotEmpty).Map(strings.ToLower).Fold("", concat)
				Expect(result).To(Equal("afewstringstoconsider"))
			})
		})
	})
})

func isNotEmpty(s string) bool {
	return s != ""
}

func concat(s, t string) string {
	return s + t
}
