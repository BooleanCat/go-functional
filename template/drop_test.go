package template_test

import (
	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DropIter", func() {
	It("drops the first n items", func() {
		iter := t.Drop(NewCounter(), 2)
		next := resultValue(iter.Next())
		Expect(next).To(Equal(2))
	})

	When("n is 0", func() {
		It("drops nothing", func() {
			iter := t.Drop(NewCounter(), 0)
			next := resultValue(iter.Next())
			Expect(next).To(Equal(0))
		})
	})

	When("n is greater than the remaining items in the Iterator", func() {
		It("drops everything", func() {
			iter := t.Drop(t.Take(NewCounter(), 5), 100)
			Expect(iter.Next().Error()).To(Equal(t.ErrNoValue))
		})
	})

	When("the underlying iterator's next fails", func() {
		It("passes the error along", func() {
			_, err := t.Collect(t.Drop(NewFailIter(), 3))
			Expect(err).To(MatchError("Oh, no."))
		})

		It("calls the underlying iterator's next method only once", func() {
			iter := NewFailIter()
			t.Collect(t.Drop(iter, 3))
			Expect(iter.NextCallCount()).To(Equal(1))
		})
	})
})
