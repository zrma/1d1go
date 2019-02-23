package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-binary-numbers/problem", func() {
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleBinaryNumbers()

			close(done)
		}()
	})
})

func ExampleBinaryNumbers() {
	BinaryNumbers(5)
	BinaryNumbers(13)
	// Output:
	// 1
	// 2
}
