package warmup_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWarmup(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Warmup Suite")
}
