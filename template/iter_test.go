package template_test

import (
	"github.com/BooleanCat/go-functional/template"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("iter", func() {
	Describe("Collect", func() {
		It("collects an iterator into a slice", func() {
			slice := template.New(NewCounter()).Take(2).Collect()
			Expect(slice).To(HaveLen(2))
			Expect(toInt(slice[0])).To(Equal(0))
			Expect(toInt(slice[1])).To(Equal(1))
		})

		When("collecting an empty iterator", func() {
			It("collects to an empty slice", func() {
				slice := template.New(NewCounter()).Take(0).Collect()
				Expect(slice).To(BeEmpty())
			})
		})
	})

	Describe("Fold", func() {
		It("applies the fold operation sequentially to iterator items", func() {
			sum := func(a, b interface{}) interface{} { return toInt(a) + toInt(b) }
			result := template.New(NewCounter()).Take(6).Fold(0, sum)
			Expect(toInt(result)).To(Equal(15))
		})

		When("folding an empty iterator", func() {
			It("returns the initial value", func() {
				alwaysTen := func(_, _ interface{}) interface{} { return 15 }
				result := template.New(NewCounter()).Take(0).Fold(7, alwaysTen)
				Expect(toInt(result)).To(Equal(7))
			})
		})
	})
})
