package university_codesprint_5

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/1d1c/hacker_rank/common/utils"
)

var _ = Describe("https://www.hackerrank.com/contests/university-codesprint-5/challenges/exceeding-speed-limit", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			exceedingTheSpeedLimit(100)
			exceedingTheSpeedLimit(140)
			exceedingTheSpeedLimit(85)
		}, []string{
			"3000 Warning",
			"25000 License removed",
			"0 No punishment",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
