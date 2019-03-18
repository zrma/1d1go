package warm_up

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/1d1c/hacker_rank/common/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/staircase/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			staircase(6)
		}, []string{
			"     #",
			"    ##",
			"   ###",
			"  ####",
			" #####",
			"######",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
