package t_30_days_of_code_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWarmUp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WarmUp Suite")
}
