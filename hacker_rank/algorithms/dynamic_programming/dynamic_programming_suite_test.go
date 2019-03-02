package dynamic_programming_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDynamicProgramming(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DynamicProgramming Suite")
}
