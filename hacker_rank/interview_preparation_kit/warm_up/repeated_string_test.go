package warm_up

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/repeated-string/problem", func() {
	It("문제를 풀었다", func() {
		actual := repeatedString("aba", 10)
		Expect(actual).Should(BeNumerically("==", 7))

		actual = repeatedString("a", 1000000000000)
		Expect(actual).Should(BeNumerically("==", 1000000000000))

		actual = repeatedString("afternoon after", 3)
		Expect(actual).Should(BeNumerically("==", 1))

		actual = repeatedString("school hometown", 5)
		Expect(actual).Should(BeNumerically("==", 0))

		actual = repeatedString("", 5)
		Expect(actual).Should(BeNumerically("==", 0))
	})
})
