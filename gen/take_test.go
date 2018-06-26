package gen_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/gen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("take", func() {
	When("string", func() {
		It("generates expected code", func() {
			expected := clean(`
				package fstring

				type TakeIter struct {
					iter Iter
					n    int
				}

				func NewTake(iter Iter, n int) *TakeIter {
					return &TakeIter{
						iter: iter,
						n:    n,
					}
				}

				func (iter *TakeIter) Next() Option {
					if iter.n <= 0 {
						return None()
					}

					iter.n--
					return iter.iter.Next()
				}
			`)

			Expect(fmt.Sprintf("%#v", gen.Take("string"))).To(Equal(expected))
		})
	})

	When("int", func() {
		It("generates expected code", func() {
			expected := clean(`
				package fint

				type TakeIter struct {
					iter Iter
					n    int
				}

				func NewTake(iter Iter, n int) *TakeIter {
					return &TakeIter{
						iter: iter,
						n:    n,
					}
				}

				func (iter *TakeIter) Next() Option {
					if iter.n <= 0 {
						return None()
					}

					iter.n--
					return iter.iter.Next()
				}
			`)

			Expect(fmt.Sprintf("%#v", gen.Take("int"))).To(Equal(expected))
		})
	})
})
