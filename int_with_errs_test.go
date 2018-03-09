package functional_test

import (
	"errors"

	functional "github.com/BooleanCat/go-functional"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoFunctional", func() {
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
				filter = func(i int) (bool, error) { return isEvenInt(i), nil }
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

		Describe("#Exclude", func() {
			var (
				slice   []int
				functor functional.IntSliceErrFunctor
				exclude func(int) (bool, error)
			)

			BeforeEach(func() {
				slice = []int{0, 1, 2, 3, 4}
				exclude = func(i int) (bool, error) { return isEvenInt(i), nil }
			})

			JustBeforeEach(func() {
				functor = functional.LiftIntSlice(slice).WithErrs().Exclude(exclude)
			})

			It("applies an exclusion filter to all members of a slice", func() {
				collection, err := functor.Collect()
				Expect(err).NotTo(HaveOccurred())
				Expect(collection).To(Equal([]int{1, 3}))
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
					_, err := functor.Exclude(fail).Collect()
					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("when the input operation returns an error", func() {
				BeforeEach(func() {
					exclude = func(i int) (bool, error) { return false, errors.New("map failed") }
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
					exclude = func(i int) (bool, error) {
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
