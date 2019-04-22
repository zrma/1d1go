package warmup

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/going/utils"
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
