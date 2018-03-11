package functional_test

import (
	functional "github.com/BooleanCat/go-functional"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoFunctional", func() {
	Describe("BoolSliceFunctor", func() {
		It("can be initialised", func() {
			slice := []bool{true, true, false, true}
			functional.LiftBoolSlice(slice)
		})

		Describe("#Collect", func() {
			It("returns the bool slice", func() {
				slice := []bool{true, true, false, true}
				functor := functional.LiftBoolSlice(slice)
				Expect(functor.Collect()).To(Equal(slice))
			})
		})

		Describe("#Map", func() {
			var (
				slice   []bool
				functor functional.BoolSliceFunctor
			)

			BeforeEach(func() {
				slice = []bool{true, true, false, true}
			})

			JustBeforeEach(func() {
				functor = functional.LiftBoolSlice(slice).Map(func(b bool) bool { return !b })
			})

			It("applies an operation to all members of a slice", func() {
				Expect(functor.Collect()).To(Equal([]bool{false, false, true, false}))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []bool{}
				})

				It("collects to an empty slice", func() {
					Expect(functor.Collect()).To(BeEmpty())
				})
			})
		})

		Describe("#Filter", func() {
			var (
				slice   []bool
				functor functional.BoolSliceFunctor
			)

			BeforeEach(func() {
				slice = []bool{true, true, false, true}
			})

			JustBeforeEach(func() {
				functor = functional.LiftBoolSlice(slice).Filter(func(b bool) bool { return b })
			})

			It("applies a filter to all members of a slice", func() {
				Expect(functor.Collect()).To(Equal([]bool{true, true, true}))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []bool{}
				})

				It("collects to an empty slice", func() {
					Expect(functor.Collect()).To(BeEmpty())
				})
			})
		})

		Describe("#Exclude", func() {
			var (
				slice   []bool
				functor functional.BoolSliceFunctor
			)

			BeforeEach(func() {
				slice = []bool{true, true, false, true}
			})

			JustBeforeEach(func() {
				functor = functional.LiftBoolSlice(slice).Exclude(func(b bool) bool { return b })
			})

			It("applies an exclusion filter to all members of a slice", func() {
				Expect(functor.Collect()).To(Equal([]bool{false}))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []bool{}
				})

				It("collects to an empty slice", func() {
					Expect(functor.Collect()).To(BeEmpty())
				})
			})
		})

		Describe("#Fold", func() {
			var (
				slice  []bool
				folded bool
			)

			BeforeEach(func() {
				slice = []bool{true, true, false, true}
			})

			JustBeforeEach(func() {
				folded = functional.LiftBoolSlice(slice).Fold(false, func(a, b bool) bool { return a || b })
			})

			It("applies a fold over all members of a slice", func() {
				Expect(folded).To(Equal(true))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []bool{}
				})

				It("returns the initial value", func() {
					Expect(folded).To(Equal(false))
				})
			})

		})

		Describe("#Take", func() {
			var (
				n       int
				slice   []bool
				functor functional.BoolSliceFunctor
			)

			BeforeEach(func() {
				n = 3
				slice = []bool{true, true, false, true}
			})

			JustBeforeEach(func() {
				functor = functional.LiftBoolSlice(slice).Take(n)
			})

			It("drops everything except the first n members of a slice", func() {
				Expect(functor.Collect()).To(Equal([]bool{true, true, false}))
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
					slice = []bool{}
				})

				It("collects to an empty slice", func() {
					Expect(functor.Collect()).To(BeEmpty())
				})
			})
		})

		Describe("#Drop", func() {
			var (
				n       int
				slice   []bool
				functor functional.BoolSliceFunctor
			)

			BeforeEach(func() {
				n = 3
				slice = []bool{true, true, false, true}
			})

			JustBeforeEach(func() {
				functor = functional.LiftBoolSlice(slice).Drop(n)
			})

			It("drops the first n members of a slice", func() {
				Expect(functor.Collect()).To(Equal([]bool{true}))
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
					Expect(functor.Collect()).To(Equal([]bool{true, true, false, true}))
				})
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []bool{}
				})

				It("collects to an empty slice", func() {
					Expect(functor.Collect()).To(BeEmpty())
				})
			})
		})
	})
})
