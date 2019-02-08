package warm_up

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/birthday-cake-candles/problem", func() {
	arr := []int32{3, 2, 1, 3}

	It("문제를 풀었다", func() {
		actual := birthdayCakeCandles(arr)
		Expect(actual).Should(BeNumerically("==", 2))
	})
})
