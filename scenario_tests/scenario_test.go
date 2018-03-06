package scenario_test

import (
	functional "BooleanCat/go-functional"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go-functional scenarios", func() {
	Describe("identifying non-directories", func() {
		// Suppose you have a slice file names, you know these files live in some
		// directory and you also know that some of them are themselves
		// directories. This test demonstrates creating a new slice of fully
		// qualified path names for those files which are not directories.

		var tempdir string

		BeforeEach(func() {
			tempdir = tempDir()
			createFile(filepath.Join(tempdir, "foo"))
			createDir(filepath.Join(tempdir, "bar"))
			createFile(filepath.Join(tempdir, "baz"))
			createDir(filepath.Join(tempdir, "boz"))
			createFile(filepath.Join(tempdir, "biz"))
		})

		AfterEach(func() {
			Expect(os.RemoveAll(tempdir)).To(Succeed())
		})

		It("correctly identifies non-directories", func() {
			names := []string{"foo", "bar", "baz", "boz", "biz"}
			makePath := func(path string) string { return filepath.Join(tempdir, path) }

			collection, err := functional.LiftStringSlice(names).Map(makePath).WithErrs().Filter(isNotDir).Collect()
			Expect(err).NotTo(HaveOccurred())
			Expect(collection).To(Equal([]string{
				filepath.Join(tempdir, "foo"),
				filepath.Join(tempdir, "baz"),
				filepath.Join(tempdir, "biz"),
			}))
		})
	})

	Describe("finding prime numbers", func() {
		// Suppose we want to find all prime numbers less than 100 and then ignore
		// the first 5. Then for some weird reason we want to add them all up.

		It("correctly identifies prime numbers and sums them", func() {
			numbers := make([]int, 100)
			for i := 0; i < 100; i++ {
				numbers[i] = i + 1
			}

			total := functional.LiftIntSlice(numbers).Filter(isPrime).Drop(5).Fold(0, add)
			Expect(total).To(Equal(1032))
		})
	})
})

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func add(a, b int) int {
	return a + b
}

func isNotDir(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return !info.IsDir(), nil
}

func tempDir() string {
	dir, err := ioutil.TempDir("", "")
	Expect(err).NotTo(HaveOccurred())
	return dir
}

func createFile(path string) {
	Expect(ioutil.WriteFile(path, nil, os.ModePerm)).To(Succeed())
}

func createDir(path string) {
	Expect(os.Mkdir(path, os.ModePerm)).To(Succeed())
}
