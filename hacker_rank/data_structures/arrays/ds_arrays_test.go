package arrays

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/arrays-ds/problem", func() {
	It("문제를 풀었다", func() {
		arr := []int32{1, 4, 3, 2}
		actual := reverseArray(arr)
		Expect(actual).Should(Equal([]int32{2, 3, 4, 1}))
	})
})
