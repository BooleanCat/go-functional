package functional_test

import (
	functional "github.com/BooleanCat/go-functional"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoFunctional", func() {
	Describe("IntSliceFunctor", func() {
		It("can be initialised", func() {
			slice := []int{0, 1, 2, 3, 4}
			functional.LiftIntSlice(slice)
		})

		Describe("#Collect", func() {
			It("returns the int slice", func() {
				slice := []int{0, 1, 2, 3, 4}
				functor := functional.LiftIntSlice(slice)
				Expect(functor.Collect()).To(Equal(slice))
			})
		})

		Describe("#Map", func() {
			var (
				slice   []int
				functor functional.IntSliceFunctor
			)

			BeforeEach(func() {
				slice = []int{0, 1, 2, 3, 4}
			})

			JustBeforeEach(func() {
				functor = functional.LiftIntSlice(slice).Map(doubleInt)
			})

			It("applies an operation to all members of a slice", func() {
				Expect(functor.Collect()).To(Equal([]int{0, 2, 4, 6, 8}))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []int{}
				})

				It("collects to an empty slice", func() {
					Expect(functor.Collect()).To(BeEmpty())
				})
			})
		})

		Describe("#Filter", func() {
			var (
				slice   []int
				functor functional.IntSliceFunctor
			)

			BeforeEach(func() {
				slice = []int{0, 1, 2, 3, 4}
			})

			JustBeforeEach(func() {
				functor = functional.LiftIntSlice(slice).Filter(isEvenInt)
			})

			It("applies a filter to all members of a slice", func() {
				Expect(functor.Collect()).To(Equal([]int{0, 2, 4}))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []int{}
				})

				It("collects to an empty slice", func() {
					Expect(functor.Collect()).To(BeEmpty())
				})
			})
		})

		Describe("#Exclude", func() {
			var (
				slice   []int
				functor functional.IntSliceFunctor
			)

			BeforeEach(func() {
				slice = []int{0, 1, 2, 3, 4}
			})

			JustBeforeEach(func() {
				functor = functional.LiftIntSlice(slice).Exclude(isEvenInt)
			})

			It("applies an exclusion filter to all members of a slice", func() {
				Expect(functor.Collect()).To(Equal([]int{1, 3}))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []int{}
				})

				It("collects to an empty slice", func() {
					Expect(functor.Collect()).To(BeEmpty())
				})
			})
		})

		Describe("#Fold", func() {
			var (
				slice  []int
				folded int
			)

			BeforeEach(func() {
				slice = []int{0, 1, 2, 3, 4}
			})

			JustBeforeEach(func() {
				folded = functional.LiftIntSlice(slice).Fold(10, sumInt)
			})

			It("applies a fold over all members of a slice", func() {
				Expect(folded).To(Equal(20))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []int{}
				})

				It("returns the initial value", func() {
					Expect(folded).To(Equal(10))
				})
			})

		})

		Describe("#Take", func() {
			var (
				n       int
				slice   []int
				functor functional.IntSliceFunctor
			)

			BeforeEach(func() {
				n = 3
				slice = []int{0, 1, 2, 3, 4, 5}
			})

			JustBeforeEach(func() {
				functor = functional.LiftIntSlice(slice).Take(n)
			})

			It("drops everything except the first n members of a slice", func() {
				Expect(functor.Collect()).To(Equal([]int{0, 1, 2}))
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
					slice = []int{}
				})

				It("collects to an empty slice", func() {
					Expect(functor.Collect()).To(BeEmpty())
				})
			})
		})

		Describe("#Drop", func() {
			var (
				n       int
				slice   []int
				functor functional.IntSliceFunctor
			)

			BeforeEach(func() {
				n = 3
				slice = []int{0, 1, 2, 3, 4}
			})

			JustBeforeEach(func() {
				functor = functional.LiftIntSlice(slice).Drop(n)
			})

			It("drops the first n members of a slice", func() {
				Expect(functor.Collect()).To(Equal([]int{3, 4}))
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
					Expect(functor.Collect()).To(Equal([]int{0, 1, 2, 3, 4}))
				})
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []int{}
				})

				It("collects to an empty slice", func() {
					Expect(functor.Collect()).To(BeEmpty())
				})
			})
		})

		Describe("a complicated chain of operations", func() {
			It("can find the sum a set of even square numbers", func() {
				slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
				total := functional.LiftIntSlice(slice).Map(squareInt).Filter(isEvenInt).Fold(0, sumInt)
				Expect(total).To(Equal(220))
			})
		})
	})
})
