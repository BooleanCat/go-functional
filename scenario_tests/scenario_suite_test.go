package scenario_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoFunctionalScenarios(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "go-functional scenario suite")
}
