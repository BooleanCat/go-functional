package template_test

import (
	"errors"
	"testing"

	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTemplate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Template Suite")
}

type Counter struct {
	i int
}

func NewCounter() *Counter {
	return new(Counter)
}

func (c *Counter) Next() t.OptionalResult {
	next := t.Some(c.i)
	c.i++
	return t.Success(next)
}

type FailIter struct {
	nextCallCount int
}

func NewFailIter() *FailIter {
	return new(FailIter)
}

func (iter *FailIter) Next() t.OptionalResult {
	iter.nextCallCount++
	return t.Failure(errors.New("Oh, no."))
}

func (iter *FailIter) NextCallCount() int {
	return iter.nextCallCount
}

func toInt(value interface{}) int {
	Expect(value).NotTo(BeNil())
	i, ok := value.(int)
	Expect(ok).To(BeTrue())
	return i
}

func resultValue(result t.OptionalResult) int {
	Expect(result.Error()).To(BeNil())
	Expect(result.Value().Present()).To(BeTrue())
	return toInt(result.Value().Value())
}

func genericNext(iter t.GenericIter) interface{} {
	value, done, err := iter.Next()
	Expect(done).To(BeFalse())
	Expect(err).NotTo(HaveOccurred())
	result, ok := value.(interface{})
	Expect(ok).To(BeTrue())
	return result
}
