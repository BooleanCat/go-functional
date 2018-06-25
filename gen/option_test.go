package gen_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/gen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Option", func() {
	When("string", func() {
		It("generates expected code", func() {
			expected := clean(`
				package fstring

				type Option struct {
					Value   string
					present bool
				}

				func Some(value string) Option {
					return Option{
						Value:   value,
						present: true,
					}
				}

				func None() Option {
					return Option{}
				}

				func (o Option) Present() bool {
					return o.present
				}
			`)

			Expect(fmt.Sprintf("%#v", gen.Option("string"))).To(Equal(expected))
		})
	})

	When("int", func() {
		It("generates expected code", func() {
			expected := clean(`
				package fint

				type Option struct {
					Value   int
					present bool
				}

				func Some(value int) Option {
					return Option{
						Value:   value,
						present: true,
					}
				}

				func None() Option {
					return Option{}
				}

				func (o Option) Present() bool {
					return o.present
				}
			`)

			Expect(fmt.Sprintf("%#v", gen.Option("int"))).To(Equal(expected))
		})
	})
})
