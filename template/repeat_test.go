package template_test

import (
	"reflect"

	"github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RepeatIter", func() {
	It("always yields the same element", func() {
		iter := template.NewRepeat("pikachu")
		expected := []interface{}{"pikachu", "pikachu", "pikachu"}
		result := template.New(iter).Take(3).Collect()
		Expect(reflect.DeepEqual(result, expected)).To(BeTrue())
	})
})
