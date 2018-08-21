package template_test

import (
	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Option", func() {
	When("instantiated with a value", func() {
		var option t.Option

		BeforeEach(func() {
			option = t.Some(5)
		})

		It("holds a value", func() {
			Expect(option.Present()).To(BeTrue())
		})

		It("holds the correct value", func() {
			Expect(option.Value()).To(Equal(5))
		})
	})

	When("instantiated as None", func() {
		It("holds no value", func() {
			option := t.None()
			Expect(option.Present()).To(BeFalse())
		})
	})

	Describe("Unwrap", func() {
		It("returns the value", func() {
			option := t.Some(4)
			Expect(option.Unwrap()).To(Equal(4))
		})

		When("the option holds no value", func() {
			It("panics", func() {
				unwrap := func() { t.None().Unwrap() }
				Expect(unwrap).To(Panic())
			})
		})
	})
})
