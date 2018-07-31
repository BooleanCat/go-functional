package template_test

import (
	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lift", func() {
	It("initialises a functor from a slice", func() {
		slice := []interface{}{1, 2}
		slice, err := t.Lift(slice).Collect()
		Expect(err).NotTo(HaveOccurred())
		expected := []interface{}{1, 2}
		Expect(slice).To(Equal(expected))
	})

	When("lifting an empty slice", func() {
		It("returns an empty iterator", func() {
			var slice []interface{}
			slice, err := t.Lift(slice).Collect()
			Expect(err).NotTo(HaveOccurred())
			Expect(slice).To(BeEmpty())
		})
	})
})
