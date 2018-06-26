package gen_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/gen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("exclude", func() {
	When("string", func() {
		It("generates expected code", func() {
			expected := clean(`
				package fstring

				type ExcludeIter struct {
					iter    Iter
					exclude func(string) bool
				}

				func NewExclude(iter Iter, exclude func(string) bool) ExcludeIter {
					return ExcludeIter{
						exclude: exclude,
						iter:    iter,
					}
				}

				func (iter ExcludeIter) Next() Option {
					for {
						if option := iter.iter.Next(); !option.Present() || iter.exclude(option.Value) {
							return option
						}
					}
				}
			`)

			Expect(fmt.Sprintf("%#v", gen.Exclude("string"))).To(Equal(expected))
		})
	})

	When("int", func() {
		It("generates expected code", func() {
			expected := clean(`
				package fint

				type ExcludeIter struct {
					iter    Iter
					exclude func(int) bool
				}

				func NewExclude(iter Iter, exclude func(int) bool) ExcludeIter {
					return ExcludeIter{
						exclude: exclude,
						iter:    iter,
					}
				}

				func (iter ExcludeIter) Next() Option {
					for {
						if option := iter.iter.Next(); !option.Present() || iter.exclude(option.Value) {
							return option
						}
					}
				}
			`)

			Expect(fmt.Sprintf("%#v", gen.Exclude("int"))).To(Equal(expected))
		})
	})
})
