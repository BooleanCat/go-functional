package gen_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/gen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("filter", func() {
	When("string", func() {
		It("generates expected code", func() {
			expected := clean(`
				package fstring

				type FilterIter struct {
					iter   Iter
					filter func(string) bool
				}

				func NewFilter(iter Iter, filter func(string) bool) FilterIter {
					return FilterIter{
						filter: filter,
						iter:   iter,
					}
				}

				func (iter FilterIter) Next() Option {
					for {
						if option := iter.iter.Next(); !option.Present() || iter.filter(option.Value) {
							return option
						}
					}
				}
			`)

			Expect(fmt.Sprintf("%#v", gen.Filter("string"))).To(Equal(expected))
		})
	})

	When("int", func() {
		It("generates expected code", func() {
			expected := clean(`
				package fint

				type FilterIter struct {
					iter   Iter
					filter func(int) bool
				}

				func NewFilter(iter Iter, filter func(int) bool) FilterIter {
					return FilterIter{
						filter: filter,
						iter:   iter,
					}
				}

				func (iter FilterIter) Next() Option {
					for {
						if option := iter.iter.Next(); !option.Present() || iter.filter(option.Value) {
							return option
						}
					}
				}
			`)

			Expect(fmt.Sprintf("%#v", gen.Filter("int"))).To(Equal(expected))
		})
	})
})
