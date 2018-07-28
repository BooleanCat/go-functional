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

	When("int", func() {
		It("generates and is importable", func() {
			cmd := goFunctionalCommand(someBinPath, "int")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import "somebin/fint"

				func main() {
					_ = []fint.T{1, 2, 3, 4}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates with Lift", func() {
			cmd := goFunctionalCommand(someBinPath, "int")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import "somebin/fint"

				func main() {
					slice := []int{1, 2, 3, 4}
					_ = fint.Lift(slice)
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates with Collect", func() {
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
					slice := []int{1, 2, 3}
					result := fint.Lift(slice).Collect()
					expected := []int{1, 2, 3}

					if !reflect.DeepEqual(expected, result) {
						panic(fmt.Sprintf("expected %v to equal %v", expected, result))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates with Drop", func() {
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
					slice := []int{1, 2, 3}
					result := fint.Lift(slice).Drop(2).Collect()
					expected := []int{3}

					if !reflect.DeepEqual(expected, result) {
						panic(fmt.Sprintf("expected %v to equal %v", expected, result))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates with Take", func() {
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
					slice := []int{1, 2, 3}
					result := fint.Lift(slice).Take(2).Collect()
					expected := []int{1, 2}

					if !reflect.DeepEqual(expected, result) {
						panic(fmt.Sprintf("expected %v to equal %v", expected, result))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates with Filter", func() {
			cmd := goFunctionalCommand(someBinPath, "int")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import (
					"fmt"
					"reflect"
					"somebin/fint"
				)

				func isOdd(i int) bool {
					return i % 2 == 1
				}

				func main() {
					slice := []int{1, 2, 3}
					result := fint.Lift(slice).Filter(isOdd).Collect()
					expected := []int{1, 3}

					if !reflect.DeepEqual(expected, result) {
						panic(fmt.Sprintf("expected %v to equal %v", expected, result))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates with Exclude", func() {
			cmd := goFunctionalCommand(someBinPath, "int")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import (
					"fmt"
					"reflect"
					"somebin/fint"
				)

				func isOdd(i int) bool {
					return i % 2 == 1
				}

				func main() {
					slice := []int{1, 2, 3}
					result := fint.Lift(slice).Exclude(isOdd).Collect()
					expected := []int{2}

					if !reflect.DeepEqual(expected, result) {
						panic(fmt.Sprintf("expected %v to equal %v", expected, result))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates with Repeat", func() {
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
					result := fint.New(fint.NewRepeat(42)).Take(3).Collect()
					expected := []int{42, 42, 42}

					if !reflect.DeepEqual(expected, result) {
						panic(fmt.Sprintf("expected %v to equal %v", expected, result))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates with Chain", func() {
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
					a := fint.NewRepeat(7)
					b := fint.NewRepeat(42)
					result := fint.New(a).Take(2).Chain(b).Take(4).Collect()
					expected := []int{7, 7, 42, 42}

					if !reflect.DeepEqual(expected, result) {
						panic(fmt.Sprintf("expected %v to equal %v", expected, result))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})

		It("generates with Map", func() {
			cmd := goFunctionalCommand(someBinPath, "int")
			Expect(cmd.Run()).To(Succeed())

			cmd = makeFunctionalSample(workDir, "somebin", clean(`
				package main

				import (
					"fmt"
					"reflect"
					"somebin/fint"
				)

				func increment(i int) int {
					return i + 1
				}

				func main() {
					slice := []int{7, 8}
					result := fint.Lift(slice).Map(increment).Collect()
					expected := []int{8, 9}

					if !reflect.DeepEqual(expected, result) {
						panic(fmt.Sprintf("expected %v to equal %v", expected, result))
					}
				}
			`))

			Expect(cmd.Run()).To(Succeed())
		})
	})
})
