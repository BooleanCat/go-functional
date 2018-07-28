package template_test

import (
	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DropIter", func() {
	It("drops the first n items", func() {
		iter := t.Drop(NewCounter(), 2)
		next := optionValue(iter.Next())
		Expect(next).To(Equal(2))
	})

	When("n is 0", func() {
		It("drops nothing", func() {
			iter := t.Drop(NewCounter(), 0)
			next := optionValue(iter.Next())
			Expect(next).To(Equal(0))
		})
	})

	When("n is greater than the remaining items in the Iterator", func() {
		It("drops everything", func() {
			iter := t.Drop(t.Take(NewCounter(), 5), 100)
			option := iter.Next()
			Expect(option.Present()).To(BeFalse())
		})
	})
})
