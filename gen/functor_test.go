package gen_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/gen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("fun", func() {
	When("string", func() {
		It("generates expected code", func() {
			expected := clean(`
				package fstring

				type Functor struct {
					iter Iter
				}

				func New(iter Iter) *Functor {
					return &Functor{iter: iter}
				}

				type Lifted struct {
					slice []string
					index int
				}

				func Lift(slice []string) *Functor {
					return &Functor{iter: &Lifted{slice: slice}}
				}

				func (f *Lifted) Next() Option {
					if f.index >= len(f.slice) {
						return None()
					}

					f.index++
					return Some(f.slice[f.index-1])
				}

				func (f *Functor) Filter(filter func(value string) bool) *Functor {
					f.iter = NewFilter(f.iter, filter)
					return f
				}

				func (f *Functor) Exclude(exclude func(value string) bool) *Functor {
					f.iter = NewExclude(f.iter, exclude)
					return f
				}

				func (f *Functor) Drop(n int) *Functor {
					f.iter = NewDrop(f.iter, n)
					return f
				}

				func (f *Functor) Take(n int) *Functor {
					f.iter = NewTake(f.iter, n)
					return f
				}

				func (f *Functor) Map(op func(string) string) *Functor {
					f.iter = NewMap(f.iter, op)
					return f
				}

				func (f *Functor) Fold(initial string, op foldOp) string {
					return Fold(f.iter, initial, op)
				}

				func (f *Functor) Collect() []string {
					return Collect(f.iter)
				}
			`)

			Expect(fmt.Sprintf("%#v", gen.Functor("string"))).To(Equal(expected))
		})
	})

	When("int", func() {
		It("generates expected code", func() {
			expected := clean(`
				package fint

				type Functor struct {
					iter Iter
				}

				func New(iter Iter) *Functor {
					return &Functor{iter: iter}
				}

				type Lifted struct {
					slice []int
					index int
				}

				func Lift(slice []int) *Functor {
					return &Functor{iter: &Lifted{slice: slice}}
				}

				func (f *Lifted) Next() Option {
					if f.index >= len(f.slice) {
						return None()
					}

					f.index++
					return Some(f.slice[f.index-1])
				}

				func (f *Functor) Filter(filter func(value int) bool) *Functor {
					f.iter = NewFilter(f.iter, filter)
					return f
				}

				func (f *Functor) Exclude(exclude func(value int) bool) *Functor {
					f.iter = NewExclude(f.iter, exclude)
					return f
				}

				func (f *Functor) Drop(n int) *Functor {
					f.iter = NewDrop(f.iter, n)
					return f
				}

				func (f *Functor) Take(n int) *Functor {
					f.iter = NewTake(f.iter, n)
					return f
				}

				func (f *Functor) Map(op func(int) int) *Functor {
					f.iter = NewMap(f.iter, op)
					return f
				}

				func (f *Functor) Fold(initial int, op foldOp) int {
					return Fold(f.iter, initial, op)
				}

				func (f *Functor) Collect() []int {
					return Collect(f.iter)
				}
			`)

			Expect(fmt.Sprintf("%#v", gen.Functor("int"))).To(Equal(expected))
		})
	})
})
