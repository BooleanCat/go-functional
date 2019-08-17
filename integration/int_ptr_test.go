package integration_test

import (
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BooleanCat/go-functional/fixtures/fpint"
	"github.com/BooleanCat/go-functional/fixtures/fpstring"
)

var _ = Describe("go-functional", func() {
	It("generates and is importable", func() {
		_ = []fpint.T{newInt(42), newInt(1)}
	})

	It("generates with Lift", func() {
		slice := []*int{newInt(1), newInt(2), newInt(3), newInt(4)}
		_ = fpint.Lift(slice)
	})

	It("generates with Collect", func() {
		slice := []*int{newInt(1), newInt(2), newInt(3)}
		result, err := fpint.Lift(slice).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]*int{newInt(1), newInt(2), newInt(3)}))
	})

	It("generates with Drop", func() {
		slice := []*int{newInt(1), newInt(2), newInt(3)}
		result := fpint.Lift(slice).Drop(2).Collapse()
		Expect(result).To(Equal([]*int{newInt(3)}))
	})

	It("generates with Take", func() {
		slice := []*int{newInt(1), newInt(2), newInt(3)}
		result := fpint.Lift(slice).Take(2).Collapse()
		Expect(result).To(Equal([]*int{newInt(1), newInt(2)}))
	})

	It("generates with Filter", func() {
		isOdd := func(i *int) bool { return *i%2 == 1 }

		slice := []*int{newInt(1), newInt(2), newInt(3)}
		result := fpint.Lift(slice).Filter(isOdd).Collapse()
		Expect(result).To(Equal([]*int{newInt(1), newInt(3)}))
	})

	It("generates with FilterErr", func() {
		isOdd := func(i *int) (bool, error) { return *i%2 == 1, nil }

		slice := []*int{newInt(1), newInt(2), newInt(3)}
		result, err := fpint.Lift(slice).FilterErr(isOdd).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]*int{newInt(1), newInt(3)}))
	})

	It("generates with Exclude", func() {
		isOdd := func(i *int) bool { return *i%2 == 1 }

		slice := []*int{newInt(1), newInt(2), newInt(3)}
		result := fpint.Lift(slice).Exclude(isOdd).Collapse()
		Expect(result).To(Equal([]*int{newInt(2)}))
	})

	It("generates with ExcludeErr", func() {
		isOdd := func(i *int) (bool, error) { return *i%2 == 1, nil }

		slice := []*int{newInt(1), newInt(2), newInt(3)}
		result, err := fpint.Lift(slice).ExcludeErr(isOdd).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]*int{newInt(2)}))
	})

	It("generates with Repeat", func() {
		result := fpint.New(fpint.Repeat(newInt(42))).Take(3).Collapse()
		Expect(result).To(Equal([]*int{newInt(42), newInt(42), newInt(42)}))
	})

	It("generates with Chain", func() {
		a := fpint.Repeat(newInt(7))
		b := fpint.Repeat(newInt(42))

		result := fpint.New(a).Take(2).Chain(b).Take(4).Collapse()
		Expect(result).To(Equal([]*int{newInt(7), newInt(7), newInt(42), newInt(42)}))
	})

	It("generates with Map", func() {
		increment := func(i *int) *int {
			j := *i + 1
			return &j
		}

		slice := []*int{newInt(7), newInt(8)}
		result := fpint.Lift(slice).Map(increment).Collapse()
		Expect(result).To(Equal([]*int{newInt(8), newInt(9)}))
	})

	It("generates with MapErr", func() {
		increment := func(i *int) (*int, error) {
			j := *i + 1
			return &j, nil
		}

		slice := []*int{newInt(7), newInt(8)}
		result, err := fpint.Lift(slice).MapErr(increment).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]*int{newInt(8), newInt(9)}))
	})

	It("generates with Fold", func() {
		sum := func(a, b *int) (*int, error) {
			c := *a + *b
			return &c, nil
		}

		slice := []*int{newInt(1), newInt(2), newInt(3), newInt(4)}
		result, err := fpint.Lift(slice).Fold(newInt(0), sum)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(newInt(10)))
	})

	It("generates with Roll", func() {
		sum := func(a, b *int) *int {
			c := *a + *b
			return &c
		}

		slice := []*int{newInt(1), newInt(2), newInt(3), newInt(4)}
		result := fpint.Lift(slice).Roll(newInt(0), sum)
		Expect(result).To(Equal(newInt(10)))
	})

	It("generates with Transmute", func() {
		v := interface{}(newInt(4))
		Expect(fpint.Transmute(v)).To(Equal(newInt(4)))

		var expectedType *int
		Expect(fpint.Transmute(v)).To(BeAssignableToTypeOf(expectedType))
	})

	It("generates with Transform", func() {
		asString := func(v interface{}) (*string, error) {
			s := strconv.Itoa(*fpint.Transmute(v))
			return &s, nil
		}

		iter := fpint.New(fpint.Repeat(newInt(42))).Blur()
		result := fpstring.New(fpstring.Transform(iter, asString)).Take(2).Collapse()
		Expect(result).To(Equal([]*string{newString("42"), newString("42")}))
	})
})
