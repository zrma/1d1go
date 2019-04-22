package warmup

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/going/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/mini-max-sum/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			miniMaxSum([]int64{1, 2, 3, 4, 5})
		}, []string{
			"10 14",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
