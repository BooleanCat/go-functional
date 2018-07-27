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

	When("int", func() {
		It("generates and is importable", func() {
			cmd := goFunctionalCommand(someBinPath, "int")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import (
					"fmt"
					"reflect"
					"somebin/fint"
				)

				func isOdd(n int) bool {
					return n % 2 == 1
				}

				func main() {
					slice := []int{1, 2, 3}
					newSlice := fint.Lift(slice).Filter(isOdd).Collect()
					expected := []int{1, 3}

					if !reflect.DeepEqual(newSlice, expected) {
						panic(fmt.Sprintf("expected %#v to equal %#v", expected, newSlice))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})
	})

	When("string", func() {
		It("generates and is importable", func() {
			cmd := goFunctionalCommand(someBinPath, "string")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import "somebin/fstring"

				func main() {
					_ = []fstring.T{"", "", "bar", ""}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates with Lift", func() {
			cmd := goFunctionalCommand(someBinPath, "string")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import "somebin/fstring"

				func main() {
					slice := []string{"", "", "bar", ""}
					_ = fstring.Lift(slice)
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates with Collect", func() {
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
					slice := []string{"bar", "foo"}
					result := fstring.Lift(slice).Collect()
					expected := []string{"bar", "foo"}

					if !reflect.DeepEqual(expected, result) {
						panic(fmt.Sprintf("expected %v to equal %v", expected, result))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates with Drop", func() {
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
					slice := []string{"bar", "foo", "baz"}
					result := fstring.Lift(slice).Drop(2).Collect()
					expected := []string{"baz"}

					if !reflect.DeepEqual(expected, result) {
						panic(fmt.Sprintf("expected %v to equal %v", expected, result))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates with Take", func() {
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
					slice := []string{"bar", "foo", "baz"}
					result := fstring.Lift(slice).Take(2).Collect()
					expected := []string{"bar", "foo"}

					if !reflect.DeepEqual(expected, result) {
						panic(fmt.Sprintf("expected %v to equal %v", expected, result))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates with Filter", func() {
			cmd := goFunctionalCommand(someBinPath, "string")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import (
					"fmt"
					"reflect"
					"somebin/fstring"
				)

				func hasLen3(s string) bool {
					return len(s) == 3
				}

				func main() {
					slice := []string{"bar", "foos", "baz"}
					result := fstring.Lift(slice).Filter(hasLen3).Collect()
					expected := []string{"bar", "baz"}

					if !reflect.DeepEqual(expected, result) {
						panic(fmt.Sprintf("expected %v to equal %v", expected, result))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates with Exclude", func() {
			cmd := goFunctionalCommand(someBinPath, "string")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import (
					"fmt"
					"reflect"
					"somebin/fstring"
				)

				func isEmpty(s string) bool {
					return s == ""
				}

				func main() {
					slice := []string{"", "foos", "baz"}
					result := fstring.Lift(slice).Exclude(isEmpty).Collect()
					expected := []string{"foos", "baz"}

					if !reflect.DeepEqual(expected, result) {
						panic(fmt.Sprintf("expected %v to equal %v", expected, result))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates with Repeat", func() {
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
					result := fstring.New(fstring.NewRepeat("foo")).Take(3).Collect()
					expected := []string{"foo", "foo", "foo"}

					if !reflect.DeepEqual(expected, result) {
						panic(fmt.Sprintf("expected %v to equal %v", expected, result))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates with Chain", func() {
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
					foos := fstring.NewRepeat("foo")
					bars := fstring.NewRepeat("bar")
					result := fstring.New(foos).Take(2).Chain(bars).Take(4).Collect()
					expected := []string{"foo", "foo", "bar", "bar"}

					if !reflect.DeepEqual(expected, result) {
						panic(fmt.Sprintf("expected %v to equal %v", expected, result))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates with Map", func() {
			cmd := goFunctionalCommand(someBinPath, "string")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import (
					"fmt"
					"reflect"
					"somebin/fstring"
				)

				func prependFoo(s string) string {
					return "foo" + s
				}

				func main() {
					slice := []string{"bar", "baz"}
					result := fstring.Lift(slice).Map(prependFoo).Collect()
					expected := []string{"foobar", "foobaz"}

					if !reflect.DeepEqual(expected, result) {
						panic(fmt.Sprintf("expected %v to equal %v", expected, result))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})
	})
})
