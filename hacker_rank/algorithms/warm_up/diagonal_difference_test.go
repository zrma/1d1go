package warm_up

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/diagonal-difference/problem", func() {
	It("문제를 풀었다", func() {
		arr := [][]int32{
			{11, 2, 4},
			{4, 5, 6},
			{10, 8, -12},
		}
		actual := diagonalDifference(arr)
		Expect(actual).Should(BeNumerically("==", 15))

		arr = [][]int32{
			{4, 2, 11},
			{4, 5, 6},
			{-12, 8, 10},
		}
		actual = diagonalDifference(arr)
		Expect(actual).Should(BeNumerically("==", 15))
	})
})
