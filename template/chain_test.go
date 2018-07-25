package template_test

import (
	"reflect"

	"github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ChainIter", func() {
	It("iterates over each iterator in turn", func() {
		third := template.NewTake(NewCounter(), 2)
		second := template.NewTake(template.NewDrop(NewCounter(), 2), 2)
		first := template.NewTake(template.NewDrop(NewCounter(), 4), 2)

		result := template.Collect(template.NewChain(first, second, third))
		expected := []interface{}{4, 5, 2, 3, 0, 1}
		Expect(reflect.DeepEqual(result, expected)).To(BeTrue())
	})

	When("there are no iterators to chain", func() {
		It("is an empty iterator", func() {
			result := template.Collect(template.NewChain())
			Expect(result).To(BeEmpty())
		})
	})
})
