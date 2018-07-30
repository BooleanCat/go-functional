package template_test

import (
	"testing"

	"github.com/BooleanCat/go-functional/template"
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
	return &Counter{}
}

func (c *Counter) Next() template.Result {
	next := template.Some(c.i)
	c.i++
	return next
}

func toInt(value interface{}) int {
	Expect(value).NotTo(BeNil())
	i, ok := value.(int)
	Expect(ok).To(BeTrue())
	return i
}

func resultValue(result template.Result) int {
	Expect(result.Error()).To(BeNil())
	return toInt(result.Value())
}
