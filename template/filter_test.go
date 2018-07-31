package template_test

import (
	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FilterIter", func() {
	greaterThanThree := func(i interface{}) bool {
		value, ok := i.(int)
		Expect(ok).To(BeTrue())
		return value > 3
	}

	It("includes items from the Iterator that pass the filter", func() {
		iter := t.Filter(NewCounter(), greaterThanThree)
		next := resultValue(iter.Next())
		Expect(next).To(Equal(4))
	})

	When("the underlying iterator's next fails", func() {
		It("passes the error along", func() {
			_, err := t.Collect(t.Filter(NewFailIter(), greaterThanThree))
			Expect(err).To(MatchError("Oh, no."))
		})
	})
})
