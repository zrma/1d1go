package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-data-types/problem", func() {
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleDataType()

			close(done)
		}()
	})
})

func ExampleDataType() {
	DataType(
		4, 12,
		4.0, 4.0,
		"HackerRank ", "is the best place to learn and practice coding!",
	)
	// Output:
	// 16
	// 8.0
	// HackerRank is the best place to learn and practice coding!
}
