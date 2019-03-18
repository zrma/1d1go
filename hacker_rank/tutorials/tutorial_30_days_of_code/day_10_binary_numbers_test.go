package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/1d1c/hacker_rank/common/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-binary-numbers/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			binaryNumbers(5)
			binaryNumbers(13)
		}, []string{
			"1",
			"2",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
