package template_test

import (
	"github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MapIter", func() {
	It("applies the map operation to each item", func() {
		double := func(value template.T) template.T {
			i, ok := value.(int)
			Expect(ok).To(BeTrue())
			return interface{}(i * 2)
		}

		iter := template.NewMap(NewCounter(), double)
		Expect(optionValue(iter.Next())).To(Equal(0))
		Expect(optionValue(iter.Next())).To(Equal(2))
		Expect(optionValue(iter.Next())).To(Equal(4))
	})
})
