package gen_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/gen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("map", func() {
	When("string", func() {
		It("generates expected code", func() {
			expected := clean(`
				package fstring

				type MapIter struct {
					iter Iter
					op   func(string) string
				}

				func NewMap(iter Iter, op func(string) string) MapIter {
					return MapIter{
						iter: iter,
						op:   op,
					}
				}

				func (iter MapIter) Next() Option {
					next := iter.iter.Next()
					if !next.Present() {
						return next
					}

					return Some(iter.op(next.Value))
				}
			`)

			Expect(fmt.Sprintf("%#v", gen.Map("string"))).To(Equal(expected))
		})
	})

	When("int", func() {
		It("generates expected code", func() {
			expected := clean(`
				package fint

				type MapIter struct {
					iter Iter
					op   func(int) int
				}

				func NewMap(iter Iter, op func(int) int) MapIter {
					return MapIter{
						iter: iter,
						op:   op,
					}
				}

				func (iter MapIter) Next() Option {
					next := iter.iter.Next()
					if !next.Present() {
						return next
					}

					return Some(iter.op(next.Value))
				}
			`)

			Expect(fmt.Sprintf("%#v", gen.Map("int"))).To(Equal(expected))
		})
	})
})
