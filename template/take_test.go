package template_test

import (
	"github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TakeIter", func() {
	It("drops all but the first n items", func() {
		iter := template.NewTake(NewCounter(), 1)
		next := optionValue(iter.Next())
		Expect(next).To(Equal(0))
		Expect(iter.Next().Present()).To(BeFalse())
	})

	When("n is 0", func() {
		It("drops all items", func() {
			iter := template.NewTake(NewCounter(), 0)
			Expect(iter.Next().Present()).To(BeFalse())
		})
	})

	When("n is greater than the remaining items in the Iterator", func() {
		It("drops no items", func() {
			iter := template.NewTake(template.NewTake(NewCounter(), 1), 100)
			Expect(optionValue(iter.Next())).To(Equal(0))
			Expect(iter.Next().Present()).To(BeFalse())
		})
	})
})
