package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-arrays/problem", func() {
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExamplePrintReverse()

			close(done)
		}()
	})
})

func ExamplePrintReverse() {
	PrintReverse([]int32{1, 4, 3, 2})
	// Output:
	// 2 3 4 1
}
