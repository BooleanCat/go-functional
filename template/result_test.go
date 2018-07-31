package template_test

import (
	"errors"

	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Result", func() {
	When("instantiated with a value", func() {
		var result t.Result

		BeforeEach(func() {
			result = t.Some(5)
		})

		It("holds a value", func() {
			Expect(result.Error()).To(BeNil())
		})

		It("holds the correct value", func() {
			Expect(resultValue(result)).To(Equal(5))
		})
	})

	When("instantiated with no value", func() {
		It("holds no value", func() {
			result := t.None()
			Expect(result.Error()).To(Equal(t.ErrNoValue))
		})
	})

	When("instantiated with an error", func() {
		It("holds the error", func() {
			result := t.Failed(errors.New("Oh, no."))
			Expect(result.Error()).To(MatchError("Oh, no."))
		})
	})
})
