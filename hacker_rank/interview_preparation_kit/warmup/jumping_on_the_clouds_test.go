package warmup

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem", func() {
	type testData struct {
		data     []int32
		expected int32
	}

	DescribeTable("문제를 풀었다.", func(d testData) {
		actual := jumpingOnClouds(d.data)
		Expect(actual).Should(BeNumerically("==", d.expected))
	},
		Entry("test0", testData{[]int32{0, 0, 1, 0, 0, 1, 0}, 4}),
		Entry("test1", testData{[]int32{0, 0, 0, 0, 1, 0}, 3}),
		Entry("test2", testData{[]int32{0, 0, 0, 0, 0}, 2}),
		Entry("test3", testData{[]int32{0, 1, 0, 0, 1, 0}, 3}),
		Entry("test4", testData{[]int32{0, 1, 1, 0, 1, 0}, -1}),
		Entry("test5", testData{[]int32{0, 0, 1, 0, 1, 0}, 3}),
		Entry("test6", testData{[]int32{0, 0, 0, 1, 0, 0}, 3}),
	)
})
