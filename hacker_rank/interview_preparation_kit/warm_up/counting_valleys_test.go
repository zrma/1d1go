package warm_up

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/counting-valleys/problem", func() {
	var n int32 = 8

	//noinspection SpellCheckingInspection
	s := "UDDDUDUU"

	It("문제를 풀었다", func() {
		actual := countingValleys(n, s)
		Expect(actual).Should(BeNumerically("==", 1))
	})
})
