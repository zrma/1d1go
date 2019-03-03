package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-loops/problem", func() {
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleLoop()

			close(done)
		}()
	})
})

func ExampleLoop() {
	Loop(2)
	// Output:
	// 2 x 1 = 2
	// 2 x 2 = 4
	// 2 x 3 = 6
	// 2 x 4 = 8
	// 2 x 5 = 10
	// 2 x 6 = 12
	// 2 x 7 = 14
	// 2 x 8 = 16
	// 2 x 9 = 18
	// 2 x 10 = 20
}
