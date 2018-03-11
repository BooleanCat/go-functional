package functional_test

import (
	"errors"

	functional "github.com/BooleanCat/go-functional"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoFunctional", func() {
	Describe("BoolSliceErrFunctor", func() {
		It("can be initialised", func() {
			slice := []bool{true, true, false, true}
			functional.LiftBoolSlice(slice).WithErrs()
		})

		Describe("#Collect", func() {
			var (
				slice      = []bool{true, true, false, true}
				functor    functional.BoolSliceErrFunctor
				collection []bool
				collectErr error
			)

			BeforeEach(func() {
				functor = functional.LiftBoolSlice(slice).WithErrs()
			})

			JustBeforeEach(func() {
				collection, collectErr = functor.Collect()
			})

			It("does not return an error", func() {
				Expect(collectErr).NotTo(HaveOccurred())
			})

			It("returns the bool slice", func() {
				Expect(collection).To(Equal(slice))
			})

			Context("when an error has previously occurred", func() {
				BeforeEach(func() {
					fail := func(bool) (bool, error) { return false, errors.New("map failed") }
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
						fail := func(bool) (bool, error) { return false, errors.New("map failed again") }
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
				slice     []bool
				functor   functional.BoolSliceErrFunctor
				operation func(bool) (bool, error)
			)

			BeforeEach(func() {
				slice = []bool{true, true, false, true}
				operation = func(b bool) (bool, error) { return !b, nil }
			})

			JustBeforeEach(func() {
				functor = functional.LiftBoolSlice(slice).WithErrs().Map(operation)
			})

			It("applies an operation to all members of a slice", func() {
				collection, err := functor.Collect()
				Expect(err).NotTo(HaveOccurred())
				Expect(collection).To(Equal([]bool{false, false, true, false}))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []bool{}
				})

				It("collects to an empty slice", func() {
					collection, err := functor.Collect()
					Expect(err).NotTo(HaveOccurred())
					Expect(collection).To(BeEmpty())
				})

				It("cannot cause collect to fail", func() {
					fail := func(b bool) (bool, error) { return false, errors.New("map failed") }
					_, err := functor.Map(fail).Collect()
					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("when the input operation returns an error", func() {
				BeforeEach(func() {
					operation = func(b bool) (bool, error) { return false, errors.New("map failed") }
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
					operation = func(b bool) (bool, error) {
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

		Describe("#Filter", func() {
			var (
				slice   []bool
				functor functional.BoolSliceErrFunctor
				filter  func(bool) (bool, error)
			)

			BeforeEach(func() {
				slice = []bool{true, true, false, true}
				filter = func(b bool) (bool, error) { return b, nil }
			})

			JustBeforeEach(func() {
				functor = functional.LiftBoolSlice(slice).WithErrs().Filter(filter)
			})

			It("applies a filter to all members of a slice", func() {
				collection, err := functor.Collect()
				Expect(err).NotTo(HaveOccurred())
				Expect(collection).To(Equal([]bool{true, true, true}))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []bool{}
				})

				It("collects to an empty slice", func() {
					collection, err := functor.Collect()
					Expect(err).NotTo(HaveOccurred())
					Expect(collection).To(BeEmpty())
				})

				It("cannot cause collect to fail", func() {
					fail := func(bool) (bool, error) { return false, errors.New("map failed") }
					_, err := functor.Filter(fail).Collect()
					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("when the input operation returns an error", func() {
				BeforeEach(func() {
					filter = func(bool) (bool, error) { return false, errors.New("map failed") }
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
					filter = func(bool) (bool, error) {
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
				slice   []bool
				functor functional.BoolSliceErrFunctor
				exclude func(bool) (bool, error)
			)

			BeforeEach(func() {
				slice = []bool{true, true, false, true}
				exclude = func(b bool) (bool, error) { return b, nil }
			})

			JustBeforeEach(func() {
				functor = functional.LiftBoolSlice(slice).WithErrs().Exclude(exclude)
			})

			It("applies an exclusion filter to all members of a slice", func() {
				collection, err := functor.Collect()
				Expect(err).NotTo(HaveOccurred())
				Expect(collection).To(Equal([]bool{false}))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []bool{}
				})

				It("collects to an empty slice", func() {
					collection, err := functor.Collect()
					Expect(err).NotTo(HaveOccurred())
					Expect(collection).To(BeEmpty())
				})

				It("cannot cause collect to fail", func() {
					fail := func(bool) (bool, error) { return false, errors.New("map failed") }
					_, err := functor.Exclude(fail).Collect()
					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("when the input operation returns an error", func() {
				BeforeEach(func() {
					exclude = func(bool) (bool, error) { return false, errors.New("map failed") }
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
					exclude = func(bool) (bool, error) {
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
