package warm_up

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/plus-minus/problem", func() {
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExamplePlusMinus()

			close(done)
		}()
	})
})

func ExamplePlusMinus() {
	arr := []int32{-4, 3, -9, 0, 4, 1}

	PlusMinus(arr)
	// Output: 0.500000
	// 0.333333
	// 0.166667
}
