package integration_test

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go-functional", func() {
	var (
		workDir     string
		someBinPath string
	)

	BeforeEach(func() {
		workDir = tempDir()
		mkdirAt(workDir, "src", "somebin")
		someBinPath = filepath.Join(workDir, "src", "somebin")
	})

	AfterEach(func() {
		Expect(os.RemoveAll(workDir)).To(Succeed())
	})

	It("succeeds", func() {
		cmd := goFunctionalCommand(someBinPath, "string")
		Expect(cmd.Run()).To(Succeed())
	})

	When("the type name is omitted", func() {
		It("fails", func() {
			cmd := goFunctionalCommand(someBinPath)
			Expect(cmd.Run()).NotTo(Succeed())
		})
	})

	It("creates a new package in the working directory", func() {
		cmd := goFunctionalCommand(someBinPath, "string")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import "somebin/fstring"

			func main() {
				_ = fstring.Some("foo")
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	Describe("Option", func() {
		It(`implements option := fint.Some(42)`, func() {
			cmd := goFunctionalCommand(someBinPath, "int")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import "somebin/fint"

				func main() {
					option := fint.Some(42)

					if !option.Present() {
						panic("expected option to hold a value")
					}

					if option.Value != 42 {
						panic("expected option to hold 42")
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It(`implements option := fint.None()`, func() {
			cmd := goFunctionalCommand(someBinPath, "int")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import "somebin/fint"

				func main() {
					option := fint.None()

					if option.Present() {
						panic("expected option not to hold a value")
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})
	})

	Describe("Functor", func() {
		It("implements fstring.Lift(slice)", func() {
			cmd := goFunctionalCommand(someBinPath, "string")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import "somebin/fstring"

				func main() {
					slice := []string{"foo", "bar"}
					_ = fstring.Lift(slice)
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("implements fstring.Lift(slice).Collect()", func() {
			cmd := goFunctionalCommand(someBinPath, "string")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import (
					"fmt"
					"reflect"
					"somebin/fstring"
				)

				func main() {
					slice := []string{"foo", "bar"}
					functor := fstring.Lift(slice)

					newSlice := functor.Collect()
					if !reflect.DeepEqual(newSlice, slice) {
						panic(fmt.Sprintf("expected %#v to equal %#v", slice, newSlice))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("implements functor.Fold(a, f)", func() {
			cmd := goFunctionalCommand(someBinPath, "int")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import (
					"fmt"
					"somebin/fint"
				)

				func main() {
					functor := fint.Lift([]int{1, 0, 14, 2})

					sum := func(a, b int) int { return a + b }

					result := functor.Fold(0, sum)
					if result != 17 {
						panic(fmt.Sprintf("expected %d to equal 17", result))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})
	})

	Describe("Drop", func() {
		It("implements functor.Drop(n)", func() {
			cmd := goFunctionalCommand(someBinPath, "string")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import (
					"fmt"
					"reflect"
					"somebin/fstring"
				)

				func main() {
					slice := []string{"foo", "bar", "baz"}
					newSlice := fstring.Lift(slice).Drop(2).Collect()
					expected := []string{"baz"}

					if !reflect.DeepEqual(newSlice, expected) {
						panic(fmt.Sprintf("expected %#v to equal %#v", expected, newSlice))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})
	})

	Describe("Take", func() {
		It("implements functor.Take(n)", func() {
			cmd := goFunctionalCommand(someBinPath, "string")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import (
					"fmt"
					"reflect"
					"somebin/fstring"
				)

				func main() {
					slice := []string{"foo", "bar", "baz"}
					newSlice := fstring.Lift(slice).Take(2).Collect()
					expected := []string{"foo", "bar"}

					if !reflect.DeepEqual(newSlice, expected) {
						panic(fmt.Sprintf("expected %#v to equal %#v", expected, newSlice))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})
	})

	Describe("Exclude", func() {
		It("implements functor.Exclude(f)", func() {
			cmd := goFunctionalCommand(someBinPath, "bool")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import (
					"fmt"
					"reflect"
					"somebin/fbool"
				)

				func main() {
					slice := []bool{true, false, true}
					isFalse := func(value bool) bool { return !value }
					newSlice := fbool.Lift(slice).Exclude(isFalse).Collect()
					expected := []bool{true, true}

					if !reflect.DeepEqual(newSlice, expected) {
						panic(fmt.Sprintf("expected %#v to equal %#v", expected, newSlice))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})
	})

	Describe("Filter", func() {
		It("implements functor.Filter(f)", func() {
			cmd := goFunctionalCommand(someBinPath, "bool")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import (
					"fmt"
					"reflect"
					"somebin/fbool"
				)

				func main() {
					slice := []bool{true, false, true}
					isFalse := func(value bool) bool { return !value }
					newSlice := fbool.Lift(slice).Filter(isFalse).Collect()
					expected := []bool{false}

					if !reflect.DeepEqual(newSlice, expected) {
						panic(fmt.Sprintf("expected %#v to equal %#v", expected, newSlice))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})
	})

	Describe("Map", func() {
		It("implements functor.Map(f)", func() {
			cmd := goFunctionalCommand(someBinPath, "int")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import (
					"fmt"
					"reflect"
					"somebin/fint"
				)

				func main() {
					slice := []int{0, 1, 2, -1}
					double := func(value int) int { return value * 2 }
					newSlice := fint.Lift(slice).Map(double).Collect()
					expected := []int{0, 2, 4, -2}

					if !reflect.DeepEqual(newSlice, expected) {
						panic(fmt.Sprintf("expected %#v to equal %#v", expected, newSlice))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})
	})

	Describe("first 10 prime numbers", func() {
		It("can be calculated using go-functional", func() {
			cmd := goFunctionalCommand(someBinPath, "int")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import (
					"fmt"
					"reflect"
					"math"
					"somebin/fint"
				)

				type counter struct {
					i int
				}

				func (c *counter) Next() fint.Option {
					option := fint.Some(c.i)
					c.i++
					return option
				}

				func isPrime(value int) bool {
					for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
						if value%i == 0 {
							return false
						}
					}
					return value > 1
				}

				func main() {
					primes := fint.New(new(counter)).Filter(isPrime).Take(10).Collect()
					expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

					if !reflect.DeepEqual(primes, expected) {
						panic(fmt.Sprintf("expected %#v to equal %#v", expected, primes))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})
	})
})
