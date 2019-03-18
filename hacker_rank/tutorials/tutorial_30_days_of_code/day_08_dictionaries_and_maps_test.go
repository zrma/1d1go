package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/1d1c/hacker_rank/common/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-dictionaries-and-maps/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			dictionariesAndMaps(3, []string{
				"sam 99912222",
				"tom 11122222",
				"harry 12299933",
				"sam",
				"edward",
				"harry",
			})
		}, []string{
			"sam=99912222",
			"Not found",
			"harry=12299933",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
