package template_test

import (
	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RepeatIter", func() {
	It("always yields the same element", func() {
		iter := t.Repeat("pikachu")
		expected := []interface{}{"pikachu", "pikachu", "pikachu"}
		result := t.New(iter).Take(3).Collect()
		Expect(result).To(Equal(expected))
	})
})
