package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/1d1c/hacker_rank/common/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-more-exceptions/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			MoreException(3, 5)
			MoreException(2, 4)
			MoreException(-1, -2)
			MoreException(-1, 3)
		}, []string{
			"243",
			"16",
			"n and p should be non-negative",
			"n and p should be non-negative",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
