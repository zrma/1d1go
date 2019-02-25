package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-abstract-classes/problem", func() {
	//noinspection SpellCheckingInspection
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleAbstractClasses()

			close(done)
		}()
	})
})

//noinspection SpellCheckingInspection
func ExampleAbstractClasses() {
	AbstractClasses("The Alchemist", "Paulo Coelho", 248)
	// Output:
	// Title: The Alchemist
	// Author: Paulo Coelho
	// Price: 248
}
