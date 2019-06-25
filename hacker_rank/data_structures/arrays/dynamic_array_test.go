package arrays

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/2d-array/problem", func() {
	It("문제를 풀었다", func() {
		arr := [][]int32{
			{1, 0, 5},
			{1, 1, 7},
			{1, 0, 3},
			{2, 1, 0},
			{2, 1, 1},
		}
		actual := dynamicArray(2, arr)
		Expect(actual).Should(Equal([]int32{7, 3}))
	})
})
