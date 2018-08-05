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
				result, err := finterface.Lift(slice).Collect()
				if err != nil {
					panic(fmt.Sprintf("expected err not to have occurred: %v", err))
				}
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
				result := finterface.Lift(slice).Drop(2).Collapse()
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
				result := finterface.Lift(slice).Take(2).Collapse()
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
				result := finterface.Lift(slice).Filter(hasLen3).Collapse()
				expected := []interface{}{"bar", "baz"}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with FilterErr", func() {
		cmd := goFunctionalCommand(someBinPath, "interface{}")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/finterface"
			)

			func hasLen3(v interface{}) (bool, error) {
				s, ok := v.(string)
				if !ok {
					panic(fmt.Sprintf("expected %v to be a string", v))
				}
				return len(s) == 3, nil
			}

			func main() {
				slice := []interface{}{"bar", "foos", "baz"}
				result, err := finterface.Lift(slice).FilterErr(hasLen3).Collect()
				if err != nil {
					panic(fmt.Sprintf("expected err not to have occurred: %v", err))
				}
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
				result := finterface.Lift(slice).Exclude(isEmpty).Collapse()
				expected := []interface{}{"foos", "baz"}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with ExcludeErr", func() {
		cmd := goFunctionalCommand(someBinPath, "interface{}")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/finterface"
			)

			func isEmpty(v interface{}) (bool, error) {
				s, ok := v.(string)
				if !ok {
					panic(fmt.Sprintf("expected %v to be a string", v))
				}
				return s == "", nil
			}

			func main() {
				slice := []interface{}{"", "foos", "baz"}
				result, err := finterface.Lift(slice).ExcludeErr(isEmpty).Collect()
				if err != nil {
					panic(fmt.Sprintf("expected err not to have occurred: %v", err))
				}
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
				result := finterface.New(finterface.Repeat("foo")).Take(3).Collapse()
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
				result := finterface.New(foos).Take(2).Chain(bars).Take(4).Collapse()
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
				result := finterface.Lift(slice).Map(prependFoo).Collapse()
				expected := []interface{}{"foobar", "foobaz"}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with MapErr", func() {
		cmd := goFunctionalCommand(someBinPath, "interface{}")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/finterface"
			)

			func prependFoo(v interface{}) (interface{}, error) {
				s, ok := v.(string)
				if !ok {
					panic(fmt.Sprintf("expected %v to be a string", v))
				}
				return "foo" + s, nil
			}

			func main() {
				slice := []interface{}{"bar", "baz"}
				result, err := finterface.Lift(slice).MapErr(prependFoo).Collect()
				if err != nil {
					panic(fmt.Sprintf("expected err not to have occurred: %v", err))
				}
				expected := []interface{}{"foobar", "foobaz"}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Fold", func() {
		cmd := goFunctionalCommand(someBinPath, "interface{}")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"somebin/finterface"
			)

			func sum(a, b interface{}) (interface{}, error) {
				x, ok := a.(int)
				if !ok {
					panic(fmt.Sprintf("expected %v to be an int", a))
				}

				y, ok := b.(int)
				if !ok {
					panic(fmt.Sprintf("expected %v to be an int", b))
				}

				return x + y, nil
			}

			func main() {
				slice := []interface{}{1, 2, 3, 4}
				result, err := finterface.Lift(slice).Fold(0, sum)
				if err != nil {
					panic(fmt.Sprintf("expected err not to have occurred: %v", err))
				}
				expected := interface{}(10)

				if result != expected {
					panic(fmt.Sprintf("expected %v to equal %d", expected, result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Roll", func() {
		cmd := goFunctionalCommand(someBinPath, "interface{}")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"somebin/finterface"
			)

			func sum(a, b interface{}) interface{} {
				x, ok := a.(int)
				if !ok {
					panic(fmt.Sprintf("expected %v to be an int", a))
				}

				y, ok := b.(int)
				if !ok {
					panic(fmt.Sprintf("expected %v to be an int", b))
				}

				return x + y
			}

			func main() {
				slice := []interface{}{1, 2, 3, 4}
				result := finterface.Lift(slice).Roll(0, sum)
				expected := interface{}(10)

				if result != expected {
					panic(fmt.Sprintf("expected %v to equal %d", expected, result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Transmute", func() {
		cmd := goFunctionalCommand(someBinPath, "interface{}")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"somebin/finterface"
			)

			func main() {
				v := interface{}(interface{}("foo"))
				result := finterface.Transmute(v)

				s, ok := result.(string)
				if !ok {
					panic(fmt.Sprintf("expected %v to be a string", result))
				}

				if s != "foo" {
					panic(fmt.Sprintf("expected %s to equal foo", s))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})
})
