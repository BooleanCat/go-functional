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
				result, err := fstring.Lift(slice).Collect()
				if err != nil {
					panic(fmt.Sprintf("expected err not to have occurred: %v", err))
				}
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
				result := fstring.Lift(slice).Drop(2).Collapse()
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
				result := fstring.Lift(slice).Take(2).Collapse()
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
				result := fstring.Lift(slice).Filter(hasLen3).Collapse()
				expected := []string{"bar", "baz"}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with FilterErr", func() {
		cmd := goFunctionalCommand(someBinPath, "string")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/fstring"
			)

			func hasLen3(s string) (bool, error) {
				return len(s) == 3, nil
			}

			func main() {
				slice := []string{"bar", "foos", "baz"}
				result, err := fstring.Lift(slice).FilterErr(hasLen3).Collect()
				if err != nil {
					panic(fmt.Sprintf("expected err not to have occurred: %v", err))
				}
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
				result := fstring.Lift(slice).Exclude(isEmpty).Collapse()
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
				result := fstring.New(fstring.Repeat("foo")).Take(3).Collapse()
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
				foos := fstring.Repeat("foo")
				bars := fstring.Repeat("bar")
				result := fstring.New(foos).Take(2).Chain(bars).Take(4).Collapse()
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
				"strings"
				"somebin/fstring"
			)

			func main() {
				slice := []string{"bar", "baz"}
				result := fstring.Lift(slice).Map(strings.Title).Collapse()
				expected := []string{"Bar", "Baz"}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with MapErr", func() {
		cmd := goFunctionalCommand(someBinPath, "string")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"strings"
				"somebin/fstring"
			)

			func title(s string) (string, error) {
				return strings.Title(s), nil
			}

			func main() {
				slice := []string{"bar", "baz"}
				result, err := fstring.Lift(slice).MapErr(title).Collect()
				if err != nil {
					panic(fmt.Sprintf("expected err not to have occurred: %v", err))
				}
				expected := []string{"Bar", "Baz"}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Fold", func() {
		cmd := goFunctionalCommand(someBinPath, "string")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"somebin/fstring"
			)

			func prepend(a, b string) string {
				return b + a
			}

			func main() {
				slice := []string{"foo", "bar", "baz"}
				result := fstring.Lift(slice).Fold("", prepend)

				if result != "bazbarfoo" {
					panic(fmt.Sprintf("expected bazbarfoo to equal %d", result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})
})
