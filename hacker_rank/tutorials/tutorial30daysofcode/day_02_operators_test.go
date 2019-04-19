package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/1d1c/hacker_rank/common/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-operators/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			operators(12.00, 20, 8)
		}, []string{
			"15",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
