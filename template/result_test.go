package template_test

import (
	"errors"

	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OptionalResult", func() {
	When("instantiated with a value", func() {
		var result t.OptionalResult

		BeforeEach(func() {
			result = t.Success(t.Some(5))
		})

		It("holds no error", func() {
			Expect(result.Error()).To(BeNil())
		})

		It("holds the correct value", func() {
			Expect(result.Value()).To(Equal(t.Some(5)))
		})
	})

	When("instantiated with an error", func() {
		It("holds the error", func() {
			result := t.Failure(errors.New("Oh, no."))
			Expect(result.Error()).To(MatchError("Oh, no."))
		})
	})

	Describe("Unwrap", func() {
		It("returns the value", func() {
			result := t.Success(t.None())
			Expect(result.Unwrap()).To(Equal(t.None()))
		})

		When("the result holds an error", func() {
			It("panics", func() {
				unwrap := func() { t.Failure(errors.New("Oh, no.")).Unwrap() }
				Expect(unwrap).To(Panic())
			})
		})
	})
})
