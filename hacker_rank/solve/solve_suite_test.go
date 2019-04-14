package solve_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSolve(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Solve Suite")
}
