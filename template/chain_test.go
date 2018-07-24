package template_test

import (
	"reflect"

	"github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ChainIter", func() {
	It("iterates over each iterator in turn", func() {
		isEven := func(n template.T) bool { return toInt(n)%2 == 0 }
		evens := template.NewTake(template.NewFilter(NewCounter(), isEven), 2)

		isOdd := func(n template.T) bool { return toInt(n)%2 == 1 }
		odds := template.NewTake(template.NewFilter(NewCounter(), isOdd), 2)

		numbers := template.Collect(template.NewChain(evens, odds))

		expected := []template.T{0, 2, 1, 3}
		Expect(reflect.DeepEqual(numbers, expected)).To(BeTrue())
	})

	When("the first iterator is empty", func() {
		It("iterates over the second iterator", func() {
			empty := template.NewTake(NewCounter(), 0)

			isOdd := func(n template.T) bool { return toInt(n)%2 == 1 }
			odds := template.NewTake(template.NewFilter(NewCounter(), isOdd), 2)

			numbers := template.Collect(template.NewChain(empty, odds))

			expected := []template.T{1, 3}
			Expect(reflect.DeepEqual(numbers, expected)).To(BeTrue())
		})
	})

	When("the second iterator is empty", func() {
		It("iterates over the first iterator", func() {
			isEven := func(n template.T) bool { return toInt(n)%2 == 0 }
			evens := template.NewTake(template.NewFilter(NewCounter(), isEven), 2)

			empty := template.NewTake(NewCounter(), 0)

			numbers := template.Collect(template.NewChain(evens, empty))

			expected := []template.T{0, 2}
			Expect(reflect.DeepEqual(numbers, expected)).To(BeTrue())
		})
	})
})
