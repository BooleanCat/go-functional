package acceptance_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BooleanCat/go-functional/fixtures/ffilemode"
	"github.com/BooleanCat/go-functional/fixtures/fstring"
)

var _ = Describe("go-functional", func() {
	It("generates and is importable", func() {
		_ = []ffilemode.T{ffilemode.T(0755)}
	})

	It("generates with Lift", func() {
		slice := []os.FileMode{0777, 0555, 0666}
		_ = ffilemode.Lift(slice)
	})

	It("generates with Collect", func() {
		slice := []os.FileMode{0777, 0555, 0666}
		result, err := ffilemode.Lift(slice).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]os.FileMode{0777, 0555, 0666}))
	})

	It("generates with Drop", func() {
		slice := []os.FileMode{0777, 0555, 0666}
		result := ffilemode.Lift(slice).Drop(2).Collapse()
		Expect(result).To(Equal([]os.FileMode{0666}))
	})

	It("generates with Take", func() {
		slice := []os.FileMode{0777, 0555, 0666}
		result := ffilemode.Lift(slice).Take(2).Collapse()
		Expect(result).To(Equal([]os.FileMode{0777, 0555}))
	})

	It("generates with Filter", func() {
		isDir := func(m os.FileMode) bool { return m.IsDir() }

		slice := []os.FileMode{0755, 0777 | os.ModeDir, 0666}
		result := ffilemode.Lift(slice).Filter(isDir).Collapse()
		Expect(result).To(Equal([]os.FileMode{0777 | os.ModeDir}))
	})

	It("generates with FilterErr", func() {
		isDir := func(m os.FileMode) (bool, error) { return m.IsDir(), nil }

		slice := []os.FileMode{0755, 0777 | os.ModeDir, 0666}
		result, err := ffilemode.Lift(slice).FilterErr(isDir).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]os.FileMode{0777 | os.ModeDir}))
	})

	It("generates with Exclude", func() {
		isDir := func(m os.FileMode) bool { return m.IsDir() }

		slice := []os.FileMode{0755, 0777 | os.ModeDir, 0666}
		result := ffilemode.Lift(slice).Exclude(isDir).Collapse()
		Expect(result).To(Equal([]os.FileMode{0755, 0666}))
	})

	It("generates with ExcludeErr", func() {
		isDir := func(m os.FileMode) (bool, error) { return m.IsDir(), nil }

		slice := []os.FileMode{0755, 0777 | os.ModeDir, 0666}
		result, err := ffilemode.Lift(slice).ExcludeErr(isDir).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]os.FileMode{0755, 0666}))
	})

	It("generates with Repeat", func() {
		result := ffilemode.New(ffilemode.Repeat(0777)).Take(3).Collapse()
		Expect(result).To(Equal([]os.FileMode{0777, 0777, 0777}))
	})

	It("generates with Chain", func() {
		a := ffilemode.Repeat(0666)
		b := ffilemode.Repeat(0411)

		result := ffilemode.New(a).Take(2).Chain(b).Take(4).Collapse()
		Expect(result).To(Equal([]os.FileMode{0666, 0666, 0411, 0411}))
	})

	It("generates with Map", func() {
		oPlusX := func(m os.FileMode) os.FileMode { return m | 0100 }

		slice := []os.FileMode{0444, 0611}
		result := ffilemode.Lift(slice).Map(oPlusX).Collapse()
		Expect(result).To(Equal([]os.FileMode{0544, 0711}))
	})

	It("generates with MapErr", func() {
		oPlusX := func(m os.FileMode) (os.FileMode, error) { return m | 0100, nil }

		slice := []os.FileMode{0444, 0611}
		result, err := ffilemode.Lift(slice).MapErr(oPlusX).Collect()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal([]os.FileMode{0544, 0711}))
	})

	It("generates with Fold", func() {
		or := func(a, b os.FileMode) (os.FileMode, error) { return a | b, nil }

		slice := []os.FileMode{0100, 0110, 0141, os.ModeDir}
		result, err := ffilemode.Lift(slice).Fold(0, or)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(0151 | os.ModeDir))
	})

	It("generates with Roll", func() {
		or := func(a, b os.FileMode) os.FileMode { return a | b }

		slice := []os.FileMode{0100, 0110, 0141, os.ModeDir}
		result := ffilemode.Lift(slice).Roll(0, or)
		Expect(result).To(Equal(0151 | os.ModeDir))
	})

	It("generates with Transmute", func() {
		v := interface{}(os.FileMode(0755))
		Expect(ffilemode.Transmute(v)).To(Equal(os.FileMode(0755)))

		var expectedType os.FileMode
		Expect(ffilemode.Transmute(v)).To(BeAssignableToTypeOf(expectedType))
	})

	It("generates with Transform", func() {
		asString := func(v interface{}) (string, error) { return ffilemode.Transmute(v).String(), nil }

		iter := ffilemode.New(ffilemode.Repeat(0755)).Blur()
		result := fstring.New(fstring.Transform(iter, asString)).Take(4).Collapse()
		Expect(result).To(Equal([]string{"-rwxr-xr-x", "-rwxr-xr-x", "-rwxr-xr-x", "-rwxr-xr-x"}))
	})
})
