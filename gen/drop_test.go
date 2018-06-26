package gen_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/gen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("option", func() {
	When("string", func() {
		It("generates expected code", func() {
			expected := clean(`
				package fstring

				type DropIter struct {
					iter Iter
					n    int
				}

				func NewDrop(iter Iter, n int) *DropIter {
					return &DropIter{
						iter: iter,
						n:    n,
					}
				}

				func (iter *DropIter) Next() Option {
					for iter.n > 0 {
						iter.n--
						iter.iter.Next()
					}

					return iter.iter.Next()
				}
			`)

			Expect(fmt.Sprintf("%#v", gen.Drop("string"))).To(Equal(expected))
		})
	})

	When("int", func() {
		It("generates expected code", func() {
			expected := clean(`
				package fint

				type DropIter struct {
					iter Iter
					n    int
				}

				func NewDrop(iter Iter, n int) *DropIter {
					return &DropIter{
						iter: iter,
						n:    n,
					}
				}

				func (iter *DropIter) Next() Option {
					for iter.n > 0 {
						iter.n--
						iter.iter.Next()
					}

					return iter.iter.Next()
				}
			`)

			Expect(fmt.Sprintf("%#v", gen.Drop("int"))).To(Equal(expected))
		})
	})
})
