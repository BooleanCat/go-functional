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
				result, err := fint.Lift(slice).Collect()
				if err != nil {
					panic(fmt.Sprintf("expected err not to have occurred: %v", err))
				}
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
				result := fint.Lift(slice).Drop(2).Collapse()
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
				result := fint.Lift(slice).Take(2).Collapse()
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
				result := fint.Lift(slice).Filter(isOdd).Collapse()
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
				result := fint.Lift(slice).Exclude(isOdd).Collapse()
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
				result := fint.New(fint.Repeat(42)).Take(3).Collapse()
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
				a := fint.Repeat(7)
				b := fint.Repeat(42)
				result := fint.New(a).Take(2).Chain(b).Take(4).Collapse()
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
				result := fint.Lift(slice).Map(increment).Collapse()
				expected := []int{8, 9}

				if !reflect.DeepEqual(expected, result) {
					panic(fmt.Sprintf("expected %v to equal %v", expected, result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})

	It("generates with Fold", func() {
		cmd := goFunctionalCommand(someBinPath, "int")
		Expect(cmd.Run()).To(Succeed())

		cmd = makeFunctionalSample(workDir, "somebin", clean(`
			package main

			import (
				"fmt"
				"somebin/fint"
			)

			func sum(a, b int) int {
				return a + b
			}

			func main() {
				slice := []int{1, 2, 3, 4}
				result := fint.Lift(slice).Fold(0, sum)

				if result != 10 {
					panic(fmt.Sprintf("expected 10 to equal %d", result))
				}
			}
		`))

		Expect(cmd.Run()).To(Succeed())
	})
})
