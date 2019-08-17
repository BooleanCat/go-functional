package integration_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BooleanCat/go-functional/fixtures/fint"
	"github.com/BooleanCat/go-functional/fixtures/fstring"
)

var _ = Describe("go-functional", func() {
	It("generates and is importable", func() {
		_ = []fstring.T{"", "", "bar", ""}
	})

	It("generates with Lift", func() {
		slice := []string{"", "", "bar", ""}
		_ = fstring.Lift(slice)
	})

	It("generates with Collect", func() {
		slice := []string{"foo", "bar"}
		result, err := fstring.Lift(slice).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]string{"foo", "bar"}))
	})

	It("generates with Drop", func() {
		slice := []string{"foo", "bar", "baz"}
		result := fstring.Lift(slice).Drop(2).Collapse()
		Expect(result).To(Equal([]string{"baz"}))
	})

	It("generates with Take", func() {
		slice := []string{"foo", "bar", "baz"}
		result := fstring.Lift(slice).Take(2).Collapse()
		Expect(result).To(Equal([]string{"foo", "bar"}))
	})

	It("generates with Filter", func() {
		hasLen3 := func(s string) bool { return len(s) == 3 }

		slice := []string{"bar", "foos", "baz"}
		result := fstring.Lift(slice).Filter(hasLen3).Collapse()
		Expect(result).To(Equal([]string{"bar", "baz"}))
	})

	It("generates with FilterErr", func() {
		hasLen3 := func(s string) (bool, error) { return len(s) == 3, nil }

		slice := []string{"bar", "foos", "baz"}
		result, err := fstring.Lift(slice).FilterErr(hasLen3).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]string{"bar", "baz"}))
	})

	It("generates with Exclude", func() {
		hasLen3 := func(s string) bool { return len(s) == 3 }

		slice := []string{"bar", "foos", "baz"}
		result := fstring.Lift(slice).Exclude(hasLen3).Collapse()
		Expect(result).To(Equal([]string{"foos"}))
	})

	It("generates with ExcludeErr", func() {
		hasLen3 := func(s string) (bool, error) { return len(s) == 3, nil }

		slice := []string{"bar", "foos", "baz"}
		result, err := fstring.Lift(slice).ExcludeErr(hasLen3).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]string{"foos"}))
	})

	It("generates with Repeat", func() {
		result := fstring.New(fstring.Repeat("foo")).Take(3).Collapse()
		Expect(result).To(Equal([]string{"foo", "foo", "foo"}))
	})

	It("generates with Chain", func() {
		foos := fstring.Repeat("foo")
		bars := fstring.Repeat("bar")

		result := fstring.New(foos).Take(2).Chain(bars).Take(4).Collapse()
		Expect(result).To(Equal([]string{"foo", "foo", "bar", "bar"}))
	})

	It("generates with Map", func() {
		slice := []string{"bar", "baz"}
		result := fstring.Lift(slice).Map(strings.Title).Collapse()
		Expect(result).To(Equal([]string{"Bar", "Baz"}))
	})

	It("generates with MapErr", func() {
		title := func(s string) (string, error) { return strings.Title(s), nil }

		slice := []string{"bar", "baz"}
		result, err := fstring.Lift(slice).MapErr(title).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]string{"Bar", "Baz"}))
	})

	It("generates with Fold", func() {
		prepend := func(a, b string) (string, error) { return b + a, nil }

		slice := []string{"foo", "bar", "baz"}
		result, err := fstring.Lift(slice).Fold("", prepend)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal("bazbarfoo"))
	})

	It("generates with Roll", func() {
		prepend := func(a, b string) string { return b + a }

		slice := []string{"foo", "bar", "baz"}
		result := fstring.Lift(slice).Roll("", prepend)
		Expect(result).To(Equal("bazbarfoo"))
	})

	It("generates with Transmute", func() {
		v := interface{}("foo")
		Expect(fstring.Transmute(v)).To(Equal("foo"))

		var expectedType string
		Expect(fstring.Transmute(v)).To(BeAssignableToTypeOf(expectedType))
	})

	It("generates with Transform", func() {
		length := func(v interface{}) (int, error) { return len(fstring.Transmute(v)), nil }

		slice := []string{"foo", "ba", "b"}
		iter := fstring.Lift(slice).Blur()
		result := fint.New(fint.Transform(iter, length)).Collapse()
		Expect(result).To(Equal([]int{3, 2, 1}))
	})
})
