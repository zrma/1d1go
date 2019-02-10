package arrays

import (
	. "github.com/onsi/ginkgo/extensions/table"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/ctci-array-left-rotation/problem", func() {
	a := []int32{1, 2, 3, 4, 5}

	DescribeTable("문제를 풀었다",
		func(d int32, orig, expected []int32) {
			actual := rotLeft(orig, d)
			Expect(actual).Should(Equal(expected))
		},
		Entry("0", int32(0), a, []int32{1, 2, 3, 4, 5}),
		Entry("1", int32(1), a, []int32{2, 3, 4, 5, 1}),
		Entry("2", int32(2), a, []int32{3, 4, 5, 1, 2}),
		Entry("3", int32(3), a, []int32{4, 5, 1, 2, 3}),
		Entry("4", int32(4), a, []int32{5, 1, 2, 3, 4}),
		Entry("5", int32(5), a, []int32{1, 2, 3, 4, 5}),
	)
})
