package template_test

import (
	"errors"

	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FilterIter", func() {
	greaterThanThree := func(i interface{}) bool {
		return toInt(i) > 3
	}

	greaterThanFour := func(i interface{}) (bool, error) {
		return toInt(i) > 4, nil
	}

	fail := func(_ interface{}) (bool, error) {
		return false, errors.New("Oh, no.")
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

	When("filter operations could fail but don't", func() {
		It("applies the map operation to each item", func() {
			expected := []interface{}{5, 6}
			result, err := t.Collect(t.FilterErr(t.Take(NewCounter(), 7), greaterThanFour))
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(expected))
		})
	})

	When("a filter operation fails", func() {
		It("it returns the error", func() {
			_, err := t.Collect(t.FilterErr(t.Take(NewCounter(), 3), fail))
			Expect(err).To(MatchError("Oh, no."))
		})
	})
})
