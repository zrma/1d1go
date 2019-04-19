package dictionariesandhashmaps

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/two-strings/problem", func() {
	It("문제를 풀었다", func() {
		actual := twoStrings("hello", "world")
		Expect(actual).Should(Equal("YES"))

		actual = twoStrings("hi", "world")
		Expect(actual).Should(Equal("NO"))
	})
})
