package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-interfaces/problem", func() {
	It("문제를 풀었다", func() {
		c := calculator{}

		Expect(c.divisorSum(1)).Should(BeNumerically("==", 1))
		Expect(c.divisorSum(6)).Should(BeNumerically("==", 12))
		Expect(c.divisorSum(20)).Should(BeNumerically("==", 42))
	})
})
