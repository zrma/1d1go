package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-loops/problem", func() {
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleLetsReview()

			close(done)
		}()
	})
})

func ExampleLetsReview() {
	LetsReview([]string{"Hacker", "Rank"})
	// Output:
	// Hce akr
	// Rn ak
}
