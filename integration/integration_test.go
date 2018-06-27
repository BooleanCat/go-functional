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
})
