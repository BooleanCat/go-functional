package template_test

import (
	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TakeIter", func() {
	It("drops all but the first n items", func() {
		slice := t.Collapse(t.Take(NewCounter(), 1))
		expected := []interface{}{0}
		Expect(slice).To(Equal(expected))
	})

	When("n is 0", func() {
		It("drops all items", func() {
			slice := t.Collapse(t.Take(NewCounter(), 0))
			Expect(slice).To(BeEmpty())
		})
	})

	When("n is greater than the remaining items in the Iterator", func() {
		It("drops no items", func() {
			slice := t.Collapse(t.Take(t.Take(NewCounter(), 1), 100))
			expected := []interface{}{0}
			Expect(slice).To(Equal(expected))
		})
	})

	When("the underlying iterator's next fails", func() {
		It("passes the error along", func() {
			_, err := t.Collect(t.Take(NewFailIter(), 5))
			Expect(err).To(MatchError("Oh, no."))
		})
	})
})
