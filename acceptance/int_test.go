package acceptance_test

import (
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BooleanCat/go-functional/fixtures/fint"
	"github.com/BooleanCat/go-functional/fixtures/fstring"
)

var _ = Describe("go-functional", func() {
	It("generates and is importable", func() {
		_ = []fint.T{1, 2, 3, 4}
	})

	It("generates with Lift", func() {
		slice := []int{1, 2, 3, 4}
		_ = fint.Lift(slice)
	})

	It("generates with Collect", func() {
		slice := []int{1, 2, 3}
		result, err := fint.Lift(slice).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]int{1, 2, 3}))
	})

	It("generates with Drop", func() {
		slice := []int{1, 2, 3}
		result := fint.Lift(slice).Drop(2).Collapse()
		Expect(result).To(Equal([]int{3}))
	})

	It("generates with Take", func() {
		slice := []int{1, 2, 3}
		result := fint.Lift(slice).Take(2).Collapse()
		Expect(result).To(Equal([]int{1, 2}))
	})

	It("generates with Filter", func() {
		isOdd := func(i int) bool { return i%2 == 1 }

		slice := []int{1, 2, 3}
		result := fint.Lift(slice).Filter(isOdd).Collapse()
		Expect(result).To(Equal([]int{1, 3}))
	})

	It("generates with FilterErr", func() {
		isOdd := func(i int) (bool, error) { return i%2 == 1, nil }

		slice := []int{1, 2, 3}
		result, err := fint.Lift(slice).FilterErr(isOdd).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]int{1, 3}))
	})

	It("generates with Exclude", func() {
		isOdd := func(i int) bool { return i%2 == 1 }

		slice := []int{1, 2, 3}
		result := fint.Lift(slice).Exclude(isOdd).Collapse()
		Expect(result).To(Equal([]int{2}))
	})

	It("generates with ExcludeErr", func() {
		isOdd := func(i int) (bool, error) { return i%2 == 1, nil }

		slice := []int{1, 2, 3}
		result, err := fint.Lift(slice).ExcludeErr(isOdd).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]int{2}))
	})

	It("generates with Repeat", func() {
		result := fint.New(fint.Repeat(42)).Take(3).Collapse()
		Expect(result).To(Equal([]int{42, 42, 42}))
	})

	It("generates with Chain", func() {
		a := fint.Repeat(7)
		b := fint.Repeat(42)

		result := fint.New(a).Take(2).Chain(b).Take(4).Collapse()
		Expect(result).To(Equal([]int{7, 7, 42, 42}))
	})

	It("generates with Map", func() {
		increment := func(i int) int { return i + 1 }

		slice := []int{7, 8}
		result := fint.Lift(slice).Map(increment).Collapse()
		Expect(result).To(Equal([]int{8, 9}))
	})

	It("generates with MapErr", func() {
		increment := func(i int) (int, error) { return i + 1, nil }

		slice := []int{7, 8}
		result, err := fint.Lift(slice).MapErr(increment).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]int{8, 9}))
	})

	It("generates with Fold", func() {
		sum := func(a, b int) (int, error) { return a + b, nil }

		slice := []int{1, 2, 3, 4}
		result, err := fint.Lift(slice).Fold(0, sum)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(10))
	})

	It("generates with Roll", func() {
		sum := func(a, b int) int { return a + b }

		slice := []int{1, 2, 3, 4}
		result := fint.Lift(slice).Roll(0, sum)
		Expect(result).To(Equal(10))
	})

	It("generates with Transmute", func() {
		v := interface{}(4)
		Expect(fint.Transmute(v)).To(Equal(4))

		var expectedType int
		Expect(fint.Transmute(v)).To(BeAssignableToTypeOf(expectedType))
	})

	It("generates with Transform", func() {
		asString := func(v interface{}) (string, error) { return strconv.Itoa(fint.Transmute(v)), nil }

		iter := fint.New(fint.Repeat(42)).Blur()
		result := fstring.New(fstring.Transform(iter, asString)).Take(4).Collapse()
		Expect(result).To(Equal([]string{"42", "42", "42", "42"}))
	})
})
