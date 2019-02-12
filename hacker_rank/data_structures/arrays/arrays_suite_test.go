package arrays_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestArrays(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Arrays Suite")
}
