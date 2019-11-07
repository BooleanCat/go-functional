package acceptance_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BooleanCat/go-functional/fixtures/fint"
	"github.com/BooleanCat/go-functional/fixtures/finterface"
)

var _ = Describe("go-functional", func() {
	It("generates and is importable", func() {
		_ = []finterface.T{"", "", "bar", ""}
	})

	It("generates with Lift", func() {
		slice := []interface{}{"", "", "bar", ""}
		_ = finterface.Lift(slice)
	})

	It("generates with Collect", func() {
		slice := []interface{}{"bar", "foo"}
		result, err := finterface.Lift(slice).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]interface{}{"bar", "foo"}))
	})

	It("generates with Drop", func() {
		slice := []interface{}{"bar", "foo", "baz"}
		result := finterface.Lift(slice).Drop(2).Collapse()
		Expect(result).To(Equal([]interface{}{"baz"}))
	})

	It("generates with Take", func() {
		slice := []interface{}{"bar", "foo", "baz"}
		result := finterface.Lift(slice).Take(2).Collapse()
		Expect(result).To(Equal([]interface{}{"bar", "foo"}))
	})

	It("generates with Filter", func() {
		hasLen3 := func(v interface{}) bool {
			s, ok := v.(string)
			Expect(ok).To(BeTrue())

			return len(s) == 3
		}

		slice := []interface{}{"bar", "foos", "baz"}
		result := finterface.Lift(slice).Filter(hasLen3).Collapse()
		Expect(result).To(Equal([]interface{}{"bar", "baz"}))
	})

	It("generates with FilterErr", func() {
		hasLen3 := func(v interface{}) (bool, error) {
			s, ok := v.(string)
			Expect(ok).To(BeTrue())

			return len(s) == 3, nil
		}

		slice := []interface{}{"bar", "foos", "baz"}
		result, err := finterface.Lift(slice).FilterErr(hasLen3).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]interface{}{"bar", "baz"}))
	})

	It("generates with Exclude", func() {
		isEmpty := func(v interface{}) bool {
			s, ok := v.(string)
			Expect(ok).To(BeTrue())

			return s == ""
		}

		slice := []interface{}{"", "foos", "baz"}
		result := finterface.Lift(slice).Exclude(isEmpty).Collapse()
		Expect(result).To(Equal([]interface{}{"foos", "baz"}))
	})

	It("generates with ExcludeErr", func() {
		isEmpty := func(v interface{}) (bool, error) {
			s, ok := v.(string)
			Expect(ok).To(BeTrue())

			return s == "", nil
		}

		slice := []interface{}{"", "foos", "baz"}
		result, err := finterface.Lift(slice).ExcludeErr(isEmpty).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]interface{}{"foos", "baz"}))
	})

	It("generates with Repeat", func() {
		result := finterface.New(finterface.Repeat("foo")).Take(3).Collapse()
		Expect(result).To(Equal([]interface{}{"foo", "foo", "foo"}))
	})

	It("generates with Chain", func() {
		foos := finterface.Repeat("foo")
		bars := finterface.Repeat("bar")

		result := finterface.New(foos).Take(2).Chain(bars).Take(4).Collapse()
		Expect(result).To(Equal([]interface{}{"foo", "foo", "bar", "bar"}))
	})

	It("generates with Map", func() {
		prependFoo := func(v interface{}) interface{} {
			s, ok := v.(string)
			Expect(ok).To(BeTrue())

			return "foo" + s
		}

		slice := []interface{}{"bar", "baz"}
		result := finterface.Lift(slice).Map(prependFoo).Collapse()
		Expect(result).To(Equal([]interface{}{"foobar", "foobaz"}))
	})

	It("generates with MapErr", func() {
		prependFoo := func(v interface{}) (interface{}, error) {
			s, ok := v.(string)
			Expect(ok).To(BeTrue())

			return "foo" + s, nil
		}

		slice := []interface{}{"bar", "baz"}
		result, err := finterface.Lift(slice).MapErr(prependFoo).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]interface{}{"foobar", "foobaz"}))
	})

	It("generates with Fold", func() {
		sum := func(a, b interface{}) (interface{}, error) {
			x, ok := a.(int)
			Expect(ok).To(BeTrue())

			y, ok := b.(int)
			Expect(ok).To(BeTrue())

			return x + y, nil
		}

		slice := []interface{}{1, 2, 3, 4}
		result, err := finterface.Lift(slice).Fold(0, sum)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(interface{}(10)))
	})

	It("generates with Roll", func() {
		sum := func(a, b interface{}) interface{} {
			x, ok := a.(int)
			Expect(ok).To(BeTrue())

			y, ok := b.(int)
			Expect(ok).To(BeTrue())

			return x + y
		}

		slice := []interface{}{1, 2, 3, 4}
		result := finterface.Lift(slice).Roll(0, sum)
		Expect(result).To(Equal(interface{}(10)))
	})

	It("generates with Transmute", func() {
		v := interface{}(interface{}("foo"))
		s, ok := finterface.Transmute(v).(string)
		Expect(ok).To(BeTrue())
		Expect(s).To(Equal("foo"))
	})

	It("generates with Transform", func() {
		length := func(v interface{}) (int, error) { return len(finterface.Transmute(v).(string)), nil }

		slice := []interface{}{"foo", "ba", "b"}
		iter := finterface.Lift(slice).Blur()
		result := fint.New(fint.Transform(iter, length)).Collapse()
		Expect(result).To(Equal([]int{3, 2, 1}))
	})
})
