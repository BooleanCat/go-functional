package gen_test

import (
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/renstrom/dedent"
)

func TestGen(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gen Suite")
}

func clean(text string) string {
	return strings.TrimSpace(dedent.Dedent(text)) + "\n"
}
