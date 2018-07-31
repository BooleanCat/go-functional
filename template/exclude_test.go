package template_test

import (
	"errors"

	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ExcludeIter", func() {
	lessThanFive := func(i interface{}) bool {
		return toInt(i) < 5
	}

	lessThanFour := func(i interface{}) (bool, error) {
		return toInt(i) < 4, nil
	}

	fail := func(_ interface{}) (bool, error) {
		return false, errors.New("Oh, no.")
	}

	It("excludes items from the Iterator that pass the exclusion check", func() {
		iter := t.Exclude(NewCounter(), lessThanFive)
		next := resultValue(iter.Next())
		Expect(next).To(Equal(5))
	})

	When("exclude operations could fail but don't", func() {
		It("applies the map operation to each item", func() {
			expected := []interface{}{4, 5, 6}
			result, err := t.Collect(t.ExcludeErr(t.Take(NewCounter(), 7), lessThanFour))
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(expected))
		})
	})

	When("an exclude operation fails", func() {
		It("it returns the error", func() {
			_, err := t.Collect(t.ExcludeErr(t.Take(NewCounter(), 7), fail))
			Expect(err).To(MatchError("Oh, no."))
		})
	})
})
