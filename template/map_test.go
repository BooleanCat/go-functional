package template_test

import (
	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MapIter", func() {
	double := func(value interface{}) interface{} {
		return interface{}(toInt(value) * 2)
	}

	It("applies the map operation to each item", func() {
		result := t.Collapse(t.Map(t.Take(NewCounter(), 3), double))
		expected := []interface{}{0, 2, 4}
		Expect(result).To(Equal(expected))
	})

	When("the underlying iterator's next fails", func() {
		It("passes the error along", func() {
			_, err := t.Collect(t.Map(NewFailIter(), double))
			Expect(err).To(MatchError("Oh, no."))
		})
	})
})
