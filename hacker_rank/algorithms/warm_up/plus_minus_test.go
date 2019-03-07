package warm_up

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/1d1c/hacker_rank/common/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/plus-minus/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() error {
			plusMinus([]int32{-4, 3, -9, 0, 4, 1})
			return nil
		}, []string{
			"0.500000",
			"0.333333",
			"0.166667",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
