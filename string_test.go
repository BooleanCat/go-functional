package functional_test

import (
	"errors"
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
				functor = functional.LiftStringSlice(slice).Exclude(isNotEmpty)
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
				result := functional.LiftStringSlice(slice).Filter(isNotEmpty).Map(strings.ToLower).Fold("", concat)
				Expect(result).To(Equal("afewstringstoconsider"))
			})
		})
	})

	Describe("StringSliceErrFunctor", func() {
		It("can be initialised", func() {
			slice := []string{"foo", "bar"}
			functional.LiftStringSlice(slice).WithErrs()
		})

		Describe("#Collect", func() {
			var (
				slice      = []string{"foo", "bar"}
				functor    functional.StringSliceErrFunctor
				collection []string
				collectErr error
			)

			BeforeEach(func() {
				functor = functional.LiftStringSlice(slice).WithErrs()
			})

			JustBeforeEach(func() {
				collection, collectErr = functor.Collect()
			})

			It("does not return an error", func() {
				Expect(collectErr).NotTo(HaveOccurred())
			})

			It("returns the string slice", func() {
				Expect(collection).To(Equal(slice))
			})

			Context("when an error has previously occurred", func() {
				BeforeEach(func() {
					fail := func(string) (string, error) { return "", errors.New("map failed") }
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
						fail := func(string) (string, error) { return "", errors.New("map failed again") }
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
				slice     []string
				functor   functional.StringSliceErrFunctor
				operation func(string) (string, error)
			)

			BeforeEach(func() {
				slice = []string{"foo", "bar"}
				operation = func(s string) (string, error) { return s + "baz", nil }
			})

			JustBeforeEach(func() {
				functor = functional.LiftStringSlice(slice).WithErrs().Map(operation)
			})

			It("applies an operation to all members of a slice", func() {
				collection, err := functor.Collect()
				Expect(err).NotTo(HaveOccurred())
				Expect(collection).To(Equal([]string{"foobaz", "barbaz"}))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []string{}
				})

				It("collects to an empty slice", func() {
					collection, err := functor.Collect()
					Expect(err).NotTo(HaveOccurred())
					Expect(collection).To(BeEmpty())
				})

				It("cannot cause collect to fail", func() {
					fail := func(string) (string, error) { return "", errors.New("map failed") }
					_, err := functor.Map(fail).Collect()
					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("when the input operation returns an error", func() {
				BeforeEach(func() {
					operation = func(string) (string, error) { return "", errors.New("map failed") }
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
					operation = func(string) (string, error) {
						count += 1
						if count > 1 {
							return "", errors.New("map failed later")
						}
						return "", nil
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
				slice   []string
				functor functional.StringSliceErrFunctor
				filter  func(string) (bool, error)
			)

			BeforeEach(func() {
				slice = []string{"foo", "", "bar"}
				filter = func(s string) (bool, error) { return s != "", nil }
			})

			JustBeforeEach(func() {
				functor = functional.LiftStringSlice(slice).WithErrs().Filter(filter)
			})

			It("applies a filter to all members of a slice", func() {
				collection, err := functor.Collect()
				Expect(err).NotTo(HaveOccurred())
				Expect(collection).To(Equal([]string{"foo", "bar"}))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []string{}
				})

				It("collects to an empty slice", func() {
					collection, err := functor.Collect()
					Expect(err).NotTo(HaveOccurred())
					Expect(collection).To(BeEmpty())
				})

				It("cannot cause collect to fail", func() {
					fail := func(string) (bool, error) { return false, errors.New("map failed") }
					_, err := functor.Filter(fail).Collect()
					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("when the input operation returns an error", func() {
				BeforeEach(func() {
					filter = func(string) (bool, error) { return false, errors.New("map failed") }
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
					filter = func(string) (bool, error) {
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
				slice   []string
				functor functional.StringSliceErrFunctor
				exclude func(string) (bool, error)
			)

			BeforeEach(func() {
				slice = []string{"foo", "", "bar"}
				exclude = func(s string) (bool, error) { return s != "", nil }
			})

			JustBeforeEach(func() {
				functor = functional.LiftStringSlice(slice).WithErrs().Exclude(exclude)
			})

			It("applies an exclusion to all members of a slice", func() {
				collection, err := functor.Collect()
				Expect(err).NotTo(HaveOccurred())
				Expect(collection).To(Equal([]string{""}))
			})

			Context("when the input slice is empty", func() {
				BeforeEach(func() {
					slice = []string{}
				})

				It("collects to an empty slice", func() {
					collection, err := functor.Collect()
					Expect(err).NotTo(HaveOccurred())
					Expect(collection).To(BeEmpty())
				})

				It("cannot cause collect to fail", func() {
					fail := func(string) (bool, error) { return false, errors.New("exclude failed") }
					_, err := functor.Exclude(fail).Collect()
					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("when the input operation returns an error", func() {
				BeforeEach(func() {
					exclude = func(string) (bool, error) { return false, errors.New("exclude failed") }
				})

				It("collects with an error", func() {
					_, err := functor.Collect()
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("exclude failed"))
				})
			})

			Context("when the input operation returns an error later", func() {
				BeforeEach(func() {
					count := 0
					exclude = func(string) (bool, error) {
						count += 1
						if count > 1 {
							return false, errors.New("exclude failed later")
						}
						return false, nil
					}
				})

				It("collects with an error", func() {
					_, err := functor.Collect()
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("exclude failed later"))
				})
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
