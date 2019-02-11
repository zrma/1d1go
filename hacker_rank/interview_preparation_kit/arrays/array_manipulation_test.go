package arrays

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/crush/problem", func() {
	It("문제를 풀었다", func() {
		arr := [][]int32{
			{1, 5, 3},
			{4, 8, 7},
			{6, 9, 1},
		}
		actual := arrayManipulation(10, arr)
		Expect(actual).Should(BeNumerically("==", 10))

		arr = [][]int32{
			{1, 2, 100},
			{2, 5, 100},
			{3, 4, 100},
		}
		actual = arrayManipulation(5, arr)
		Expect(actual).Should(BeNumerically("==", 200))

		arr = [][]int32{
			{2, 6, 8},
			{3, 5, 7},
			{1, 8, 1},
			{5, 9, 15},
		}
		actual = arrayManipulation(10, arr)
		Expect(actual).Should(BeNumerically("==", 31))

		arr = [][]int32{
			{3, 5, 7},
			{2, 6, 8},
			{1, 8, 1},
			{5, 9, 15},
		}
		actual = arrayManipulation(10, arr)
		Expect(actual).Should(BeNumerically("==", 31))
	})
})
