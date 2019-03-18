package dictionaries_and_hashmaps

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/1d1c/hacker_rank/common/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/ctci-ransom-note/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			checkMagazine(
				[]string{"give", "me", "one", "grand", "today", "night"},
				[]string{"give", "one", "grand", "today"},
			)
			checkMagazine(
				[]string{"two", "times", "three", "is", "not", "four"},
				[]string{"two", "times", "two", "is", "four"},
			)
			checkMagazine(
				[]string{"ive", "got", "a", "lovely", "brunch", "of", "coconuts"},
				[]string{"ive", "got", "some", "coconuts"},
			)
		}, []string{
			"Yes",
			"No",
			"No",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
