package template_test

import (
	"github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lift", func() {
	It("initialises a functor from a slice", func() {
		slice := []interface{}{1, 2}
		collected := template.Lift(slice).Collect()
		Expect(collected).To(HaveLen(2))
		Expect(toInt(collected[0])).To(Equal(1))
		Expect(toInt(collected[1])).To(Equal(2))
	})

	When("lifting an empty slice", func() {
		It("returns an empty iterator", func() {
			var slice []interface{}
			collected := template.Lift(slice).Collect()
			Expect(collected).To(BeEmpty())
		})
	})
})
