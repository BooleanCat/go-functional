package template_test

import (
	"github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ExcludeIter", func() {
	It("excludes items from the Iterator that pass the exclusion check", func() {
		lessThanFive := func(i template.T) bool {
			value, ok := i.(int)
			Expect(ok).To(BeTrue())
			return value < 5
		}

		iter := template.NewExclude(NewCounter(), lessThanFive)
		next := optionValue(iter.Next())
		Expect(next).To(Equal(5))
	})
})
