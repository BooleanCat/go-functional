package template_test

import (
	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ChainIter", func() {
	It("iterates over each iterator in turn", func() {
		third := t.Take(NewCounter(), 2)
		second := t.Take(t.Drop(NewCounter(), 2), 2)
		first := t.Take(t.Drop(NewCounter(), 4), 2)

		result := t.Collect(t.Chain(first, second, third))
		expected := []interface{}{4, 5, 2, 3, 0, 1}
		Expect(result).To(Equal(expected))
	})

	When("there are no iterators to chain", func() {
		It("is an empty iterator", func() {
			result := t.Collect(t.Chain())
			Expect(result).To(BeEmpty())
		})
	})
})
