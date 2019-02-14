package dictionaries_and_hashmaps

import (
	. "github.com/onsi/ginkgo/extensions/table"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//noinspection SpellCheckingInspection
var _ = XDescribe("https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem", func() {
	DescribeTable("문제를 풀었다",
		func(s string, i int) {
			actual := sherlockAndAnagrams(s)
			Expect(actual).Should(BeNumerically("==", i))
		},
		Entry("test_0", "abba", 4),
		Entry("test_1", "abcd", 0),
		Entry("test_2", "ifailuhkqq", 3),
		Entry("test_3", "kkkk", 10),
		Entry("test_4", "cdcd", 5),
	)
})
