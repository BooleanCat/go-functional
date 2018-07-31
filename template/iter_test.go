package template_test

import (
	t "github.com/BooleanCat/go-functional/template"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("iter", func() {
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
	})

	Describe("Fold", func() {
		It("applies the fold operation sequentially to iterator items", func() {
			sum := func(a, b interface{}) interface{} { return toInt(a) + toInt(b) }
			result, err := t.New(NewCounter()).Take(6).Fold(0, sum)
			Expect(err).NotTo(HaveOccurred())
			Expect(toInt(result)).To(Equal(15))
		})

		When("folding an empty iterator", func() {
			It("returns the initial value", func() {
				alwaysTen := func(_, _ interface{}) interface{} { return 15 }
				result, err := t.New(NewCounter()).Take(0).Fold(7, alwaysTen)
				Expect(err).NotTo(HaveOccurred())
				Expect(toInt(result)).To(Equal(7))
			})
		})
	})
})
