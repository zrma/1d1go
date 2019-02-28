package arrays

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/2d-array/problem", func() {
	It("문제를 풀었다", func() {
		arr := [][]int32{
			{-9, -9, -9, 1, 1, 1},
			{0, -9, 0, 4, 3, 2},
			{-9, -9, -9, 1, 2, 3},
			{0, 0, 8, 6, 6, 0},
			{0, 0, 0, -2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		}
		actual := HourGlassSum(arr)
		Expect(actual).Should(BeNumerically("==", 28))

		arr = [][]int32{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 0, 2, 4, 4, 0},
			{0, 0, 0, 2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		}
		actual = HourGlassSum(arr)
		Expect(actual).Should(BeNumerically("==", 19))
	})

	It("예외에도 잘 동작한다.", func() {
		actual := HourGlassSum([][]int32{})
		Expect(actual).Should(BeNumerically("==", 0))

		actual = HourGlassSum([][]int32{
			{1, 2}, {3, 4}, {5, 6}, {7, 8},
		})
	})
})
