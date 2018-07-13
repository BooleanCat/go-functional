package template_test

import (
	"github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Option", func() {
	When("instantiated with a value", func() {
		var option template.Option

		BeforeEach(func() {
			option = template.Some(5)
		})

		It("holds a value", func() {
			Expect(option.Present()).To(BeTrue())
		})

		It("holds the correct value", func() {
			Expect(optionValue(option)).To(Equal(5))
		})
	})

	When("instantiated with no value", func() {
		It("holds no value", func() {
			option := template.None()
			Expect(option.Present()).To(BeFalse())
		})
	})
})
