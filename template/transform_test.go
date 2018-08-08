package template_test

import (
	"errors"
	"fmt"

	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transform", func() {
	Describe("Blur", func() {
		It("wraps represents an iter's values as interfaces", func() {
			iter := t.Blur(t.Drop(NewCounter(), 4))
			Expect(toInt(genericNext(iter))).To(Equal(4))
		})

		When("the underlying iter is exhausted", func() {
			It("returns true", func() {
				iter := t.Blur(t.Take(NewCounter(), 0))
				_, done, _ := iter.Next()
				Expect(done).To(BeTrue())
			})
		})

		When("the underlying iter fails", func() {
			It("propogates the error", func() {
				iter := t.Blur(new(FailIter))
				_, _, err := iter.Next()
				Expect(err).To(MatchError("Oh, no."))
			})
		})
	})

	Describe("Transform", func() {
		It("creates an Iter from a GenericIter", func() {
			slice := t.Collapse(t.Take(t.Transform(fooGenericIter{}, addBar), 3))
			expected := []interface{}{"foobar", "foobar", "foobar"}
			Expect(slice).To(Equal(expected))
		})

		When("the underlying iter fails", func() {
			It("propgates the error", func() {
				result := t.Transform(failGenericIter{}, addBar).Next()
				Expect(result.Error()).To(MatchError("Oh, no."))
			})
		})

		When("the transform function fails", func() {
			It("propogates the error", func() {
				fail := func(_ interface{}) (interface{}, error) { return "", errors.New("Oh, no.") }
				result := t.Transform(fooGenericIter{}, fail).Next()
				Expect(result.Error()).To(MatchError("Oh, no."))
			})
		})
	})

	Describe("Transmute", func() {
		It("type asserts an interface to the alias of T", func() {
			value := interface{}(interface{}(6))
			result := t.Transmute(value)
			Expect(toInt(result)).To(Equal(6))
		})
	})
})

type fooGenericIter struct{}

func (iter fooGenericIter) Next() (interface{}, bool, error) {
	return "foo", false, nil
}

type failGenericIter struct{}

func (iter failGenericIter) Next() (interface{}, bool, error) {
	return "", false, errors.New("Oh, no.")
}

func addBar(v interface{}) (interface{}, error) {
	s, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("expected string, got %v", v)
	}
	return s + "bar", nil
}
