package gen_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/gen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("iter", func() {
	When("string", func() {
		It("generates expected code", func() {
			expected := clean(`
				package fstring

				type Iter interface {
					Next() Option
				}

				func Collect(iter Iter) []string {
					slice := []string{}
					for {
						option := iter.Next()
						if !option.Present() {
							return slice
						}

						slice = append(slice, option.Value)
					}
				}

				func Fold(iter Iter, initial string, op foldOp) string {
					result := initial
					for {
						next := iter.Next()
						if !next.Present() {
							return result
						}

						result = op(result, next.Value)
					}
				}

				type foldOp func(string, string) string
			`)

			Expect(fmt.Sprintf("%#v", gen.Iter("string"))).To(Equal(expected))
		})
	})

	When("int", func() {
		It("generates expected code", func() {
			expected := clean(`
				package fint

				type Iter interface {
					Next() Option
				}

				func Collect(iter Iter) []int {
					slice := []int{}
					for {
						option := iter.Next()
						if !option.Present() {
							return slice
						}

						slice = append(slice, option.Value)
					}
				}

				func Fold(iter Iter, initial int, op foldOp) int {
					result := initial
					for {
						next := iter.Next()
						if !next.Present() {
							return result
						}

						result = op(result, next.Value)
					}
				}

				type foldOp func(int, int) int
			`)

			Expect(fmt.Sprintf("%#v", gen.Iter("int"))).To(Equal(expected))
		})
	})
})
