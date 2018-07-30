package template_test

import (
	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FilterIter", func() {
	It("includes items from the Iterator that pass the filter", func() {
		greaterThanThree := func(i interface{}) bool {
			value, ok := i.(int)
			Expect(ok).To(BeTrue())
			return value > 3
		}

		iter := t.Filter(NewCounter(), greaterThanThree)
		next := resultValue(iter.Next())
		Expect(next).To(Equal(4))
	})
})
