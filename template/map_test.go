package template_test

import (
	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MapIter", func() {
	It("applies the map operation to each item", func() {
		double := func(value interface{}) interface{} {
			return interface{}(toInt(value) * 2)
		}

		iter := t.Map(t.Take(NewCounter(), 3), double)
		result := t.Collect(iter)
		expected := []interface{}{0, 2, 4}
		Expect(result).To(Equal(expected))
	})
})
