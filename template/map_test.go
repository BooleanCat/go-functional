package template_test

import (
	"errors"

	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MapIter", func() {
	double := func(value interface{}) interface{} {
		return interface{}(toInt(value) * 2)
	}

	triple := func(value interface{}) (interface{}, error) {
		return interface{}(toInt(value) * 3), nil
	}

	fail := func(_ interface{}) (interface{}, error) {
		return interface{}(0), errors.New("Oh, no.")
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

	When("map operations could fail but don't", func() {
		It("applies the map operation to each item", func() {
			expected := []interface{}{0, 3, 6}
			result, err := t.Collect(t.MapErr(t.Take(NewCounter(), 3), triple))
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(expected))
		})
	})

	When("a map operation fails", func() {
		It("it returns the error", func() {
			_, err := t.Collect(t.MapErr(t.Take(NewCounter(), 3), fail))
			Expect(err).To(MatchError("Oh, no."))
		})
	})
})
