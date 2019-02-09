package warm_up

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/mini-max-sum/problem", func() {
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleMiniMaxSum()

			close(done)
		}()
	})
})

func ExampleMiniMaxSum() {
	arr := []int64{1, 2, 3, 4, 5}

	MiniMaxSum(arr)
	// Output: 10 14
}
