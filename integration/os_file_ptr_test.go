package integration_test

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BooleanCat/go-functional/fixtures/fpfile"
	"github.com/BooleanCat/go-functional/fixtures/fstring"
)

var _ = Describe("go-functional", func() {
	var (
		workDir string
		files   []*os.File
	)

	BeforeEach(func() {
		workDir = tempDir()
		writeFileString(filepath.Join(workDir, "foo"), "bar", 0666)
		writeFileString(filepath.Join(workDir, "bar"), "foo", 0660)
		mkdir(filepath.Join(workDir, "baz"), os.ModePerm)
		mkdir(filepath.Join(workDir, "biz"), os.ModePerm)

		files = append(
			[]*os.File{},
			open(filepath.Join(workDir, "foo")),
			open(filepath.Join(workDir, "bar")),
			open(filepath.Join(workDir, "baz")),
			open(filepath.Join(workDir, "biz")),
		)
	})

	AfterEach(func() {
		for _, file := range files {
			Expect(file.Close()).To(Succeed())
		}
		Expect(os.RemoveAll(workDir)).To(Succeed())
	})

	It("generates and is importable", func() {
		_ = fpfile.T(new(os.File))
	})

	It("generates with Lift", func() {
		slice := []*os.File{new(os.File), new(os.File)}
		_ = fpfile.Lift(slice)
	})

	It("generates with Collect", func() {
		result, err := fpfile.Lift(files).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(files))
	})

	It("generates with Drop", func() {
		result := fpfile.Lift(files).Drop(2).Collapse()
		Expect(result).To(Equal(files[2:]))
	})

	It("generates with Take", func() {
		result := fpfile.Lift(files).Take(2).Collapse()
		Expect(result).To(Equal(files[0:2]))
	})

	It("generates with Filter", func() {
		isDir := func(f *os.File) bool {
			s, err := f.Stat()
			Expect(err).NotTo(HaveOccurred())
			return s.IsDir()
		}

		result := fpfile.Lift(files).Filter(isDir).Collapse()
		Expect(result).To(Equal(files[2:]))
	})

	It("generates with FilterErr", func() {
		isDir := func(f *os.File) (bool, error) {
			s, err := f.Stat()
			if err != nil {
				return false, err
			}

			return s.IsDir(), nil
		}

		result, err := fpfile.Lift(files).FilterErr(isDir).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(files[2:]))
	})

	It("generates with Exclude", func() {
		isDir := func(f *os.File) bool {
			s, err := f.Stat()
			Expect(err).NotTo(HaveOccurred())
			return s.IsDir()
		}

		result := fpfile.Lift(files).Exclude(isDir).Collapse()
		Expect(result).To(Equal(files[0:2]))
	})

	It("generates with ExcludeErr", func() {
		isDir := func(f *os.File) (bool, error) {
			s, err := f.Stat()
			if err != nil {
				return false, err
			}

			return s.IsDir(), nil
		}

		result, err := fpfile.Lift(files).ExcludeErr(isDir).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(files[0:2]))
	})

	It("generates with Repeat", func() {
		result := fpfile.New(fpfile.Repeat(files[1])).Take(3).Collapse()
		Expect(result).To(Equal([]*os.File{files[1], files[1], files[1]}))
	})

	It("generates with Chain", func() {
		a := fpfile.Repeat(files[1])
		b := fpfile.Repeat(files[3])

		result := fpfile.New(a).Take(2).Chain(b).Take(4).Collapse()
		Expect(result).To(Equal([]*os.File{files[1], files[1], files[3], files[3]}))
	})

	It("generates with Map", func() {
		aPlusX := func(f *os.File) *os.File {
			s, err := f.Stat()
			Expect(err).NotTo(HaveOccurred())
			Expect(f.Chmod(s.Mode() | 0111)).To(Succeed())

			t, err := f.Stat()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Mode() | 0111).To(Equal(t.Mode()))

			return f
		}

		result := fpfile.Lift(files).Map(aPlusX).Collapse()
		Expect(result).To(Equal(files))
	})

	It("generates with MapErr", func() {
		aPlusX := func(f *os.File) (*os.File, error) {
			s, err := f.Stat()
			if err != nil {
				return nil, err
			}

			if err = f.Chmod(s.Mode() | 0111); err != nil {
				return nil, err
			}

			t, err := f.Stat()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Mode() | 0111).To(Equal(t.Mode()))

			return f, nil
		}

		result, err := fpfile.Lift(files).MapErr(aPlusX).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(files))
	})

	It("generates with Fold", func() {
		var names []string

		getNames := func(a, b *os.File) (*os.File, error) {
			names = append(names, filepath.Base(b.Name()))
			return b, nil
		}

		result, err := fpfile.Lift(files).Fold(nil, getNames)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(files[3]))
		Expect(names).To(Equal([]string{"foo", "bar", "baz", "biz"}))
	})

	It("generates with Roll", func() {
		var names []string

		getNames := func(a, b *os.File) *os.File {
			names = append(names, filepath.Base(b.Name()))
			return b
		}

		result := fpfile.Lift(files).Roll(nil, getNames)
		Expect(result).To(Equal(files[3]))
		Expect(names).To(Equal([]string{"foo", "bar", "baz", "biz"}))
	})

	It("generates with Transmute", func() {
		v := interface{}(files[2])
		Expect(fpfile.Transmute(v)).To(Equal(files[2]))

		var expectedType *os.File
		Expect(fpfile.Transmute(v)).To(BeAssignableToTypeOf(expectedType))
	})

	It("generates with Transform", func() {
		basename := func(v interface{}) (string, error) {
			f := fpfile.Transmute(v)
			return filepath.Base(f.Name()), nil
		}

		iter := fpfile.Lift(files).Blur()
		result := fstring.New(fstring.Transform(iter, basename)).Collapse()
		Expect(result).To(Equal([]string{"foo", "bar", "baz", "biz"}))
	})
})
