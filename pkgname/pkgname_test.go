package pkgname_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BooleanCat/go-functional/pkgname"
)

var _ = Describe("pkgname", func() {
	It("generates a name for a simple type", func() {
		Expect(pkgname.Name("int")).To(Equal("fint"))
	})

	It("generates a name for interface{}", func() {
		Expect(pkgname.Name("interface{}")).To(Equal("finterface"))
	})

	It("generates a name for a pointer", func() {
		Expect(pkgname.Name("*int")).To(Equal("fpint"))
	})

	It("generates a name for a types with mixed case as lower", func() {
		Expect(pkgname.Name("FileInfo")).To(Equal("ffileinfo"))
	})
})
