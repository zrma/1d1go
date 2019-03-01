package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-exceptions-string-to-integer/problem", func() {
	//noinspection SpellCheckingInspection
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleStringToInteger()

			close(done)
		}()
	})
})

//noinspection SpellCheckingInspection
func ExampleStringToInteger() {
	StringToInteger("3")
	StringToInteger("za")
	// Output:
	// 3
	// Bad String
}
