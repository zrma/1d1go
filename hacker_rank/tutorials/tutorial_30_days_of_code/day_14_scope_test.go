package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-scope/problem", func() {
	//noinspection SpellCheckingInspection
	It("문제를 풀었다", func() {
		difference := NewDifference([]int{1, 2, 5})
		Expect(difference.computeDifference()).Should(BeNumerically("==", 4))
	})
})
