package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-operators/problem", func() {
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleOperators()

			close(done)
		}()
	})
})

func ExampleOperators() {
	Operators(12.00, 20, 8)
	// Output:
	// 15
}
