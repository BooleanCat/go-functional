package functional_test

import (
	functional "BooleanCat/go-functional"

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
				functor = functional.LiftIntSlice(slice).Map(double)
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

			Context("when the input slice is arbitrarily large", func() {
				BeforeEach(func() {
					slice = make([]int, 50000)
					for i := range slice {
						slice[i] = i
					}
				})

				It("applies an operation to all members of a slice", func() {
					expected := make([]int, 50000)
					for i := range expected {
						expected[i] = i * 2
					}
					Expect(functor.Collect()).To(Equal(expected))
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
				functor = functional.LiftIntSlice(slice).Filter(isEven)
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

			Context("when the input slice is arbitrarily large", func() {
				BeforeEach(func() {
					slice = make([]int, 50000)
					for i := range slice {
						slice[i] = i
					}
				})

				It("applies a filter to all members of a slice", func() {
					expected := make([]int, 25000)
					for i := range expected {
						expected[i] = i * 2
					}
					Expect(functor.Collect()).To(Equal(expected))
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
				folded = functional.LiftIntSlice(slice).Fold(10, sum)
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

			Context("when the input slice is arbitrarily large", func() {
				BeforeEach(func() {
					slice = make([]int, 50000)
					for i := range slice {
						slice[i] = i
					}
				})

				It("applies a fold over all members of a slice", func() {
					expected := 10
					for i := 0; i < 50000; i++ {
						expected += i
					}
					Expect(folded).To(Equal(expected))
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

		Describe("a complicated chain of operations", func() {
			It("can find the sum a set of even square numbers", func() {
				slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
				total := functional.LiftIntSlice(slice).Map(square).Filter(isEven).Fold(0, sum)
				Expect(total).To(Equal(220))
			})
		})
	})
})

func double(i int) int {
	return i * 2
}

func square(i int) int {
	return i * i
}

func isEven(i int) bool {
	return i%2 == 0
}

func sum(a, b int) int {
	return a + b
}
