package template_test

import (
	"github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lift", func() {
	It("initialises a functor from a slice", func() {
		slice := []template.T{1, 2}
		collected := template.Lift(slice).Collect()
		Expect(collected).To(HaveLen(2))
		Expect(toInt(collected[0])).To(Equal(1))
		Expect(toInt(collected[1])).To(Equal(2))
	})

	When("lifting an empty slice", func() {
		It("returns an empty iterator", func() {
			var slice []template.T
			collected := template.Lift(slice).Collect()
			Expect(collected).To(BeEmpty())
		})
	})

	Describe("TFrom", func() {
		It("copies a slice of concrete values to a slice of T", func() {
			slice := []interface{}{7}
			collected := template.Lift(template.TFrom(slice)).Collect()
			Expect(collected).To(HaveLen(1))
			Expect(toInt(collected[0])).To(Equal(7))
		})
	})
})
