package arrays

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/new-year-chaos/problem", func() {
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleMinimumBribes()

			close(done)
		}()
	})
})

func ExampleMinimumBribes() {
	MinimumBribes([]int32{2, 1, 5, 3, 4})
	MinimumBribes([]int32{2, 5, 1, 3, 4})
	MinimumBribes([]int32{5, 1, 2, 3, 7, 8, 6, 4})
	MinimumBribes([]int32{1, 2, 5, 3, 7, 8, 6, 4})
	MinimumBribes([]int32{1, 2, 5, 3, 4, 7, 8, 6})
	// Output:
	// 3
	// Too chaotic
	// Too chaotic
	// 7
	// 4
}
