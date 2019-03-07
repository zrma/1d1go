package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/1d1c/hacker_rank/common/utils"
)

//noinspection SpellCheckingInspection
var _ = Describe("https://www.hackerrank.com/challenges/30-abstract-classes/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() error {
			AbstractClasses("The Alchemist", "Paulo Coelho", 248)
			return nil
		}, []string{
			"Title: The Alchemist",
			"Author: Paulo Coelho",
			"Price: 248",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
