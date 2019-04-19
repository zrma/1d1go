package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/1d1c/hacker_rank/common/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-linked-list/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			displayLinkedList([]int{2, 3, 4, 1})
		}, []string{
			"2",
			"3",
			"4",
			"1",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
