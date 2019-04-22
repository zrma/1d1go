package arrays

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/going/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/new-year-chaos/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			minimumBribes([]int32{2, 1, 5, 3, 4})
			minimumBribes([]int32{2, 5, 1, 3, 4})
			minimumBribes([]int32{5, 1, 2, 3, 7, 8, 6, 4})
			minimumBribes([]int32{1, 2, 5, 3, 7, 8, 6, 4})
			minimumBribes([]int32{1, 2, 5, 3, 4, 7, 8, 6})
		}, []string{
			"3",
			"Too chaotic",
			"Too chaotic",
			"7",
			"4",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
