package warmup

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/counting-valleys/problem", func() {
	It("문제를 풀었다", func() {
		//noinspection SpellCheckingInspection
		s := "UDDDUDUU"
		actual := countingValleys(int32(len(s)), s)
		Expect(actual).Should(BeNumerically("==", 1))

		s = ""
		actual = countingValleys(int32(len(s)), s)
		Expect(actual).Should(BeNumerically("==", 0))

		s = "ABC"
		actual = countingValleys(int32(len(s)), s)
		Expect(actual).Should(BeNumerically("==", 0))

		s = "ABC"
		actual = countingValleys(int32(len(s)+1), s)
		Expect(actual).Should(BeNumerically("==", 0))
	})
})
