package template_test

import (
	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lift", func() {
	It("initialises a functor from a slice", func() {
		slice := []interface{}{1, 2}
		collected := t.Lift(slice).Collect()
		expected := []interface{}{1, 2}
		Expect(collected).To(Equal(expected))
	})

	When("lifting an empty slice", func() {
		It("returns an empty iterator", func() {
			var slice []interface{}
			collected := t.Lift(slice).Collect()
			Expect(collected).To(BeEmpty())
		})
	})
})
