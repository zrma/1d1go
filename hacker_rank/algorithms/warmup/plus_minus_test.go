package warmup

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/going/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/plus-minus/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			plusMinus([]int32{-4, 3, -9, 0, 4, 1})
		}, []string{
			"0.500000",
			"0.333333",
			"0.166667",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
