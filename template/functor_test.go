package template_test

import (
	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lift", func() {
	It("initialises a functor from a slice", func() {
		slice := []interface{}{1, 2}
		result := t.Lift(slice).Collapse()
		Expect(result).To(Equal(slice))
	})

	When("lifting an empty slice", func() {
		It("returns an empty iterator", func() {
			result := t.Lift([]interface{}{}).Collapse()
			Expect(result).To(BeEmpty())
		})
	})
})
