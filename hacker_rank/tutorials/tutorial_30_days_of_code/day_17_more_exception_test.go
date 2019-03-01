package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-more-exceptions/problem", func() {
	//noinspection SpellCheckingInspection
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleMoreException()

			close(done)
		}()
	})
})

//noinspection SpellCheckingInspection
func ExampleMoreException() {
	MoreException(3, 5)
	MoreException(2, 4)
	MoreException(-1, -2)
	MoreException(-1, 3)
	// Output:
	// 243
	// 16
	// n and p should be non-negative
	// n and p should be non-negative
}
