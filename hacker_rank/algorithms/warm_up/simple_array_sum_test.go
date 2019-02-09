package warm_up

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/simple-array-sum/problem", func() {
	It("문제를 풀었다", func() {
		arr := []int32{1, 2, 3, 4, 10, 11}
		actual := simpleArraySum(arr)
		Expect(actual).Should(BeNumerically("==", 31))
	})
})
