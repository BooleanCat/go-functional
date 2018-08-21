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
		cmd := goFunctionalCommand(someBinPath, "*int")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import "somebin/fpint"

			func main() {
				_ = []fpint.T{newInt(1), newInt(2)}
			}

			func newInt(i int) *int {
				return &i
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Lift", func() {
		cmd := goFunctionalCommand(someBinPath, "*int")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import "somebin/fpint"

			func main() {
				slice := []*int{newInt(1), newInt(2), newInt(3), newInt(4)}
				_ = fpint.Lift(slice)
			}

			func newInt(i int) *int {
				return &i
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Collect", func() {
		cmd := goFunctionalCommand(someBinPath, "*int")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/fpint"
			)

			func main() {
				slice := []*int{newInt(1), newInt(2), newInt(3)}
				result, err := fpint.Lift(slice).Collect()
				if err != nil {
					panic(fmt.Sprintf("expected err not to have occurred: %v", err))
				}
				expected := []*int{newInt(1), newInt(2), newInt(3)}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}

			func newInt(i int) *int {
				return &i
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Drop", func() {
		cmd := goFunctionalCommand(someBinPath, "*int")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/fpint"
			)

			func main() {
				slice := []*int{newInt(1), newInt(2), newInt(3)}
				result := fpint.Lift(slice).Drop(2).Collapse()
				expected := []*int{newInt(3)}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}

			func newInt(i int) *int {
				return &i
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Take", func() {
		cmd := goFunctionalCommand(someBinPath, "*int")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/fpint"
			)

			func main() {
				slice := []*int{newInt(1), newInt(2), newInt(3)}
				result := fpint.Lift(slice).Take(2).Collapse()
				expected := []*int{newInt(1), newInt(2)}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}

			func newInt(i int) *int {
				return &i
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Filter", func() {
		cmd := goFunctionalCommand(someBinPath, "*int")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/fpint"
			)

			func isOdd(i *int) bool {
				return *i % 2 == 1
			}

			func main() {
				slice := []*int{newInt(1), newInt(2), newInt(3)}
				result := fpint.Lift(slice).Filter(isOdd).Collapse()
				expected := []*int{newInt(1), newInt(3)}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}

			func newInt(i int) *int {
				return &i
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with FilterErr", func() {
		cmd := goFunctionalCommand(someBinPath, "*int")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/fpint"
			)

			func isOdd(i *int) (bool, error) {
				return *i % 2 == 1, nil
			}

			func main() {
				slice := []*int{newInt(1), newInt(2), newInt(3)}
				result, err := fpint.Lift(slice).FilterErr(isOdd).Collect()
				if err != nil {
					panic(fmt.Sprintf("expected err not to have occurred: %v", err))
				}
				expected := []*int{newInt(1), newInt(3)}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}

			func newInt(i int) *int {
				return &i
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Exclude", func() {
		cmd := goFunctionalCommand(someBinPath, "*int")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/fpint"
			)

			func isOdd(i *int) bool {
				return *i % 2 == 1
			}

			func main() {
				slice := []*int{newInt(1), newInt(2), newInt(3)}
				result := fpint.Lift(slice).Exclude(isOdd).Collapse()
				expected := []*int{newInt(2)}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}

			func newInt(i int) *int {
				return &i
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with ExcludeErr", func() {
		cmd := goFunctionalCommand(someBinPath, "*int")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/fpint"
			)

			func isOdd(i *int) (bool, error) {
				return *i % 2 == 1, nil
			}

			func main() {
				slice := []*int{newInt(1), newInt(2), newInt(3)}
				result, err := fpint.Lift(slice).ExcludeErr(isOdd).Collect()
				if err != nil {
					panic(fmt.Sprintf("expected err not to have occurred: %v", err))
				}
				expected := []*int{newInt(2)}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}

			func newInt(i int) *int {
				return &i
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Repeat", func() {
		cmd := goFunctionalCommand(someBinPath, "*int")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/fpint"
			)

			func main() {
				result := fpint.New(fpint.Repeat(newInt(42))).Take(3).Collapse()
				expected := []*int{newInt(42), newInt(42), newInt(42)}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}

			func newInt(i int) *int {
				return &i
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Chain", func() {
		cmd := goFunctionalCommand(someBinPath, "*int")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/fpint"
			)

			func main() {
				a := fpint.Repeat(newInt(7))
				b := fpint.Repeat(newInt(42))
				result := fpint.New(a).Take(2).Chain(b).Take(4).Collapse()
				expected := []*int{newInt(7), newInt(7), newInt(42), newInt(42)}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}

			func newInt(i int) *int {
				return &i
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Map", func() {
		cmd := goFunctionalCommand(someBinPath, "*int")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/fpint"
			)

			func increment(i *int) *int {
				j := *i + 1
				return &j
			}

			func main() {
				slice := []*int{newInt(7), newInt(8)}
				result := fpint.Lift(slice).Map(increment).Collapse()
				expected := []*int{newInt(8), newInt(9)}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}

			func newInt(i int) *int {
				return &i
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with MapErr", func() {
		cmd := goFunctionalCommand(someBinPath, "*int")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"somebin/fpint"
			)

			func increment(i *int) (*int, error) {
				j := *i + 1
				return &j, nil
			}

			func main() {
				slice := []*int{newInt(7), newInt(8)}
				result, err := fpint.Lift(slice).MapErr(increment).Collect()
				if err != nil {
					panic(fmt.Sprintf("expected err not to have occurred: %v", err))
				}
				expected := []*int{newInt(8), newInt(9)}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}

			func newInt(i int) *int {
				return &i
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Fold", func() {
		cmd := goFunctionalCommand(someBinPath, "*int")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"somebin/fpint"
			)

			func sum(a, b *int) (*int, error) {
				c := *a + *b
				fmt.Println(c)
				return &c, nil
			}

			func main() {
				slice := []*int{newInt(1), newInt(2), newInt(3), newInt(4)}
				result, err := fpint.Lift(slice).Fold(newInt(0), sum)
				if err != nil {
					panic(fmt.Sprintf("expected err not to have occurred: %v", err))
				}

				if *result != 10 {
					panic(fmt.Sprintf("expected 10 to equal %d", result))
				}
			}

			func newInt(i int) *int {
				return &i
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Roll", func() {
		cmd := goFunctionalCommand(someBinPath, "*int")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"somebin/fpint"
			)

			func sum(a, b *int) *int {
				c := *a + *b
				return &c
			}

			func main() {
				slice := []*int{newInt(1), newInt(2), newInt(3), newInt(4)}
				result := fpint.Lift(slice).Roll(newInt(0), sum)

				if *result != 10 {
					panic(fmt.Sprintf("expected 10 to equal %d", result))
				}
			}

			func newInt(i int) *int {
				return &i
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Transmute", func() {
		cmd := goFunctionalCommand(someBinPath, "*int")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"somebin/fpint"
			)

			func main() {
				v := interface{}(newInt(4))
				i := fpint.Transmute(v)
				if *i != 4 {
					panic(fmt.Sprintf("expected %d to equal 4", i))
				}
			}

			func newInt(i int) *int {
				return &i
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Transform", func() {
		cmd := goFunctionalCommand(someBinPath, "*int")
		Expect(cmd.Run()).To(Succeed())

		cmd = goFunctionalCommand(someBinPath, "*string")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"reflect"
				"strconv"
				"somebin/fpstring"
				"somebin/fpint"
			)

			type Counter struct {
				i fpint.T
			}

			func NewCounter() *Counter {
				return &Counter{i: newInt(0)}
			}

			func (iter *Counter) Next() fpint.OptionalResult {
				next := *iter.i
				i := next + 1
				iter.i = &i
				return fpint.Success(fpint.Some(&next))
			}

			func asString(v interface{}) (*string, error) {
				s := strconv.Itoa(*fpint.Transmute(v))
				return &s, nil
			}

			func main() {
				iter := fpint.New(NewCounter()).Blur()
				numbers := fpstring.New(fpstring.Transform(iter, asString)).Take(4).Collapse()

				expected := []*string{newString("0"), newString("1"), newString("2"), newString("3")}
				if !reflect.DeepEqual(expected, numbers) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, numbers))
				}
			}

			func newString(s string) *string {
				return &s
			}

			func newInt(i int) *int {
				return &i
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})
})
