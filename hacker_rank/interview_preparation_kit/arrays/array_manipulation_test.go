package arrays

import (
	. "github.com/onsi/ginkgo/extensions/table"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/crush/problem", func() {
	DescribeTable("문제를 풀었다",
		func(n int, queries [][]int32, expected int) {
			actual := arrayManipulation(int32(n), queries)
			Expect(actual).Should(Equal(int64(expected)))
		},
		Entry("0", 10, [][]int32{
			{1, 5, 3},
			{4, 8, 7},
			{6, 9, 1},
		}, 10),
		Entry("1", 5, [][]int32{
			{1, 2, 100},
			{2, 5, 100},
			{3, 4, 100},
		}, 200),
		Entry("2", 10, [][]int32{
			{2, 6, 8},
			{3, 5, 7},
			{1, 8, 1},
			{5, 9, 15},
		}, 31),
		Entry("3", 10, [][]int32{
			{3, 5, 7},
			{2, 6, 8},
			{1, 8, 1},
			{5, 9, 15},
		}, 31),
	)
})
