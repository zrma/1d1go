package warm_up

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/sock-merchant/problem", func() {
	arr := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}

	It("문제를 풀었다", func() {
		actual := sockMerchant(arr)
		Expect(actual).Should(BeNumerically("==", 3))
	})
})
