package template_test

import (
	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repeat", func() {
	It("always yields the same element", func() {
		iter := t.Repeat("pikachu")
		expected := []interface{}{"pikachu", "pikachu", "pikachu"}
		slice := t.New(iter).Take(3).Collapse()
		Expect(slice).To(Equal(expected))
	})
})
