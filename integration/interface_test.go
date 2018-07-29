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

	It("generates and is importable", func() {
		cmd := goFunctionalCommand(someBinPath, "interface{}")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import "somebin/finterface"

			func main() {
				_ = []finterface.T{"", "", "bar", ""}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Lift", func() {
		cmd := goFunctionalCommand(someBinPath, "interface{}")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import "somebin/finterface"

			func main() {
				slice := []interface{}{"", "", "bar", ""}
				_ = finterface.Lift(slice)
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Collect", func() {
		cmd := goFunctionalCommand(someBinPath, "interface{}")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/finterface"
			)

			func main() {
				slice := []interface{}{"bar", "foo"}
				result := finterface.Lift(slice).Collect()
				expected := []interface{}{"bar", "foo"}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Drop", func() {
		cmd := goFunctionalCommand(someBinPath, "interface{}")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/finterface"
			)

			func main() {
				slice := []interface{}{"bar", "foo", "baz"}
				result := finterface.Lift(slice).Drop(2).Collect()
				expected := []interface{}{"baz"}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Take", func() {
		cmd := goFunctionalCommand(someBinPath, "interface{}")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/finterface"
			)

			func main() {
				slice := []interface{}{"bar", "foo", "baz"}
				result := finterface.Lift(slice).Take(2).Collect()
				expected := []interface{}{"bar", "foo"}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Filter", func() {
		cmd := goFunctionalCommand(someBinPath, "interface{}")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/finterface"
			)

			func hasLen3(v interface{}) bool {
				s, ok := v.(string)
				if !ok {
					panic(fmt.Sprintf("expected %v to be a string", v))
				}
				return len(s) == 3
			}

			func main() {
				slice := []interface{}{"bar", "foos", "baz"}
				result := finterface.Lift(slice).Filter(hasLen3).Collect()
				expected := []interface{}{"bar", "baz"}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Exclude", func() {
		cmd := goFunctionalCommand(someBinPath, "interface{}")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/finterface"
			)

			func isEmpty(v interface{}) bool {
				s, ok := v.(string)
				if !ok {
					panic(fmt.Sprintf("expected %v to be a string", v))
				}
				return s == ""
			}

			func main() {
				slice := []interface{}{"", "foos", "baz"}
				result := finterface.Lift(slice).Exclude(isEmpty).Collect()
				expected := []interface{}{"foos", "baz"}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Repeat", func() {
		cmd := goFunctionalCommand(someBinPath, "interface{}")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/finterface"
			)

			func main() {
				result := finterface.New(finterface.Repeat("foo")).Take(3).Collect()
				expected := []interface{}{"foo", "foo", "foo"}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Chain", func() {
		cmd := goFunctionalCommand(someBinPath, "interface{}")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/finterface"
			)

			func main() {
				foos := finterface.Repeat("foo")
				bars := finterface.Repeat("bar")
				result := finterface.New(foos).Take(2).Chain(bars).Take(4).Collect()
				expected := []interface{}{"foo", "foo", "bar", "bar"}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Map", func() {
		cmd := goFunctionalCommand(someBinPath, "interface{}")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/finterface"
			)

			func prependFoo(v interface{}) interface{} {
				s, ok := v.(string)
				if !ok {
					panic(fmt.Sprintf("expected %v to be a string", v))
				}
				return "foo" + s
			}

			func main() {
				slice := []interface{}{"bar", "baz"}
				result := finterface.Lift(slice).Map(prependFoo).Collect()
				expected := []interface{}{"foobar", "foobaz"}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})
})
