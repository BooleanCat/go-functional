package functional_test

import (
	"strings"

	functional "github.com/BooleanCat/go-functional"
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
				functor = functional.LiftStringSlice(slice).Filter(isNotEmptyString)
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

		})

		Describe("#Exclude", func() {
			var (
				slice   []string
				functor functional.StringSliceFunctor
			)

			BeforeEach(func() {
				slice = []string{"foo", "", "bar"}
			})

			JustBeforeEach(func() {
				functor = functional.LiftStringSlice(slice).Exclude(isNotEmptyString)
			})

			It("applies an exclusion to all members of a slice", func() {
				Expect(functor.Collect()).To(Equal([]string{""}))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []string{}
				})

				It("collects to an empty slice", func() {
					Expect(functor.Collect()).To(BeEmpty())
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
				folded = functional.LiftStringSlice(slice).Fold("please", concatString)
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

		})

		Describe("#Take", func() {
			var (
				n       int
				slice   []string
				functor functional.StringSliceFunctor
			)

			BeforeEach(func() {
				n = 3
				slice = []string{"a", "few", "words", "to", "say"}
			})

			JustBeforeEach(func() {
				functor = functional.LiftStringSlice(slice).Take(n)
			})

			It("drops everything except the first n members of a slice", func() {
				Expect(functor.Collect()).To(Equal([]string{"a", "few", "words"}))
			})

			Context("when n exceeds the number of members", func() {
				BeforeEach(func() {
					n = 6
				})

				It("takes all members", func() {
					Expect(functor.Collect()).To(Equal(slice))
				})
			})

			Context("when taking 0", func() {
				BeforeEach(func() {
					n = 0
				})

				It("collects to an empty slice", func() {
					Expect(functor.Collect()).To(BeEmpty())
				})
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []string{}
				})

				It("collects to an empty slice", func() {
					Expect(functor.Collect()).To(BeEmpty())
				})
			})
		})

		Describe("#Drop", func() {
			var (
				n       int
				slice   []string
				functor functional.StringSliceFunctor
			)

			BeforeEach(func() {
				n = 3
				slice = []string{"a", "few", "words", "to", "say"}
			})

			JustBeforeEach(func() {
				functor = functional.LiftStringSlice(slice).Drop(n)
			})

			It("drops the first n members of a slice", func() {
				Expect(functor.Collect()).To(Equal([]string{"to", "say"}))
			})

			Context("when n exceeds the number of members", func() {
				BeforeEach(func() {
					n = 6
				})

				It("drops all members", func() {
					Expect(functor.Collect()).To(BeEmpty())
				})
			})

			Context("when dropping 0", func() {
				BeforeEach(func() {
					n = 0
				})

				It("collects to the original underlying slice", func() {
					Expect(functor.Collect()).To(Equal([]string{"a", "few", "words", "to", "say"}))
				})
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []string{}
				})

				It("collects to an empty slice", func() {
					Expect(functor.Collect()).To(BeEmpty())
				})
			})
		})

		Describe("a complicated chain of operations", func() {
			It("can filter out empty strings, convert the remaining to lower case and concatenate them", func() {
				slice := []string{"", "a", "FEW", "strings", "", "", "to", "CoNsIdEr"}
				result := functional.LiftStringSlice(slice).Filter(isNotEmptyString).Map(strings.ToLower).Fold("", concatString)
				Expect(result).To(Equal("afewstringstoconsider"))
			})
		})
	})
})
