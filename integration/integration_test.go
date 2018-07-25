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

				func isOdd(n fint.T) bool {
					return n % 2 == 1
				}

				func main() {
					slice := []int{1, 2, 3}
					newSlice := fint.Lift(slice).Filter(isOdd).Collect()
					expected := []fint.T{1, 3}

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

				import (
					"fmt"
					"reflect"
					"somebin/fstring"
				)

				func isEmpty(s fstring.T) bool {
					return s == ""
				}

				func prependFoo(s string) string {
					return "foo" + s
				}

				func main() {
					slice := []string{"", "", "bar", ""}
					newSlice := fstring.Lift(slice).Exclude(isEmpty).Map(prependFoo).Collect()
					expected := []fstring.T{"foobar"}

					if !reflect.DeepEqual(newSlice, expected) {
						panic(fmt.Sprintf("expected %#v to equal %#v", expected, newSlice))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates a type wrapper for fold functions", func() {
			cmd := goFunctionalCommand(someBinPath, "string")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import (
					"fmt"
					"somebin/fstring"
				)

				func prepend(a, b string) string {
					return a + b
				}

				func main() {
					slice := []string{"foo", "bar"}
					result := fstring.Lift(slice).Fold("", fstring.Π(prepend))

					if result != "foobar" {
						panic(fmt.Sprintf("expected %s to equal foobar", result))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})
	})
})
