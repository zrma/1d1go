package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/1d1c/hacker_rank/common/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-loops/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() error {
			letsReview("Hacker")
			letsReview("Rank")
			return nil
		}, []string{
			"Hce akr",
			"Rn ak",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
