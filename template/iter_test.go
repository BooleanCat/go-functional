package template_test

import (
	"errors"

	t "github.com/BooleanCat/go-functional/template"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("iter", func() {
	add := func(a, b interface{}) (interface{}, error) {
		return toInt(a) + toInt(b), nil
	}

	sum := func(a, b interface{}) interface{} {
		return toInt(a) + toInt(b)
	}

	Describe("Collect", func() {
		It("collects an iterator into a slice", func() {
			slice, err := t.New(NewCounter()).Take(2).Collect()
			Expect(err).NotTo(HaveOccurred())
			expected := []interface{}{0, 1}
			Expect(slice).To(Equal(expected))
		})

		When("collecting an empty iterator", func() {
			It("collects to an empty slice", func() {
				slice, err := t.New(NewCounter()).Take(0).Collect()
				Expect(err).NotTo(HaveOccurred())
				Expect(slice).To(BeEmpty())
			})
		})

		When("collecting encounters an error", func() {
			It("returns the error", func() {
				_, err := t.New(NewFailIter()).Collect()
				Expect(err).To(MatchError("Oh, no."))
			})
		})
	})

	Describe("Collapse", func() {
		It("collects an iterator into a slice", func() {
			slice := t.New(NewCounter()).Take(2).Collapse()
			expected := []interface{}{0, 1}
			Expect(slice).To(Equal(expected))
		})

		When("collecting an empty iterator", func() {
			It("collects to an empty slice", func() {
				slice := t.New(NewCounter()).Take(0).Collapse()
				Expect(slice).To(BeEmpty())
			})
		})

		When("collapsing encounters an error", func() {
			It("panics", func() {
				collapse := func() { t.New(NewFailIter()).Collapse() }
				Expect(collapse).To(Panic())
			})
		})
	})

	Describe("Fold", func() {
		It("applies the fold operation sequentially to iterator items", func() {
			result, err := t.New(NewCounter()).Take(6).Fold(0, add)
			Expect(err).NotTo(HaveOccurred())
			Expect(toInt(result)).To(Equal(15))
		})

		When("folding an empty iterator", func() {
			It("returns the initial value", func() {
				alwaysTen := func(_, _ interface{}) (interface{}, error) { return 15, nil }
				result, err := t.New(NewCounter()).Take(0).Fold(7, alwaysTen)
				Expect(err).NotTo(HaveOccurred())
				Expect(toInt(result)).To(Equal(7))
			})
		})

		When("folding encounters an error", func() {
			It("returns the error", func() {
				_, err := t.New(NewFailIter()).Fold(0, add)
				Expect(err).To(MatchError("Oh, no."))
			})
		})

		When("the fold op fails", func() {
			It("returns the error", func() {
				fail := func(_, _ interface{}) (interface{}, error) {
					return interface{}(nil), errors.New("Oh, no.")
				}

				_, err := t.New(NewCounter()).Take(5).Fold(nil, fail)
				Expect(err).To(MatchError("Oh, no."))
			})
		})
	})

	Describe("Roll", func() {
		It("applies the roll operation sequentially to iterator items", func() {
			result := t.New(NewCounter()).Take(6).Roll(0, sum)
			Expect(toInt(result)).To(Equal(15))
		})

		When("rolling an empty iterator", func() {
			It("returns the initial value", func() {
				alwaysTen := func(_, _ interface{}) interface{} { return 15 }
				result := t.New(NewCounter()).Take(0).Roll(7, alwaysTen)
				Expect(toInt(result)).To(Equal(7))
			})
		})

		When("rolling encounters an error", func() {
			It("panics", func() {
				roll := func() { t.New(NewFailIter()).Roll(0, sum) }
				Expect(roll).To(Panic())
			})
		})
	})
})
