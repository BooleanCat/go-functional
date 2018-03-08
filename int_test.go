package functional_test

import (
	"errors"

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
				total := functional.LiftIntSlice(slice).Map(square).Filter(isEven).Fold(0, sum)
				Expect(total).To(Equal(220))
			})
		})
	})

	Describe("IntSliceErrFunctor", func() {
		It("can be initialised", func() {
			slice := []int{0, 1, 2, 3, 4}
			functional.LiftIntSlice(slice).WithErrs()
		})

		Describe("#Collect", func() {
			var (
				slice      = []int{0, 1, 2, 3, 4}
				functor    functional.IntSliceErrFunctor
				collection []int
				collectErr error
			)

			BeforeEach(func() {
				functor = functional.LiftIntSlice(slice).WithErrs()
			})

			JustBeforeEach(func() {
				collection, collectErr = functor.Collect()
			})

			It("does not return an error", func() {
				Expect(collectErr).NotTo(HaveOccurred())
			})

			It("returns the int slice", func() {
				Expect(collection).To(Equal(slice))
			})

			Context("when an error has previously occurred", func() {
				BeforeEach(func() {
					fail := func(int) (int, error) { return 0, errors.New("map failed") }
					functor = functor.Map(fail)
				})

				It("returns an error", func() {
					Expect(collectErr).To(MatchError("map failed"))
				})

				It("returns an empty slice", func() {
					Expect(collection).To(BeEmpty())
				})

				Context("and another error would have occurred", func() {
					BeforeEach(func() {
						fail := func(int) (int, error) { return 0, errors.New("map failed again") }
						functor = functor.Map(fail)
					})

					It("does not run the second map", func() {
						Expect(collectErr).To(HaveOccurred())
						Expect(collectErr).NotTo(MatchError("map failed again"))
					})
				})
			})
		})

		Describe("#Map", func() {
			var (
				slice     []int
				functor   functional.IntSliceErrFunctor
				operation func(int) (int, error)
			)

			BeforeEach(func() {
				slice = []int{0, 1, 2, 3, 4}
				operation = func(i int) (int, error) { return i * 2, nil }
			})

			JustBeforeEach(func() {
				functor = functional.LiftIntSlice(slice).WithErrs().Map(operation)
			})

			It("applies an operation to all members of a slice", func() {
				collection, err := functor.Collect()
				Expect(err).NotTo(HaveOccurred())
				Expect(collection).To(Equal([]int{0, 2, 4, 6, 8}))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []int{}
				})

				It("collects to an empty slice", func() {
					collection, err := functor.Collect()
					Expect(err).NotTo(HaveOccurred())
					Expect(collection).To(BeEmpty())
				})

				It("cannot cause collect to fail", func() {
					fail := func(i int) (int, error) { return 0, errors.New("map failed") }
					_, err := functor.Map(fail).Collect()
					Expect(err).NotTo(HaveOccurred())
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
					collection, err := functor.Collect()
					Expect(err).NotTo(HaveOccurred())
					Expect(collection).To(Equal(expected))
				})
			})

			Context("when the input operation returns an error", func() {
				BeforeEach(func() {
					operation = func(i int) (int, error) { return 0, errors.New("map failed") }
				})

				It("collects with an error", func() {
					_, err := functor.Collect()
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("map failed"))
				})
			})

			Context("when the input operation returns an error later", func() {
				BeforeEach(func() {
					count := 0
					operation = func(i int) (int, error) {
						count += 1
						if count > 1 {
							return 0, errors.New("map failed later")
						}
						return 0, nil
					}
				})

				It("collects with an error", func() {
					_, err := functor.Collect()
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("map failed later"))
				})
			})
		})

		Describe("#Filter", func() {
			var (
				slice   []int
				functor functional.IntSliceErrFunctor
				filter  func(int) (bool, error)
			)

			BeforeEach(func() {
				slice = []int{0, 1, 2, 3, 4}
				filter = func(i int) (bool, error) { return isEven(i), nil }
			})

			JustBeforeEach(func() {
				functor = functional.LiftIntSlice(slice).WithErrs().Filter(filter)
			})

			It("applies a filter to all members of a slice", func() {
				collection, err := functor.Collect()
				Expect(err).NotTo(HaveOccurred())
				Expect(collection).To(Equal([]int{0, 2, 4}))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []int{}
				})

				It("collects to an empty slice", func() {
					collection, err := functor.Collect()
					Expect(err).NotTo(HaveOccurred())
					Expect(collection).To(BeEmpty())
				})

				It("cannot cause collect to fail", func() {
					fail := func(i int) (bool, error) { return false, errors.New("map failed") }
					_, err := functor.Filter(fail).Collect()
					Expect(err).NotTo(HaveOccurred())
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
					expected := make([]int, 25000)
					for i := range expected {
						expected[i] = i * 2
					}
					collection, err := functor.Collect()
					Expect(err).NotTo(HaveOccurred())
					Expect(collection).To(Equal(expected))
				})
			})

			Context("when the input operation returns an error", func() {
				BeforeEach(func() {
					filter = func(i int) (bool, error) { return false, errors.New("map failed") }
				})

				It("collects with an error", func() {
					_, err := functor.Collect()
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("map failed"))
				})
			})

			Context("when the input operation returns an error later", func() {
				BeforeEach(func() {
					count := 0
					filter = func(i int) (bool, error) {
						count += 1
						if count > 1 {
							return false, errors.New("map failed later")
						}
						return false, nil
					}
				})

				It("collects with an error", func() {
					_, err := functor.Collect()
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("map failed later"))
				})
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
