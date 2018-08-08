package template_test

import (
	"errors"
	"testing"

	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGen(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gen Suite")
}

type Counter struct {
	i int
}

func NewCounter() *Counter {
	return new(Counter)
}

func (c *Counter) Next() t.Result {
	next := t.Some(c.i)
	c.i++
	return next
}

type FailIter struct {
	nextCallCount int
}

func NewFailIter() *FailIter {
	return new(FailIter)
}

func (iter *FailIter) Next() t.Result {
	iter.nextCallCount++
	return t.Failed(errors.New("Oh, no."))
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

func resultValue(result t.Result) int {
	Expect(result.Error()).To(BeNil())
	return toInt(result.Value())
}

func genericNext(iter t.GenericIter) interface{} {
	value, done, err := iter.Next()
	Expect(done).To(BeFalse())
	Expect(err).NotTo(HaveOccurred())
	result, ok := value.(interface{})
	Expect(ok).To(BeTrue())
	return result
}
