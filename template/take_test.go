package template_test

import (
	"github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TakeIter", func() {
	It("drops all but the first n items", func() {
		slice := template.Collect(template.NewTake(NewCounter(), 1))
		expected := []interface{}{0}
		Expect(slice).To(Equal(expected))
	})

	When("n is 0", func() {
		It("drops all items", func() {
			slice := template.Collect(template.NewTake(NewCounter(), 0))
			Expect(slice).To(BeEmpty())
		})
	})

	When("n is greater than the remaining items in the Iterator", func() {
		It("drops no items", func() {
			iter := template.NewTake(template.NewTake(NewCounter(), 1), 100)
			slice := template.Collect(iter)
			expected := []interface{}{0}
			Expect(slice).To(Equal(expected))
		})
	})
})
